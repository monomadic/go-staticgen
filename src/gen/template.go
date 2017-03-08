package main

import (
  "bytes"
  "text/template"
  "path/filepath"
  "strings"
  "fmt"
)

func processTemplate(from string, dir string) (bytes.Buffer, error) {
  var doc bytes.Buffer
  var siteName = filepathToSitename(from)

  funcMap := template.FuncMap {
    "copy": func (src string) string { return helperCopyFile(src, from) },
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

func helperCopyFile(from string, src string) string {
  copy_from := "sites/_shared/" + from
  copy_to := strings.Replace(convertSrcToDestPath(copy_from), "_shared", filepathToSitename(src), 1)

  if err := makeDirIfMissing(filepath.Dir(copy_to)); err != nil { createError(copy_from, err) }

  if err := cp(copy_from, copy_to); err != nil {
    createError(copy_from, err)
  } else {
    consoleSuccess(fmt.Sprintf("\t%s\n", copy_to))
  }

  return from
}