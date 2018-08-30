package cg

import (
	"ipc"
)

type Message struct {
	From    string "from"
	to      string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player

	rooms []*Room
}
