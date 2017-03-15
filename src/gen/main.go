package main

import (
)

func main() {
  var cfg config
  cfg.Host = "localhost"
  cfg.Port = "9000"
  cfg.Name = "go-staticgen"
  cfg.SiteDir = "sites"

  cfg.processArgs()
}
