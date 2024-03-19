package database

import (
	"log"

	"github.com/avidito/mp-billboard-core-api/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	doMigration(db)

	return db
}

func doMigration(db *gorm.DB) {
	db.AutoMigrate(
		&domain.Billboard{},
		&domain.Design{},
		&domain.Status{},
		&domain.BillboardVersion{},
		&domain.ContractType{},
		&domain.Contract{},
	)
}
