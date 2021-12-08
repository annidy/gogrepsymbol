package seacher

import (
	"bufio"
	"bytes"
	"strings"
)

func ContainsI(a string, b string) bool {
	return strings.Contains(
		strings.ToLower(a),
		strings.ToLower(b),
	)
}

func TextSearch(txt []byte, pattern string) ([]string, bool) {
	s := bufio.NewScanner(bytes.NewReader(txt))
	s.Split(bufio.ScanWords)
	r := make([]string, 0)
	succ := false
	for s.Scan() {
		word := s.Text()
		if ContainsI(word, pattern) {
			r = append(r, word)
			succ = true
		}
	}
	return r, succ
}

func LineSearch(txt []byte, pattern string) ([]string, bool) {
	s := bufio.NewScanner(bytes.NewReader(txt))
	s.Split(bufio.ScanLines)
	r := make([]string, 0)
	succ := false
	for s.Scan() {
		word := s.Text()
		if ContainsI(word, pattern) {
			r = append(r, word)
			succ = true
		}
	}
	return r, succ
}
