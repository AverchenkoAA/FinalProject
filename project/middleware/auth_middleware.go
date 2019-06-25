package middleware

import (
	"fmt"
	"net/http"
	"project/domain"
	"project/infrastracture"
)

type AuthMiddleware interface {
	CheckAuth(next http.Handler) http.Handler
	CheckAdminRole(next http.Handler) http.Handler
}

type authMiddleware struct {
	usersRep infrastracture.UserRepository
}

func NewAuthMiddleware(ur infrastracture.UserRepository) AuthMiddleware {
	return &authMiddleware{
		usersRep: ur,
	}
}

func (middleware *authMiddleware) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("checking auth")
		nameID := r.Header.Get("id")
		user, err := middleware.usersRep.FindByID(nameID)
		if err != nil {
			fmt.Println("Access denied")
			w.WriteHeader(404)
			return
		}
		if user == nil {
			fmt.Println("Access denied")
			w.WriteHeader(404)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (middleware *authMiddleware) CheckAdminRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("checking admin role")
		nameID := r.Header.Get("id")
		user, err := middleware.usersRep.FindByID(nameID)
		if err != nil {
			fmt.Println("Access denied")
			w.WriteHeader(404)
			return
		}
		if user == nil {
			fmt.Println("Access denied")
			w.WriteHeader(404)
			return
		}
		if user.User.UserRights == domain.AdminRights {
			next.ServeHTTP(w, r)
		}

	})
}
