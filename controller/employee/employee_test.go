package employee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"micobianParty/domain/entity"
	employeeRepo "micobianParty/domain/repository/employee"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type employeeRepositoryMock struct {
	err               error
	ListOfEmployees   []entity.Employee
	UpdatedEmployee   entity.Employee
	FilteredEmployees []entity.Employee
}

func (e employeeRepositoryMock) Create(employee *entity.Employee) error {
	return e.err
}

func (e employeeRepositoryMock) GetAll() ([]entity.Employee, error) {
	return e.ListOfEmployees, e.err
}

func (e employeeRepositoryMock) Update(ctx *gin.Context) (*entity.Employee, error) {
	return &e.UpdatedEmployee, e.err
}

func (e employeeRepositoryMock) Delete(id uint) error {
	return e.err
}

func (e employeeRepositoryMock) GetWithCustomFilters(conditions map[string]interface{}, event_id uint) ([]entity.Employee, error) {
	return e.FilteredEmployees, e.err
}

func MockJsonPost(c *gin.Context /* the test context */, content interface{}) {
	c.Request.Method = "POST" // or PUT
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestRegisterEmployee(t *testing.T) {

	testcases := []struct {
		description string
		status      int
		err         error
	}{
		{
			description: "A",
			status:      400,
			err:         nil,
		},
		{
			description: "B",
			status:      201,
			err:         nil,
		},
		{
			description: "C",
			err:         fmt.Errorf("error"),
			status:      500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			repo := &employeeRepositoryMock{}
			repo.err = tc.err

			employeeRepo.Repository = repo
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = &http.Request{
				Header: make(http.Header),
			}

			MockJsonPost(ctx, map[string]interface{}{"first_name": "a", "last_name": "b", "birth_day": "2021-02-02T00:00:00Z", "gender": "male"})

			if tc.description == "A" {
				MockJsonPost(ctx, map[string]interface{}{"first_name": "a"})

			}
			Controller.RegisterEmployee(ctx)
			// MyHandler(ctx)
			assert.EqualValues(t, tc.status, w.Code)

		})
	}

}

func TestGetAll(t *testing.T) {
	testcases := []struct {
		description string
		status      int
		err         error
	}{

		{
			description: "A",
			status:      201,
			err:         nil,
		},
		{
			description: "B",
			err:         fmt.Errorf("error"),
			status:      500,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			repo := &employeeRepositoryMock{}
			repo.err = tc.err

			employeeRepo.Repository = repo
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			ctx.Request = &http.Request{
				Header: make(http.Header),
			}

			Controller.GetAllEmployees(ctx)
			// MyHandler(ctx)
			assert.EqualValues(t, tc.status, w.Code)

		})
	}

}

/*
	test cases are not completed and for the rest of the functions unit tests
	should be written :((

*/
