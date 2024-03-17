package pkg

import (
	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	_billboards_repository "github.com/avidito/mp-billboard-core-api/pkg/modules/billboards/repository/postgres"
	_billboards_usecase "github.com/avidito/mp-billboard-core-api/pkg/modules/billboards/usecase"
	_v1_http_router "github.com/avidito/mp-billboard-core-api/pkg/routers/v1/http"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Repository
type Repository struct {
	billboards_repository domain.BillboardsRepository
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{
		billboards_repository: _billboards_repository.NewBillboardsPostgresRepositoryImpl(db),
	}
}

// Usecase
type Usecase struct {
	billboards_usecase domain.BillboardsUsecase
}

func NewUsecase(r Repository) Usecase {
	return Usecase{
		billboards_usecase: _billboards_usecase.NewBillboardsUsecaseImpl(r.billboards_repository),
	}
}

// Router
func RegisterRouter(app *fiber.App, u Usecase) {
	_v1_http_router.RegisterBillboardsHttpRouter(app, u.billboards_usecase)
}

// Init
func InitServices(app *fiber.App, db *gorm.DB) {
	repository := NewRepository(db)
	usecase := NewUsecase(repository)
	RegisterRouter(app, usecase)
}
