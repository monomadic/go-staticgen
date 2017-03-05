package main

import (
  "github.com/yosssi/ace"
  "html/template"
  "bytes"
  "strings"
  "fmt"
)

func compileAce(filename string) error {
  var doc bytes.Buffer

  funcMap := template.FuncMap{
    "image_tag": func(s string) template.HTML {
      return template.HTML("<img src='"+s+"'>")
    },
    "img": func(s string) template.HTML {
      return template.HTML(s)
    },
  }

  if tpl, err := ace.Load(trimExt(filename), "", &ace.Options{
    FuncMap: funcMap,
    DynamicReload: true,
    }); err == nil {

    if err := tpl.Execute(&doc, nil); err != nil {
      return err
    }

    toMake := strings.Replace(filename, "sites", "public", 1)
    toMake = strings.Replace(toMake, "ace", "html", 1)

    writeStringToFile(toMake, doc.String())

    consoleSuccess(fmt.Sprintf("[ACE]: " + filename + " => " + toMake + "\n"))
  } else {
    return err
  }
  return nil
}
