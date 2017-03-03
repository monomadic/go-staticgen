package main

import (
  "flag"
  "fmt"
  "os"
)

func processArgs() {
  flag.Usage = func() {
  fmt.Println(`

  NAME:
    go-staticgen

  DESCRIPTION:
    An opinionated multi-site static generator written in golang.

  COMMANDS:
    new <sitename>       Creates a new site scaffolding.
    build <sitename>     Process a specific site only.
    build                Process all sites.
    serve                Serve your site locally.
`[1:])
  }
  flag.Parse()

  switch {
  case flag.Arg(0) == "new":
    p := flag.Arg(1)
    if p == "" {
      flag.Usage()
      os.Exit(1)
    }
    scaffold(flag.Arg(1))

  case flag.Arg(0) == "build":
    p := flag.Arg(1)
    if p == "" {
      err := processSites(); if err != nil {
        consoleError(err)
      }
    } else {
      // processSite(flag.Arg(1))
    }

  // case flag.Arg(0) == "dupe":
  //   p := flag.Arg(1)
  //   if p == "" {
  //     duplicateSite(p)
  //   } else {
  //     // processSite(flag.Arg(1))
  //   }

  case flag.Arg(0) == "serve":
    serve()

  default:
    flag.Usage()
  }
}
