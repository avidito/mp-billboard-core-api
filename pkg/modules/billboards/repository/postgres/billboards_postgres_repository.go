package postgres

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type BillboardsRepositoryImpl struct {
	db *gorm.DB
}

func NewBillboardsPostgresRepositoryImpl(db *gorm.DB) domain.BillboardsRepository {
	return &BillboardsRepositoryImpl{
		db: db,
	}
}

// Implementation
func (r BillboardsRepositoryImpl) Create(billboard domain.Billboard) (domain.Billboard, error) {
	query := `
		INSERT INTO billboards (
			billboard,
			created_dt,
			modified_dt
		)
		VALUES (?, NOW() AT TIME ZONE 'Asia/Jakarta', NOW() AT TIME ZONE 'Asia/Jakarta')
		RETURNING *
	`

	row := r.db.Raw(
		query,
		billboard.Billboard,
	).Row()

	var createdBillboard domain.Billboard
	err := row.Scan(
		&createdBillboard.ID,
		&createdBillboard.Billboard,
		&createdBillboard.CreatedDt,
		&createdBillboard.ModifiedDt,
	)

	if err != nil {
		return domain.Billboard{}, err
	}

	return createdBillboard, nil
}

func (r BillboardsRepositoryImpl) GetByID(id int64) (domain.BillboardRead, error) {
	query := `
		SELECT
			id,
			billboard
		FROM billboards
		WHERE id = ?
	`

	row := r.db.Raw(
		query,
		id,
	).Row()

	var billboardRead domain.BillboardRead
	err := row.Scan(
		&billboardRead.ID,
		&billboardRead.Billboard,
	)

	if err != nil {
		return domain.BillboardRead{}, err
	}

	return billboardRead, nil
}

func (r BillboardsRepositoryImpl) Fetch() ([]domain.BillboardRead, error) {
	query := `
		SELECT
			id,
			billboard
		FROM billboards
		ORDER BY created_dt ASC
	`

	rows, err := r.db.Raw(
		query,
	).Rows()
	if err != nil {
		return nil, err
	}

	var billboardReadList []domain.BillboardRead
	var tmpBillboardRead domain.BillboardRead
	for rows.Next() {
		rows.Scan(
			&tmpBillboardRead.ID,
			&tmpBillboardRead.Billboard,
		)

		billboardReadList = append(billboardReadList, tmpBillboardRead)
	}

	return billboardReadList, nil
}

func (r BillboardsRepositoryImpl) Update(id int64, billboard domain.Billboard) (domain.Billboard, error) {
	query := `
		UPDATE billboards
		SET
			billboard = ?,
			modified_dt = NOW() AT TIME ZONE 'Asia/Jakarta'
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		billboard.Billboard,
		id,
	).Row()

	var updatedBillboard domain.Billboard
	err := row.Scan(
		&updatedBillboard.ID,
		&updatedBillboard.Billboard,
		&updatedBillboard.CreatedDt,
		&updatedBillboard.ModifiedDt,
	)

	if err != nil {
		return domain.Billboard{}, err
	}

	return updatedBillboard, nil
}

func (r BillboardsRepositoryImpl) Delete(id int64) (domain.Billboard, error) {
	query := `
		DELETE FROM billboards
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		id,
	).Row()

	var deletedBillboard domain.Billboard
	err := row.Scan(
		&deletedBillboard.ID,
		&deletedBillboard.Billboard,
		&deletedBillboard.CreatedDt,
		&deletedBillboard.ModifiedDt,
	)

	if err != nil {
		return domain.Billboard{}, err
	}

	return deletedBillboard, nil
}
