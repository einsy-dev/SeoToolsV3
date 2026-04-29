package socket

import (
	"encoding/json"
	"log"
	"net/http"

	t "github.com/einsy-dev/NetSail/pkg/knot/types"
	"github.com/gorilla/websocket"
	"github.com/lucsky/cuid"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // Allow all origins for simplicity
}

var Clients = make(t.Clients)

func Handler(ch chan t.SocketEvent) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		var id string
		id = cuid.New()
		Clients[id] = conn

		if ms, err := json.Marshal(t.SocketAction{Connect: &t.ConnectAction{ID: id}}); err == nil {
			err = conn.WriteMessage(websocket.TextMessage, ms)
		}

		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			var ms = t.SocketEvent{}
			err = json.Unmarshal([]byte(p), &ms)
			if err != nil {
				log.Println(err)
				return
			}

			ch <- ms
		}
	}
}
