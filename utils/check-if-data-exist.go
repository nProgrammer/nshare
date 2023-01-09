package utils

import "os"

func CheckDataDir(path string) bool {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}
