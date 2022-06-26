package app

import (
	"fmt"
	"log"
	zapLogger "micobianParty/client/logger"
	"micobianParty/client/postgres"
	"micobianParty/config"
	"micobianParty/controller/employee"
	"micobianParty/controller/event"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	group "github.com/oklog/oklog/pkg/group"
	"go.uber.org/zap"
)

// StartApplication func
func (a *App) StartApplication() {
	fmt.Println("\n\n--------------------------------")
	// if go code crashed we get error and line
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// init configs
	dir, _ := os.Getwd()
	if err := Base.initConfigs(dir + "/config.yaml"); err != nil {
		log.Println(err)
		return
	}

	// init zap logger
	Base.initLogger(config.Confs)

	if err := Base.initPostgres(); err != nil {
		log.Println(err)
		return
	}
	defer postgres.Storage.Close()

	// listen and serve
	Base.initServer()

	// create service
	g := Base.createService()
	if err := g.Run(); err != nil {
		zapLogger.Prepare(logger).Development().Level(zap.ErrorLevel).Commit("server stopped")
	}
}

// init zap logger
func (a *App) initLogger(confs config.Config) {
	defer fmt.Printf("zap logger is available \n")
	logger = zapLogger.GetZapLogger(confs.GetDebug())
}

// init configs
func (a *App) initConfigs(path string) error {

	// Current working directory

	if err := config.Confs.Load(path); err != nil {
		return err
	}

	fmt.Printf("configs loaded from file successfully \n")
	return nil
}

// init cancle Interrupt
func (a *App) initCancelInterrupt(g *group.Group, c chan os.Signal) {
	cancelInterrupt = make(chan struct{})
	g.Add(func() error {
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(err error) {
		if err != nil {
			close(cancelInterrupt)
		}
	})
}

// init postgres database
func (a *App) initPostgres() error {
	if err := postgres.Storage.Connect(config.Confs); err != nil {
		return err
	}

	fmt.Printf("postgres database loaded successfully \n")
	return nil
}

func (a *App) createService() (g *group.Group) {
	g = &group.Group{}
	// init cancel
	c := make(chan os.Signal, 1)
	Base.initCancelInterrupt(g, c)

	fmt.Printf("--------------------------------\n\n")
	return g
}

func (a *App) initServer() {
	engine := gin.Default()
	g := engine.Group("/api")

	// if we want to add some middlewares we can act like this
	// g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{
		gEvent := v1.Group("events")
		{
			gEvent.GET("/", event.Controller.GetAllEvents)
			gEvent.GET("/:event_id", event.Controller.GetEventByID)
			gEvent.GET("/:event_id/employees", employee.Controller.GetEmployeesByEventID)
		}

		gEmployee := v1.Group("employees")
		{
			gEmployee.POST("/", employee.Controller.RegisterEmployee)
			gEmployee.GET("/", employee.Controller.GetAllEmployees)
			gEmployee.PUT("/:id", employee.Controller.UpdateEmployee)
			gEmployee.DELETE("/:id", employee.Controller.DeleteEmployee)
		}

	}

	engine.Run()
}
