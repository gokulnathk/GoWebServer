package controllers

import (
	"GoWebServer/controllers/util"
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
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()
	index.templates.Execute(responseWriter, vm)
}
