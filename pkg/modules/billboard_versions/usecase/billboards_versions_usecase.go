package usecase

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
)

// Define
type BillboardVersionsUsecaseImpl struct {
	repo domain.BillboardVersionsRepository
}

func NewBillboardVersionsUsecaseImpl(r domain.BillboardVersionsRepository) domain.BillboardVersionsUsecase {
	return &BillboardVersionsUsecaseImpl{
		repo: r,
	}
}

// Implementation
func (u BillboardVersionsUsecaseImpl) Create(billboardVersion domain.BillboardVersion) (domain.BillboardVersion, error) {
	createdBillboardVersion, err := u.repo.Create(billboardVersion)
	return createdBillboardVersion, err
}

func (u BillboardVersionsUsecaseImpl) GetByID(id int64) (domain.BillboardVersionRead, error) {
	billboardVersionRead, err := u.repo.GetByID(id)
	return billboardVersionRead, err
}

func (u BillboardVersionsUsecaseImpl) Fetch(billboard_id int64) ([]domain.BillboardVersionRead, error) {
	billboardVersionReadList, err := u.repo.Fetch(billboard_id)
	return billboardVersionReadList, err
}

func (u BillboardVersionsUsecaseImpl) Update(id int64, billboardVersion domain.BillboardVersion) (domain.BillboardVersion, error) {
	updatedBillboardVersion, err := u.repo.Update(id, billboardVersion)
	return updatedBillboardVersion, err
}

func (u BillboardVersionsUsecaseImpl) Delete(id int64) (domain.BillboardVersion, error) {
	deletedBillboardVersion, err := u.repo.Delete(id)
	return deletedBillboardVersion, err
}
