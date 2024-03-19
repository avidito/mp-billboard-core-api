package payloads

import (
	"github.com/avidito/mp-billboard-core-api/pkg/common/types"
)

type BillboardRequest struct {
	Billboard string `json:"billboard"`
}

type BillboardVersionRequest struct {
	BillboardID int64              `json:"billboard_id"`
	DesignID    int64              `json:"design_id"`
	StatusID    int64              `json:"status_id"`
	Name        string             `json:"name"`
	Notes       string             `json:"notes"`
	PeriodStart types.DateStandard `json:"period_start"`
	PeriodEnd   types.DateStandard `json:"period_end"`
	Address     string             `json:"address"`
	Amount      int64              `json:"amount"`
}
