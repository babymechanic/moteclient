package components

import (
	"github.com/babymechanic/moteclient/utils"
	"github.com/babymechanic/motecommon/messages"
	"net/rpc"
	"reflect"
)

type Gyroscope struct {
	invoker *utils.RPCInvoker
}

func (gyroscope *Gyroscope) Initialize(client *rpc.Client) {
	gyroscope.invoker = utils.NewRPCInvoker(client)
}

func (gyroscope *Gyroscope) Orientation() messages.Orientation {
	var response messages.Orientation
	gyroscope.invoker.Invoke("Gyroscope.Orientation", struct{}{}, &response)
	return response
}

func init() {
	Register("Gyroscope.Orientation", &Gyroscope{}, reflect.TypeOf(nil), reflect.TypeOf(messages.Orientation{}))
}
