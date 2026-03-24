package extension

import (
	"strings"

	"github.com/einsy-dev/SEOTools/internal/utils"
)

type index struct {
	Links   []string
	Current int
}

func (e *index) Next() {
	Action.OpenTab("https://google.com")
	if e.Current < len(e.Links)-1 {
		e.Current++
	}
}

func (e *index) Reload() {}

func (e *index) Prev() {
	if e.Current > 0 {
		e.Current--
	}
}

func (e *index) First() {
	e.Current = 0
}

func (e *index) Last() {
	e.Current = len(e.Links) - 1
}

func (e *index) Upload(str string) {
	if str == "" {
		e.Links = []string{}
	} else {
		e.Links = strings.Split(str, "\n")
	}
	e.Current = 0
}

func (e *index) State() State {
	if len(e.Links) == 0 {
		return State{Links: []string{}}
	}

	return State{Links: e.Links, Index: e.Current, Current: e.Links[e.Current], CurrentDomain: utils.GetDomain(e.Links[e.Current])}
}

type State struct {
	Links         []string
	Index         int
	Current       string
	CurrentDomain string
}

var Bind = &index{}
