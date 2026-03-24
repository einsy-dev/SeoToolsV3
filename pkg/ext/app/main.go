package app

import (
	t "github.com/einsy-dev/SEOTools/pkg/ext/types"
)

var Clients = make(t.Clients)
var Ch chan t.EventMessage = make(chan t.EventMessage)
