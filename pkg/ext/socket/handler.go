package socket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/einsy-dev/SEOTools/pkg/ext/app"
	t "github.com/einsy-dev/SEOTools/pkg/ext/types"
	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins for simplicity
}

func Handler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	var id string
	id = cuid.New()
	app.Clients[id] = conn

	if ms, err := json.Marshal(t.ActionMessage{Action: t.ActionConnect, Data: map[string]any{"id": id}}); err == nil {
		err = conn.WriteMessage(websocket.TextMessage, ms)
		fmt.Println("err or", err)
	}

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		var ms = t.EventMessage{}
		err = json.Unmarshal([]byte(p), &ms)
		if err != nil {
			log.Println(err)
			return
		}

		switch ms.Event {
		default:
			app.Ch <- ms
		}
	}
}
