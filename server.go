package main

import (
  "log"
  "net/http"
)

func serve() {
  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  log.Println("Listening on http://localhost:9000/")
  http.ListenAndServe(":9000", nil)
}
