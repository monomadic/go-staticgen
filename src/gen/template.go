package main

import (
  "bytes"
  "strings"
  "fmt"
  "text/template"
)

func processTemplate(from string, dir string) bytes.Buffer {
  var doc bytes.Buffer

  funcMap := template.FuncMap {
      "title": strings.Title,
  }

  tpl := template.Must(template.New("main.sass").Funcs(funcMap).ParseGlob(dir))
  // tplVars := map[string]string {
  //     "Title": "Hello world",
  //     "Content": "Hi there",
  // }
  // tpl.ExecuteTemplate(os.Stdout, tpl.Name(), tplVars)
  err := tpl.Execute(&doc, nil)
  if err != nil { consoleError(err) }

  consoleSuccess(fmt.Sprintf("[TPL]: " + from + "\n"))
  return doc
}