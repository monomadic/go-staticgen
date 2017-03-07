package main

import (
  "bytes"
  "strings"
  "text/template"
  "path/filepath"
)

func processTemplate(from string, dir string) (bytes.Buffer, error) {
  var doc bytes.Buffer
  var siteName = filepathToSitename(from)

  funcMap := template.FuncMap {
    "title": strings.Title,
  }

  baseName := filepath.Base(from)
  globbedFiles, _ := PartialGlob("sites/"+siteName+"/styles", ".sass")

  tpl := template.New(baseName).Funcs(funcMap)

  if parsedTpl, err := tpl.ParseFiles(globbedFiles...); err != nil {
    return doc, err
  } else {
    tpl = parsedTpl
  }

  err := tpl.Execute(&doc, nil)
  return doc, err
}