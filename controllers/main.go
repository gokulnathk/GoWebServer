package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"
	"webserver/viewmodels"
)

/* Register... function to register the HTTP request */
func Register(templates *template.Template) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		template := templates.Lookup(requestedFile + ".html")

		var context interface{}

		switch requestedFile {
		case "index":
			context = viewmodels.GetIndex()
		case "brands":
			context = viewmodels.GetBrands()
		}

		if template != nil {
			template.Execute(w, context)
		} else {
			w.WriteHeader(404)
		}
	})

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
