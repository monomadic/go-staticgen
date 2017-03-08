package main

import (
  "bytes"
  "text/template"
  "path/filepath"
)

func processTemplate(from string, dir string) (bytes.Buffer, error) {
  var doc bytes.Buffer
  var siteName = filepathToSitename(from)

  funcMap := template.FuncMap {
    "shared_file": helperCopyFile,
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

func helperCopyFile(from string) string {
  shared_from := "sites/_shared/" + from
  shared_to_dir := convertSrcToDestPath(filepath.Dir(shared_from))
  if err := makeDirIfMissing(shared_to_dir); err != nil { createError(shared_from, err) }
  err := copyFile(shared_from)
  if err != nil { createError(shared_from, err) }
  return from
}