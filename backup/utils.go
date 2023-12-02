package main

import (
	"fmt"
	"os"
)

func exit(code int, msg string) {
    fmt.Fprintf(os.Stderr, "%s\n", msg)
    os.Exit(code)
}

func makeDirAll(p string) {
    err := os.MkdirAll(p, os.ModeDir|os.ModePerm)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        os.Exit(1)
    }
}
