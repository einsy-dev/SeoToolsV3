package socket

import (
	"encoding/json"

	t "github.com/einsy-dev/NetSail/pkg/knot/types"
	"github.com/gorilla/websocket"
)

func SendMessageAll(message t.ActionMessage) {
	for id := range Clients {
		SendMessage(id, message)
	}
}

func SendMessage(id string, message t.ActionMessage) {
	if ms, err := json.Marshal(message); err == nil {
		err = Clients[id].WriteMessage(websocket.TextMessage, []byte(ms))
		if err != nil {
			delete(Clients, id)
		}
	}
}
