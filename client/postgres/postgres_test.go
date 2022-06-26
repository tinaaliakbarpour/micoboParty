package postgres

import (
	"fmt"
	"micobianParty/config"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	// Storage.Connect(config)
	tests := []struct {
		step string
		conf config.Config
		err  error
	}{
		{
			step: "A",
			conf: config.Config{
				Debug:   true,
				Service: config.Service{},
				POSTGRES: config.Database{
					Username: "admin",
					Password: "password",
					Host:     "127.0.0.1",
					Port:     "5432",
					DB:       "",
					SslMode:  "disable",
					TimeZone: "",
				},
			},
			err: fmt.Errorf("failed to connect to `host=127.0.0.1 user=admin database=port=5432`: dial error (dial tcp 127.0.0.1:5432: connect: connection refused)"),
		},
		{
			step: "B",
			conf: config.Config{
				Debug:   true,
				Service: config.Service{},
				POSTGRES: config.Database{
					Username: "admin",
					Password: "password",
					Host:     "postgres",
					Port:     "5432",
					DB:       "",
					SslMode:  "disable",
					TimeZone: "",
				},
			},
			err: nil,
		}}

	for _, tc := range tests {
		t.Run(tc.step, func(t *testing.T) {
			once = sync.Once{}
			err := Storage.Connect(tc.conf)
			if tc.err != nil {
				assert.Equal(t, tc.err.Error(), err.Error())
			}
		})
	}
}

func TestDB(t *testing.T) {
	db := Storage.DB()
	if db == nil {
		assert.Error(t, fmt.Errorf("error in database get DB"))
	}
}

func TestGet(t *testing.T) {
	db := Storage.DB()
	if db == nil {
		assert.Error(t, fmt.Errorf("error in database get DB"))
	}
}

func TestClose(t *testing.T) {
	err := Storage.Close()
	if err != nil {
		assert.Error(t, fmt.Errorf("error in close DB"))
	}
}
