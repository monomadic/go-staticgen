// command line router

package main

import (
	"flag"
	"fmt"
	"os"
)

func processArgs() {
	flag.Usage = func() {
		fmt.Printf(`

  NAME:
    %s

  DESCRIPTION:
    An opinionated multi-site static generator written in golang. 🥒

  COMMANDS:
    new <sitename>       Creates a new site scaffolding.
    build <sitename>     Process a specific site only.
    build                Process all sites.
    serve                Serve your site locally.

`[1:], cfg.Name)
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

	case flag.Arg(0) == "build" || flag.Arg(0) == "b":
		siteName := flag.Arg(1)
		if siteName == "" {
			err := processSites()
			if err != nil {
				consoleError(err)
			}
		} else {
			err := processSite(siteName)
			if err != nil {
				consoleError(err)
			}
		}

	// case flag.Arg(0) == "dupe":
	//   p := flag.Arg(1)
	//   if p == "" {
	//     duplicateSite(p)
	//   } else {
	//     // processSite(flag.Arg(1))
	//

	case flag.Arg(0) == "serve" || flag.Arg(0) == "s":
		serve()

	default:
		flag.Usage()
	}
}
