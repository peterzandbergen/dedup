package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	// "hash"
	"path/filepath"
	"time"

	"github.com/peterzandbergen/dedup"
	"github.com/peterzandbergen/dedup/hasher"
	"github.com/peterzandbergen/dedup/model"
)


// FindIdenticals searches the root and all subfolders for files with the same hash.
func FindIdenticals(scanRoot string, largerThan int64) {
	// Scan the root.
	var p = scanRoot
	if s, err := filepath.Abs(p); err == nil {
		p = s
	}

	root := &model.ScanRoot{
		Path: p,
	}

	var st, et time.Time
	var d time.Duration

	fmt.Printf("Scanning %s...", root.Path)
	st = time.Now()
	dedup.ScanTree(root, largerThan)
	et = time.Now()
	d = et.Sub(st)
	fmt.Printf("done\nScan took %v to run.\n", d.String())

	// Sort and search for equal sized files.
	fmt.Print("done.\nLooking for doubles...")
	st = time.Now()
	gn := model.EqualSizes(root.Files)
	et = time.Now()
	d = et.Sub(st)
	fmt.Printf("done\nLooking for doubles took %v to run.\n", d.String())

	// Generate the hashes for the equal sized files.
	hashEquals(gn)

	// Again search for equal files hashEquals.
	eh := model.EqualHashes(root.Files)
	fmt.Printf("Equal Hashes: \n%s\n", model.EqualSizeList(eh).String())
}

func hashEquals(equals []model.EqualSize) {
	h := sha1.New()
	for _, es := range equals {
		for _, fs := range es.Files {
			if id, err := hasher.HashFile(h, fs.Name); err == nil {
				fs.Hash = id
			}
		}
	}
}

const (
	meg = 1000000
)

var (
	scanRoot   = flag.String("root", ".", "Root to scan.")
	largerThan = flag.Int64("lt", 1*meg, "Include file larger than this value.")
)

func main() {
	flag.Parse()
	FindIdenticals(*scanRoot, *largerThan)
}
