package main

import (
  // "os"
  "github.com/yosssi/gcss"
  "bytes"
  "strings"
  "fmt"
  "text/template"
)

func compileGcss(from string) {
  var doc bytes.Buffer

  to := strings.Replace(convertSrcToDestPath(from), "sass", "css", 1)

  if tpl, err := template.ParseFiles(from); err == nil {
    var sassdoc bytes.Buffer

    if err := tpl.Execute(&doc, nil); err != nil {
      consoleError(err)
    }

    print(doc.String())

    print("=====")


    // fromstream, err := os.Open(from)

    if _, err = gcss.Compile(&sassdoc, &doc) ; err == nil {
      print(sassdoc.String())
      writeStringToFile(to, sassdoc.String())
      consoleSuccess(fmt.Sprintf("[SASS]: " + from + " => " + to + "\n"))
    } else {
      consoleError(err)
    }
  } else {
    consoleError(err)
  }
}

// func compileGcss(filename string) {
//   var doc bytes.Buffer

//   toMake := strings.Replace(filename, "sites", "public", 1)
//   toMake = strings.Replace(toMake, "sass", "css", 1)

//   from, err := os.Open(filename)
//   _, err = gcss.Compile(&doc, from)

//   if err != nil {
//     consoleError(err)
//   } else {
//     // writeStringToFile(toMake, doc.String())
//     writeStringToFile(toMake, doc.String())
//     consoleSuccess(fmt.Sprintf("[SASS]: " + filename + " => " + toMake + "\n"))
//   }
// }

// func processAsTemplate(file string) string {
//   var doc bytes.Buffer
//   t := template.New("css")
//   if t, err := t.Parse(file); err != nil {
//     consoleError(err)
//   } else {
//     t.Execute(&doc, nil)
//     println(doc.String())
//   }
//   return doc.String()
// }
