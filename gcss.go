package main

import (
  "github.com/yosssi/gcss"
  "bytes"
  "strings"
  "fmt"
)

func compileGcss(from string) {
  var fromdoc bytes.Buffer
  var todoc bytes.Buffer

  to := strings.Replace(convertSrcToDestPath(from), "sass", "css", 1)

  fromdoc = processTemplate(from, "sites/robsaunders/styles/**")

  if _, err := gcss.Compile(&todoc, &fromdoc) ; err == nil {
    writeStringToFile(to, todoc.String())
    consoleSuccess(fmt.Sprintf("[SASS]: " + from + " => " + to + "\n"))
  } else {
    consoleError(err)
  }
}

// func ccompileGcss(from string) {
//   var doc bytes.Buffer
//   println("ok")
//   funcMap := template.FuncMap {
//       "title": strings.Title,
//   }

//   tpl := template.Must(template.New("main.sass").Funcs(funcMap).ParseGlob("sites/robsaunders/styles/**"))
//   // tplVars := map[string]string {
//   //     "Title": "Hello world",
//   //     "Content": "Hi there",
//   // }
//   println("executing: " + tpl.Name())
//   // tpl.ExecuteTemplate(os.Stdout, tpl.Name(), tplVars)
//   err := tpl.Execute(&doc, nil)
//   if err != nil { consoleError(err) }
//   println("here:")
//   println(doc.String())
// }

// func cocmpileGcss(from string) {
//   // var doc bytes.Buffer
//   // // var tmplGet = template.Must(template.ParseFiles(from)).Funcs(funcMap)
//   // // tmplGet.Execute(&doc, nil)

//   // // var tpl = template.Must(template.New("new").Parse(from))

//   // funcMap := template.FuncMap{
//   //   "title": strings.Title,
//   // }

//   // tmpl, err := template.New(n).Funcs(funcMap).ParseFiles(from)
//   // if err != nil { consoleError(err) }

//   // err = tmpl.Execute(&doc, "the go programming language")
//   // if err != nil { consoleError(err) }

//   // // consoleSuccess(fmt.Sprintf("[SASS]: " + from + " => " + to + "\n"))
//   // println("ok")
//   // print(doc.String())



//   // if err := tmplGet.Execute(&doc, nil); err != nil {
//   //   consoleError(err)
//   // } else {
//   //   print(doc.String())
//   // }
// }

// func _compileGcss(from string) {
//   var doc bytes.Buffer

//   to := strings.Replace(convertSrcToDestPath(from), "sass", "css", 1)

//   // t, err := template.New("").Funcs(funcMap).ParseGlob("templates/*.tmpl"‌​)

//   if tpl, err := template.ParseFiles(from); err == nil {
//     var sassdoc bytes.Buffer

//     funcMap := template.FuncMap{"poop": TemplateInclude}

//     tpl.Funcs(funcMap)

//     if err := tpl.Execute(&doc, nil); err != nil {
//       consoleError(err)
//     }

//     print(doc.String())

//     print("=====")

//     if _, err = gcss.Compile(&sassdoc, &doc) ; err == nil {
//       print(sassdoc.String())
//       writeStringToFile(to, sassdoc.String())
//       consoleSuccess(fmt.Sprintf("[SASS]: " + from + " => " + to + "\n"))
//     } else {
//       consoleError(err)
//     }
//   } else {
//     consoleError(err)
//   }
// }

