package seacher

type SearchType int

const (
	UnknowType SearchType = iota
	LibType
	FileType
)

func Search(file string, symbol string, searchType SearchType) ([]string, bool) {
	switch searchType {
	case LibType:
		return NMSearch(file, symbol)
	case FileType:
		return FileSearch(file, symbol)
	}
	return nil, false
}
