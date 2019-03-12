package main

import (
	"fmt"
	"path/filepath"
)

type Config struct {
	Host string
	Port string
	Name string

	SiteDir  string
	BuildDir string
	ImageDir string
	PageDir  string
	StyleDir string

	SrcDir  string
	DestDir string
}

func (cfg *Config) ServerURL() string {
	return fmt.Sprintf("http://%s:%s/", cfg.Host, cfg.Port)
}

func (cfg *Config) PagesSrcDir(site string) string {
	return filepath.Join(cfg.SiteDir, site, cfg.PageDir)
}

func (cfg *Config) ErrorFile() string {
	return filepath.Join(cfg.BuildDir, "error.html")
}
