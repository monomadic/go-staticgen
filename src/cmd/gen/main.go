package main

func main() {
	var cfg config
	cfg.Host = "localhost"
	cfg.Port = "9000"
	cfg.Name = "go-staticgen"
	cfg.SiteDir = "sites"
	cfg.BuildDir = "public"
	cfg.ImageDir = "images"
	cfg.PageDir = "pages"
	cfg.StyleDir = "styles"

	cfg.SrcDir = "sites"
	cfg.DestDir = "public"

	cfg.processArgs()
}
