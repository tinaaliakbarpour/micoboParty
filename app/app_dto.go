package app

import (
	"micobianParty/config"
	"os"

	group "github.com/oklog/oklog/pkg/group"
	"go.uber.org/zap"
)

var (
	logger          *zap.Logger
	Base            Application = &App{}
	cancelInterrupt chan struct{}
)

// Application interface for start application
type Application interface {
	StartApplication()
	initLogger(config.Config)
	initConfigs(path string) error
	initCancelInterrupt(g *group.Group, c chan os.Signal)
	createService() (g *group.Group)
	initPostgres() error
	initServer()
}

type App struct{}
