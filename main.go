package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my Awesome Site!</h1>")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <ahref=\"mailto:support@pixellocked.com\">support@pixellocked.com</a>.")
}

func Faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Frequently Asked Questions</h1><p>Here is a list of questions that our users commonly ask.</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry,but we couldn't find the page you were looking for</h1>")
}
func main() {
	r := mux.NewRouter()

	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", Home)
	r.HandleFunc("/contact", Contact)
	r.HandleFunc("/faq", Faq)

	http.ListenAndServe(":3000", r)
}
