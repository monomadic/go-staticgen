package main

import (
)

type Processor interface {
  dstfile()     string
  compile() error
}
