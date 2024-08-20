package impl_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	userControllerImpl "github.com/gab-rod23/minitweeter/users/controller/impl"
	"github.com/gab-rod23/minitweeter/users/entities/dto"
	userUsecaseMock "github.com/gab-rod23/minitweeter/users/usecase/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_HandleCreateNewUser_When_User_Is_Created_OK(t *testing.T) {
	ass := assert.New(t)

	// Mock building
	user := "user92"
	name := "Juan Perez"
	mail := "mail@gmail.com"
	request := &dto.CreateUserRequestDTO{
		Username: &user,
		Name:     &name,
		Mail:     &mail,
	}
	usecaseMock := userUsecaseMock.NewUserUsecaseMock()
	usecaseMock.PatchCreateNewUser(request, nil)

	testController := userControllerImpl.NewUserControllerWithMocks(usecaseMock)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	marshaledRequest, _ := json.Marshal(request)
	c.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(marshaledRequest)),
	}

	// Method execution
	testController.HandlerCreateNewUser(c)

	// Result assertion
	ass.Equal(http.StatusCreated, w.Result().StatusCode)
}

// func Test_HandleCreateNewUser_When_Request_Is_Invalid(t *testing.T) {
// 	ass := assert.New(t)

// 	testController := userControllerImpl.NewUserControllerWithMocks(nil)

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)

// 	c.Request = &http.Request{}

// 	// Method execution
// 	testController.HandlerCreateNewUser(c)

// 	// Result assertion
// 	ass.Equal(http.StatusBadRequest, w.Result().StatusCode)
// 	ass.Equal("null", w.Body.String())
// }
