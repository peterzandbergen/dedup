package main

import (
	"fmt"
	"github.com/peterzandbergen/dedup"
	"github.com/peterzandbergen/dedup/model"

	"path/filepath"
	"time"
)

func main() {
	var p = `Z:\`
	if s, err := filepath.Abs(p); err == nil {
		p = s
	}

	root := &model.ScanRoot{
		Path: p,
	}

	st := time.Now()
	err := dedup.ScanTree(root)
	et := time.Now()

	d := et.Sub(st)
	fmt.Printf("Scan took %v to run.\n", d.String())

	if err != nil {
		fmt.Printf("Error during scan: %s\n", err.Error())
		return
	}
	fmt.Printf("Found %d file under %s.\n", len(root.Files), root.Path)

	fmt.Print(root.String())
}
