package postgres

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type ContractsRepositoryImpl struct {
	db *gorm.DB
}

func NewContractsPostgresRepositoryImpl(db *gorm.DB) domain.ContractsRepository {
	return &ContractsRepositoryImpl{
		db: db,
	}
}

// Implementation
func (r ContractsRepositoryImpl) Create(contract domain.Contract) (domain.Contract, error) {
	query := `
		INSERT INTO contracts (
		  billboard_version_id,
		  contract_type_id,
		  name,
		  type,
		  description,
		  filepath,
		  created_dt,
		  modified_dt
		)
		VALUES (?, ?, ?, ?, ?, ?, NOW() AT TIME ZONE 'Asia/Jakarta', NOW() AT TIME ZONE 'Asia/Jakarta')
		RETURNING *
	`

	row := r.db.Raw(
		query,
		contract.BillboardVersionID,
		contract.ContractTypeID,
		contract.Name,
		contract.Type,
		contract.Description,
		contract.Filepath,
	).Row()

	var createdContract domain.Contract
	err := row.Scan(
		&createdContract.ID,
		&createdContract.BillboardVersionID,
		&createdContract.ContractTypeID,
		&createdContract.Name,
		&createdContract.Type,
		&createdContract.Description,
		&createdContract.Filepath,
		&createdContract.CreatedDt,
		&createdContract.ModifiedDt,
	)

	if err != nil {
		return domain.Contract{}, err
	}

	return createdContract, nil
}

func (r ContractsRepositoryImpl) GetByID(id int64) (domain.ContractRead, error) {
	query := `
		SELECT
		  c.id,
		  bv.name AS version_name,
		  ct.contract_type,
		  c.name,
		  c.type,
		  c.description,
		  c.filepath
		FROM contracts AS c
		JOIN billboard_versions AS bv
		  ON c.billboard_version_id = bv.id
		JOIN contract_types AS ct
		  ON c.contract_type_id = ct.id
		WHERE c.id = ?
	`

	row := r.db.Raw(
		query,
		id,
	).Row()

	var contractRead domain.ContractRead
	err := row.Scan(
		&contractRead.ID,
		&contractRead.VersionName,
		&contractRead.ContractType,
		&contractRead.Name,
		&contractRead.Type,
		&contractRead.Description,
		&contractRead.Filepath,
	)

	if err != nil {
		return domain.ContractRead{}, err
	}

	return contractRead, nil
}

func (r ContractsRepositoryImpl) Fetch(billboard_version_id int64) ([]domain.ContractRead, error) {
	query := `
		SELECT
		  c.id,
		  bv.name AS version_name,
		  ct.contract_type,
		  c.name,
		  c.type,
		  c.description,
		  c.filepath
		FROM contracts AS c
		JOIN billboard_versions AS bv
		  ON c.billboard_version_id = bv.id
		JOIN contract_types AS ct
		  ON c.contract_type_id = ct.id
		WHERE TRUE
	`

	values := make([]interface{}, 0)
	if billboard_version_id != 0 {
		query += "\n  AND bv.id = ?"
		values = append(values, billboard_version_id)
	}

	rows, err := r.db.Raw(
		query,
		values...,
	).Rows()
	if err != nil {
		return nil, err
	}

	var contractReadList []domain.ContractRead
	var contractRead domain.ContractRead
	for rows.Next() {
		rows.Scan(
			&contractRead.ID,
			&contractRead.VersionName,
			&contractRead.ContractType,
			&contractRead.Name,
			&contractRead.Type,
			&contractRead.Description,
			&contractRead.Filepath,
		)

		contractReadList = append(contractReadList, contractRead)
	}

	return contractReadList, nil
}

func (r ContractsRepositoryImpl) Update(id int64, contract domain.Contract) (domain.Contract, error) {
	query := `
		UPDATE contracts
		SET
		  billboard_version_id = ?,
		  contract_type_id = ?,
		  name = ?,
		  type = ?,
		  description = ?,
		  filepath = ?,
		  modified_dt = NOW() AT TIME ZONE 'Asia/Jakarta'
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		contract.BillboardVersionID,
		contract.ContractTypeID,
		contract.Name,
		contract.Type,
		contract.Description,
		contract.Filepath,
		id,
	).Row()

	var updatedContract domain.Contract
	err := row.Scan(
		&updatedContract.ID,
		&updatedContract.BillboardVersionID,
		&updatedContract.ContractTypeID,
		&updatedContract.Name,
		&updatedContract.Type,
		&updatedContract.Description,
		&updatedContract.Filepath,
		&updatedContract.CreatedDt,
		&updatedContract.ModifiedDt,
	)

	if err != nil {
		return domain.Contract{}, err
	}

	return updatedContract, nil
}

func (r ContractsRepositoryImpl) Delete(id int64) (domain.Contract, error) {
	query := `
		DELETE FROM contracts
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		id,
	).Row()

	var deletedContract domain.Contract
	err := row.Scan(
		&deletedContract.ID,
		&deletedContract.BillboardVersionID,
		&deletedContract.ContractTypeID,
		&deletedContract.Name,
		&deletedContract.Type,
		&deletedContract.Description,
		&deletedContract.Filepath,
		&deletedContract.CreatedDt,
		&deletedContract.ModifiedDt,
	)

	if err != nil {
		return domain.Contract{}, err
	}

	return deletedContract, nil
}
