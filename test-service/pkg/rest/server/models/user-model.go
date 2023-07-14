package models

type User struct {
	Id int64 `json:"id,omitempty"`

	Password `json:"password,omitempty"`

	Username `json:"username,omitempty"`
}
