package clientMod

import (
	"testing"
	"net/http"
	"github.com/stretchr/testify/assert"
)

var testURL="htttp://somewebsite/?"
var addField ="testID"
var addValue ="SomeTestValue"

func TestAddParamToURL(t *testing.T){
	exp, err := http.NewRequest(http.MethodGet, testURL, nil) 
	exp.Header.Add(addField,addValue)
	if err != nil {
		t.Fatal(err)
	}
	
	res:=GetBodyWithHeader(testURL,addField,addValue)
	assert.Equal(t,exp,res)
}

func TestPostBodyWithHeader(t *testing.T){
	exp, err := http.NewRequest(http.MethodPost, testURL, nil) 
	exp.Header.Add(addField,addValue)
	if err != nil {
		t.Fatal(err)
	}
	
	res:=PostBodyWithHeader(testURL,addField,addValue,"","")
	assert.Equal(t,exp.Header.Get(addField),res.Header.Get(addField))
}
func TestPutWithHeader(t *testing.T){
	exp, err := http.NewRequest(http.MethodPut, testURL, nil) 
	exp.Header.Add(addField,addValue)
	if err != nil {
		t.Fatal(err)
	}
	
	res:=PutWithHeader(testURL,addField,addValue)
	assert.Equal(t,exp,res)
}
func TestDeleteWithHeader(t *testing.T){
	exp, err := http.NewRequest(http.MethodDelete, testURL, nil) 
	exp.Header.Add(addField,addValue)
	if err != nil {
		t.Fatal(err)
	}
	
	res:=DeleteWithHeader(testURL,addField,addValue)
	assert.Equal(t,exp,res)
}