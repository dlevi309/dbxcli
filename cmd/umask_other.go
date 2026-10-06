//go:build !(darwin || dragonfly || freebsd || illumos || linux || netbsd || openbsd || solaris)

package cmd

var processUmask = 0o022
