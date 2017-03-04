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
  return filepath.Walk("sites/"+sitename, func(name string, info os.FileInfo, err error) error {
    if info == nil { return err }

    from := filepath.ToSlash(name)
    dot := filepath.Base(name)[0]

    if dot != '.' && dot != '_' {
      if info.IsDir() {
        if makeDirIfMissing(convertSrcToDestPath(from)); err != nil {
          return err
        }
      } else {
        consoleInfo("Processing file: " + from)
        if err := processFile(from, sitename); err != nil {
          return err
        }
      }
      return nil
    } else {
      return nil
    }
    return err
  })
}

func processFile(filename string, sitename string) error {
  switch filepath.Ext(filename) {
  case ".ace":
    return compileAce(filename)
  case ".sass":
    return compileGcss(filename, sitename)
  default:
    return copyFile(filename)
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

func makeDirIfMissing(dir string) error {
  if _, err := os.Stat(dir); err != nil {
    if os.IsNotExist(err) {
      consoleSuccess("[MKDIR] " + dir)
      os.MkdirAll(dir, os.ModePerm)
    }
    return nil
  } else {
    return err
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