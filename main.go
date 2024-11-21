package main

import (
	"log"
	"net/http"

	"hastaka/goths-starter/app/layout"
	"hastaka/goths-starter/app/pages"

	"github.com/a-h/templ"
)

func handlePage(c templ.Component) *templ.ComponentHandler {
	page := layout.Layout(c)
	return templ.Handler(page)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handlePage(pages.Error404(r.URL.Path)).ServeHTTP(w, r)
	} else {
		handlePage(pages.Index()).ServeHTTP(w, r)
	}
}

func main() {
	// Set up static routing
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up root handling (including 404)
	http.HandleFunc("/", handleRoot)

	// Set up handling of individual pages
	http.Handle("/index", handlePage(pages.Index()))

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
