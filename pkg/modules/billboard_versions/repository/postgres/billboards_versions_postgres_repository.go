package postgres

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	"gorm.io/gorm"
)

// Define
type BillboardVersionsRepositoryImpl struct {
	db *gorm.DB
}

func NewBillboardVersionsPostgresRepositoryImpl(db *gorm.DB) domain.BillboardVersionsRepository {
	return &BillboardVersionsRepositoryImpl{
		db: db,
	}
}

// Implementation
func (r BillboardVersionsRepositoryImpl) Create(billboardVersion domain.BillboardVersion) (domain.BillboardVersion, error) {
	query := `
		INSERT INTO billboard_versions (
		  billboard_id,
		  design_id,
		  status_id,
		  name,
		  notes,
		  period_start,
		  period_end,
		  address,
		  amount,
		  created_dt,
		  modified_dt
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW() AT TIME ZONE 'Asia/Jakarta', NOW() AT TIME ZONE 'Asia/Jakarta')
		RETURNING *
	`

	row := r.db.Raw(
		query,
		billboardVersion.BillboardID,
		billboardVersion.DesignID,
		billboardVersion.StatusID,
		billboardVersion.Name,
		billboardVersion.Notes,
		billboardVersion.PeriodStart,
		billboardVersion.PeriodEnd,
		billboardVersion.Address,
		billboardVersion.Amount,
	).Row()

	var createdBillboardVersion domain.BillboardVersion
	err := row.Scan(
		&createdBillboardVersion.ID,
		&createdBillboardVersion.BillboardID,
		&createdBillboardVersion.DesignID,
		&createdBillboardVersion.StatusID,
		&createdBillboardVersion.Name,
		&createdBillboardVersion.Notes,
		&createdBillboardVersion.PeriodStart,
		&createdBillboardVersion.PeriodEnd,
		&createdBillboardVersion.Address,
		&createdBillboardVersion.Amount,
		&createdBillboardVersion.CreatedDt,
		&createdBillboardVersion.ModifiedDt,
	)

	if err != nil {
		return domain.BillboardVersion{}, err
	}

	return createdBillboardVersion, nil
}

func (r BillboardVersionsRepositoryImpl) GetByID(id int64) (domain.BillboardVersionRead, error) {
	query := `
		SELECT
		  bv.id,
		  b.billboard,
		  d.design,
		  s.status,
		  bv.name,
		  bv.period_start,
		  bv.period_end,
		  bv.address,
		  bv.amount
		FROM billboard_versions AS bv
		JOIN billboards AS b
		  ON bv.billboard_id = b.id
		JOIN designs AS d
		  ON bv.design_id = d.id
		JOIN statuses AS s
		  ON bv.status_id = s.id
		WHERE bv.id = ?
	`

	row := r.db.Raw(
		query,
		id,
	).Row()

	var billboardVersionRead domain.BillboardVersionRead
	err := row.Scan(
		&billboardVersionRead.ID,
		&billboardVersionRead.Billboard,
		&billboardVersionRead.Design,
		&billboardVersionRead.Status,
		&billboardVersionRead.Status,
		&billboardVersionRead.PeriodStart,
		&billboardVersionRead.PeriodEnd,
		&billboardVersionRead.Address,
		&billboardVersionRead.Amount,
	)

	if err != nil {
		return domain.BillboardVersionRead{}, err
	}

	return billboardVersionRead, nil
}

