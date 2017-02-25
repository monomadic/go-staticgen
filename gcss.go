package main

import (
  "github.com/yosssi/gcss"
  "bytes"
  "strings"
  "fmt"
)

func compileGcss(from string, sitename string) {
  var fromdoc bytes.Buffer
  var todoc bytes.Buffer

  to := strings.Replace(convertSrcToDestPath(from), "sass", "css", 1)

  fromdoc = processTemplate(from, "sites/"+ sitename +"/styles/**")

  if _, err := gcss.Compile(&todoc, &fromdoc) ; err == nil {
    writeStringToFile(to, todoc.String())
    consoleSuccess(fmt.Sprintf("[SASS]: " + from + " => " + to + "\n"))
  } else {
    consoleError(err)
  }
}


