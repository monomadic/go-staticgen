package main

import (
  "github.com/yosssi/gcss"
  "bytes"
  "strings"
  "fmt"
)

func compileGcss(from string, sitename string) error {
  var fromdoc bytes.Buffer
  var todoc bytes.Buffer
  var err error

  to := strings.Replace(convertSrcToDestPath(from), "sass", "css", 1)

  if fromdoc, err = processTemplate(from, "sites/"+ sitename +"/styles/**"); err != nil {
    return err
  }

  if _, err := gcss.Compile(&todoc, &fromdoc) ; err == nil {
    writeStringToFile(to, todoc.String())
    consoleSuccess(fmt.Sprintf("[SASS]: " + from + " => " + to + "\n"))
  } else {
    return err
  }
  return nil
}
