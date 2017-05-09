package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func (cfg *config) processSites() error {
	sites, err := cfg.Sites()
	if err != nil {
		return err
	}

	for _, site := range sites {
		consoleInfo("\nProcessing Site: " + cfg.ServerURL() + site)
		if err := makeDirIfMissing(filepath.Join(cfg.DestDir, site)); err != nil {
			return err
		}
		if err := cfg.processSite(site); err != nil {
			return err
		}
	}
	return err
}

func (cfg *config) processSite(sitename string) error {
	os.RemoveAll(cfg.ErrorFile())
	os.RemoveAll(filepath.Join(cfg.DestDir, sitename, "*.*"))

	if err := cfg.processPages(sitename); err != nil {
		return err
	}

	if err := cfg.processStyles(sitename); err != nil {
		return err
	}

	if fileExists(filepath.Join(cfg.SrcDir, sitename, cfg.ImageDir)) {
		if err := cfg.processImages(sitename); err != nil {
			return err
		}
	}

	return nil
}

func (cfg *config) processImages(sitename string) error {
	var err error
	var files []string

	if files, err = RecursiveGlob(filepath.Join(cfg.SrcDir, sitename, cfg.ImageDir)); err == nil {
		for _, name := range files {
			err = copyFile(name)
		}
	}
	return err
}

// func (cfg *config) pprocessPages(sitename string) error {
// 	var err error
// 	var files []string

// 	if files, err = FileTypeGlob(filepath.Join(cfg.SrcDir, sitename), ".ace"); err == nil {
// 		for _, name := range files {
// 			err = compileAce(name)
// 			if err == nil {
// 				consoleSuccess(fmt.Sprintf("\t" + aceOutputFilePath(name) + "\n"))
// 			}
// 		}
// 	}
// 	return err
// }

func (cfg *config) processPages(sitename string) error {

	if err := makeDirIfMissing(filepath.Join(cfg.DestDir, sitename, cfg.PageDir)); err != nil {
		return err
	}

	return cfg.processDir(filepath.Join(cfg.SrcDir, sitename, cfg.PageDir), ".ace", &AceProcessor{})
}

func (cfg *config) processStyles(sitename string) error {
	if err := makeDirIfMissing(filepath.Join(cfg.DestDir, sitename, cfg.StyleDir)); err != nil {
		return err
	}

	// var processor = &GcssProcessor{}
	// var processor = &SassProcessor{}
	return cfg.processDir(filepath.Join(cfg.SrcDir, sitename, cfg.StyleDir), ".sass", &GcssProcessor{})
}

func (cfg *config) processDir(srcdir string, filetype string, processor Processor) error {
	var err error

	println("processing " + srcdir)

	if files, err := FileTypeGlob(srcdir, filetype); err == nil {
		for _, name := range files {
			tpl := NewTemplateWriter(name, processor.dstfile(name))
			if tpl.err != nil {
				return tpl.err
			}

			processor.compile(tpl)

			if tpl.err == nil {
				consoleSuccess(fmt.Sprintf("\t" + processor.dstfile(name) + "\n"))
			} else {
				print(name)
				return tpl.err
			}
		}
	}

	return err
}
