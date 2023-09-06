package utils

import (
	"fmt"
	"os/exec"
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
	cmd := exec.Command("magick", p+".{jpg,png}", "-quality", "100", "-density", "150", p+".pdf")
	_, err := cmd.Output()
	return err
}
