package main

import (
	"fmt"
	"console-client/helper"
	"console-client/clientMod"

	"encoding/json"
	"strings"

)


func main(){
	var login, password string

	fmt.Println("You must enter login and password.")

	fmt.Println("Enter login: ")
	fmt.Scan(&login)

	fmt.Println("Enter password: ")
	fmt.Scan(&password)

	password=helper.GetHash(password)
	
	url:="http://localhost:8081/login?"
	url=helper.AddParamToURL(url,"login",login)
	url=helper.AddParamToURL(url,"password",password)

	answer:=string(clientMod.DoRequest(clientMod.GetBodyWithHeader(url,"_","_")))
	var user struct {
		UserID,
		UserLogin,
		UserRights string
	}
	json.NewDecoder(strings.NewReader(answer)).Decode(&user)

	clientMod.GetClientMod(user.UserRights, user.UserID)

}