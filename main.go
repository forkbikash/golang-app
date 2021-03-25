package main

import (
	"github.com/forkbikash/golang-app/models"
	"github.com/forkbikash/golang-app/routes"
	"github.com/forkbikash/golang-app/utils"
	"net/http"
)

func main() {

	models.Init()
	utils.LoadTemplates("templates/*.html")

	r := routes.NewRouter()

	// registering the gorilla/mux router as default handler
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
