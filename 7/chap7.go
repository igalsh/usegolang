package main

import (
	"net/http"

	"github.com/igalsh/usegolang/7/controllers"
	"github.com/igalsh/usegolang/7/views"

	"log"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
var p404View *views.View

// var signupView *views.View

func home(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	// fmt.Println(homeView.Layout)
	homeView.Render(w, nil)

}

func contact(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	contactView.Render(w, nil)
}

func faq(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/html")
	// fmt.Println(faqView.Layout)
	faqView.Render(w, nil)

}

func notFound(w http.ResponseWriter, r *http.Request) {
	p404View.Render(w, nil)

	// fmt.Fprint(w, "<h1>404</h1>")
	// index := strings.Index(r.URL.Path, "/")

}

// func signup(w http.ResponseWriter, r *http.Request) {
// 	signupView.Render(w, nil)
//
// }
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	faqView = views.NewView("bootstrap", "views/faq.gohtml")
	p404View = views.NewView("bootstrap", "views/404.gohtml")
	// signupView = views.NewView("bootstrap", "views/signup.gohtml")
	usersC := controllers.NewUsers()

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/faq", faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.ListenAndServe(":3000", r)

}

//Comment
