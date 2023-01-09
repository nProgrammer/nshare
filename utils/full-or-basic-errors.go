package utils

import (
	"os"
	"strconv"
)

func FBErr(fullError string, bError string) string {
	if fe, _ := strconv.ParseBool(os.Getenv("FULL_ERRORS")); fe {
		return fullError
	} else {
		return bError
	}
}
