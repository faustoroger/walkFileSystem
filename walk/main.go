package main

import (
	"flag"
)

type config struct {
	//  extension to filter out
	ext string
	// min file size
	size int64
	// list files
	list bool
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	// Action options
	list := flag.Bool("list", false, "List files only")
	// Filter options
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()
}
