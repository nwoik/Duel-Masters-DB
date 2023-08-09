package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Read in the template with main webpage
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("error parsing template.")

	}

	// Render the template
	t.Execute(w, nil)

	// Done.
	fmt.Fprintf(w, "Finished HTTP request at %s", r.URL.Path)

}
