package main

import (
  "io"
  "os"
)

type TemplateWriter struct {
  buffer      io.Reader
  writer      io.Writer
  // processor   Processor
  // src         string
  // dst         string
  err         error
}

func NewTemplateWriter(src string, dst string) *TemplateWriter {
  t := TemplateWriter{}
  t.SetReader(src)
  t.SetWriter(dst)
  return &t
}

func (w *TemplateWriter) SetReader(filename string) error {
  w.buffer, w.err = os.Open(filename)
  return w.err
}

func (w *TemplateWriter) SetWriter(filename string) error {
  w.writer, w.err = os.Create(filename)
  return w.err
}
