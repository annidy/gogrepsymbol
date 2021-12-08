package seacher

import (
	"fmt"
	"os"
	"os/exec"
)

func NMSearch(file string, symbol string) ([]string, bool) {
	out, err := exec.Command("nm", file).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil, false
	}
	return LineSearch(out, symbol)
}
