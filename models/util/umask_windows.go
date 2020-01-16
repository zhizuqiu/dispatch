// +build windows

package util

import (
	"errors"
)

func Umask(mask int) (int, error) {
	return 0, errors.New("platform and architecture is not supported")
}
