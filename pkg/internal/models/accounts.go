package models

import (
	"time"

	"gorm.io/datatypes"
)

type Account struct {
	BaseModel

	Name         string            `json:"name" gorm:"uniqueIndex"`
	Nick         string            `json:"nick"`
	Description  string            `json:"description"`
	Avatar       string            `json:"avatar"`
	Banner       string            `json:"banner"`
	EmailAddress string            `json:"email"`
	ConfirmedAt  *time.Time        `json:"confirmed_at"`
	SuspendedAt  *time.Time        `json:"suspended_at"`
	PermNodes    datatypes.JSONMap `json:"perm_nodes"`
}
