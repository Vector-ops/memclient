package main

import (
	"context"
	"fmt"
	"os"

	"github.com/danvergara/gocui"
	"github.com/vector-ops/memclient/client"
	"github.com/vector-ops/memclient/gui"
)

const (
	GET  = "GET"
	SET  = "SET"
	PING = "PING"
	GETA = "GETA"
	DEL  = "DEL"
	UPD  = "UPD"
)

func main() {

	g, err := gocui.NewGui(gocui.Output256)
	if err != nil {
		fmt.Println(err)
	}

	args := os.Args[1:]

	client, err := client.New(":5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(args) == 0 {
		cgui := gui.New(g, client)
		err = cgui.Run()
		if err != nil {
			fmt.Printf("\n\n\n\n\n\n\n\n\n")
			panic(err)
		}
		return
	}

	switch args[0] {
	case PING:
		pong, err := client.Ping(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(pong)
	case SET:
		if len(args) < 3 {
			fmt.Println("Not enough arguments found.")
			return
		}
		err = client.Set(context.Background(), args[1], args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s stored in database\n", args[1])

	case GET:
		if len(args) < 2 {
			fmt.Println("Not enough arguments found.")
			return
		}
		value, err := client.Get(context.Background(), args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		if value == "" {
			fmt.Println("Key not found.")
		}
		fmt.Println(value)
	case DEL:
		if len(args) < 2 {
			fmt.Println("Not enough arguments found.")
			return
		}
		err := client.Del(context.Background(), args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Deleted key", args[1])

	case UPD:
		if len(args) < 3 {
			fmt.Println("Not enough arguments found.")
			return
		}
		err = client.Upd(context.Background(), args[1], args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s updated in database\n", args[1])

	case GETA:
		data, err := client.GetAll(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}

		Bubbletea(data)

	}

}
