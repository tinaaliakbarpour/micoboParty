package employee

import (
	"micobianParty/domain/entity"

	"github.com/gin-gonic/gin"
)

var Repository EmployeeRepository = &employee{}

type EmployeeRepository interface {
	Create(employee *entity.Employee) error
	GetAll() ([]entity.Employee, error)
	Update(ctx *gin.Context) (*entity.Employee, error)
	Delete(id uint) error
	GetWithCustomFilters(conditions map[string]interface{}, event_id uint) ([]entity.Employee, error)
}

type employee struct{}
