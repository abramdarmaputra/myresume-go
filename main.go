package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	parsedTemplate.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func main() {
	http.HandleFunc("/", homeHandler)

	// Serve static files (CSS, images, icons)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
