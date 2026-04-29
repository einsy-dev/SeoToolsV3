package server

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Api struct{}

func (s *Api) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {

	return nil
}
