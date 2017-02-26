package main

import (
  "github.com/yosssi/ace"
  "bytes"
  "strings"
  "fmt"
)

func compileAce(filename string) {
  var doc bytes.Buffer

  if tpl, err := ace.Load(trimExt(filename), "", nil); err == nil {

    err := tpl.Execute(&doc, nil)
    if err != nil { panic(err) }

    toMake := strings.Replace(filename, "sites", "public", 1)
    toMake = strings.Replace(toMake, "ace", "html", 1)

    writeStringToFile(toMake, doc.String())

    consoleSuccess(fmt.Sprintf("[ACE]: " + filename + " => " + toMake + "\n"))
  } else {
    consoleError(err)
  }
}
