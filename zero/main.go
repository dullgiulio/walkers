package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	// XXX: Until Go 1.5
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type file struct {
	os.FileInfo
	path string
}

func makeFile(f os.FileInfo, path string) file {
	return file{f, path}
}

func main() {
	flag.Parse()
	root := filepath.Clean(filepath.ToSlash(flag.Arg(0)))
	sink := make(chan file)

	go func() {
		filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			sink <- makeFile(info, path)
			return nil
		})
		close(sink)
	}()

	for f := range sink {
		fmt.Printf("%s\n", f)
	}
}
