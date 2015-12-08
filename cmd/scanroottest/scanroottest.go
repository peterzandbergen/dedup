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

func TestOne() {
	// var p = `C:\Users\Peter\Dropbox (Moven)`
	var p = `C:\Users\Peter\Documents`
	if s, err := filepath.Abs(p); err == nil {
		p = s
	}

	root := &model.ScanRoot{
		Path: p,
	}

	st := time.Now()
	err := dedup.ScanTree(root, -1)
	et := time.Now()

	d := et.Sub(st)
	fmt.Printf("Scan took %v to run.\n", d.String())

	if err != nil {
		fmt.Printf("Error during scan: %s\n", err.Error())
		return
	}
	fmt.Printf("Found %d file under %s.\n", len(root.Files), root.Path)

	// fmt.Print(root.String())

	sorted := model.SortSizeName(root.Files)

	fmt.Print(sorted.String())
}

var testList = model.FileInfoList{
	&model.FileInfo{
		Name: "aaa",
		Size: 200,
	},
	&model.FileInfo{
		Name: "bbb",
		Size: 200,
	},
	&model.FileInfo{
		Name: "ccc",
		Size: 200,
	},
	&model.FileInfo{
		Name: "ddd",
		Size: 300,
	},
	&model.FileInfo{
		Name: "eee",
		Size: 300,
	},
	&model.FileInfo{
		Name: "fff",
		Size: 200,
	},
}

func TestTwo() {
	// sn := model.SortSizeName(testList)

	gn := model.EqualSizes(testList)

	fmt.Printf("%s\n", model.EqualSizeList(gn).String())
}

// TestThree scans the tree starting in r and produces a list with the
// files that have the same length.
func TestThree(scanRoot string, largerThan int64) {
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

	fmt.Print("done.\nLooking for doubles...")
	st = time.Now()
	gn := model.EqualSizes(root.Files)
	et = time.Now()
	d = et.Sub(st)
	fmt.Printf("done\nLooking for doubles took %v to run.\n", d.String())
	fmt.Printf("\n%s\n", model.EqualSizeList(gn).String())
}

func TestFour(scanRoot string, largerThan int64) {
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
	HashEquals(gn)

	// Again search for equal files HashEquals.
	eh := model.EqualHashes(root.Files)
	fmt.Printf("Equal Hashes: \n%s\n", model.EqualSizeList(eh).String())
}

func HashEquals(equals []model.EqualSize) {
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
	TestFour(*scanRoot, *largerThan)
}
