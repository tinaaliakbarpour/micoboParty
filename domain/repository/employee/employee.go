package employee

import (
	"micobianParty/client/logger"
	"micobianParty/client/postgres"
	"micobianParty/config"
	"micobianParty/domain/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//Create will add a new employee to db
func (employee) Create(employee *entity.Employee) error {
	zaplogger := logger.GetZapLogger(config.Confs.Debug)
	db := postgres.Storage.DB()

	result := db.Create(employee)
	if result.Error != nil && result.RowsAffected != 1 {
		logger.Prepare(zaplogger).Development().Level(zap.ErrorLevel).
			Commit("failed to create an employee with error : " + result.Error.Error())
		return result.Error
	}

	return nil

}

//GetAll will return all employees
func (employee) GetAll() ([]entity.Employee, error) {
	var employees []entity.Employee
	zaplogger := logger.GetZapLogger(config.Confs.Debug)

	db := postgres.Storage.DB()
	if err := db.Find(&employees).Error; err != nil {
		logger.Prepare(zaplogger).Development().Level(zap.ErrorLevel).
			Commit("failed to fetched all employees with error : " + err.Error())
		return nil, err
	}
	return employees, nil
}

//Update will update a specific employee
func (employee) Update(ctx *gin.Context) (*entity.Employee, error) {
	var updatingEmployee entity.Employee
	zaplogger := logger.GetZapLogger(config.Confs.Debug)

	input := ctx.Param("id")

	id, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		logger.Prepare(zaplogger).Development().Level(zap.ErrorLevel).
			Commit("failed to convert string to uint with error : " + err.Error())
		return nil, err
	}

	db := postgres.Storage.DB()
	if err := db.Where("id = ?", id).First(&updatingEmployee).Error; err != nil {
		logger.Prepare(zaplogger).Development().Level(zap.ErrorLevel).
			Commit("failed to find employee with error : " + err.Error())
		return nil, err
	}
	ctx.BindJSON(&updatingEmployee)

	if err := db.Omit("id").Updates(&updatingEmployee).Error; err != nil {
		logger.Prepare(zaplogger).Development().Level(zap.ErrorLevel).
			Commit("failed to save employee with error : " + err.Error())
		return nil, err
	}
	return &updatingEmployee, nil
}

//Delete will delete a specific employee
func (employee) Delete(id uint) error {
	var deletingEmployee entity.Employee
	zaplogger := logger.GetZapLogger(config.Confs.Debug)

	db := postgres.Storage.DB()

	if err := db.Where("id = ?", id).Delete(&deletingEmployee).Error; err != nil {
		logger.Prepare(zaplogger).Development().Level(zap.ErrorLevel).
			Commit("failed to delete employee with error : " + err.Error())
		return err
	}
	return nil
}

//GetWithCustomFilters will return a list of employees with dynamic filters and
//specific event_id
func (employee) GetWithCustomFilters(conditions map[string]interface{}, event_id uint) ([]entity.Employee, error) {
	var employees []entity.Employee
	zaplogger := logger.GetZapLogger(config.Confs.Debug)

	db := postgres.Storage.DB()

	if err := db.Where("event_id = ?", event_id).Where(conditions).
		Find(&employees).Error; err != nil {
		logger.Prepare(zaplogger).Development().Level(zap.ErrorLevel).
			Commit("failed to fetched employees with custom filters with error : " + err.Error())
		return nil, err
	}

	return employees, nil
}
