package client

import (
	"fmt"
	"github.com/babymechanic/moteclient/components"
	"github.com/babymechanic/moteclient/requestprocessor"
	"github.com/golang/glog"
	"net/rpc"
)

type MoteClient struct {
	serverIp         string
	serverPort       int
	requestProcessor *requestprocessor.RequestProcessor
}

func New(serverIp string, serverPort int, listenPort int) *MoteClient {
	return &MoteClient{serverIp: serverIp, serverPort: serverPort, requestProcessor: requestprocessor.New(listenPort)}
}

func (moteClient *MoteClient) Start() {
	client := moteClient.getServerConnection()
	moteClient.initializeComponents(client)
	moteClient.handleRequests()
}

func (moteClient *MoteClient) handleRequests() {
	moteClient.requestProcessor.ProcessRequests()
}

func (moteClient *MoteClient) initializeComponents(client *rpc.Client) {
	for _, component := range components.AllComponents() {
		component.Initialize(client)
	}
}

func (moteClient *MoteClient) getServerConnection() *rpc.Client {
	serverDetails := fmt.Sprintf("%v:%v", moteClient.serverIp, moteClient.serverPort)
	fmt.Println("connecting to server on : ", serverDetails)
	client, err := rpc.DialHTTP("tcp", serverDetails)
	if err != nil {
		glog.Fatalf("dialing:", err)
		fmt.Println("dialing:", err)
	}
	fmt.Println("started client and connected to server on", moteClient.serverIp, ":", moteClient.serverPort)
	return client
}
