package usecase

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
)

// Define
type BillboardsUsecaseImpl struct {
	repo domain.BillboardsRepository
}

func NewBillboardsUsecaseImpl(r domain.BillboardsRepository) domain.BillboardsUsecase {
	return &BillboardsUsecaseImpl{
		repo: r,
	}
}

// Implementation
func (u BillboardsUsecaseImpl) Create(billboard domain.Billboard) (domain.Billboard, error) {
	createdBillboard, err := u.repo.Create(billboard)
	return createdBillboard, err
}

func (u BillboardsUsecaseImpl) GetByID(id int64) (domain.BillboardRead, error) {
	billboardRead, err := u.repo.GetByID(id)
	return billboardRead, err
}

func (u BillboardsUsecaseImpl) Fetch() ([]domain.BillboardRead, error) {
	billboardReadList, err := u.repo.Fetch()
	return billboardReadList, err
}

func (u BillboardsUsecaseImpl) Update(id int64, billboard domain.Billboard) (domain.Billboard, error) {
	updatedBillboard, err := u.repo.Update(id, billboard)
	return updatedBillboard, err
}

func (u BillboardsUsecaseImpl) Delete(id int64) (domain.Billboard, error) {
	deletedBillboard, err := u.repo.Delete(id)
	return deletedBillboard, err
}
