package domain

import (
	"time"

	"github.com/avidito/mp-billboard-core-api/pkg/common/types"
)

// Entity
type BillboardVersion struct {
	ID          int64              `json:"id"`
	BillboardID int64              `json:"billboard_id"`
	DesignID    int64              `json:"design_id"`
	StatusID    int64              `json:"status_id"`
	Name        string             `json:"name"`
	Notes       string             `json:"notes"`
	PeriodStart types.DateStandard `json:"period_start"`
	PeriodEnd   types.DateStandard `json:"period_end"`
	Address     string             `json:"address"`
	Amount      int64              `json:"amount"`
	CreatedDt   time.Time          `json:"created_dt"`
	ModifiedDt  time.Time          `json:"modified_dt"`
	Billboard   Billboard          `gorm:"foreignKey:BillboardID;references:ID"`
	Design      Design             `gorm:"foreignKey:DesignID;references:ID"`
	Status      Status             `gorm:"foreignKey:StatusID;references:ID"`
}

type BillboardVersionRead struct {
	ID          int64              `json:"id"`
	Billboard   string             `json:"billboard"`
	Design      string             `json:"design"`
	Status      string             `json:"status"`
	Name        string             `json:"name"`
	Notes       string             `json:"notes"`
	PeriodStart types.DateStandard `json:"period_start"`
	PeriodEnd   types.DateStandard `json:"period_end"`
	Address     string             `json:"address"`
	Amount      int64              `json:"amount"`
}

// Repository
type BillboardVersionsRepository interface {
	Create(billboardVersion BillboardVersion) (BillboardVersion, error)
	GetByID(id int64) (BillboardVersionRead, error)
	Fetch(billboard_id int64) ([]BillboardVersionRead, error)
	Update(id int64, billboard BillboardVersion) (BillboardVersion, error)
	Delete(id int64) (BillboardVersion, error)
}

// Usecase
type BillboardVersionsUsecase interface {
	Create(billboardVersion BillboardVersion) (BillboardVersion, error)
	GetByID(id int64) (BillboardVersionRead, error)
	Fetch(billboard_id int64) ([]BillboardVersionRead, error)
	Update(id int64, billboard BillboardVersion) (BillboardVersion, error)
	Delete(id int64) (BillboardVersion, error)
}
