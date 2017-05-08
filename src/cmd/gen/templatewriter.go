package main

import (
	"io"
	"os"
)

// TemplateWriter - golang template adaptor.
type TemplateWriter struct {
	buffer io.Reader
	writer io.Writer
	err    error
}

// NewTemplateWriter - creates a TemplateWriter instance.
func NewTemplateWriter(src string, dst string) *TemplateWriter {
	t := TemplateWriter{}
	t.SetReader(src)
	t.SetWriter(dst)
	return &t
}

// SetReader - sets the reader internally (this should not be exported)
func (w *TemplateWriter) SetReader(filename string) error {
	w.buffer, w.err = os.Open(filename)
	return w.err
}

// SetWriter - sets the writer internally (this should not be exported)
func (w *TemplateWriter) SetWriter(filename string) error {
	w.writer, w.err = os.Create(filename)
	return w.err
}
