package components

import (
	"github.com/babymechanic/moteclient/utils"
	"github.com/babymechanic/motecommon/messages"
)

type Display struct {
	invoker *utils.RPCInvoker
}

func NewDisplay(invoker *utils.RPCInvoker) *Display {
	return &Display{invoker}
}

func (display *Display) Resolution() messages.Resolution {
	var response messages.Resolution
	display.invoker.Invoke("Display.Resolution", struct{}{}, &response)
	return response
}

func init() {

}
