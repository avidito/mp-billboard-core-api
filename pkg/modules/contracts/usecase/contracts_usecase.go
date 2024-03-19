package usecase

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
)

// Define
type ContractsUsecaseImpl struct {
	repo domain.ContractsRepository
}

func NewContractsUsecaseImpl(r domain.ContractsRepository) domain.ContractsUsecase {
	return &ContractsUsecaseImpl{
		repo: r,
	}
}

// Implementation
func (u ContractsUsecaseImpl) Create(contract domain.Contract) (domain.Contract, error) {
	createdContract, err := u.repo.Create(contract)
	return createdContract, err
}

func (u ContractsUsecaseImpl) GetByID(id int64) (domain.ContractRead, error) {
	contractRead, err := u.repo.GetByID(id)
	return contractRead, err
}

func (u ContractsUsecaseImpl) Fetch(billboard_version_id int64) ([]domain.ContractRead, error) {
	contractReadList, err := u.repo.Fetch(billboard_version_id)
	return contractReadList, err
}

func (u ContractsUsecaseImpl) Update(id int64, contract domain.Contract) (domain.Contract, error) {
	updatedContract, err := u.repo.Update(id, contract)
	return updatedContract, err
}

func (u ContractsUsecaseImpl) Delete(id int64) (domain.Contract, error) {
	deletedContract, err := u.repo.Delete(id)
	return deletedContract, err
}
