package models

type UserToLogin struct {
	Id int64 `json:"id,omitempty"`

	Username string `json:"username,omitempty"`
}
