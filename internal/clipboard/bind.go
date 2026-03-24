package clipboard

import "golang.design/x/clipboard"

type index struct{}

func (*index) CopyText(text string) {
	var err = clipboard.Init()
	if err != nil {
		panic(err)
	}
	clipboard.Write(clipboard.FmtText, []byte(text))
}

var Bind = &index{}
