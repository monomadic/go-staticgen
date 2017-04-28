// go templates are serious aids and people who seriously think they're good (eg most go programmers) are potatoes.

package main

import (
  "bytes"
  "text/template"
  "path/filepath"
  "strings"
  "fmt"

  "os"
)

func parseTemplate(tpl *TemplateWriter) error {
  buf := new(bytes.Buffer)
  buf.ReadFrom(tpl.buffer)

  t := template.New("newTemplate")
  t.Parse(buf.String())

  println(buf.String())

  println("sdfdsf")
  t.Execute(os.Stdout, nil)

  println("sdfdsf")

  return nil
}

func processTemplate(from string, dir string) (bytes.Buffer, error) {
  var doc bytes.Buffer
  var siteName = filepathToSitename(from)

  funcMap := template.FuncMap {
    "copy": func (rel string) string { return helperCopyFile(rel, findSharedFile(siteName, rel), siteName) },
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

func helperCopyFile(rel string, src string, sitename string) string {
  if src == "" { createError(rel, nil) }

  dest := strings.Replace(convertSrcToDestPath(src), "_shared", sitename, 1)

  if err := makeDirIfMissing(filepath.Dir(dest)); err != nil { createError(src, err) }

  if err := cp(src, dest); err != nil {
    createError(src, err)
  } else {
    consoleSuccess(fmt.Sprintf("\t%s\n", dest))
  }

  return rel + "?checksum=" + checksum(src)
}

func findSharedFile(site string, from string) string {
  if fileExists("sites/" + site + "/" + from) {
    return "sites/" + site + "/" + from
  } else {
    if fileExists("sites/_shared/" + from) {
      return "sites/_shared/" + from
    } else { return "" }
  }
}