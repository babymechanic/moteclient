package main

import (
	"fmt"
	"github.com/babymechanic/client/components"
	"github.com/babymechanic/client/utils"
	"github.com/golang/glog"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8080")
	if err != nil {
		glog.Fatalf("dialing:", err)
	}
	display := components.NewDisplay(utils.NewRPCInvoker(client))
	resolution := display.Resolution()
	fmt.Printf("Resolution: %d x %d", resolution.Width, resolution.Height)
}
