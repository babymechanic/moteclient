package components

import (
	"github.com/babymechanic/moteclient/utils"
	"github.com/babymechanic/motecommon/messages"
	"net/rpc"
	"reflect"
)

type Accelerometer struct {
	invoker *utils.RPCInvoker
}

func (accelerometer *Accelerometer) Initialize(client *rpc.Client) {
	accelerometer.invoker = utils.NewRPCInvoker(client)
}

func (accelerometer *Accelerometer) Heading() messages.Heading {
	var response messages.Heading
	accelerometer.invoker.Invoke("Accelerometer.Heading", struct{}{}, &response)
	return response
}

func init() {
	Register("Accelerometer.Heading", &Accelerometer{}, reflect.TypeOf(nil), reflect.TypeOf(messages.Heading{}))
}
