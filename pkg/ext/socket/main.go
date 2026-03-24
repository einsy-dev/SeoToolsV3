package socket

import (
	"encoding/json"

	"github.com/einsy-dev/SEOTools/pkg/ext/app"
	t "github.com/einsy-dev/SEOTools/pkg/ext/types"
	"github.com/gorilla/websocket"
)

func SendMessageAll(message t.ActionMessage) {
	for id := range app.Clients {
		SendMessage(id, message)
	}
}

func SendMessage(id string, message t.ActionMessage) {
	if ms, err := json.Marshal(message); err == nil {
		err = app.Clients[id].WriteMessage(websocket.TextMessage, []byte(ms))
		if err != nil {
			delete(app.Clients, id)
		}
	}
}
