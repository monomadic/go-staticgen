// todo: rename to styles.go, change to be library agnostic.

package main

import (
  "github.com/yosssi/gcss"
  "strings"
)

// func TransformGcss(c *TemplateWriter) error {
//   var err error
//   return err
// }

type GcssProcessor struct {
}

func (p GcssProcessor) compile(tpl *TemplateWriter) error {
  _, err := gcss.Compile(tpl.writer, tpl.buffer)
  return err
}

func (p GcssProcessor) dstfile(filename string) string {
  return strings.Replace(convertSrcToDestPath(filename), "sass", "css", 1)
}

func compileGcss(tpl *TemplateWriter) error {
  _, err := gcss.Compile(tpl.writer, tpl.buffer)
  return err
}

// func compileGcss(from string) error {
//   var fromdoc bytes.Buffer
//   var todoc bytes.Buffer
//   var err error

//   sitename := filepathToSitename(from)

//   if fromdoc, err = processTemplate(from, filepath.Join("sites", sitename, "styles", "**")); err != nil {
//     return err
//   }

//   if _, err := gcss.Compile(&todoc, &fromdoc) ; err == nil {
//     writeStringToFile(gcssOutputFilePath(from), todoc.String())
//   } else {
//     return err
//   }
//   return nil
// }

func gcssOutputFilePath(from string) string {
  return strings.Replace(convertSrcToDestPath(from), "sass", "css", 1)
}
