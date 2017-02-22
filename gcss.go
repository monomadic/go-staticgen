package main

import (
  "os"
  "github.com/yosssi/gcss"
  "bytes"
  "strings"
  "fmt"
  // "text/template"
)

func compileGcss(filename string) {
  var doc bytes.Buffer

  toMake := strings.Replace(filename, "sites", "public", 1)
  toMake = strings.Replace(toMake, "sass", "css", 1)

  from, err := os.Open(filename)
  _, err = gcss.Compile(&doc, from)

  if err != nil {
    writeStringToFile(toMake, doc.String())
    consoleSuccess(fmt.Sprintf("[SASS]: " + filename + " => " + toMake + "\n"))
  } else {
    consoleError(err)
  }
}

// func processAsTemplate(filename string) bytes.Buffer {
//   var buf bytes.Buffer
//   err := tmpl[name].ExecuteTemplate(&buf, filename, data)

// }
