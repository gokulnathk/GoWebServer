package main

import (
  "net/http"
)

func main() {
  http.ListenAndServe(":8800", http.FileServer(http.Dir("public_html")))
}
