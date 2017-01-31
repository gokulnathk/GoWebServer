package main

import (
  "net/http"
  "os"
  "text/template"
  "bufio"
  "strings"
)

func main() {
  templates := populateTemplates()

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    requestedFile := r.URL.Path[1:]
    template := templates.Lookup(requestedFile + ".html")
    if template != nil {
      template.Execute(w, nil)
    } else {
      w.WriteHeader(404)
    }
  })

  http.HandleFunc("/images/", serveStaticResource)
  http.HandleFunc("/css/", serveStaticResource)
  http.HandleFunc("/js/", serveStaticResource)

  http.ListenAndServe(":8800", nil)
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


func populateTemplates() *template.Template {
  // Creating new template with name templates
  result := template.New("templates")
  // setting path for the templates location (folder name)
  templatePath := "templates"
  // reading the template folder content from the path
  templateFolder, _ := os.Open(templatePath)
  // close the template folder once read complete using the defer key
  defer templateFolder.Close()

  // read template folder's content will be a slice
  templatePathsRaw, _ := templateFolder.Readdir(-1)

  templatePaths := new([]string)

  for _, pathInfo := range templatePathsRaw {
    if !pathInfo.IsDir() {
      *templatePaths = append(*templatePaths, templatePath + "/" + pathInfo.Name())
    }
  }

  /**
    parse the files present in the template paths, ParseFiles will accept a variadic
    parameter adding ... to the end of the slice variable will convert the slice to comma separated strings
   */
  result.ParseFiles(*templatePaths...)
  return result
}
