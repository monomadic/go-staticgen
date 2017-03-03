package main

import (
  "bytes"
  "strings"
  "fmt"
  "text/template"
)

func processTemplate(from string, dir string) (bytes.Buffer, error) {
  var doc bytes.Buffer

  funcMap := template.FuncMap {
      "title": strings.Title,
  }

  tpl := template.
    Must(template.New("main.sass").
      Funcs(funcMap).
      ParseGlob(dir))

  err := tpl.Execute(&doc, nil)
  if err != nil { return doc, err }

  consoleSuccess(fmt.Sprintf("[TPL]: " + from + "\n"))
  return doc, nil
}