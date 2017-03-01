package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

//Register func to register the HTTP request
func Register(templates *template.Template) {
	router := mux.NewRouter()
	ic := new(indexController)
	ic.templates = templates.Lookup("index.html")
	router.HandleFunc("/index", ic.get)

	bc := new(brandsController)
	bc.templates = templates.Lookup("brands.html")
	router.HandleFunc("/brands", bc.get)

	rc := new(registerController)
	rc.templates = templates.Lookup("register.html")
	router.HandleFunc("/register", rc.handle)

	http.Handle("/", router)

	http.HandleFunc("/images/", serveStaticResource)
	http.HandleFunc("/css/", serveStaticResource)
	http.HandleFunc("/js/", serveStaticResource)
}

func serveStaticResource(w http.ResponseWriter, r *http.Request) {
	path := "public_html/html" + r.URL.Path
	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)
		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".jpg") {
			contentType = "image/jpeg"
		} else {
			contentType = "text/plain"
		}

		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
