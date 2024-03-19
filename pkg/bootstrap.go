package pkg

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	_billboard_versions_repository "github.com/avidito/mp-billboard-core-api/pkg/modules/billboard_versions/repository/postgres"
	_billboard_versions_usecase "github.com/avidito/mp-billboard-core-api/pkg/modules/billboard_versions/usecase"
	_billboards_repository "github.com/avidito/mp-billboard-core-api/pkg/modules/billboards/repository/postgres"
	_billboards_usecase "github.com/avidito/mp-billboard-core-api/pkg/modules/billboards/usecase"
	_contracts_repository "github.com/avidito/mp-billboard-core-api/pkg/modules/contracts/repository/postgres"
	_contracts_usecase "github.com/avidito/mp-billboard-core-api/pkg/modules/contracts/usecase"
	_v1_http_router "github.com/avidito/mp-billboard-core-api/pkg/routers/v1/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository
type Repository struct {
	billboards_repository         domain.BillboardsRepository
	billboard_versions_repository domain.BillboardVersionsRepository
	contracts_repository          domain.ContractsRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		billboards_repository:         _billboards_repository.NewBillboardsPostgresRepositoryImpl(db),
		billboard_versions_repository: _billboard_versions_repository.NewBillboardVersionsPostgresRepositoryImpl(db),
		contracts_repository:          _contracts_repository.NewContractsPostgresRepositoryImpl(db),
	}
}

// Usecase
type Usecase struct {
	billboards_usecase         domain.BillboardsUsecase
	billboard_versions_usecase domain.BillboardVersionsUsecase
	contracts_usecase          domain.ContractsUsecase
}

func NewUsecase(r Repository) Usecase {
	return Usecase{
		billboards_usecase:         _billboards_usecase.NewBillboardsUsecaseImpl(r.billboards_repository),
		billboard_versions_usecase: _billboard_versions_usecase.NewBillboardVersionsUsecaseImpl(r.billboard_versions_repository),
		contracts_usecase:          _contracts_usecase.NewContractsUsecaseImpl(r.contracts_repository),
	}
}

// Router
func RegisterRouter(app *fiber.App, u Usecase) {
	_v1_http_router.RegisterBillboardsHttpRouter(app, u.billboards_usecase)
	_v1_http_router.RegisterBillboardVersionsHttpRouter(app, u.billboard_versions_usecase)
	_v1_http_router.RegisterContractsHttpRouter(app, u.contracts_usecase)
}

// Init
func InitServices(app *fiber.App, db *gorm.DB) {
	repository := NewRepository(db)
	usecase := NewUsecase(repository)
	RegisterRouter(app, usecase)
}
