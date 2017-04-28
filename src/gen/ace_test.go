package main

import(
  "testing"
)

func TestaceOutputFilePath(t *testing.T) {
  actual := aceOutputFilePath("/sites/pages/test.ace")
  expected := "/public/est.htm"
  if actual != expected { t.Error("expected " + expected + " but got " + actual) }
}
