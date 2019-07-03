package clientMod

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"bytes"
)
const CONST_URL string = "http://localhost:8081"

func GetClientMod(r, id string){
	usermod:=NewUserComm(id)
	adminmod:=NewAdminComm(id)
	switch r{
	case "admin":{
		fmt.Println("Admin mod activated")
		adminmod.StartAdminMod()
	}
	case "user":{
		fmt.Println("User mod activated")
		usermod.StartUserMod()
	}
	}	
}

func DoRequest(req *http.Request)[]byte{
	client := &http.Client{}
	resp, err := client.Do(req) 
	if err != nil { 
		log.Fatal(err)
	} 
	defer resp.Body.Close() 
	
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func GetBodyWithHeader(url, key, value string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil) 
	req.Header.Add(key,value)
	if err != nil {
		log.Fatal(err)
	}

	return req
}
func PostBodyWithHeader(urlapi, key, value, dataKey, datastr string) *http.Request {
	buffer := bytes.Buffer{}
	params := url.Values{}
	params.Set(dataKey,datastr)
	buffer.WriteString(params.Encode())

	req, err := http.NewRequest(http.MethodPost, urlapi, &buffer)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add(key,value)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")


	return req
}
func PutWithHeader(url, key, value  string) *http.Request {
	req, err := http.NewRequest(http.MethodPut, url, nil) 
	req.Header.Add(key,value)
	if err != nil {
		log.Fatal(err)
	}

	return req
}
func DeleteWithHeader(url, key, value string) *http.Request {
	req, err := http.NewRequest(http.MethodDelete, url, nil) 
	req.Header.Add(key,value)
	if err != nil {
		log.Fatal(err)
	}

	return req
}

