package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"github.com/wellington/go-libsass"
)

func processSASS(filename string) error {
	// open file for reading
	reader, err := os.Open(filename)
	if err != nil {
		return err
	}

	// open file for writing
	writer, err := os.Create(strings.Replace(convertSrcToDestPath(filename), "sass", "css", 1))
	if err != nil {
		return err
	}
	println("opened: " + strings.Replace(convertSrcToDestPath(filename), "sass", "css", 1))

	// create empty buffer
	buffer := new(bytes.Buffer)

	// convert sass to scss
	err = libsass.ToScss(reader, buffer)
	if err != nil {
		return err
	}

	// set up compiler
	compiler, err := libsass.New(writer, buffer)
	if err != nil {
		return err
	}

	// configure @import paths
	srcDir, _ := filepath.Split(filename)
	includePaths := []string{"sites/_shared/styles", srcDir} // TODO: use cfg variables for shared path/s
	compiler.Option(libsass.IncludePaths(includePaths))

	// run compiler
	return compiler.Run()
}

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
