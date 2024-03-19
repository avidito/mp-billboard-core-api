package domain

import (
	"time"
)

// Entity
type ContractType struct {
	ID           int64     `json:"id" gorm:"primaryKey"`
	ContractType string    `json:"contract_type"`
	CreatedDt    time.Time `json:"created_dt"`
	ModifiedDt   time.Time `json:"modified_dt"`
}
