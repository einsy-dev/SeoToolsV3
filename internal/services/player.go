package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/einsy-dev/NetSail/internal/utils"
	"github.com/einsy-dev/NetSail/pkg/ext"
	"github.com/wailsapp/wails/v3/pkg/application"
)

var Action ext.Action
var Ch <-chan ext.EventMessage

// Action, Ch = ext.Startup()
type Player struct {
	Links   []string
	Current int
}

func (s *Player) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	fmt.Println("Player startup")
	Action, Ch = ext.Startup()
	return nil
}

func (e *Player) Next() {
	// Action.OpenTab("https://google.com")
	if e.Current < len(e.Links)-1 {
		e.Current++
	}
}

func (e *Player) Reload() {}

func (e *Player) Prev() {
	if e.Current > 0 {
		e.Current--
	}
}

func (e *Player) First() {
	e.Current = 0
}

func (e *Player) Last() {
	e.Current = len(e.Links) - 1
}

func (e *Player) Upload(str string) {
	if str == "" {
		e.Links = []string{}
	} else {
		e.Links = strings.Split(str, "\n")
	}
	e.Current = 0
}

func (e *Player) State() State {
	if len(e.Links) == 0 {
		return State{Links: []string{}}
	}

	return State{Links: e.Links, Player: e.Current, Current: e.Links[e.Current], CurrentDomain: utils.FormatDomain(e.Links[e.Current])}
}

type State struct {
	Links         []string
	Player        int
	Current       string
	CurrentDomain string
}
