package handlers

import (
	"encoding/json"
	"github.com/babymechanic/moteclient/components"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"reflect"
)

type OperationHandler struct {
	executor *CommandExecutor
}

func (operationHandler *OperationHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	jsonMap := getJsonMap(request)
	operationName := getOperationName(jsonMap)
	requestObject := getRequestObject(jsonMap, operationName)
	operationHandler.executor.Execute(operationName, requestObject)
}

func getJsonMap(request *http.Request) map[string]interface{} {
	defer request.Body.Close()
	body, _ := ioutil.ReadAll(request.Body)
	var unMarshalledJson interface{}
	json.Unmarshal(body, &unMarshalledJson)
	jsonMap, _ := unMarshalledJson.(map[string]interface{})
	return jsonMap
}

func getRequestObject(jsonMap map[string]interface{}, operationName string) interface{} {
	operationDetails := components.GetOperationsDetails(operationName)
	if operationDetails.Request == reflect.TypeOf(nil) {
		return nil
	}
	newObject := reflect.New(operationDetails.Request).Interface()
	mapstructure.Decode(jsonMap[operationName], &newObject)
	return newObject
}

func getOperationName(jsonMap map[string]interface{}) string {
	var operationName string
	for key, _ := range jsonMap {
		operationName = key
	}
	return operationName
}

func init() {
	RegisterHandler("/execute", &OperationHandler{&CommandExecutor{}})
}
