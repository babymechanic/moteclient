package utils

import (
	"fmt"
	"github.com/golang/glog"
	"net/rpc"
)

type RPCInvoker struct {
	client *rpc.Client
}

func NewRPCInvoker(client *rpc.Client) *RPCInvoker {
	return &RPCInvoker{client}
}

func (invoker RPCInvoker) Invoke(operation string, request interface{}, response interface{}) error {
	err := invoker.client.Call(operation, request, response)
	if err != nil {
		fmt.Println("error response: %s", err)
		glog.Fatalf("error response: %s", err)
	}
	return err
}
