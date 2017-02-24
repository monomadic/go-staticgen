package main

import (
  "flag"
  "fmt"
)

func processArgs() {
  flag.Usage = func() {
  fmt.Println(`

  NAME:
    go-staticgen

  DESCRIPTION:
    An opinionated multi-site static generator written in golang.

  COMMANDS:
    new <sitename>       Creates a new site scaffold in the current directory.
    build                Process files from sites directory into public directory.
    serve                Serve your site locally
`[1:])
  }
  flag.Parse()

  switch {
  case flag.Arg(0) == "build":
    processSites()
  default:
    flag.Usage()
  }
}
