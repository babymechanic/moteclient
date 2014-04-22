package handlers

import (
	"github.com/babymechanic/moteclient/components"
	"reflect"
	"strings"
)

type CommandExecutor struct {
}

func (executor *CommandExecutor) Execute(operation string, request interface{}) interface{} {
	operationDetails := components.GetOperationsDetails(operation)
	valueOfComponent := reflect.ValueOf(operationDetails.Component)
	methodName := strings.Split(operation, ".")[1]
	method := valueOfComponent.MethodByName(methodName)
	inputs := make([]reflect.Value, 0)
	if request != nil {
		inputs = make([]reflect.Value, 1)
		inputs[0] = reflect.ValueOf(request)
	}
	ouputs := method.Call(inputs)
	if len(ouputs) == 0 {
		return nil
	}
	return ouputs[0].Interface()
}
