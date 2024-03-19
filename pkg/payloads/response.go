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

type ContractResponse struct {
	ID                 int64     `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	BillboardVersionID int64     `json:"billboard_version_id"`
	ContractTypeID     int64     `json:"contract_type_id"`
	Name               string    `json:"name"`
	Type               string    `json:"type"`
	Description        string    `json:"description"`
	Filepath           string    `json:"filepath"`
	CreatedDt          time.Time `json:"created_dt"`
	ModifiedDt         time.Time `json:"modified_dt"`
}

type ContractReadResponse struct {
	ID           int64  `json:"id"`
	VersionName  string `json:"version_name"`
	ContractType string `json:"contract_type"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Filepath     string `json:"filepath"`
}
