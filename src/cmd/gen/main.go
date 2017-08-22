package main

var cfg = Config {
	Host: "localhost",
	Port: "9000",
	Name: "go-staticgen",
	SiteDir: "sites",
	BuildDir: "public",
	ImageDir: "images",
	PageDir: "pages",
	StyleDir: "styles",
	SrcDir: "sites",
	DestDir: "public",
}

func main() {
	processArgs()
}
