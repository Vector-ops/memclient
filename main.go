package main

import (
	"context"
	"fmt"
	"os"
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
	args := os.Args[1:]

	client, err := New(":5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(args) == 0 {
		fmt.Println("No command found.")
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
