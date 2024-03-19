package payloads

import (
	"time"

	"github.com/avidito/mp-billboard-core-api/pkg/common/types"
)

type BillboardResponse struct {
	ID         int64     `json:"id"`
	Billboard  string    `json:"billboard"`
	CreatedDt  time.Time `json:"created_dt"`
	ModifiedDt time.Time `json:"modified_dt"`
}

type BillboardReadResponse struct {
	ID        int64  `json:"id"`
	Billboard string `json:"billboard"`
}

type BillboardVersionResponse struct {
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
}

type BillboardVersionReadResponse struct {
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
