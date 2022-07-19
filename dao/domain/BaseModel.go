package domain

import (
	"time"
)

/**
* 公共字段
 */
type BaseModel struct {
	ID         int       `json:"id" gorm:"primary_key"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeleteFlag uint8     `json:"deleteFlag"`
}
