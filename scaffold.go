package main

import (
)

func scaffold(name string) {
  consoleInfo("Generating new site scaffold: " + name)
  ensureDirectoryStructure()
  makeDirIfMissing("sites/" + name)
}

func ensureDirectoryStructure() {
  // ensures the most basic directory structure is in place
  makeDirIfMissing("public/")
  makeDirIfMissing("sites/")
}
