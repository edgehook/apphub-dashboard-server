package v1

import ()

type ProjectData struct {
	Name        string `form:"name" json:"name,omitempty"`
	Description string `form:"description" json:"description,omitempty"`
	Content     string `form:"content" json:"content,omitempty"`
	State       *int   `form:"state" json:"state,omitempty"`
}
