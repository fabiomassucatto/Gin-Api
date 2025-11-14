package main

import (
	"github.com/fabiomassucatto/gin-go-api/database"
	"github.com/fabiomassucatto/gin-go-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
