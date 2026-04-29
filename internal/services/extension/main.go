package extension

import (
	"context"
	"fmt"

	"github.com/einsy-dev/NetSail/pkg/knot"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type Extension struct{}

func (s *Extension) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	knot.Startup()
	go func() {
		for value := range knot.Ch {
			fmt.Println("data", value.Data, "ping", value.Ping, "status", value.Status)
		}
	}()

	return nil
}
