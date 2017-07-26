package main

import (
	"bytes"
	"path/filepath"
	"strings"

	"github.com/wellington/go-libsass"
)

// SassProcessor : template wrapper for the libsass library.
type SassProcessor struct{}

func (p SassProcessor) compile(tpl *TemplateWriter) error {
	buffer := new(bytes.Buffer)
	err := libsass.ToScss(tpl.buffer, buffer)
	if err != nil {
		return err
	}

	compiler, err := libsass.New(tpl.writer, buffer)
	if err != nil {
		return err
	}

	// configure @import paths
	srcDir, _ := filepath.Split(tpl.src)
	includePaths := []string{"sites/_shared/styles", srcDir}
	compiler.Option(libsass.IncludePaths(includePaths))
	if err != nil {
		return err
	}

	return compiler.Run()
}

func (p SassProcessor) dstfile(filename string) string {
	return strings.Replace(convertSrcToDestPath(filename), "sass", "css", 1)
}
