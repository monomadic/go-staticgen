package main

import (
  "io/ioutil"
  "crypto/sha256"
  "log"
  "encoding/hex"
)

func checksum(file string) string {
  hasher := sha256.New()
  s, err := ioutil.ReadFile(file)    
  hasher.Write(s)

  if err != nil { log.Fatal(err) }

  return hex.EncodeToString(hasher.Sum(nil))
}