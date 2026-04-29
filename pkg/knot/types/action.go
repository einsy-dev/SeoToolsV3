package types

type SocketAction struct {
	Connect *ConnectAction `json:"connect,omitempty"`
	Tab     *TabActions    `json:"tab,omitempty"`
}

type ConnectAction struct {
	ID string `json:"id"`
}

type TabActions struct {
	Get    *GetTabAction    `json:"get,omitempty"`
	GetAll []int            `json:"getAll,omitempty"`
	Create *CreateTabAction `json:"create,omitempty"`
	Update *UpdateTabAction `json:"update,omitempty"`
	Close  *CloseTabAction  `json:"close,omitempty"`
}

type GetTabAction struct {
	TabID int `json:"tabId"`
}

type CreateTabAction struct {
	URL    string `json:"url"`
	Active bool   `json:"active"`
}

type UpdateTabAction struct {
	TabID  int    `json:"tabId"`
	URL    string `json:"url"`
	Active bool   `json:"active"`
}

type CloseTabAction struct {
	TabID int `json:"tabId"`
}
