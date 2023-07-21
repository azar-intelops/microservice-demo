package models

type Device struct {
	Id int64 `gorm:"primaryKey;autoIncrement"`

	Year uint32 `json:"year,omitempty"`
}
