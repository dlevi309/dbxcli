//go:build darwin || dragonfly || freebsd || illumos || linux || netbsd || openbsd || solaris

package cmd

import "syscall"

var processUmask = readUmask()

func readUmask() int {
	mask := syscall.Umask(0)
	syscall.Umask(mask)
	return mask
}
