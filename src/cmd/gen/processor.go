package main

import (
)

type Processor interface {
  dstfile(string)           string
  compile(*TemplateWriter)  error
}
