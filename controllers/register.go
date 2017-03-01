package controllers

import (
	"GoWebServer/controllers/util"
	"GoWebServer/viewmodels"
	"net/http"
	"text/template"
)

type registerController struct {
	templates *template.Template
}

func (rcon *registerController) handle(w http.ResponseWriter, r *http.Request) {
	responseWriter := util.GetResponseWriter(w, r)
	defer responseWriter.Close()

	vm := viewmodels.GetRegistration()

	responseWriter.Header().Add("Content-type", "text/html")
	rcon.templates.Execute(responseWriter, vm)
}
