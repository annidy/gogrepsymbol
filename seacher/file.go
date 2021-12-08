package seacher

import (
	"fmt"
	"io/ioutil"
	"os"
)

func FileSearch(file string, symbol string) ([]string, bool) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, false
	}
	return TextSearch(bytes, symbol)
}
