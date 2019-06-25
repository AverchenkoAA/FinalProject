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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
func TestGetRightType_ReturnInt(t *testing.T){
	answer := getRightType("computer.hddvolume","140")
	assert.Equal(t,140,answer)
}
func TestGetRightType_ReturnFloat64(t *testing.T){
	answer := getRightType("computer.core.frequency","3.6")
	assert.Equal(t,3.6,answer)
}
func TestGetRightType_ReturnString(t *testing.T){
	answer := getRightType("computervendor","APC")
	assert.Equal(t,"APC",answer)
}

func TestHandlerGetPCByID_ShouldEndWithError(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=infrastracture.NewComputerDB()

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	
	compMock.On("FindByID", mock.Anything).Return(pc,errors.New("humpful error"))

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.GetPCByID)
	handler.ServeHTTP(rr, req)
	
	compMock.AssertExpectations(t)
}
func TestHandlerGetPCByID_ShouldEndWithoutErrors(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=infrastracture.NewComputerDB()
	pc.Computer.InventoryNumber=999

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	compMock.On("FindByID", mock.Anything).Return(pc,nil)

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.GetPCByID)
	
	handler.ServeHTTP(rr, req)
	compMock.AssertExpectations(t)
}

func TestHandlerGetPCByField_ShouldEndWithError(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=[]*infrastracture.ComputerDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	compMock.On("FindComputer", mock.Anything, mock.Anything).Return(pc,errors.New("humpful error"))

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.GetPCByField)
	
	handler.ServeHTTP(rr, req)
	compMock.AssertExpectations(t)
}
func TestHandlerGetPCByField_ShouldEndWithoutErrors(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=[]*infrastracture.ComputerDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	compMock.On("FindComputer", mock.Anything, mock.Anything).Return(pc,nil)

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.GetPCByField)
	
	handler.ServeHTTP(rr, req)
	compMock.AssertExpectations(t)
}

func TestHandlerGetAllPC_ShouldEndWithError(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=[]*infrastracture.ComputerDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	compMock.On("FindAllComputer").Return(pc,errors.New("humpful error"))

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.GetAllPC)
	
	handler.ServeHTTP(rr, req)
	compMock.AssertExpectations(t)
}
func TestHandlerGetAllPC_ShouldEndWithoutErrors(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=[]*infrastracture.ComputerDB{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	compMock.On("FindAllComputer").Return(pc,nil)

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.GetAllPC)
	
	handler.ServeHTTP(rr, req)
	compMock.AssertExpectations(t)
}

func TestHandlerAddNewPC_ShouldEndWithError(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=infrastracture.NewComputerDB()

	b:=bytes.NewBufferString("")
	err := json.NewEncoder(b).Encode(*pc)
	if err != nil { 
		t.Fatal(err)
	} 

	buffer := bytes.Buffer{}
	params := url.Values{}
	params.Set("pc",b.String())
	buffer.WriteString(params.Encode())

	req, err := http.NewRequest(http.MethodPost, "", &buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	compMock.On("InsertComputer", mock.Anything).Return(errors.New("humpful error"))

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.AddNewPC)
	
	handler.ServeHTTP(rr, req)
	compMock.AssertExpectations(t)
}
func TestHandlerAddNewPC_ShouldEndWithoutErrors(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	pc:=infrastracture.NewComputerDB()

	b:=bytes.NewBufferString("")
	err := json.NewEncoder(b).Encode(*pc)
	if err != nil { 
		t.Fatal(err)
	} 

	buffer := bytes.Buffer{}
	params := url.Values{}
	params.Set("pc",b.String())
	buffer.WriteString(params.Encode())

	req, err := http.NewRequest(http.MethodPost, "", &buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()

	compMock.On("InsertComputer", mock.Anything).Return(nil)

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.AddNewPC)
	
	handler.ServeHTTP(rr, req)
	compMock.AssertExpectations(t)
}

func TestHandlerUpdatePCByID_ShouldEndWithError(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	
	compMock.On("UpdateComputerByID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("humpful error"))

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.UpdatePCByID)
	handler.ServeHTTP(rr, req)
	
	compMock.AssertExpectations(t)
}
func TestHandlerUpdatePCByID_ShouldEndWithoutErrors(t *testing.T) {
	compMock:=&infrastracture.CompRepoMock{}

	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
        t.Fatal(err)
    }
	rr := httptest.NewRecorder()
	
	compMock.On("UpdateComputerByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	compHandler:=NewComputerHandler(compMock)
	handler := http.HandlerFunc(compHandler.UpdatePCByID)
	handler.ServeHTTP(rr, req)
	
	compMock.AssertExpectations(t)
}