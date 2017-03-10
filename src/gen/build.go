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
  // if err := makeDirIfMissing("public"); err != nil { return err }

  for _, file := range files {
    if file.IsDir() {
      dot := filepath.Base(file.Name())[0]
      if dot != '.' && dot != '_' {
        consoleInfo("\nProcessing Site: "+ file.Name())
        if err := makeDirIfMissing("public/" + file.Name()); err != nil { return err }
        if err := processSite(file.Name()); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

func processSite(sitename string) error {
  os.RemoveAll("public/error.html")
  os.RemoveAll("public/"+sitename)
  // if err := makeDirIfMissing("public/"+sitename); err != nil { return err }
  if err := processPages(sitename); err != nil { return err }
  if err := processStyles(sitename); err != nil { return err }
  if fileExists("sites/"+sitename+"/images") {
    if err := processImages(sitename); err != nil { return err }
  }

  return nil
}

func processImages(sitename string) error {
  var err error
  var files []string

  if err := makeDirIfMissing("public/"+sitename+"/images"); err != nil { return err }

  if files, err = RecursiveGlob("sites/"+sitename+"/images"); err == nil {
    for _, name := range files {
      err = copyFile(name)
    }
  }
  return err
}

func processPages(sitename string) error {
  var err error
  var files []string

  if files, err = FileTypeGlob("sites/"+sitename, ".ace"); err == nil {
    for _, name := range files {
      err = compileAce(name)
      if err == nil { consoleSuccess(fmt.Sprintf("\t" + aceOutputFilePath(name) + "\n")) }
    }
  }
  return err
}

func processStyles(sitename string) error {
  var err error
  var files []string

  if err := makeDirIfMissing("public/"+sitename+"/styles"); err != nil { return err }

  if files, err = FileTypeGlob("sites/"+sitename+"/styles", ".sass"); err == nil {
    for _, name := range files {
      // println(name)
      err = compileGcss(name)
      if err == nil { consoleSuccess(fmt.Sprintf("\t" + gcssOutputFilePath(name) + "\n")) }
    }
  }
  return err
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
      consoleSuccess("\t" + dir)
      os.MkdirAll(dir, os.ModePerm)
    }
    return nil
  } else {
    return err
  }
}

func removeFileIfExists(filename string) {
  os.RemoveAll("public/error.html")
}

func fileExists(filename string) bool {
  if _, err := os.Stat(filename); err == nil {
    return true
  } else {
    return false
  }
}

func writeStringToFile(path string, content string) error {
  if err := makeDirIfMissing(filepath.Dir(path)); err != nil { return err }

  if fo, err := os.Create(path); err != nil {
    return err
  } else {
    defer fo.Close()
    if _, err := fmt.Fprintf(fo, content); err != nil {
      return err
    }
  }
  return nil
}

func filepathToSitename(filepath string) string {
  return strings.Split(filepath, "/")[1]
}

func PartialGlob(dirpath string, extMask string) ([]string, error) {
  var files []string

  if allFiles, err := RecursiveGlob(dirpath); err == nil {
    for _, name := range allFiles {
      dot := filepath.Base(name)[0]
      ext := filepath.Ext(name)
      if dot != '.' && ext == extMask {
        files = append(files, name)
      }
    }
  } else {
    return nil, err
  }
  return files, nil
}

func FileTypeGlob(dirpath string, extMask string) ([]string, error) {
  var files []string

  if allFiles, err := RecursiveGlob(dirpath); err == nil {
    for _, name := range allFiles {
      dot := filepath.Base(name)[0]
      ext := filepath.Ext(name)
      if dot != '.' && dot != '_' && ext == extMask {
        files = append(files, name)
      }
    }
  } else {
    return nil, err
  }
  return files, nil
}

func RecursiveGlob(dirpath string) ([]string, error) {
  var paths []string
  err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
    if err != nil { return err }

    if !info.IsDir() {
      filename := filepath.Base(path)
      dot := filepath.Base(path)[0]

      if dot != '.' && filename != ".DS_Store" {
        paths = append(paths, path)
      }
    } else {
      if err := makeDirIfMissing(path); err != nil { return err }
    }

    return nil
  })
  if err != nil {
    return nil, err
  }
  return paths, nil
}
