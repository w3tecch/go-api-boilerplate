package app

import (
	"github.com/gorilla/mux"
	"w3tec.ch/boilerplate/app/ctrl"
)

// NewRouter ...
func NewRouter() *mux.Router {

	//Create main router
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.KeepContext = true

	/**
	 * meta-data
	 */
	mainRouter.Methods("GET").Path("/api/info").HandlerFunc(ctrl.GetAPIInfo)

	/**
	 * /users
	 */
	// usersRouter.HandleFunc("/", l.Use(c.GetAllUsersHandler, m.SaySomething())).Methods("GET")
	mainRouter.Methods("GET").Path("/api/users").HandlerFunc(ctrl.GetAllUsersHandler)
	mainRouter.Methods("POST").Path("/api/users").HandlerFunc(ctrl.CreateUserHandler)
	mainRouter.Methods("GET").Path("/api/users/{id}").HandlerFunc(ctrl.GetUserByIdHandler)
	mainRouter.Methods("PUT").Path("/api/users/{id}").HandlerFunc(ctrl.UpdateUserHandler)
	mainRouter.Methods("DELETE").Path("/api/users/{id}").HandlerFunc(ctrl.DeleteUserHandler)

	return mainRouter
}
