package main

import (
	"html/template"
	"net/http"
)

func PageAbout(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.New("").ParseFiles("../templates/app_about/site_about.html", "../templates/site_base.html")
	_ = templ.ExecuteTemplate(w, "base", nil)
}
