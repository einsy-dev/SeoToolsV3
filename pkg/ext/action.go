package ext

import (
	"github.com/einsy-dev/SEOTools/pkg/ext/socket"
	t "github.com/einsy-dev/SEOTools/pkg/ext/types"
)

type Action struct{}

func (a *Action) OpenTab(url string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionCreate, Data: t.Data{"url": url}})
}

func (a *Action) UpdateTab(tabId string, url string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionUpdate, Data: t.Data{"tabId": tabId, "url": url}})
}

func (a *Action) CloseTab(tabId string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionClose, Data: t.Data{"tabId": tabId}})
}

func (a *Action) ReloadTab(tabId string, url string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionUpdate, Data: t.Data{"tabId": tabId, "url": url}})
}

func (a *Action) GetTabList() {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionGetTabList})
}

func (a *Action) GetTab(tabId string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionGetTabList, Data: t.Data{"tabId": tabId}})
}
