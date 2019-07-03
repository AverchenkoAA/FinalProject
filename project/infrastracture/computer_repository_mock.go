package infrastracture

import (
	"github.com/stretchr/testify/mock"
	//"errors"
)

type CompRepoMock struct{
	mock.Mock
}


func (mock *CompRepoMock) InsertComputer(comp *ComputerDB) error {
	arg:=mock.Called(comp)
	return  arg.Error(0)
}

func (mock *CompRepoMock) DeleteComputerByID(idValue string) error {
	arg:=mock.Called(idValue)
	return arg.Error(0)
}

func (mock *CompRepoMock) UpdateComputerByID(idValue, updateField string, updateValue interface{}) error {
	arg:=mock.Called(idValue,updateField,updateValue)
	return arg.Error(0)
}

func (mock *CompRepoMock) FindOneComputer(filterField string, filterValue interface{}) (*ComputerDB, error) {
	arg:=mock.Called(filterField,filterValue)
	return	arg.Get(0).(*ComputerDB), arg.Error(1)
}

func (mock *CompRepoMock) FindAllComputer() ([]*ComputerDB, error) {
	arg:=mock.Called()
	return	arg.Get(0).([]*ComputerDB), arg.Error(1)
}

func (mock *CompRepoMock) FindComputer(filterField string, filterValue interface{}) ([]*ComputerDB, error) {
	arg:=mock.Called(filterField,filterValue)
	return	arg.Get(0).([]*ComputerDB), arg.Error(1)
}

func (mock *CompRepoMock) FindByID(idValue string) (*ComputerDB, error) {
	arg:=mock.Called(idValue)
	return	arg.Get(0).(*ComputerDB), arg.Error(1)
}
