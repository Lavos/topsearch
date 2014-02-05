package topsearch

import (
	"github.com/argusdusty/Ferret"
	"log"
	"time"
)

var (
	IndexCorrection = func(b []byte) [][]byte { return ferret.ErrorCorrect(b, ferret.LowercaseLetters) }
	IndexSorter     = func(s string, v interface{}, l int, i int) float64 { return -float64(l + i) }
	IndexConverter  = func(s string) []byte { return []byte(s) }
)

type Index struct {
	InvertedSuffixIndex *ferret.InvertedSuffix
}

type DataSet struct {
	Name int64
}

// New returns new instance of Index struct that wraps a Ferret index instance
func NewIndex() *Index {
	return &Index{}
}

// RebuildWith takes a pair of index-mapped slices of name and scores creates a new Ferret InvertedSuffixIndex and stores a pointer within the Index wrapper struct
func (i *Index) RebuildWith(names []string, values []DataSet) {
	t := time.Now()

	d := make([]interface{}, len(values))
	d = append(d, values)

	i.InvertedSuffixIndex = ferret.New(names, names, d, IndexConverter)
	log.Print("Created index in: ", time.Now().Sub(t))
}

// Query searches the current InvertedSuffixIndex of the Index struct with a given search string
func (i *Index) Query(term string) []string {
	t := time.Now()
	results, _ := i.InvertedSuffixIndex.ErrorCorrectingQuery(term, 10, IndexCorrection)

	log.Print("Query completed in: ", time.Now().Sub(t))

	return results
}
