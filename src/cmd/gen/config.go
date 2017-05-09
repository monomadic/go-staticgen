package main

import (
	"fmt"
	"path/filepath"
)

type config struct {
	Host     string
	Port     string
	Name     string
	SiteDir  string
	BuildDir string
	ImageDir string
	PageDir  string
	StyleDir string

	SrcDir  string
	DestDir string
}

func (cfg *config) ServerURL() string {
	return fmt.Sprintf("http://%s:%s/", cfg.Host, cfg.Port)
}

func (cfg *config) PagesSrcDir(site string) string {
	return filepath.Join(cfg.SiteDir, site, cfg.PageDir)
}

func (cfg *config) ErrorFile() string {
	return filepath.Join(cfg.BuildDir, "error.html")
}
