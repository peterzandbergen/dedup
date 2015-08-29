package dedup

import (
	"github.com/peterzandbergen/dedup/model"

	"os"
	"path/filepath"
	"time"
)

func ScanTree(root *model.ScanRoot) error {
	// Define the walk func.
	// Appends the found files to the files list.
	var w = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.Mode().IsRegular() {
			fi := &model.FileInfo{
				Name:    path,
				Size:    info.Size(),
				ModTime: info.ModTime(),
				Parent:  root,
			}
			root.Files = append(root.Files, fi)
		}
		return nil
	}

	root.ScanDate = time.Now()
	root.Files = nil
	filepath.Walk(root.Path, w)
	return nil
}
