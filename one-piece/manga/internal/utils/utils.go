package utils

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

const (
	JPEG = "image/jpeg"
	PNG  = "image/png"
)

func GetExtension(t string) (ext string, err error) {
	switch t {
	case PNG:
		ext = "png"
	case JPEG:
		ext = "jpg"
	default:
		ext = "unknown"
		err = fmt.Errorf("unknown extension from type \"%s\"", t)
	}

	return ext, err
}

func ConvertImagesToPDF(p string) error {
	log.Printf("[INFO] Run: `%s %s %s %s`", "magick", filepath.Join(p, "*.{jpg,png}"), "-quality 100 -density 150", p+".pdf")
	cmd := exec.Command("magick", filepath.Join(p, "*.{jpg,png}"), "-quality", "100", "-density", "150", p+".pdf")
	_, err := cmd.Output()
	return err
}
