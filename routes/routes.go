package routes

import (
	"log"
	"net/http"

	c "github.com/nobruin/api-go-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", c.Home)
	http.HandleFunc("/api/personalities", c.Personalities)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
