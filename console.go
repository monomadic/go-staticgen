package main

import (
  "github.com/fatih/color"
)

func checkFatal(err error) {
  if err != nil {
    color.Red("[ERROR]: %s\n", err)
  }
}

func consoleError(err error) {
  color.Red("[ERROR]: %s\n", err)
}

func consoleInfo(info string) {
  color.Yellow(info)
}

func consoleSuccess(success string) {
  color.Green(success)
}
