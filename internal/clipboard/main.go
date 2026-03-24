package clipboard

import (
	"golang.design/x/clipboard"
)

func Startup() {
	var err = clipboard.Init()
	if err != nil {
		panic(err)
	}
}
