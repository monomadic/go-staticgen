package main

import (
  "github.com/yosssi/ace"
  "github.com/fatih/color"
  "bytes"
  "strings"
)

func compileAce(filename string) {
  var doc bytes.Buffer

  // if tpl, err := ace.Load(findLayoutFile(filename), trimExt(filename), &ace.Options{DynamicReload: true}); err == nil {
  if tpl, err := ace.Load(trimExt(filename), "", nil); err == nil {

    tpl.Execute(&doc, nil)

    toMake := strings.Replace(filename, "sites", "public", 1)
    toMake = strings.Replace(toMake, "ace", "html", 1)

    writeStringToFile(toMake, doc.String())

    color.Green("[ACE]: " + filename + " => " + toMake + "\n")
  } else {
    color.Red("[ERROR]: %s\n", err)
  }
}
