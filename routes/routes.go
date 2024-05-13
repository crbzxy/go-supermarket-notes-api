package routes

import (
    "github.com/gorilla/mux"
    "go-supermarket-notes-api/controllers"
)

const (
    listIDPath  = "/lists/{id}"
    noteIDPath  = "/notes/{id}"
    userIDPath  = "/users/{id}"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/lists", controllers.CreateList).Methods("POST")
    router.HandleFunc("/lists", controllers.GetLists).Methods("GET")
    router.HandleFunc(listIDPath, controllers.GetList).Methods("GET")
    router.HandleFunc(listIDPath, controllers.UpdateList).Methods("PUT")
    router.HandleFunc(listIDPath, controllers.DeleteList).Methods("DELETE")

    router.HandleFunc("/notes", controllers.CreateNote).Methods("POST")
    router.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
    router.HandleFunc(noteIDPath, controllers.GetNote).Methods("GET")
    router.HandleFunc(noteIDPath, controllers.UpdateNote).Methods("PUT")
    router.HandleFunc(noteIDPath, controllers.DeleteNote).Methods("DELETE")

    router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
    router.HandleFunc(userIDPath, controllers.GetUser).Methods("GET")
    router.HandleFunc(userIDPath, controllers.UpdateUser).Methods("PUT")
    router.HandleFunc(userIDPath, controllers.DeleteUser).Methods("DELETE")
}
