package types

type SocketEvent struct {
	Id     string             `json:"id"`
	Status *SocketEventStatus `json:"status,omitempty"`
	Data   *SocketEventStatus `json:"data,omitempty"`
	Ping   *string            `json:"ping,omitempty"`
}

type SocketEventStatus struct {
	Active bool
	Status string
	TabId  int
}
