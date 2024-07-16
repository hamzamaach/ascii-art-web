package ascii_art_web

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders the specified template with data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {

	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
    
	err = t.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		fmt.Println("Error executing template: ", err)
		return
	}
}
