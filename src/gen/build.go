package main

import (
  "fmt"
  "os"
  "path/filepath"
  // "strings"
)

func (cfg *config) processSites() error {
  sites, err := cfg.Sites()
  if err != nil { return err }

  for _, site := range sites {
    consoleInfo("\nProcessing Site: " + cfg.ServerURL() + site)
    if err := makeDirIfMissing(filepath.Join(cfg.BuildDir, site)); err != nil { return err }
    if err := processSite(site); err != nil { return err }
  }
  return err
}

func processSite(sitename string) error {
  os.RemoveAll(filepath.Join("public", "error.html"))
  os.RemoveAll(filepath.Join("public", sitename, "*.*"))
  // if err := makeDirIfMissing("public/"+sitename); err != nil { return err }
  if err := processPages(sitename); err != nil { return err }
  if err := processStyles(sitename); err != nil { return err }
  if fileExists(filepath.Join("sites", sitename, "images")) {
    if err := processImages(sitename); err != nil { return err }
  }

  return nil
}

func processImages(sitename string) error {
  var err error
  var files []string

  // if err := makeDirIfMissing("public/"+sitename+"/images"); err != nil { return err }

  if files, err = RecursiveGlob(filepath.Join("sites", sitename, "images")); err == nil {
    for _, name := range files {
      err = copyFile(name)
    }
  }
  return err
}

func processPages(sitename string) error {
  var err error
  var files []string

  if files, err = FileTypeGlob(filepath.Join("sites", sitename), ".ace"); err == nil {
    for _, name := range files {
      err = compileAce(name)
      if err == nil { consoleSuccess(fmt.Sprintf("\t" + aceOutputFilePath(name) + "\n")) }
    }
  }
  return err
}

func processStyles(sitename string) error {
  var filetype = ".sass"
  var srcdir = filepath.Join("sites", sitename, "styles")
  var dstdir = filepath.Join("public", sitename, "styles")
  var err error

  var processor = &GcssProcessor{}

  if err := makeDirIfMissing(dstdir); err != nil { return err }

  if files, err := FileTypeGlob(srcdir, filetype); err == nil {
    for _, name := range files {
      tpl := NewTemplateWriter(name, processor.dstfile(name))

      processor.compile(tpl)

      if tpl.err == nil {
        consoleSuccess(fmt.Sprintf("\t" + processor.dstfile(name) + "\n"))
      } else { print(name) ; return tpl.err }
    }
  }

  return err
}

// func processStyles(sitename string) error {
//   var err error
//   var files []string

//   if err := makeDirIfMissing(filepath.Join("public", sitename, "styles")); err != nil { return err }

//   if files, err = FileTypeGlob(filepath.Join("sites", sitename, "styles"), ".sass"); err == nil {
//     for _, name := range files {
//       // println(name)
//       err = compileGcss(name)
//       if err == nil { consoleSuccess(fmt.Sprintf("\t" + gcssOutputFilePath(name) + "\n")) }
//     }
//   }
//   return err
// }

