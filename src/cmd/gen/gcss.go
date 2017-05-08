package main

import (
	"strings"

	"github.com/yosssi/gcss"
)

// GcssProcessor : template wrapper for the gcss library.
type GcssProcessor struct{}

func (p GcssProcessor) compile(tpl *TemplateWriter) error {
	_, err := gcss.Compile(tpl.writer, tpl.buffer)
	return err
}

func (p GcssProcessor) dstfile(filename string) string {
	return strings.Replace(convertSrcToDestPath(filename), "sass", "css", 1)
}
