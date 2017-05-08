package main

// Processor : interface for processor library wrappers.
type Processor interface {
	dstfile(string) string
	compile(*TemplateWriter) error
}
