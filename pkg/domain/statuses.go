package domain

import (
	"time"
)

// Entity
type Status struct {
	ID         int64     `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Status     string    `json:"status" gorm:"unique"`
	CreatedDt  time.Time `json:"created_dt"`
	ModifiedDt time.Time `json:"modified_dt"`
}
