package helper

import (
	"errors"
	"strings"
)

func CheckFile(filename string) (string, error) {
	formatfile := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])
	if formatfile != "jpg" && formatfile != "jpeg" && formatfile != "png" {
		return "", errors.New("Forbidden File Tipe")
	}

	return formatfile, nil
}

func CheckSize(size int64) error {
	if size == 0 {
		return errors.New("Illegal File Size")
	}
	if size > 3000000 {
		return errors.New("File Size is to big")
	}

	return nil
}
