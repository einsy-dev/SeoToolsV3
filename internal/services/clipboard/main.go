package clipboard

import (
	"context"

	"github.com/wailsapp/wails/v3/pkg/application"
	"golang.design/x/clipboard"
)

type Clipboard struct{}

func (s *Clipboard) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	var err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *Clipboard) CopyText(text string) {
	var err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
}
