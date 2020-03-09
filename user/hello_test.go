package main

import (
  "testing"
)

func TestHello(t *testing.T) {
  actual := Hello()
  expected := "Hello User!"

  if actual != expected {
    t.Errorf("Actual: %q. Expected: %q", actual, expected)
  }
}
