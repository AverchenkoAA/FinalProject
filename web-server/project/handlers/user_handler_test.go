package handlers


import (
	"testing"
	"net/http"
	"net/http/httptest"
	"project/infrastracture"
	"errors"
	"encoding/json"
	"bytes"
	"net/url"

	"github.com/stretchr/testify/mock"
)

func TestHandlerLoginUser_ShouldEndWithError(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}
	
	user:=infrastracture.NewUserDB()
	userMock.On("FindOneUser", mock.Anything, mock.Anything).Return(user,errors.New("something went wrong. error just for testing"))

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.LoginUser)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}
func TestHandlerLoginUser_ShouldEndWithoutErrors(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=infrastracture.NewUserDB()
	
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	userMock.On("FindOneUser", mock.Anything, mock.Anything).Return(user,nil)

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.LoginUser)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}

func TestHandlerGetUserByID_ShouldEndWithError(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=infrastracture.NewUserDB()

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	
	userMock.On("FindByID", mock.Anything).Return(user,errors.New("something went wrong. error just for testing"))

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.GetUserByID)
	handler.ServeHTTP(rr, req)
	
	userMock.AssertExpectations(t)
}
func TestHandlerGetUserByID_ShouldEndWithoutErrors(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=infrastracture.NewUserDB()

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	userMock.On("FindByID", mock.Anything).Return(user,nil)

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.GetUserByID)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}

func TestHandlerGetUserByField_ShouldEndWithError(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=[]*infrastracture.UserDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	userMock.On("FindUser", mock.Anything, mock.Anything).Return(user,errors.New("something went wrong. error just for testing"))

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.GetUsersByField)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}
func TestHandlerGetUserByField_ShouldEndWithoutErrors(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=[]*infrastracture.UserDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	userMock.On("FindUser", mock.Anything, mock.Anything).Return(user,nil)

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.GetUsersByField)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}

func TestHandlerGetAllUsers_ShouldEndWithError(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=[]*infrastracture.UserDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	userMock.On("FindAllUser").Return(user,errors.New("something went wrong. error just for testing"))

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.GetAllUsers)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}
func TestHandlerGetAllUsers_ShouldEndWithoutErrors(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=[]*infrastracture.UserDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	userMock.On("FindAllUser").Return(user,nil)

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.GetAllUsers)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}

func TestHandlerAddNewUser_ShouldEndWithError(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=infrastracture.NewUserDB()

	b:=bytes.NewBufferString("")
	err := json.NewEncoder(b).Encode(*user)
	if err != nil { 
		t.Fatal(err)
	} 

	buffer := bytes.Buffer{}
	params := url.Values{}
	params.Set("user",b.String())
	buffer.WriteString(params.Encode())

	req, err := http.NewRequest(http.MethodPost, "", &buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	userMock.On("InsertUser", mock.Anything).Return(errors.New("something went wrong. error just for testing"))

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.AddNewUser)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}
func TestHandlerAddNewUser_ShouldEndWithoutErrors(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	user:=infrastracture.NewUserDB()

	b:=bytes.NewBufferString("")
	err := json.NewEncoder(b).Encode(*user)
	if err != nil { 
		t.Fatal(err)
	} 

	buffer := bytes.Buffer{}
	params := url.Values{}
	params.Set("user",b.String())
	buffer.WriteString(params.Encode())

	req, err := http.NewRequest(http.MethodPost, "", &buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	userMock.On("InsertUser", mock.Anything).Return(nil)

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.AddNewUser)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}

func TestHandlerUpdateUserByID_ShouldEndWithError(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	
	userMock.On("UpdateUserByID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("something went wrong. error just for testing"))

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.UpdateUserByID)
	handler.ServeHTTP(rr, req)
	
	userMock.AssertExpectations(t)
}
func TestHandlerUpdateUserByID_ShouldEndWithoutErrors(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	
	userMock.On("UpdateUserByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.UpdateUserByID)
	handler.ServeHTTP(rr, req)
	
	userMock.AssertExpectations(t)
}

func TestHandlerDeleteUserByID_ShouldStopDeletingAdminUser(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	user:=infrastracture.NewUserDB()
	user.User.Login="ADMIN"
	userMock.On("FindByID", mock.Anything).Return(user,nil)
	
	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.DeleteUserByID)
	handler.ServeHTTP(rr, req)
	
	userMock.AssertExpectations(t)
}
func TestHandlerDeleteUserByID_ShouldEndWithError(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	user:=infrastracture.NewUserDB()

	userMock.On("FindByID", mock.Anything).Return(user,nil)
	userMock.On("DeleteUserByID", mock.Anything).Return(errors.New("something went wrong. error just for testing"))

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.DeleteUserByID)
	handler.ServeHTTP(rr, req)
	
	userMock.AssertExpectations(t)
}
func TestHandlerDeleteUserByID_ShouldEndWithoutErrors(t *testing.T) {
	userMock:=&infrastracture.UserRepoMock{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	user:=infrastracture.NewUserDB()

	userMock.On("FindByID", mock.Anything).Return(user,nil)
	userMock.On("DeleteUserByID", mock.Anything).Return(nil)

	userHandler:=NewUserHandler(userMock)
	handler := http.HandlerFunc(userHandler.DeleteUserByID)
	
	handler.ServeHTTP(rr, req)
	userMock.AssertExpectations(t)
}