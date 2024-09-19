package app

import (
	"L0/internal/repository"
	"html/template"
	"log"
	"net/http"
)

func OrderServer(cache *repository.Cache, htmlfile string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(htmlfile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id := r.FormValue("Id")
		if id == "" {
			tmpl.Execute(w, nil)
			return
		}

		order, exists := cache.Get(id)
		if !exists {
			tmpl.Execute(w, nil)
			return
		}

		tmpl.Execute(w, order)
	})

	log.Println("Server is listening on port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
