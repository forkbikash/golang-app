package main

import (
	"html/template"
	"net/http"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var templates *template.Template

// the byte array passed here is being used as a key to sign our cookies
// "secret-key" is a random string here
var store = sessions.NewCookieStore([]byte("secret-key"))

var client *redis.Client

func main() {
	// creating a redis client
	client = redis.NewClient(&redis.Options{
		Addr : "localhost:6379", // 6379 is the default port for redis
	})

	// parse the code from the folder templates
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", indexGetHandler).Methods("GET")
	r.HandleFunc("/", indexPostHandler).Methods("POST")
	r.HandleFunc("/login", loginGetHandler).Methods("GET")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")
	r.HandleFunc("/test", testSessionHandler).Methods("GET")

	// instantiating file server object
	fs := http.FileServer(http.Dir("./static/"))

	// path with prefix /static/ should be served by file server
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// registering the gorilla/mux router as default handler
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func indexGetHandler(w http.ResponseWriter, r *http.Request) {

	// getting data from redis server
	comments, err := client.LRange("comments", 0, 10).Result()

	if(err != nil) {
		return
	}
	
	templates.ExecuteTemplate(w, "index.html", comments)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.PostForm.Get("comment")
	client.LPush("comments", comment)
	http.Redirect(w, r, "/", http.StatusFound)
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	session, _ := store.Get(r, "session")
	session.Values["username"] = username
	session.Save(r, w)
}

func testSessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	untyped, ok := session.Values["username"]
	if !ok {
		return
	}
	username, ok := untyped.(string)
	if !ok {
		return
	}
	w.Write([]byte(username))
}