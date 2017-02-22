package main

import (
  "github.com/fatih/color"
  // "strings"
)

func consoleError(err error) {
  color.Red("[ERROR]: %s\n", err)
}

func consoleSuccess(success string) {
  color.Green(success)
}
