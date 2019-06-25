package clientMod
import (
	"console-client/domain"
	"console-client/helper"
	"encoding/json"
	"strings"
	"fmt"
	"bytes"
)
type UsersCommands interface{
	StartUserMod()
	ListenUserModInput(input string)

	ShowAllComputers()
	ShowComputerByID(idValue string)
	ShowComputersByParam(field,value string)

	GetAllComputers() *[]domain.PC
	GetComputerByField(key,value string) *[]domain.PC
	GetComputerByID(idValue string) *domain.PC
	AddComputer(comp *domain.Computer)
	UpdateComputerByID(id, key,value string) string
	DeleteComputerByID(idValue string) string
}

type userCommands struct{
	ID string
}

func NewUserComm(id string) UsersCommands{
	return &userCommands{
		ID: id,
	}
}
var UsersListOfCommands = []string{
	"\nshow-all-pc --- Show all computers.\n",
	"\nshow-pc/id --- Show one computer after enter ID of computer.\n",
	"\nshow-pc/p --- Show one computer after enter parameters.\n   Every field parameter starts with 'computer.[your field]' .\n   Information about owners and cores returns by commands\n 'computer.owner.[your field]' or 'computer.core.[your field]'\n",
	"\nadd-pc --- Add one computer after enter all fields of the computer.\n",
	"\nupdate-pc --- Update one computer after enter ID and updating fields of the computer.\n   Every field parameter starts with 'computer.[your field]' .\n   Information about owners and cores returns by commands\n 'computer.owner.[your field]' or 'computer.core.[your field]'\n",
	"\ndelete-pc --- Delete one computer after enter ID of the computer.\n",
	"\nexit --- Logout and exit.\n",
}
func(uc *userCommands) ShowComputersByParam(field,value string){
	allpc:=uc.GetComputerByField(field,value)
	for _,pc:=range *allpc{
		helper.PrintPC(&pc)		
	}
}
func(uc *userCommands) ShowComputerByID(idValue string){
	pc:=uc.GetComputerByID(idValue)
	if pc.ID!=""{
		helper.PrintPC(pc)		
	}
		
}
func(uc *userCommands) ShowAllComputers(){
	allpc:=uc.GetAllComputers()
	for _,pc:=range *allpc{
		helper.PrintPC(&pc)		
	}
}
func(uc *userCommands) GetAllComputers() *[]domain.PC{
	var pc []domain.PC
	url:=CONST_URL+"/pc"
	answer:=string(DoRequest(GetBodyWithHeader(url,"id",uc.ID)))
	json.NewDecoder(strings.NewReader(answer)).Decode(&pc)

	return &pc
}
func(uc *userCommands) GetComputerByID(idValue string) *domain.PC{
	var pc domain.PC
	url:=CONST_URL+"/pc/"
	url=helper.AddParamToURL(url,"_id",idValue)
	answer:=string(DoRequest(GetBodyWithHeader(url,"id",uc.ID)))
	json.NewDecoder(strings.NewReader(answer)).Decode(&pc)

	return &pc
}
func(uc *userCommands) AddComputer(comp *domain.Computer){
	url:=CONST_URL+"/pc/"
	compstr:=""
	b:=bytes.NewBufferString(compstr)
	err := json.NewEncoder(b).Encode(*comp)
	if err != nil { 
		fmt.Println(err) 
	} 
	answer:=string(DoRequest(PostBodyWithHeader(url,"id",uc.ID, "pc", b.String())))
	fmt.Println(answer) 

	
}
func(uc *userCommands) GetComputerByField(key,value string) *[]domain.PC{
	var pc []domain.PC
	url:=CONST_URL+"/pc/"
	url=helper.AddParamToURL(url,"search",key)
	url=helper.AddParamToURL(url,"searchValue",value)
	answer:=string(DoRequest(GetBodyWithHeader(url,"id",uc.ID)))
	json.NewDecoder(strings.NewReader(answer)).Decode(&pc)

	return &pc
}
func(uc *userCommands) UpdateComputerByID(id, key,value string) string{
	url:=CONST_URL+"/pc/"
	url=helper.AddParamToURL(url,"_id",id)
	url=helper.AddParamToURL(url,"field",key)
	url=helper.AddParamToURL(url,"value",value)
	answer:=string(DoRequest(PutWithHeader(url,"id",uc.ID)))

	return answer
}
func(uc *userCommands) DeleteComputerByID(idValue string) string{
	url:=CONST_URL+"/pc/"
	url=helper.AddParamToURL(url,"_id",idValue)
	answer:=string(DoRequest(DeleteWithHeader(url,"id",uc.ID)))

	return answer
}
func(uc *userCommands) StartUserMod(){
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
		}
		default:{
			uc.ListenUserModInput(input)
		}
		}	
	}
}
func(uc *userCommands) ListenUserModInput(input string) {
		switch input{
		case "show-all-pc":{
			fmt.Println("List of all computers")
			uc.ShowAllComputers()
		}
		case "show-pc/id":{
			fmt.Println("Enter ID: ")
			fmt.Scan(&input)
			uc.ShowComputerByID(input)
		}
		case "show-pc/p":{
			field, value:="",""
			fmt.Println("Enter searhing field: ")
			fmt.Scan(&field)
			fmt.Println("Enter searhing value: ")
			fmt.Scan(&value)
			uc.ShowComputersByParam(field,value)
		}
		case "add-pc":{
			comp:=domain.NewComputer()

			fmt.Println("Enter invetory number: ")
			fmt.Scan(&comp.InventoryNumber)
			fmt.Println("Enter HDD volume value: ")
			fmt.Scan(&comp.HDDVolume)
			fmt.Println("Enter RAM volume value: ")
			fmt.Scan(&comp.RAMVolume)
			fmt.Println("Enter vendor value: ")
			fmt.Scan(&comp.Vendor)

			fmt.Println("Enter core frequency: ")
			fmt.Scan(&comp.Core.Frequency)
			fmt.Println("Enter core model: ")
			fmt.Scan(&comp.Core.Model)
			fmt.Println("Enter core vendor value: ")
			fmt.Scan(&comp.Core.CoreVendor)

			fmt.Println("Enter owner's first name: ")
			fmt.Scan(&comp.Owner.FirstName)
			fmt.Println("Enter owner's first name: ")
			fmt.Scan(&comp.Owner.LastName)
			fmt.Println("Enter owner's room number: ")
			fmt.Scan(&comp.Owner.RoomNumber)

			uc.AddComputer(&comp)
		}
		case "update-pc":{
			fmt.Println("Enter ID: ")
			fmt.Scan(&input)
			uc.ShowComputerByID(input)

			field, value:="",""
			fmt.Println("Enter updating field: ")
			fmt.Scan(&field)
			fmt.Println("Enter updating value: ")
			fmt.Scan(&value)
			fmt.Println(uc.UpdateComputerByID(input,field,value))
		}
		case "delete-pc":{
			fmt.Println("Enter ID: ")
			fmt.Scan(&input)
			uc.ShowComputerByID(input)

			fmt.Println(uc.DeleteComputerByID(input))
		}
		default:{
			fmt.Println("Unknown command for user.")
		}
		}
}