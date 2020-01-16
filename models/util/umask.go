// +build !windows

package util

import "syscall"

func Umask(mask int) (old int, err error) {
	return syscall.Umask(mask), nil
}
