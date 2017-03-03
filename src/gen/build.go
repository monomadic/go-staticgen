package main

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
  "io/ioutil"
)

func processSites() error {
  files, err := ioutil.ReadDir("sites")
  if err != nil { return err }

  os.RemoveAll("public")
  makeDirIfMissing("public")

  for _, file := range files {
    if file.IsDir() {
      dot := filepath.Base(file.Name())[0]
      if dot != '.' && dot != '_' {
        consoleInfo("\nProcessing Site: "+ file.Name())
        if err := processSite(file.Name()); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

func processSite(sitename string) error {
  err := filepath.Walk("sites/"+sitename, func(name string, info os.FileInfo, err error) error {
    if info == nil { return err }

    from := filepath.ToSlash(name)
    dot := filepath.Base(name)[0]
    ext := filepath.Ext(name)

    if info.IsDir() {
      if dot == '.' || dot == '_' {
        return filepath.SkipDir
      } else {
        makeDirIfMissing(convertSrcToDestPath(from))
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
          if err := compileAce(from); err != nil {
            return err
          }
        case ".sass":
          if err := compileGcss(from, sitename); err != nil {
            return err
          }
        default:
          if err := copyFile(from); err != nil {
            return err
          }
        }
      }
    }
    return nil
  })

  if err != nil {
    return err
  } else {
    return nil
  }
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
  consoleSuccess("[MKDIR] " + dir)
  if _, err := os.Stat(dir); os.IsNotExist(err) {
    os.MkdirAll(dir, os.ModePerm)
  }
}

func writeStringToFile(filepath string, content string) error {
  if fo, err := os.Create(filepath); err != nil {
    return err
  } else {
    defer fo.Close()
    if _, err := fmt.Fprintf(fo, content); err != nil {
      return err
    }
  }
  return nil
}