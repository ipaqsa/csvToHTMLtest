package main

import (
	"TestTask/pkg/server"
	"TestTask/pkg/utils"
)

func init() {
	utils.MyCreateFunc("./data")
	utils.MyCreateFunc("./data/downloads")
	utils.MyCreateFunc("./data/static")
	utils.MyCreateFunc("./data/static/js")
	utils.MyCreateFunc("./data/static/css")
	utils.MyCreateFunc("./data/static/templates")
	utils.MyCreateFunc("./data/static/html")
	utils.MyCreateFunc("./data/static/html/clients")
	utils.CreateIndexHTML()
	utils.CreateMainHTML()
	utils.CreateTemplateCSS()
	utils.CreateTemplateHTMl()
	utils.CreateMainCSS()
	utils.CreateMainJS()
}

func main() {
	server.Start(":8080")
}
