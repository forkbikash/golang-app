package main

import (
	"net/http"
	"text/template"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

var templates *template.Template
var client *redis.Client

func main() {
	// creating a redis client
	client = redis.NewClient(&redis.Options{
		Addr : "localhost:6379", // 6379 is the default port for redis
	})

	// parse the code from the folder templates
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")

	// registering the gorilla/mux router as default handler
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	// getting data from redis server
	comments, err := client.LRange("comments", 0, 10).Result()

	if(err != nil) {
		return
	}
	
	templates.ExecuteTemplate(w, "index.html", comments)
}