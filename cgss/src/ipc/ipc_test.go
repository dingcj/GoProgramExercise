package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, request string) *Response {

	resp := &Response{"200", "ECHO:" + request}

	return resp
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	resp1, _ := client1.Call("", "Hello")
	resp2, _ := client2.Call("", "World")

	if resp1.Code != "200" || resp1.Body != "ECHO:Hello" ||
		resp2.Code != "200" || resp2.Body != "ECHO:World" {
		t.Error("IpcClient.Call failed. resp1: ", resp1, ", resp2:", resp2)
	}

	client1.Close()
	client2.Close()
}
