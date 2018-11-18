package ticktick

import "time"

// A Project is a named collection of Tasks.
type Project struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	IsOwner      bool      `json:"isOwner"`
	Color        string    `json:"color"`
	InAll        bool      `json:"inAll"`
	UserCount    int       `json:"userCount"`
	ModifiedTime time.Time `json:"modifiedTime"`
	Closed       bool      `json:"closed"`
	GroupID      string    `json:"groupId"`
}
