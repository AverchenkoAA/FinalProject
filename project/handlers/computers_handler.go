package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/domain"
	"project/infrastracture"
	"strconv"
	"strings"
)

type ComputersHandler interface {
	GetPCByID(w http.ResponseWriter, r *http.Request)
	GetAllPC(w http.ResponseWriter, r *http.Request)
	GetPCByField(w http.ResponseWriter, r *http.Request)

	AddNewPC(w http.ResponseWriter, r *http.Request)
	UpdatePCByID(w http.ResponseWriter, r *http.Request)
	DeletePCByID(w http.ResponseWriter, r *http.Request)
}

type computersHandler struct {
	computerRep infrastracture.ComputerRepository
}

func NewComputerHandler(cr infrastracture.ComputerRepository) ComputersHandler {
	return &computersHandler{
		computerRep: cr,
	}
}
func getRightType(field, value string) interface{} {
	if field == "computer.inventorynumber" ||
		field == "computer.hddvolume" ||
		field == "computer.ramvolume" ||
		field == "computer.owner.roomnumber" {
		i, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return i
	}
	if field == "computer.core.frequency" {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return f
	}
	return value
}

func (ch *computersHandler) GetPCByField(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	findField := r.URL.Query().Get("search")
	findValue := r.URL.Query().Get("searchValue")
	v := getRightType(findField, findValue)
	comp, err := ch.computerRep.FindComputer(findField, v)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(comp)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w)
}
func (ch *computersHandler) GetPCByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param1 := r.URL.Query().Get("_id")
	fmt.Println(param1)
	comp, err := ch.computerRep.FindByID(param1)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(comp)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w)
}
func (ch *computersHandler) GetAllPC(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	computers, err := ch.computerRep.FindAllComputer()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(computers)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w)
}
func (ch *computersHandler) AddNewPC(w http.ResponseWriter, r *http.Request) {
	a := r.PostFormValue("pc")
	inputPC := domain.NewComputer()
	err := json.NewDecoder(strings.NewReader(a)).Decode(&inputPC)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	comp := infrastracture.NewComputerDB()
	comp.Computer = inputPC
	ch.computerRep.InsertComputer(comp)
	fmt.Println(comp)
	fmt.Fprintf(w, "Insert new computer succeed:%v", comp)
}
func (ch *computersHandler) UpdatePCByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("_id")
	updateFieldParam := r.URL.Query().Get("field")
	updateValueParam := r.URL.Query().Get("value")
	err := ch.computerRep.UpdateComputerByID(idParam, updateFieldParam, updateValueParam)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Update secceed")
	fmt.Fprintf(w, "Update computer succeed: %v", idParam)

}
func (ch *computersHandler) DeletePCByID(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("_id")
	err := ch.computerRep.DeleteComputerByID(param1)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Delete secceed")
	fmt.Fprintf(w, "Delete computer succeed: %v", param1)
}
