package domain

import (
	"time"
)

/**
* 公共字段
 */
type BaseModel struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeleteFlag uint8
}
