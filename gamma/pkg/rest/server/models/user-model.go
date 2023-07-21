package models

type User struct {
	Id int64 `gorm:"primaryKey;autoIncrement"`

	Username string `json:"username,omitempty"`
}
