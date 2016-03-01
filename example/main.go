package main

import (
	"fmt"
	"os"

	"github.com/tcharding/bak-go-pkg-optarg/go-pkg-optarg"
)

func main() {
	optarg.Add("a", "all", "do all the stuff", false)
	optarg.Add("h", "help", "show this help list", false)

	usageInfo := fmt.Sprintf("Usage: %s [options] arg", os.Args[0])
	optarg.SetUsageInfo(usageInfo)

	var aFlag, hFlag bool
	const numArgs = 1

	for opt := range optarg.Parse() {
		switch opt.ShortName {
		case "a":
			aFlag = opt.Bool()
		case "h":
			hFlag = opt.Bool()
		}
	}

	if len(optarg.Remainder) != numArgs || hFlag {
		optarg.Usage()
		os.Exit(1)
	}

	fmt.Printf("aFlag: %t\nhFlag: %t\nargument: %s\n", aFlag, hFlag, optarg.Remainder[0])
}
