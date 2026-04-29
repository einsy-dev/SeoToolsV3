package actions

import (
	"github.com/einsy-dev/NetSail/pkg/knot/socket"
	t "github.com/einsy-dev/NetSail/pkg/knot/types"
)

type SocketActions struct{}

func (a *SocketActions) OpenTab(url string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionCreate, Data: t.Data{"url": url}})
}

func (a *SocketActions) UpdateTab(tabId string, url string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionUpdate, Data: t.Data{"tabId": tabId, "url": url}})
}

func (a *SocketActions) CloseTab(tabId string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionClose, Data: t.Data{"tabId": tabId}})
}

func (a *SocketActions) ReloadTab(tabId string, url string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionUpdate, Data: t.Data{"tabId": tabId, "url": url}})
}

func (a *SocketActions) GetTabList() {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionGetTabList})
}

func (a *SocketActions) GetTab(tabId string) {
	socket.SendMessageAll(t.ActionMessage{Action: t.ActionGetTabList, Data: t.Data{"tabId": tabId}})
}
