package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/codeskyblue/go-sh"
)

var (
    src string
    dst string
)

func main() {
    parseFlags()
    checkSource()
    checkDestination()
    runBackup()
}

func parseFlags() {
    flag.StringVar(&src, "src", src, "source path to backup")
    flag.StringVar(&dst, "dst", dst, "destination directory")

    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [-quiet] -src <path> -dst <path>\n", os.Args[0])
        flag.PrintDefaults()
    }

    flag.Parse()
}

func checkSource() {
    if src == "" {
        exit(1, "Missing source `-src <path>`")
    }

    var err error

    src, err = filepath.Abs(src)
    if err != nil {
        exit(1, fmt.Sprintf("Invalid source path: %s", err.Error()))
    }

    _, err = os.Stat(src)
    if err != nil {
        exit(1, fmt.Sprintf("source path error: %s", err.Error()))
    }
}

func checkDestination() {
    if dst == "" {
        exit(1, "Missing destination directory for backups `-dst <path>`\n")
    }

    var err error

    dst, err = filepath.Abs(dst)
    if err != nil {
        exit(1, fmt.Sprintf("Invalid destination path: %s", err.Error()))
    }

    info, err := os.Stat(dst)
    if err != nil {
        makeDirAll(dst)
    } else if !info.IsDir() {
        exit(1, fmt.Sprintf("destination path have to be a directory: \"%s\"", dst))
    }
}

func runBackup() {
    session := sh.NewSession()

    // TODO: run the backup using tar command

    err := session.Run()
    if err != nil {
        exit(1, err.Error())
    }
}