func (r BillboardVersionsRepositoryImpl) Fetch(billboard_id int64) ([]domain.BillboardVersionRead, error) {
	query := `
		SELECT
		  bv.id,
		  b.billboard,
		  d.design,
		  s.status,
		  bv.name,
		  bv.period_start,
		  bv.period_end,
		  bv.address,
		  bv.amount
		FROM billboard_versions AS bv
		JOIN billboards AS b
		  ON bv.billboard_id = b.id
		JOIN designs AS d
	      ON bv.design_id = d.id
		JOIN statuses AS s
		  ON bv.status_id = s.id
		WHERE TRUE
	`

	values := make([]interface{}, 0)
	if billboard_id != 0 {
		query += "\n  AND b.id = ?"
		values = append(values, billboard_id)
	}

	rows, err := r.db.Raw(
		query,
		values...,
	).Rows()
	if err != nil {
		return nil, err
	}

	var billboardVersionReadList []domain.BillboardVersionRead
	var tmpBillboardVersionRead domain.BillboardVersionRead
	for rows.Next() {
		rows.Scan(
			&tmpBillboardVersionRead.ID,
			&tmpBillboardVersionRead.Billboard,
			&tmpBillboardVersionRead.Design,
			&tmpBillboardVersionRead.Status,
			&tmpBillboardVersionRead.Notes,
			&tmpBillboardVersionRead.PeriodStart,
			&tmpBillboardVersionRead.PeriodEnd,
			&tmpBillboardVersionRead.Address,
			&tmpBillboardVersionRead.Amount,
		)

		billboardVersionReadList = append(billboardVersionReadList, tmpBillboardVersionRead)
	}

	return billboardVersionReadList, nil
}

func (r BillboardVersionsRepositoryImpl) Update(id int64, billboardVersion domain.BillboardVersion) (domain.BillboardVersion, error) {
	query := `
		UPDATE billboard_versions
		SET
		  billboard_id = ?,
		  design_id = ?,
		  status_id = ?,
		  name = ?,
		  notes = ?,
		  period_start = ?,
		  period_end = ?,
		  address = ?,
		  amount = ?,
		  modified_dt = NOW() AT TIME ZONE 'Asia/Jakarta'
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		billboardVersion.BillboardID,
		billboardVersion.DesignID,
		billboardVersion.StatusID,
		billboardVersion.Name,
		billboardVersion.Notes,
		billboardVersion.PeriodStart,
		billboardVersion.PeriodEnd,
		billboardVersion.Address,
		billboardVersion.Amount,
		id,
	).Row()

	var updatedBillboardVersion domain.BillboardVersion
	err := row.Scan(
		&updatedBillboardVersion.ID,
		&updatedBillboardVersion.BillboardID,
		&updatedBillboardVersion.DesignID,
		&updatedBillboardVersion.StatusID,
		&updatedBillboardVersion.Name,
		&updatedBillboardVersion.Notes,
		&updatedBillboardVersion.PeriodStart,
		&updatedBillboardVersion.PeriodEnd,
		&updatedBillboardVersion.Address,
		&updatedBillboardVersion.Amount,
		&updatedBillboardVersion.CreatedDt,
		&updatedBillboardVersion.ModifiedDt,
	)

	if err != nil {
		return domain.BillboardVersion{}, err
	}

	return updatedBillboardVersion, nil
}

func (r BillboardVersionsRepositoryImpl) Delete(id int64) (domain.BillboardVersion, error) {
	query := `
		DELETE FROM billboard_versions
		WHERE id = ?
		RETURNING *
	`

	row := r.db.Raw(
		query,
		id,
	).Row()

	var deletedBillboardVersion domain.BillboardVersion
	err := row.Scan(
		&deletedBillboardVersion.ID,
		&deletedBillboardVersion.BillboardID,
		&deletedBillboardVersion.DesignID,
		&deletedBillboardVersion.StatusID,
		&deletedBillboardVersion.Name,
		&deletedBillboardVersion.Notes,
		&deletedBillboardVersion.PeriodStart,
		&deletedBillboardVersion.PeriodEnd,
		&deletedBillboardVersion.Address,
		&deletedBillboardVersion.Amount,
		&deletedBillboardVersion.CreatedDt,
		&deletedBillboardVersion.ModifiedDt,
	)

	if err != nil {
		return domain.BillboardVersion{}, err
	}

	return deletedBillboardVersion, nil
}
