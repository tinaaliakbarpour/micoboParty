package migrations

import (
	"micobianParty/client/postgres"
	"micobianParty/domain/entity"
)

func AutoMigrateDB() error {
	// Auto migrate database
	db := postgres.Storage.DB()
	// Add required models here
	err := db.AutoMigrate(&entity.Employee{}, &entity.Event{})
	return err
}
