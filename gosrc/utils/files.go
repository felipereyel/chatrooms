package utils

import (
	"os"
)

func CheckFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
