package main

import (
	"github.com/babymechanic/moteclient/client"
)

func main() {
	client.New("127.0.0.1", 8080).Start()
}
