package main

import (
	"fmt"

	"github.com/nobruin/api-go-rest/database"
	"github.com/nobruin/api-go-rest/routes"
)

func main() {
	database.Connect()
	fmt.Println("start rest api")
	routes.HandleRequest()
}
