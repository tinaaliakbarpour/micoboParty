package postgres

import (
	"database/sql"
	"fmt"
	zapLogger "micobianParty/client/logger"
	"micobianParty/config"
	"sync"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (*psql) GetDSN(cnf config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cnf.POSTGRES.Host,
		cnf.POSTGRES.Username, cnf.POSTGRES.Password, cnf.POSTGRES.DB,
		cnf.POSTGRES.Port, cnf.POSTGRES.SslMode, cnf.POSTGRES.TimeZone)
}

// Connect method job is connect to postgres database and check migration
func (p *psql) Connect(cnf config.Config) error {
	var err error
	var db *gorm.DB
	var sqlDB *sql.DB

	once.Do(func() {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  p.GetDSN(cnf),
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{
			QueryFields: true,
		})

		if err != nil {
			zapLogger.Prepare(zapLogger.GetZapLogger(cnf.GetDebug())).Development().Level(zap.ErrorLevel).Commit("failed to connect to postgres with error : " + err.Error())
		}
		p.db = db
	})

	if err != nil {
		return err
	}

	// Create the connection pool

	sqlDB, err = db.DB()
	if err != nil {
		zapLogger.Prepare(zapLogger.GetZapLogger(cnf.GetDebug())).Development().Level(zap.ErrorLevel).Commit("failed to create connection Pool with error : " + err.Error())
		return err
	}

	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}

func (p *psql) DB() *gorm.DB {
	return p.db
}

func (p *psql) Close() error {
	once = sync.Once{}
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
