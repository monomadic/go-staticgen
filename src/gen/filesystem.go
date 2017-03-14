// filesystem operations

package main

import (
  "io"
  "os"
  "fmt"
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
  if err != nil {
    return nil, err
  }
  return paths, nil
}
