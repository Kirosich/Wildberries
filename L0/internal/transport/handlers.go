package transport

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	data, exists := CacheHe.store[id]
	if exists {
		file := "web/ui/html/home.page.tmpl"
		tmpl, _ := template.ParseFiles(file)
		tmpl.Execute(w, data)
	} else {
		fmt.Fprintf(w, "Error, key is not finded")
	}
}
