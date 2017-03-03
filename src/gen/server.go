package main

import (
  "log"
  "net/http"

  "github.com/jaschaephraim/lrserver"
  "github.com/dietsche/rfsnotify"
  "time"
)

func serve() {
  processSites()
  
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
        
        if err := processSites(); err != nil {
          println("WHAT")
          lr.Reload("guppo")
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

  fs := http.FileServer(http.Dir("public"))
  http.Handle("/", fs)

  log.Println("Listening on http://localhost:9000/")
  http.ListenAndServe(":9000", nil)
}