package domain

import (
	"time"
)

// Entity
type Billboard struct {
	ID         int64     `json:"id"`
	Billboard  string    `json:"billboard"`
	CreatedDt  time.Time `json:"created_dt"`
	ModifiedDt time.Time `json:"modified_dt"`
}

type BillboardRead struct {
	ID        int64  `json:"id"`
	Billboard string `json:"billboard"`
}

// Repository
type BillboardsRepository interface {
	Create(billboard Billboard) (Billboard, error)
	GetByID(id int64) (BillboardRead, error)
	Fetch() ([]BillboardRead, error)
	Update(id int64, billboard Billboard) (Billboard, error)
	Delete(id int64) (Billboard, error)
}

// Usecase
type BillboardsUsecase interface {
	Create(billboard Billboard) (Billboard, error)
	GetByID(id int64) (BillboardRead, error)
	Fetch() ([]BillboardRead, error)
	Update(id int64, billboard Billboard) (Billboard, error)
	Delete(id int64) (Billboard, error)
}
