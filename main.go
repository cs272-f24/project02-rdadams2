package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	fmt.Println("Starting server on :3000")
	
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		//searchTerm := r.URL.Query().Get("term")
		w.Header().Set("Content-Type", "text/html")
		hits := []string{"First tmpl hit", "Second tmpl hit"}
		tmplBody := "<ol> {{range .}} <li>{{.}}</li> {{end}} </ol>"		// Use templates to display hits on the client side page(Response)
		tmpl, err := template.New("demo").Parse(tmplBody)
		if err != nil {
			fmt.Printf("template.Parse returned %v\n", err)
		}
		tmpl.Execute(w, hits)
	})
	
	err := http.ListenAndServe(":3000", nil) // go ListenAndServe... "Go_Routine" (1 for main) (2 go_routines for crawling & searching)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
