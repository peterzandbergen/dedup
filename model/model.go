// model.go
package model

import (
	"bytes"
	"fmt"
	"time"
)

// ScanRoot is the root of a scan.
type ScanRoot struct {
	Path     string
	ScanDate time.Time
	Files    FileInfoList
}

// String returns a printable presentation of the scanroot, including the files.
func (sr *ScanRoot) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Path: %s\nFiles:\n%s", sr.Path, sr.Files.String())
	return b.String()
}

// FileInfo contains the
type FileInfo struct {
	Parent *ScanRoot `json:"-"`

	Name    string
	Size    int64
	ModTime time.Time
}

func (fi *FileInfo) String() string {
	return fi.Name
}

type FileInfoList []*FileInfo

func (l FileInfoList) String() string {
	var buf bytes.Buffer

	if l == nil {
		return ""
	}

	for _, fi := range l {
		var f string
		if buf.Len() > 0 {
			f = "\n%s"
		} else {
			f = "%s"
		}
		fmt.Fprintf(&buf, f, fi.String())
	}
	return buf.String()
}
