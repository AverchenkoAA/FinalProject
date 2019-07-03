package handlers

import (
	"fmt"
	"net/http"
	"project/infrastracture"
)
type MainHandlers interface{
	MainComputerHandler(w http.ResponseWriter, r *http.Request)
	MainUserHandler(w http.ResponseWriter, r *http.Request)
}
type mainHandlers struct{
	computerRep infrastracture.ComputerRepository
	usersRep infrastracture.UserRepository
}

func NewMainHandler(cr infrastracture.ComputerRepository,ur infrastracture.UserRepository) MainHandlers{
	return &mainHandlers{
		computerRep: cr,
		usersRep: ur,
	}
}

func(mh *mainHandlers) MainComputerHandler(w http.ResponseWriter, r *http.Request) {
	compHandlers:=NewComputerHandler(mh.computerRep)

	switch r.Method{
	case http.MethodGet:{
		fmt.Println("MethodGet")
		if r.URL.Query().Get("_id")!=""{
			compHandlers.GetPCByID(w,r)
		}else if r.URL.Query().Get("search")!=""{
			compHandlers.GetPCByField(w,r)
		}
	}
	case http.MethodPost:{
		fmt.Println("MethodPost")
		compHandlers.AddNewPC(w,r)
	}
	case http.MethodPut:{
		fmt.Println("MethodPut")
		compHandlers.UpdatePCByID(w,r)
	}
	case http.MethodDelete:{
		fmt.Println("MethodDelete")
		compHandlers.DeletePCByID(w,r)
	}
	}
}

func(mh *mainHandlers) MainUserHandler(w http.ResponseWriter, r *http.Request) {
	userHandlers:=NewUserHandler(mh.usersRep)

	switch r.Method{
	case http.MethodGet:{
		fmt.Println("MethodGet")
		if r.URL.Query().Get("_id")!=""{
			userHandlers.GetUserByID(w,r)
		}else if r.URL.Query().Get("search")!=""{
			userHandlers.GetUsersByField(w,r)
		}
	}
	case http.MethodPost:{
		fmt.Println("MethodPost")
		userHandlers.AddNewUser(w,r)
	}
	case http.MethodPut:{
		fmt.Println("MethodPut")
		userHandlers.UpdateUserByID(w,r)
	}
	case http.MethodDelete:{
		fmt.Println("MethodDelete")
		userHandlers.DeleteUserByID(w,r)
	}
	}
}