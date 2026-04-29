package types

import "github.com/gorilla/websocket"

type Clients map[string]*websocket.Conn

// Action
type Data = map[string]any

type ActionMessage struct {
	Action string `json:"action"`
	Data   Data   `json:"data"`
}

const (
	ActionConnect    = "connect"
	ActionCreate     = "create"
	ActionUpdate     = "update"
	ActionClose      = "close"
	ActionGetTab     = "getTab"
	ActionGetTabList = "getTabList"
)

// Event
type EventMessage struct {
	Event string
	Data  map[string]any
}

const (
	EventUpdate     = "update"
	EventGetTab     = "getTab"
	EventGetTabList = "getTabList"
)
