// filesystem operations

package main

import (
  "io"
  "os"
  "fmt"
  "strings"
  "path/filepath"
)

// todo: this really doesn't need to be here.
func copyFile(from string) error {
  if err := cp(from, convertSrcToDestPath(from)); err != nil {
    return err
  } else {
    consoleSuccess(fmt.Sprintf("\t%s\n", convertSrcToDestPath(from)))
  }
  return nil
}

// cp: copy a file from src to dst
func cp(src string, dst string) error {
  s, err := os.Open(src)
  if err != nil {
    return err
  }
  // no need to check errors on read only file, we already got everything
  // we need from the filesystem, so nothing can go wrong now.
  defer s.Close()
  d, err := os.Create(dst)
  if err != nil {
    return err
  }
  if _, err := io.Copy(d, s); err != nil {
    d.Close()
    return err
  }
  return d.Close()
}

// check if file exists
func fileExists(filename string) bool {
  if _, err := os.Stat(filename); err == nil {
    return true
  } else {
    return false
  }
}

// returns all files and directories within a given directory, recursively
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
  return paths, err
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

func convertSrcToDestPath(filename string) string {
  return strings.Replace(filename, "sites", "public", 1)
}

func findLayoutFile(filename string) string {
  return filepath.Dir(filename) + "/_layout"
}

func trimExt(filename string) string {
  return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func removeFileIfExists(filename string) {
  os.RemoveAll("public/error.html")
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
