// filesystem operations

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// todo: this really doesn't need to be here.
func copyFile(from string) error {
	if err := cp(from, convertSrcToDestPath(from)); err != nil {
		return err
	}
	consoleSuccess(fmt.Sprintf("\t%s\n", convertSrcToDestPath(from)))
	return nil
}

// cp: copy a file from src to dst
func cp(src string, dst string) error {
	s, err := os.Open(src)
	defer s.Close()
	if err != nil {
		return err
	}

	d, err := os.Create(dst)
	defer d.Close()

	if err != nil {
		return err
	}

	if _, err := io.Copy(d, s); err != nil {
		return err
	}

	return err
}

// check if file exists
func fileExists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}
	return false
}

func (cfg *config) Sites() ([]string, error) {
	var sites []string

	files, err := ioutil.ReadDir(cfg.SiteDir)

	for _, file := range files {
		if file.IsDir() {
			dot := filepath.Base(file.Name())[0]
			if dot != '.' && dot != '_' {
				println(filepath.Base(file.Name()))
				sites = append(sites, filepath.Base(file.Name()))
			}
		}
	}

	return sites, err
}

// RecursiveGlob : returns all files and directories within a given directory, recursively
func RecursiveGlob(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			filename := filepath.Base(path)
			dot := filepath.Base(path)[0]

			if dot != '.' && filename != ".DS_Store" {
				paths = append(paths, path)
			}
		} else {
			if err := makeDirIfMissing(path); err != nil {
				return err
			}
		}

		return nil
	})
	return paths, err
}

func makeDirIfMissing(dir string) error {
	var err error
	if _, err = os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			consoleSuccess("\t" + dir)
			os.MkdirAll(dir, os.ModePerm)
		}
		return nil
	}
	return err
}

func filepathToSitename(filepath string) string {
	return strings.Split(filepath, "/")[1]
}

// PartialGlob : returns a string array of directories within a path.
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

// FileTypeGlob : returns a string array of files within a path.
func FileTypeGlob(dirpath string, extMask string) ([]string, error) {
	var files []string

	if allFiles, err := RecursiveGlob(dirpath); err == nil {
		for _, name := range allFiles {
			dot := filepath.Base(name)[0]
			ext := filepath.Ext(name)
			if dot != '.' && dot != '_' && name != ".DS_Store" && ext == extMask {
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
	var err error

	if err = makeDirIfMissing(filepath.Dir(path)); err != nil {
		return err
	}

	fo, err := os.Create(path)
	defer fo.Close()

	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(fo, content)

	return err
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
