package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/annidy/gogrepsymbol/seacher"
	"github.com/facebookgo/symwalk"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stderr, os.Args[0], " PATTER [FOLDER [FOLDER...]]")
		flag.PrintDefaults()
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			flag.Usage()
		}
	}()

	flag.Parse()
	args := flag.Args()
	var symbol string
	var folders []string
	if len(args) == 0 {
		panic("need a seach")
	}
	symbol = args[0]
	if len(args) > 1 {
		folders = args[1:]
	} else {
		folders = []string{"."}
	}

	symbolChecker := SymbolChecker{true, true, true}

	for _, folder := range folders {

		folder, err := filepath.Abs(folder)
		if err != nil {
			panic(err)
		}

		err = symwalk.Walk(folder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return nil
			}
			if seachType, _ := symbolChecker.GetSearchType(path); seachType != seacher.UnknowType {
				if outs, succ := seacher.Search(path, symbol, seachType); succ {
					fmt.Println(path)
					for _, out := range outs {
						fmt.Println("\t", out)
					}
				}
			}

			return nil
		})
		if err != nil {
			panic(err)
		}
	}
}
