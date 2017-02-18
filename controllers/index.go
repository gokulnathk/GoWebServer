package controllers

import (
	"GoWebServer/viewmodels"
	"net/http"
	"text/template"
)

type indexController struct {
	templates *template.Template
}

func (index *indexController) get(w http.ResponseWriter, r *http.Request) {
	vm := viewmodels.GetIndex()
	w.Header().Add("content-type", "text/html")
	index.templates.Execute(w, vm)
}
