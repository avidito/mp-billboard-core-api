package payloads

import (
	"time"
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
