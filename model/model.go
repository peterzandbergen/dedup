// model.go
package model

import (
	_ "os"
)

type FileInfo struct {
	Parent *FileInfo
	Path   string
	Size   int64
}
