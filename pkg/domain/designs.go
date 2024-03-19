package domain

import (
	"time"
)

// Entity
type Design struct {
	ID         int64     `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Design     string    `json:"name" gorm:"unique"`
	CreatedDt  time.Time `json:"created_dt"`
	ModifiedDt time.Time `json:"modified_dt"`
}
