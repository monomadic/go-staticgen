package main

import (
  "fmt"
)

func scaffold(name string) {
  consoleInfo("Generating new site scaffold: " + name)
  ensureDirectoryStructure()
  makeDirIfMissing("sites/" + name)
  makeDirIfMissing("sites/" + name + "/scripts/")
  makeDirIfMissing("sites/" + name + "/styles/")
  makeDirIfMissing("sites/" + name + "/images/")
  writeStringToFile("sites/" + name + "/index.ace", indexAceTemplate(name))
  consoleSuccess(fmt.Sprintf("[ACE] sites/%s/index.ace", name))
  writeStringToFile("sites/" + name + "/_head.ace", indexAceHead(name))
  consoleSuccess(fmt.Sprintf("[ACE] sites/%s/_head.ace", name))
  writeStringToFile("sites/" + name + "/styles/main.sass", sassStyle())
  consoleSuccess(fmt.Sprintf("[SASS] sites/%s/styles/main.sass", name))
  writeStringToFile("sites/" + name + "/styles/_partial.sass", sassPartial())
  consoleSuccess(fmt.Sprintf("[SASS] sites/%s/styles/_partial.sass", name))
}

func ensureDirectoryStructure() {
  // ensures the most basic directory structure is in place
  makeDirIfMissing("public/")
  makeDirIfMissing("sites/")
}

func indexAceTemplate(name string) string {
  return `= doctype html

html lang=en
  = include sites/` + name + `/_head

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
  return `{{ template "_partial.sass" }}

html
  font-family: Helvetica
  text-align: center
  padding-top: 40px

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
}`
}

func sassPartial() string {
  return `body
  background-color: red
`
}