package main

import (
	"os"
)

func main() {
	client := NewClient()
	err := client.PostMessage(os.Getenv("CHANNEL_ID"), "Hello World", true)
	if err != nil {
		panic(err)
	}
}
