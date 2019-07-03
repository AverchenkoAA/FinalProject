package infrastracture

import (
	"github.com/stretchr/testify/mock"
	//"errors"
)

type UserRepoMock struct{
	mock.Mock
}


func (mock *UserRepoMock) InsertUser(user *UserDB) error {
	arg:=mock.Called(user)
	return  arg.Error(0)
}

func (mock *UserRepoMock) DeleteUserByID(idValue string) error {
	arg:=mock.Called(idValue)
	return arg.Error(0)
}

func (mock *UserRepoMock) UpdateUserByID(idValue, updateField string, updateValue interface{}) error {
	arg:=mock.Called(idValue,updateField,updateValue)
	return arg.Error(0)
}

func (mock *UserRepoMock) FindOneUser(filterField string, filterValue interface{}) (*UserDB, error) {
	arg:=mock.Called(filterField,filterValue)
	return	arg.Get(0).(*UserDB), arg.Error(1)
}

func (mock *UserRepoMock) FindAllUser() ([]*UserDB, error) {
	arg:=mock.Called()
	return	arg.Get(0).([]*UserDB), arg.Error(1)
}

func (mock *UserRepoMock) FindUser(filterField string, filterValue interface{}) ([]*UserDB, error) {
	arg:=mock.Called(filterField,filterValue)
	return	arg.Get(0).([]*UserDB), arg.Error(1)
}

func (mock *UserRepoMock) FindByID(idValue string) (*UserDB, error) {
	arg:=mock.Called(idValue)
	return	arg.Get(0).(*UserDB), arg.Error(1)
}
