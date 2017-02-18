package controllers

import (
	"GoWebServer/viewmodels"
	"net/http"
	"text/template"
)

type brandsController struct {
	templates *template.Template
}

func (brands *brandsController) get(w http.ResponseWriter, r *http.Request) {
	bm := viewmodels.GetBrands()
	w.Header().Add("content-type", "text/html")
	brands.templates.Execute(w, bm)
}

type brandController struct {
	templates *template.Template
}

func (brand brandController) get(w http.ResponseWriter, r *http.Request) {

}
