package main

import (
  "github.com/wellington/go-libsass"
  "strings"
)

type SassProcessor struct {}

func (p SassProcessor) compile(tpl *TemplateWriter) error {
  _, err := libsass.New(tpl.writer, tpl.buffer)
  return err
}

func (p SassProcessor) dstfile(filename string) string {
  return strings.Replace(convertSrcToDestPath(filename), "sass", "css", 1)
}
