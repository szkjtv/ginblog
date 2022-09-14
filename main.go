package main

import (
	"example.com/m/controller"
	"example.com/m/router"
)

func main() {
	controller.Dbinit()
	router.Router()
}
