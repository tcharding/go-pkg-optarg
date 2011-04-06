# This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
# license. Its contents can be found at:
# http://creativecommons.org/publicdomain/zero/1.0/ and
# http://creativecommons.org/publicdomain/zero/1.0/legalcode

include $(GOROOT)/src/Make.inc

TARG = github.com/jteeuwen/go-pkg-optarg
GOFILES = optarg.go string.go

include $(GOROOT)/src/Make.pkg
