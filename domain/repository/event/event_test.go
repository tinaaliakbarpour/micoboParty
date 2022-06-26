package event

import (
	"database/sql"
	"fmt"
	pClient "micobianParty/client/postgres"
	"micobianParty/config"
	"micobianParty/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type sqlMock struct {
	err error
	db  *gorm.DB
}

func (*sqlMock) GetDSN(cnf config.Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cnf.POSTGRES.Host,
		cnf.POSTGRES.Username, cnf.POSTGRES.Password, cnf.POSTGRES.DB,
		cnf.POSTGRES.Port, cnf.POSTGRES.SslMode, cnf.POSTGRES.TimeZone)
}

// Connect method job is connect to postgres database and check migration
func (p *sqlMock) Connect(cnf config.Config) error {
	var err error
	var db *gorm.DB
	var sqlDB *sql.DB

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.GetDSN(cnf),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		QueryFields: true,
	})

	if err != nil {
		fmt.Errorf("failed to connect to postgres with error : " + err.Error())
	}
	p.db = db

	if err != nil {
		return err
	}

	// Create the connection pool

	sqlDB, err = db.DB()
	if err != nil {
		fmt.Errorf("failed to create connection Pool with error : " + err.Error())
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

func (p *sqlMock) DB() *gorm.DB {
	return p.db
}

func (p *sqlMock) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func TestGetAll(t *testing.T) {
	db := sqlMock{}

	testcases := []struct {
		description string
		err         error
	}{
		{
			description: "A",
			err:         nil,
		},
		{
			description: "B",
			err:         fmt.Errorf("ERROR: relation \"events\" does not exist (SQLSTATE 42P01)"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			db.err = tc.err
			pClient.Storage = &db

			if err := pClient.Storage.Connect(config.Config{
				POSTGRES: config.Database{
					Username: "admin",
					Password: "password",
					Host:     "postgres",
					Port:     "5432",
					DB:       "micobo",
					SslMode:  "disable",
					TimeZone: "Europe/Berlin",
				},
			}); err != nil {
				fmt.Println(err.Error())
				return
			}

			if tc.description == "B" {
				if err := db.db.Migrator().DropTable(&entity.Employee{}, &entity.Event{}); err != nil {
					fmt.Errorf("failed migration with error : %v", err)
				}
			}

			if _, err := Repository.GetAll(); err != nil {
				if tc.err != nil {
					assert.Equal(t, tc.err.Error(), err.Error())
					return
				}
			}

		})
	}

}

func TestGet(t *testing.T) {
	db := sqlMock{}

	testcases := []struct {
		description string
		input       uint
		err         error
	}{
		{
			description: "A",
			input:       1,
			err:         nil,
		},
		{
			description: "B",
			input:       1,
			err:         fmt.Errorf("ERROR: relation \"events\" does not exist (SQLSTATE 42P01)"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			db.err = tc.err
			pClient.Storage = &db

			if err := pClient.Storage.Connect(config.Config{
				POSTGRES: config.Database{
					Username: "admin",
					Password: "password",
					Host:     "postgres",
					Port:     "5432",
					DB:       "micobo",
					SslMode:  "disable",
					TimeZone: "Europe/Berlin",
				},
			}); err != nil {
				fmt.Println(err.Error())
				return
			}

			if tc.description == "B" {
				if err := db.db.Migrator().DropTable(&entity.Employee{}, &entity.Event{}); err != nil {
					fmt.Errorf("failed migration with error : %v", err)
				}
			}

			if _, err := Repository.Get(tc.input); err != nil {
				if tc.err != nil {
					assert.Equal(t, tc.err.Error(), err.Error())
					return
				}
			}

		})
	}

}
