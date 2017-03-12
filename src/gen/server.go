package main

import (
  "log"
  "net/http"
  "os"
  "github.com/jaschaephraim/lrserver"
  "github.com/dietsche/rfsnotify"
  "time"
)

func serve() {
  if err := processSites(); err != nil { createError("", err); os.Exit(1) }
  
  // Create file watcher
  watcher, err := rfsnotify.NewWatcher()
  if err != nil { log.Fatalln(err) }
  defer watcher.Close()

  // Start LiveReload server
  lr := lrserver.New(lrserver.DefaultName, lrserver.DefaultPort)
  if err != nil { log.Fatalln(err) }
  go lr.ListenAndServe()

  // Start goroutine that requests reload upon watcher event
  // done := make(chan bool)
  go func() {
    for {
      select {
      case event := <-watcher.Events:
        consoleInfo("[FSNotify] changes detected: " + event.Name + " " + time.Now().Format(time.RFC3339))
        consoleInfo("\nProcessing Site: "+ filepathToSitename(event.Name))
        if err := processSite(filepathToSitename(event.Name)); err != nil {
          createError(event.Name, err)
          lr.Reload(event.Name)
        } else {
          lr.Reload(event.Name)
        }
      case err := <-watcher.Errors:
        log.Println(err)
      }
    }
  }()
  // Watch dir
  err = watcher.AddRecursive("sites")
  if err != nil { log.Fatalln(err) }
  // <-done

  log.Println("Listening on http://localhost:9000/")
  serveStatic()
}

func createError(filename string, err error) {
  var templateError = ""

  if filename == "" {
    templateError = err.Error()
  } else {
    templateError = "<strong>Error processing: " + filename + "</strong>" + "<p>" + err.Error() + "</p>"
  }
  if writeErr := writeStringToFile("public/error.html", templateError); writeErr != nil {
    consoleError(writeErr)
  }
  consoleError(err)
}

func serveStatic() {
  http.HandleFunc("/fonts", fontPreview)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    consoleInfo("[HTTP] requested: " + r.URL.Path[1:])
    if (fileExists("public/error.html")) {
      http.ServeFile(w, r, "public/error.html")
    } else {
      http.ServeFile(w, r, "public/" + r.URL.Path[1:])
    }
  })

  http.ListenAndServe(":9000", nil)
}

func fontPreview(w http.ResponseWriter, r *http.Request) {
}
