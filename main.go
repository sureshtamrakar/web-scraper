package main

import (
	"github.com/sureshtamrakar/web-scraper/routes"
)

func main() {
	r := routes.AddRoutes()
	r.Run()

}
