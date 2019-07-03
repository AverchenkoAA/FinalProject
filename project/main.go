package main

import (
	"fmt"
	"net/http"
	"project/handlers"
	"project/middleware"
	"project/infrastracture"
	"time"
)

func main() {
	compRep := infrastracture.NewCompRep()
	userRep := infrastracture.NewUserRep()
	compHandlers:=handlers.NewComputerHandler(compRep)
	userHandlers:=handlers.NewUserHandler(userRep)
	mainHandlers:=handlers.NewMainHandler(compRep,userRep)
	authMiddleware:=middleware.NewAuthMiddleware(userRep)

	adminMux:=http.NewServeMux()
	adminMux.HandleFunc("/admin/user",userHandlers.GetAllUsers)
	adminMux.HandleFunc("/admin/user/",mainHandlers.MainUserHandler)
	adminAuthHandler := authMiddleware.CheckAuth(adminMux)
	adminRoleHandler := authMiddleware.CheckAdminRole(adminAuthHandler)

	pcMux:=http.NewServeMux()
	pcMux.HandleFunc("/pc/",mainHandlers.MainComputerHandler)
	pcMux.HandleFunc("/pc",compHandlers.GetAllPC)
	userAuthHandler := authMiddleware.CheckAuth(pcMux)

	mux := http.NewServeMux()
	mux.Handle("/", userAuthHandler)
	mux.Handle("/admin/", adminRoleHandler)
	mux.HandleFunc("/login", userHandlers.LoginUser)

	server := http.Server{
		Addr:         ":8081",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("start on 8081")
	server.ListenAndServe()
}
