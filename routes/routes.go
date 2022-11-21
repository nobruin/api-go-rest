package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	c "github.com/nobruin/api-go-rest/controllers"
)

const (
	GET_METHOD  = "get"
	POST_METHOD = "post"
)

func HandleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", c.Home)
	router.HandleFunc("/api/personalities", c.Personalities).Methods(GET_METHOD)
	router.HandleFunc("/api/personalities/{id}", c.FindPersonality).Methods(GET_METHOD)

	log.Fatal(http.ListenAndServe(":8000", router))
}
