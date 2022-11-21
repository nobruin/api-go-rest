package main

import (
	"fmt"

	"github.com/nobruin/api-go-rest/models"
	"github.com/nobruin/api-go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Name1", History: "History1"},
		{Name: "Name2", History: "History2"},
	}
	fmt.Println("start rest api")
	routes.HandleRequest()
}
