package extension

import (
	"context"
	"strings"
	"sync"

	"github.com/einsy-dev/NetSail/internal/utils"
	"github.com/einsy-dev/NetSail/pkg/knot"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type Player struct {
	Links   []string
	Current int
	mu      sync.RWMutex
}

func (s *Player) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	s.Links = []string{}
	s.Current = 0
	return nil
}

func (e *Player) Next() {
	if len(e.Links) == 0 {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	if e.Current < len(e.Links)-1 {
		e.Current++
	}

	knot.SocketActions.OpenTab(e.Links[e.Current])
}

func (e *Player) Reload() {
	if len(e.Links) == 0 {
		return
	}
	knot.SocketActions.ReloadTab("", e.Links[e.Current])
}

func (e *Player) Prev() {
	if len(e.Links) == 0 {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	if e.Current > 0 {
		e.Current--
	}
	knot.SocketActions.OpenTab(e.Links[e.Current])
}

func (e *Player) First() {
	if len(e.Links) == 0 {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.Current = 0
	knot.SocketActions.OpenTab(e.Links[e.Current])
}

func (e *Player) Last() {
	if len(e.Links) == 0 {
		return
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.Current = len(e.Links) - 1
	knot.SocketActions.OpenTab(e.Links[e.Current])
}

func (e *Player) Upload(str string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if str == "" {
		e.Links = []string{}
	} else {
		e.Links = strings.Split(str, "\n")
	}
	e.Current = 0
}

func (e *Player) State() State {
	e.mu.Lock()
	defer e.mu.Unlock()

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
