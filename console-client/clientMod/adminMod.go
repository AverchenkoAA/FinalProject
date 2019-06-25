package clientMod
import (
	"console-client/domain"
	"console-client/helper"
	"encoding/json"
	"strings"
	"fmt"
	"bytes"
)
type AdminsCommands interface{
	StartAdminMod()
	ListenAdminModInput(input string)
	
	ShowUsersByParam(field,value string)
	ShowUserByID(idValue string)
	ShowAllUsers()

	GetAllUsers() *[]domain.UserID
	GetUserByID(idValue string) *domain.UserID
	AddUser(comp *domain.User)
	GetUserByField(key,value string) *[]domain.UserID
	UpdateUserByID(id, key,value string) string
	DeleteUserByID(idValue string) string
}

type adminCommands struct{
	ID string
}

func NewAdminComm(id string) AdminsCommands{
	return &adminCommands{
		ID: id,
	}
}
var AdminsListOfCommands = []string{
	"\nshow-all-users --- Show all users.\n",
	"\nshow-user/id --- Show one user after enter ID of user.\n",
	"\nshow-user/p --- Show one user after enter parameters.\n   Every field parameter starts with 'user.[your field]' .\n",
	"\nadd-user --- Add one user after enter all fields of the user.\n",
	"\nupdate-user --- Update one user after enter ID and updating fields of the user.\n   Every field parameter starts with 'user.[your field]' .\n",
	"\ndelete-user --- Delete one user after enter ID of the user.\n",
}
func(ac *adminCommands) ShowUsersByParam(field,value string){
	allusers:=ac.GetUserByField(field,value)
	for _,user:=range *allusers{
		helper.PrintUser(&user)		
	}
}
func(ac *adminCommands) ShowUserByID(idValue string){
	user:=ac.GetUserByID(idValue)
	if user.ID!=""{
		helper.PrintUser(user)		
	}
		
}
func(ac *adminCommands) ShowAllUsers(){
	allusers:=ac.GetAllUsers()
	for _,user:=range *allusers{
		helper.PrintUser(&user)		
	}
}
func(ac *adminCommands) GetAllUsers() *[]domain.UserID{
	var user []domain.UserID
	url:=CONST_URL+"/admin/user"
	answer:=string(DoRequest(GetBodyWithHeader(url,"id",ac.ID)))
	json.NewDecoder(strings.NewReader(answer)).Decode(&user)

	return &user
}
func(ac *adminCommands) GetUserByID(idValue string) *domain.UserID{
	var user domain.UserID
	url:=CONST_URL+"/admin/user/"
	url=helper.AddParamToURL(url,"_id",idValue)
	answer:=string(DoRequest(GetBodyWithHeader(url,"id",ac.ID)))
	json.NewDecoder(strings.NewReader(answer)).Decode(&user)

	return &user
}
func(ac *adminCommands) AddUser(user *domain.User){
	url:=CONST_URL+"/admin/user/"
	userstr:=""
	b:=bytes.NewBufferString(userstr)
	err := json.NewEncoder(b).Encode(*user)
	if err != nil { 
		fmt.Println(err) 
	} 
	answer:=string(DoRequest(PostBodyWithHeader(url,"id",ac.ID, "user", b.String())))
	fmt.Println(answer) 
}
func(ac *adminCommands) GetUserByField(key,value string) *[]domain.UserID{
	var user []domain.UserID
	url:=CONST_URL+"/admin/user/"
	url=helper.AddParamToURL(url,"search",key)
	url=helper.AddParamToURL(url,"searchValue",value)
	answer:=string(DoRequest(GetBodyWithHeader(url,"id",ac.ID)))
	json.NewDecoder(strings.NewReader(answer)).Decode(&user)

	return &user
}
func(ac *adminCommands) UpdateUserByID(id, key,value string) string{
	url:=CONST_URL+"/admin/user/"
	url=helper.AddParamToURL(url,"_id",id)
	url=helper.AddParamToURL(url,"field",key)
	url=helper.AddParamToURL(url,"value",value)
	answer:=string(DoRequest(PutWithHeader(url,"id",ac.ID)))

	return answer
}
func(ac *adminCommands) DeleteUserByID(idValue string) string{
	url:=CONST_URL+"/admin/user/"
	url=helper.AddParamToURL(url,"_id",idValue)
	answer:=string(DoRequest(DeleteWithHeader(url,"id",ac.ID)))

	return answer
}
func(ac *adminCommands) StartAdminMod(){
	uc:=NewUserComm(ac.ID)
	for{
		input:=""
		fmt.Scan(&input)
		switch input{
		case "exit":{
			fmt.Println("Logout.Bye!")
			return
		}
		case "help":{
			for _,msg:=range UsersListOfCommands{
				fmt.Print(msg)
			}
			fmt.Println("\nSpecial admin's commands:")
			for _,msg:=range AdminsListOfCommands{
				fmt.Print(msg)
			}
		}
		default:{
			uc.ListenUserModInput(input)
			ac.ListenAdminModInput(input)	
		}
		}
	}
}
func(ac *adminCommands) ListenAdminModInput(input string){
	switch input{
	case "show-all-users":{
		fmt.Println("List of all users")
		ac.ShowAllUsers()
	}
	case "show-user/id":{
		fmt.Println("Enter ID: ")
		fmt.Scan(&input)
		ac.ShowUserByID(input)
	}
	case "show-user/p":{
		field, value:="",""
		fmt.Println("Enter searhing field: ")
		fmt.Scan(&field)
		fmt.Println("Enter searhing value: ")
		fmt.Scan(&value)
		ac.ShowUsersByParam(field,value)
	}
	case "add-user":{
		user:=domain.NewUser()
		password:=""
		fmt.Println("Enter user's login: ")
		fmt.Scan(&user.Login)
		fmt.Println("Enter user's password: ")
		fmt.Scan(&password)
		(&user).Password=helper.GetHash(password)
		fmt.Println("Enter user's rights: ")
		fmt.Scan(&user.UserRights)
		
		ac.AddUser(&user)
	}
	case "update-user":{
		fmt.Println("Enter user's ID: ")
		fmt.Scan(&input)
		ac.ShowUserByID(input)

		field, value:="",""
		fmt.Println("Enter updating field: ")
		fmt.Scan(&field)
		fmt.Println("Enter updating value: ")
		fmt.Scan(&value)
		if field=="user.password"{
			value=helper.GetHash(value)
		}
		fmt.Println(ac.UpdateUserByID(input,field,value))
	}
	case "delete-user":{
		fmt.Println("Enter ID: ")
		fmt.Scan(&input)
		ac.ShowUserByID(input)

		fmt.Println(ac.DeleteUserByID(input))
	}
	default:{
		fmt.Println("Unknown command for admin.")
	}
	}
}