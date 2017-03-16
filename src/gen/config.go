package main

import (
  "fmt"
)

type config struct {
  Host string
  Port string
  Name string
  SiteDir string
  BuildDir string
}

func (cfg *config) ServerURL() string {
  return fmt.Sprintf("http://%s:%s/", cfg.Host, cfg.Port)
}
