package services

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v3/pkg/application"
	"golang.design/x/clipboard"
)

type Clipboard struct{}

func (s *Clipboard) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	fmt.Println("Clipboard startup")
	var err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	// Initialisation
	return nil
}

func (s *Clipboard) CopyText(text string) {
	var err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
}
