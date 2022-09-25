package utils

import (
	"strconv"
)

func Str2Uint32(str string) (uint32, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}
