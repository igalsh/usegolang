package main

import (
	"net/http"

	"github.com/igalsh/usegolang/views"

	"log"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View
var faqView *views.View
var p404View *views.View

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
	// w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "<h1>404</h1>")
	// index := strings.Index(r.URL.Path, "/")
	//
	// fmt.Fprint(w, html.EscapeString(r.URL.Path[index+1:]))
}
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

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = http.HandlerFunc(notFound)

	http.ListenAndServe(":3000", r)

}

//Comment
