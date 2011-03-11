// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/ and
// http://creativecommons.org/publicdomain/zero/1.0/legalcode

package optarg

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Args = []string{ // manually rebuild os.Args for testing purposes.
		os.Args[0],
		"--bin", "/a/b/foo/bin",
		"--arch", os.Getenv("GOARCH"),
		"-nps", "/a/b/foo/src",
		"foo.go", "bar.go",
	}

	// Add some flags
	Add("s", "source", "Path to the source folder. Here is some added description information which is completely useless, but it makes sure we can pimp our sexy Usage() output when dealing with lenghty, multi-line description texts.", "")
	Add("b", "bin", "Path to the binary folder.", "")
	Add("a", "arch", "Target architecture.", os.Getenv("GOARCH"))
	Add("n", "noproc", "Skip pre/post processing.", false)
	Add("p", "purge", "Clean compiled packages after linking is complete.", false)

	// These will hold the option values.
	var src, bin, arch string
	var noproc, purge bool

	// Parse os.Args
	for opt := range Parse() {
		switch opt.ShortName {
		case "s":
			src = opt.String()
		case "b":
			bin = opt.String()
		case "a":
			arch = opt.String()
		case "p":
			purge = opt.Bool()
		case "n":
			noproc = opt.Bool()
		}
	}

	// Make sure everything went ok.

	if arch != os.Getenv("GOARCH") {
		t.Errorf("Parse(): incorrect value for arch: %s", arch)
	}

	if bin != "/a/b/foo/bin" {
		t.Errorf("Parse(): incorrect value for bin: %s", bin)
	}

	if src != "/a/b/foo/src" {
		t.Errorf("Parse(): incorrect value for src: %s", src)
	}

	if !purge {
		t.Errorf("Parse(): purge is not set")
	}

	if !noproc {
		t.Errorf("Parse(): noproc is not set")
	}

	if len(Remainder) != 2 { // should contain: foo.go, bar.go
		t.Errorf("Parse(): incorrect number of remaining arguments. Expected 2. got %d", len(Remainder))
	}

	// This outputs the usage information. No need to do this in a test case.
	Usage()
}
