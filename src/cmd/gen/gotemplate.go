package main

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"

	"os"
)

func parseTemplate(tpl *TemplateWriter) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(tpl.buffer)

	t := template.New("newTemplate")
	t.Parse(buf.String())

	println(buf.String())

	println("sdfdsf")
	t.Execute(os.Stdout, nil)

	println("sdfdsf")

	return nil
}

func processTemplate(from string, dir string) (bytes.Buffer, error) {
	var doc bytes.Buffer
	var siteName = filepathToSitename(from)
	var err error

	funcMap := template.FuncMap{
	// "copy": func(rel string) string { return helperCopyFile(rel, findSharedFile(siteName, rel), siteName) },
	}

	baseName := filepath.Base(from)
	globbedFiles, _ := PartialGlob("sites/"+siteName+"/styles", ".sass")

	tpl := template.New(baseName).Funcs(funcMap)

	if parsedTpl, err := tpl.ParseFiles(globbedFiles...); err == nil {
		tpl = parsedTpl
		err = tpl.Execute(&doc, nil)
	}

	return doc, err
}

func helperCopyFile(rel string, sitename string) string {
	src := findFile(rel, sitename)
	dest := filepath.Clean(filepath.Join("public", sitename, rel))

	if err := cp(src, dest); err != nil {
		createError(src, err)
	} else {
		consoleSuccess(fmt.Sprintf("\t%s\n", dest))
	}

	return rel + "?checksum=" + checksum(src)
}

func findFile(from string, sitename string) string {
	var sharedFile = filepath.Clean(filepath.Join("sites", "_shared", from))

	if fileExists(sharedFile) {
		return sharedFile
	}

	return filepath.Clean(filepath.Join("sites", sitename, from))
}
