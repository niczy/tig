package service

type RootTable struct {
	rootNodeHash []string
}

var (
	rootTable = &RootTable{rootNodeHash: make([]string, 0, 0)}
)
