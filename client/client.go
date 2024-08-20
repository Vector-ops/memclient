package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net"

	"github.com/tidwall/resp"
)

type Client struct {
	addr string
	conn net.Conn
}

func New(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		addr: addr,
		conn: conn,
	}, nil
}

func (c *Client) Ping(ctx context.Context) (string, error) {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("PING"),
	})
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return "", err
	}

	b := make([]byte, 1024)
	n, err := c.conn.Read(b)
	return string(b[:n]), err
}

func (c *Client) Set(ctx context.Context, key string, val string) error {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("SET"),
		resp.StringValue(key),
		resp.StringValue(val),
	})
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Get(ctx context.Context, key string) (string, error) {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("GET"),
		resp.StringValue(key),
	})
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return "", err
	}

	b := make([]byte, 1024)
	n, err := c.conn.Read(b)
	return string(b[:n]), err
}

func (c *Client) Del(ctx context.Context, key string) error {
	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("DEL"),
		resp.StringValue(key),
	})
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}
	b := make([]byte, 1024)
	n, err := c.conn.Read(b)
	if err != nil {
		return err
	}
	if string(b[:n]) == "key not found" {
		return errors.New(string(b[:n]))
	}
	return nil
}

func (c *Client) Upd(ctx context.Context, key, value string) error {
	err := c.Del(ctx, key)
	if err != nil {
		return err
	}
	err = c.Set(ctx, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAll(ctx context.Context) (map[string]string, error) {

	var buf bytes.Buffer
	wr := resp.NewWriter(&buf)
	wr.WriteArray([]resp.Value{
		resp.StringValue("GETA"),
	})

	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	b := make([]byte, 2048)
	in := make(map[string][]byte)
	n, err := c.conn.Read(b)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b[:n], &in)
	if err != nil {
		return nil, err
	}

	out := make(map[string]string)
	for k, v := range in {
		out[k] = string(v)
	}
	return out, err

	// data := map[string]string{
	// 	"user:1001":               "John Doe",
	// 	"user:1002":               "Jane Smith",
	// 	"product:2001":            "Laptop",
	// 	"product:2002":            "Smartphone",
	// 	"company:name":            "TechCorp",
	// 	"user:1001:name":          "John Doe",
	// 	"user:1001:age":           "28",
	// 	"user:1001:email":         "johndoe@example.com",
	// 	"user:1002:name":          "Jane Smith",
	// 	"user:1002:age":           "34",
	// 	"user:1002:email":         "janesmith@example.com",
	// 	"product:2001:name":       "Laptop",
	// 	"product:2001:brand":      "TechBrand",
	// 	"product:2001:price":      "1200",
	// 	"product:2002:name":       "Smartphone",
	// 	"product:2002:brand":      "PhoneCo",
	// 	"product:2002:price":      "800",
	// 	"company:info:name":       "TechCorp",
	// 	"company:info:founded":    "2008",
	// 	"company:info:employees":  "500",
	// 	"notifications:1001:0":    "You have a new message",
	// 	"notifications:1001:1":    "Your order has been shipped",
	// 	"notifications:1002:0":    "Password changed successfully",
	// 	"notifications:1002:1":    "Welcome to our service",
	// 	"recent_searches:1001:0":  "headphones",
	// 	"recent_searches:1001:1":  "gaming mouse",
	// 	"recent_searches:1002:0":  "smartwatch",
	// 	"recent_searches:1002:1":  "tablet",
	// 	"user:1001:favorites:0":   "product:2001",
	// 	"user:1001:favorites:1":   "product:2002",
	// 	"user:1002:favorites:0":   "product:2002",
	// 	"user:1001:friends:0":     "user:1002",
	// 	"user:1002:friends:0":     "user:1001",
	// 	"leaderboard:user:1001":   "1000",
	// 	"leaderboard:user:1002":   "1500",
	// 	"sales:2024:product:2001": "2000",
	// 	"sales:2024:product:2002": "3000",
	// }
	// return data, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
