package main

import (
  "github.com/yosssi/ace"
  "bytes"
  "strings"
  "fmt"
)

func compileAce(filename string) error {
  var doc bytes.Buffer

  if tpl, err := ace.Load(trimExt(filename), "", &ace.Options{DynamicReload: true}); err == nil {

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
