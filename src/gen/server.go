package main

import (
  "log"
  "net/http"
  "os"
)

func serve() {
  if _, err := os.Stat("public"); os.IsNotExist(err) {
    consoleError(err)
    println("Directory ./public does not exist. Please run gen build first.")
    os.Exit(1)
  }

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  log.Println("Listening on http://localhost:9000/")
  http.ListenAndServe(":9000", nil)
}
