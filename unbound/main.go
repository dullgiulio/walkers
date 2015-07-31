package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
)

func init() {
	// XXX: Until Go 1.5
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	flag.Parse()
	root := filepath.Clean(filepath.ToSlash(flag.Arg(0)))

	// Start scanning the directory
	idx := newIndexer()
	idx.scan(root)

	sink := idx.sink()
	for f := range sink {
		fmt.Printf("%s\n", f)
	}
}
