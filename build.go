package main

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

func processSites() {

  filepath.Walk("sites/robsaunders", func(name string, info os.FileInfo, err error) error {
    if info == nil || name == "sites/robsaunders" {
      return err
    }

    from := filepath.ToSlash(name)
    dot := filepath.Base(name)[0]
    ext := filepath.Ext(name)

    if info.IsDir() {
      consoleInfo("[MKDIR] " + from + "\n")
      makeDirIfMissing(from)
      if from == "public/robsaunders" || dot == '.' || dot == '_' {
        return filepath.SkipDir
      }
    } else {
      // for _, exclude := range [] {
      //   if strings.HasSuffix(from, exclude) {
      //     return err
      //   }
      // }
      if dot != '.' && dot != '_' {
        switch ext {
        case ".ace":
          compileAce(from)
        case ".sass":
          compileGcss(from)
        default:
          copyFile(from)
        }
      }
    }
    return err
  })
}

func convertSrcToDestPath(filename string) string {
  return strings.Replace(filename, "sites", "public", 1)
}

func findLayoutFile(filename string) string {
  return filepath.Dir(filename) + "/_layout"
}

func trimExt(filename string) string {
  return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func makeDirIfMissing(dir string) {
  toMake := strings.Replace(dir, "sites", "public", 1)
  if _, err := os.Stat(toMake); os.IsNotExist(err) {
    os.MkdirAll(toMake, os.ModePerm)
  }
}

func writeStringToFile(filepath string, content string) {
  fo, err := os.Create(filepath)
  if err != nil {
    panic(err)
  }
  defer fo.Close()

  fmt.Fprintf(fo, content)
}