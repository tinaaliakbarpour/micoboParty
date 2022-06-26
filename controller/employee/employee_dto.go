package employee

import "github.com/gin-gonic/gin"

var Controller EmployeeController = &employee{}

type EmployeeController interface {
	RegisterEmployee(ctx *gin.Context)
	GetAllEmployees(ctx *gin.Context)
	UpdateEmployee(ctx *gin.Context)
	DeleteEmployee(ctx *gin.Context)
	GetEmployeesByEventID(ctx *gin.Context)
}

type employee struct{}
