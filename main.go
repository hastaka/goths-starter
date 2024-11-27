package main

import (
	"io"
	"log"
	"net/http"

	"hastaka/goths-starter/app/layout"
	"hastaka/goths-starter/app/pages"

	"github.com/a-h/templ"
)

var routes = map[string]templ.Component{
	"/": pages.Index(),
	// Other routes go here
}

func handlePage(path string, c templ.Component) *templ.ComponentHandler {
	page := layout.Layout(path, c)
	return templ.Handler(page)
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	page, exists := routes[r.URL.Path]
	if exists {
		handlePage(r.URL.Path, page).ServeHTTP(w, r)
	} else {
		handlePage(r.URL.Path, pages.Error404(r.URL.Path)).ServeHTTP(w, r)
	}
}

func handleAPI(template func(string) templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ensure post
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Render the template with the parsed data
		component := template(string(body))
		component.Render(r.Context(), w)
	}
}

func main() {
	// Set up static routing
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Set up route handling (including 404)
	http.HandleFunc("/", handleRoute)

	// Set up handling of api endpoints
	// e.g. http.Handle("/api/search", handleAPI(partial.Search))

	// Set up handling of dynamic components
	// e.g. http.Handle("/component/element", handleAPI(components.Element))

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
