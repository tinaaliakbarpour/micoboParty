package postgres

import (
	"micobianParty/config"
	"sync"

	"gorm.io/gorm"
)

var (
	Storage Store = &psql{}
	once    sync.Once
)

// store interface is interface for store things into postgres
type Store interface {
	GetDSN(cnf config.Config) string
	Connect(cnf config.Config) error
	DB() *gorm.DB
	Close() error
}

// postgres struct
type psql struct {
	db *gorm.DB
}
