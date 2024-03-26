package routes

import (
	"github.com/arvind910/momgo/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/delete/allmovie", controller.DeleteAllMovie).Methods("DELETE")


	return router


}