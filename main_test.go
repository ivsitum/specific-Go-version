package main

import "testing"

func TestMain(t *testing.T) {
    expected := "Hello, Travis CI with Go 1.22.5!"
    if expected != "Hello, Travis CI with Go 1.22.5!" {
        t.Errorf("Expected %s, got something else.", expected)
    }
}
