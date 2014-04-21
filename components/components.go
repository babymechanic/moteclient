package components

import (
	"net/rpc"
	"reflect"
)

type Component interface {
	Initialize(client *rpc.Client)
}

type OperationDetails struct {
	Component Component
	Request   reflect.Type
	Response  reflect.Type
}

type Operations map[string]OperationDetails

var operations = Operations{}

func Register(operation string, component Component, requestType reflect.Type, responseType reflect.Type) {
	operations[operation] = OperationDetails{Component: component, Request: requestType, Response: responseType}
}

func AllComponents() []Component {
	components := make([]Component, len(operations))
	componentIndex := 0
	for _, value := range operations {
		components[componentIndex] = value.Component
		componentIndex++
	}
	return components
}

func GetOperationsDetails(operation string) OperationDetails {
	return operations[operation]
}
