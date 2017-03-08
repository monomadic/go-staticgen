package main

import (
  "github.com/yosssi/ace"
  "html/template"
  "bytes"
  "strings"
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
    "shared_file": helperCopyFile,
  }

  if tpl, err := ace.Load(aceInputFilePath(filename), "", &ace.Options{
    FuncMap: funcMap,
    DynamicReload: true,
    BaseDir: "sites/" + filepathToSitename(filename) + "/pages",
    }); err == nil {

    if err := tpl.Execute(&doc, nil); err != nil {
      return err
    }

    writeStringToFile(aceOutputFilePath(filename), doc.String())
    
  } else {
    return err
  }
  return nil
}

func aceInputFilePath(from string) string {
  return trimExt(strings.Replace(from, "sites/" + filepathToSitename(from) + "/pages", "", 1))
}

func aceOutputFilePath(from string) string {
  toMake := strings.Replace(from, "sites", "public", 1)
  toMake = strings.Replace(toMake, "ace", "html", 1)
  return strings.Replace(toMake, "/pages/", "/", 1)
}
