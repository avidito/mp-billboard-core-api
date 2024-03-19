package domain

import (
	"time"
)

// Entity
type Contract struct {
	ID                 int64            `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	BillboardVersionID int64            `json:"billboard_version_id"`
	ContractTypeID     int64            `json:"contract_type_id"`
	Name               string           `json:"name"`
	Type               string           `json:"type"`
	Description        string           `json:"description"`
	Filepath           string           `json:"filepath"`
	CreatedDt          time.Time        `json:"created_dt"`
	ModifiedDt         time.Time        `json:"modified_dt"`
	BillboardVersion   BillboardVersion `gorm:"foreignKey:BillboardVersionID;references:ID"`
	ContractType       ContractType     `gorm:"foreignKey:ContractTypeID;references:ID"`
}

type ContractRead struct {
	ID           int64  `json:"id"`
	VersionName  string `json:"version_name"`
	ContractType string `json:"contract_type"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Description  string `json:"description"`
	Filepath     string `json:"filepath"`
}

// Repository
type ContractsRepository interface {
	Create(contract Contract) (Contract, error)
	GetByID(id int64) (ContractRead, error)
	Fetch(billboard_version_id int64) ([]ContractRead, error)
	Update(id int64, contract Contract) (Contract, error)
	Delete(id int64) (Contract, error)
}

// Usecase
type ContractsUsecase interface {
	Create(contract Contract) (Contract, error)
	GetByID(id int64) (ContractRead, error)
	Fetch(billboard_version_id int64) ([]ContractRead, error)
	Update(id int64, contract Contract) (Contract, error)
	Delete(id int64) (Contract, error)
}
