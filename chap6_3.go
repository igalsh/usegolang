package main

// import (
// 	"fmt"
// 	"html"
// 	"net/http"
// 	"strings"
//
// 	"html/template"
//
// 	"github.com/gorilla/mux"
// )
//
// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	homeTemplate.Execute(w, nil)
// }
//
// func contact(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "To get in touch, please send an email "+
// 		"to <a href=\"mailto:support@usegolang.com\">"+
// 		"support@usegolang.com</a>.")
// }
//
// func faq(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>This is FAQ page</h1>")
// }
//
// func notFound(w http.ResponseWriter, r *http.Request) {
//
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprint(w, "<h1>404</h1>")
// 	index := strings.Index(r.URL.Path, "/")
//
// 	fmt.Fprint(w, html.EscapeString(r.URL.Path[index+1:]))
// }
//
// var homeTemplate *template.Template
// var contactTemplate *template.Template
//
// func main() {
// 	var err error
// 	homeTemplate, err = template.ParseFiles(
// 		"views/home.gohtml",
// 		"views/layouts/footer.gohtml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	contactTemplate, err = template.ParseFiles(
// 		"views/contact.gohtml",
// 		"views/layouts/footer.gohtml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", home)
// 	r.HandleFunc("/contact", contact)
// 	r.HandleFunc("/faq", faq)
// 	r.NotFoundHandler = http.HandlerFunc(notFound)
// 	http.ListenAndServe(":3000", r)
//
// }

//Comment
