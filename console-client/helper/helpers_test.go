package helper

import (
	"testing"
	"console-client/domain"
	"github.com/stretchr/testify/assert"
)
func TestAddParamToURL(t *testing.T){
	testURL:="htttp://somewebsite/?"
	addField:="testID"
	addValue:="SomeTestValue"
	res:=AddParamToURL(testURL,addField,addValue)
	assert.Equal(t,"htttp://somewebsite/?testID=SomeTestValue",res)
}
func TestGetHash(t *testing.T){
	res:=GetHash("1")
	assert.NotEqual(t,"1",res)
}
func TestPrintPC(t *testing.T){
	pc:=domain.NewPC()
	PrintPC(pc)
}
func TestPrintUser(t *testing.T){
	user:=domain.NewUserID()
	PrintUser(user)
}
