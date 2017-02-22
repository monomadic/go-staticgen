package main

import (
  "os"
  "github.com/yosssi/gcss"
  "github.com/fatih/color"
  "bytes"
  "strings"
)

func compileGcss(filename string) {
  var doc bytes.Buffer

  toMake := strings.Replace(filename, "sites", "public", 1)
  toMake = strings.Replace(toMake, "sass", "css", 1)

  from, err := os.Open(filename)
  _, err = gcss.Compile(&doc, from)

  if err != nil {
  }

  writeStringToFile(toMake, doc.String())

  color.Green("[SASS]: " + filename + " => " + toMake + "\n")
}
