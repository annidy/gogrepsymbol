package main

import (
	"errors"
	"path/filepath"

	"github.com/annidy/gogrepsymbol/seacher"
)

type SymbolChecker struct {
	source    bool
	framework bool
	archive   bool
}

func NewSymbolChecker() *SymbolChecker {
	return &SymbolChecker{true, true, true}
}

func (checker *SymbolChecker) GetSearchType(path string) (seacher.SearchType, error) {
	if checker.source {
		sourceExts := []string{".c", ".cpp", ".m", ".mm"}
		ext := filepath.Ext(path)
		for _, v := range sourceExts {
			if v == ext {
				return seacher.FileType, nil
			}
		}
	}
	if checker.framework {
		dir, file := filepath.Split(path)
		_, dirExt := filepath.Split(filepath.Base(dir))
		if dirExt == file+".framework" {
			return seacher.LibType, nil
		}
	}

	if checker.archive {
		ext := filepath.Ext(path)
		if ext == ".a" {
			return seacher.LibType, nil
		}
	}
	return seacher.UnknowType, errors.New("unknow filetype")
}
