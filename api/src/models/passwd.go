package models

// Represents the format of the update password
type Passwd struct {
	Old_passwd string `json:"oldpasswd,omitempty"`
	New_passwd string `json:"newpasswd,omitempty"`
}
