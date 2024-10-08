package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/codeskyblue/go-sh"
)

var (
	src    string
	dst    string
	system bool
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
	flag.BoolVar(&system, "system", system, "using bsdtar for backup")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -src <path> -dst <path>\n", os.Args[0])
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
	date := time.Now()

	cwd := filepath.Dir(src)
	base := filepath.Base(src)

	out := filepath.Join(
		dst,
		fmt.Sprintf(
			"%s-%d-%02d-%02d.tar.gz",
			base, date.Year(), date.Month(), date.Day(),
		),
	)

	session := sh.NewSession()

	if system {
		session.Command(
			"bsdtar",
			"--exclude="+dst,
			"--acls",
			"--xattrs",
			"-cpvaf",
			out,
			"-C",
			cwd,
			base,
		)
	} else {
		session.Command(
			"tar",
			"--create",
			"--gzip",
			"--preserve-permissions",
			"--file",
			out,
			"-C",
			cwd,
			base,
		)
	}

	err := session.Run()
	if err != nil {
		exit(1, err.Error())
	}
}
