package components

import (
	"github.com/babymechanic/moteclient/utils"
	"github.com/babymechanic/motecommon/messages"
	"net/rpc"
	"reflect"
)

type Magnetometer struct {
	invoker *utils.RPCInvoker
}

func (magnetometer *Magnetometer) Initialize(client *rpc.Client) {
	magnetometer.invoker = utils.NewRPCInvoker(client)
}

func (magnetometer *Magnetometer) Heading() messages.Heading {
	var response messages.Heading
	magnetometer.invoker.Invoke("Magnetometer.Heading", struct{}{}, &response)
	return response
}

func init() {
	Register("Magnetometer.Heading", &Magnetometer{}, reflect.TypeOf(nil), reflect.TypeOf(messages.Heading{}))
}
