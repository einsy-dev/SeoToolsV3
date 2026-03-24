package extension

import (
	"fmt"

	"github.com/einsy-dev/SEOTools/pkg/ext"
)

type Extension struct{}

var Ext = &Extension{}

var Action ext.Action
var Ch <-chan ext.EventMessage

func Startup() {
	Action, Ch = ext.Startup()

	for v := range Ch {
		fmt.Println(v) // Process received value v
	}
}
