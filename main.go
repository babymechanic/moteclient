package main

import (
	"flag"
	"github.com/babymechanic/moteclient/client"
)

func main() {
	flag.Parse()
	client.New("127.0.0.1", 8080, 8081).Start()
}
