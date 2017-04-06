package main

import (
  "github.com/yosssi/gcss"
  "bytes"
  "strings"
  "path/filepath"
)

func compileGcss(from string) error {
  var fromdoc bytes.Buffer
  var todoc bytes.Buffer
  var err error

  sitename := filepathToSitename(from)

  if fromdoc, err = processTemplate(from, filepath.Join("sites", sitename, "styles", "**")); err != nil {
    return err
  }

  if _, err := gcss.Compile(&todoc, &fromdoc) ; err == nil {
    writeStringToFile(gcssOutputFilePath(from), todoc.String())
  } else {
    return err
  }
  return nil
}

func gcssOutputFilePath(from string) string {
  return strings.Replace(convertSrcToDestPath(from), "sass", "css", 1)
}
