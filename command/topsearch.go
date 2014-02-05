package main

import (
	"github.com/Lavos/topsearch"
	"os"
	"encoding/csv"
	"log"
)

func main () {
	index := topsearch.NewIndex()
	csv := csv.NewReader(os.Stdin)

	log.Printf("%#v", csv)

	names, values := topsearch.Import(csv)

	// log.Printf("%#v", names)
	// log.Printf("%#v", values)

	index.RebuildWith(names, values)

	log.Printf("%#v", index.Query("ehow"))
}
