package main

import (
	"github.com/felipemarinhodev/api-go-gin/database"
	"github.com/felipemarinhodev/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
