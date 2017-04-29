// scaffold.go
//  templates used during the 'gen init' phase.

package main

import (
)

func scaffold(name string) error {
  var err error
  consoleInfo("Generating new site scaffold: " + name)
  err = ensureDirectoryStructure(name)
  err = writeStringToFile("sites/" + name + "/pages/index.ace", indexAceTemplate())
  err = writeStringToFile("sites/" + name + "/pages/_head.ace", indexAceHead(name))
  err = writeStringToFile("sites/" + name + "/styles/main.sass", sassStyle())
  err = writeStringToFile("sites/" + name + "/styles/_settings.sass", sassPartial())

  if err != nil { consoleError(err) }
  return err
}

func ensureDirectoryStructure(name string) error {
  var err error
  err = makeDirIfMissing("sites/" + name + "/scripts/")
  err = makeDirIfMissing("sites/" + name + "/styles/")
  err = makeDirIfMissing("sites/" + name + "/images/")
  return err
}

func indexAceTemplate() string {
  return `= doctype html

html lang=en
  = include _head

  body
    h1 Welcome to the site
`
}

func indexAceHead(name string) string {
  return `head
  title `+name+`
  link href="styles/main.css" rel="stylesheet" type="text/css"
`
}

func sassStyle() string {
  return `{{ template "_settings.sass" }}

html
  font-family: Helvetica
  text-align: center
  padding-top: 40px
  background-color: $background-color

h1
  color: #00F4D1
  letter-spacing: 0.3em
  font-weight: normal

h2
  color: #00E7EE
  font-weight: normal
  letter-spacing: 0.2em

a
  color: #00E7EE
  border-bottom: 5px solid #FAFF9D
  text-decoration: none
`
}

func sassPartial() string {
  return `$background-color: white
$text-color: #222
`
}