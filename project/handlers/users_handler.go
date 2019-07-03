package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/domain"
	"project/infrastracture"
	"strings"
)

type UsersHandler interface {
	LoginUser(w http.ResponseWriter, r *http.Request)

	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUsersByField(w http.ResponseWriter, r *http.Request)

	AddNewUser(w http.ResponseWriter, r *http.Request)
	UpdateUserByID(w http.ResponseWriter, r *http.Request)
	DeleteUserByID(w http.ResponseWriter, r *http.Request)
}

type usersHandler struct {
	usersRep infrastracture.UserRepository
}

func NewUserHandler(ur infrastracture.UserRepository) UsersHandler {
	return &usersHandler{
		usersRep: ur,
	}
}

func (ch *usersHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LoginPage")
	userName := r.URL.Query().Get("login")
	userPassword := r.URL.Query().Get("password")
	user, err := ch.usersRep.FindOneUser("user.login", userName)
	if err != nil {
		fmt.Println("some login error", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	} else if user == nil || user.User.Password != userPassword {
		fmt.Println("no auth at", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		userinfo := userInfo{
			UserID:     user.ID.Hex(),
			UserLogin:  user.User.Login,
			UserRights: user.User.UserRights,
		}
		err = json.NewEncoder(w).Encode(userinfo)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(404)
			return
		}
		fmt.Printf("\nAccess permit for user %v", user.User.Login)
		fmt.Fprint(w)
	}
}

func (ch *usersHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param1 := r.URL.Query().Get("_id")
	fmt.Println(param1)
	user, err := ch.usersRep.FindByID(param1)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w)
}
func (ch *usersHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := ch.usersRep.FindAllUser()
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}
	fmt.Fprint(w)
}
func (ch *usersHandler) GetUsersByField(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	findField := r.URL.Query().Get("search")
	findValue := r.URL.Query().Get("searchValue")
	v := getRightType(findField, findValue)
	user, err := ch.usersRep.FindUser(findField, v)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w)
}
func (ch *usersHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	a := r.PostFormValue("user")
	inputUser := domain.NewUser()
	json.NewDecoder(strings.NewReader(a)).Decode(&inputUser)

	user := infrastracture.NewUserDB()
	user.User = inputUser

	err := ch.usersRep.InsertUser(user)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}
	fmt.Println(user)
	fmt.Fprintf(w, "%v", user)
}
func (ch *usersHandler) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("_id")
	updateFieldParam := r.URL.Query().Get("field")
	updateValueParam := r.URL.Query().Get("value")

	err := ch.usersRep.UpdateUserByID(idParam, updateFieldParam, updateValueParam)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}
	fmt.Println("Update secceed")
	fmt.Fprintf(w, "Update user succeed: %v", idParam)

}
func (ch *usersHandler) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("_id")

	user, err := ch.usersRep.FindByID(idParam)
	if err != nil {
		fmt.Println("There is no such user")
		w.WriteHeader(404)
		return
	}
	if user.User.Login == "ADMIN" {
		fmt.Println("You can't delete ADMIN user.")
		fmt.Fprintf(w, "You can't delete %v user.", user.User.Login)
		return
	}
	err = ch.usersRep.DeleteUserByID(idParam)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(404)
		return
	}
	fmt.Println("Delete secceed")
	fmt.Fprintf(w, "Delete user succeed: %v", idParam)
}
