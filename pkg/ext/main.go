package ext

import (
	"log"
	"net/http"

	"github.com/einsy-dev/SEOTools/pkg/ext/app"
	"github.com/einsy-dev/SEOTools/pkg/ext/socket"
)

func Startup() (Action, <-chan EventMessage) {
	http.HandleFunc("/", socket.Handler)

	go func() {
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	return Action{}, app.Ch
}
