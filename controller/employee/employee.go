package employee

import (
	"fmt"
	"micobianParty/domain/entity"
	employeerepo "micobianParty/domain/repository/employee"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//RegisterEmployee will register and store a new employee in db
func (employee) RegisterEmployee(ctx *gin.Context) {
	var newEmployee entity.Employee
	if err := ctx.ShouldBindJSON(&newEmployee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newEmployee.FirstName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "first_name field is empty"})
		return
	}

	if newEmployee.LastName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "last_name field is empty"})
		return
	}

	if newEmployee.Birthday.IsZero() {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bith_day field is empty"})
		return
	}

	if (entity.Gender(newEmployee.Gender) != entity.MALE) && (entity.Gender(newEmployee.Gender) != entity.FEMALE) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "gender field not available"})
		return
	}

	if err := employeerepo.Repository.Create(&newEmployee); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed creating a new user with error : " + err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})

}

//GetAllEmployees will return a list of all employees in micobo
func (employee) GetAllEmployees(ctx *gin.Context) {

	employees, err := employeerepo.Repository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed fetching all employees in micobo with error : " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "fetched all employees successfully",
		"employees": employees,
	})

}

//UpdateEmployee will update an employee with specific id
func (employee) UpdateEmployee(ctx *gin.Context) {

	UpdatedEmployee, err := employeerepo.Repository.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed updating employee  with error : %s ", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":          "updated  employee successfully",
		"updated_employee": UpdatedEmployee,
	})

}

//DeleteEmployee will remove an employee from db
func (employee) DeleteEmployee(ctx *gin.Context) {
	input := ctx.Param("id")

	id, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed converting string to uint with error : " + err.Error(),
		})
		return
	}

	if err := employeerepo.Repository.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed deleting employee with id %d with error : %s ", id, err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted  employee successfully",
		"id":      id,
	})
}

//GetEmployeesByEventID will fetch all the employees that belonged to one
// event and also provides dynamic filtering
func (employee) GetEmployeesByEventID(ctx *gin.Context) {

	input := ctx.Param("event_id")

	event_id, err := strconv.ParseUint(input, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "failed converting string to uint with error : " + err.Error(),
		})
		return
	}

	conditions := make(map[string]interface{})
	params := ctx.Request.URL.Query()["filter"]

	for _, value := range params {
		seeds := strings.Split(value, ":")
		conditions[seeds[0]] = seeds[1]
	}

	employees, err := employeerepo.Repository.GetWithCustomFilters(conditions, uint(event_id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("failed fetching employee with event_id %d with error : %s ", event_id, err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "fetching  employees with custom filters and specific event id successfully",
		"employees": employees,
	})
}
