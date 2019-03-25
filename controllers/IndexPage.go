package controllers

import (
	"net/http"
	"html/template"
	"../utils"
	"../settings"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		redirectTo("404", w, r)
		return
	}

	tmpl, err := template.ParseFiles(settings.PATH_VIEWS + "index.html")
    utils.CheckError(err)
    tmpl.Execute(w, nil)
}
