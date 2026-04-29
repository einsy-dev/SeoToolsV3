package knot

import (
	"log"
	"net/http"

	"github.com/einsy-dev/NetSail/pkg/knot/actions"
	"github.com/einsy-dev/NetSail/pkg/knot/socket"
	"github.com/einsy-dev/NetSail/pkg/knot/types"
)

var SocketActions actions.SocketActions
var Ch <-chan types.SocketEvent

type EventMessage = types.EventMessage

func Startup() {
	SocketActions = actions.SocketActions{}
	ch := make(chan types.SocketEvent)
	Ch = ch

	http.HandleFunc("/", socket.Handler(ch))

	go func() {
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

}
