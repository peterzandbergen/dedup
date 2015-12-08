// group.go contains the sort methods for the FileInfoList.
// Group by size, to find doubles.
package model

import (
	"bytes"
	"fmt"
	// "log"
)

// EqualSize has the size and the files with that size.
type EqualSize struct {
	Size  int64
	Hash  string
	Files FileInfoList
}

// EqualSizeList is a list of with files of equal sizes.
type EqualSizeList []EqualSize

// EqualSizes returns a slice with files that have equal sizes.
func EqualSizes(fil FileInfoList) []EqualSize {
	// Sort by size/name
	es := SortSizeName(fil)

	return equalSizes(es)
}

func (e EqualSizeList) String() string {
	var b bytes.Buffer

	for _, es := range e {
		fmt.Fprintf(&b, "%s\n", es.String())
	}
	return b.String()
}

func (e *EqualSize) String() string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "Size: %v\n%s", e.Size, e.Files.String())
	return b.String()
}

func equalSizes(fil FileInfoList) []EqualSize {
	var res = make([]EqualSize, 0)

	var size int64
	var ofSize int

	for i, fi := range fil {
		if size > 0 && size == fi.Size {
			// Same size found.
			// log.Printf("Found equal, i=%d, size=%d, name=%s\n", i, size, fi.Name)
			switch {
			case ofSize == 0:
				// Second of equal size found.
				// Add current and previous to result list.
				// Add new equal size.
				res = append(res, EqualSize{
					Size: fi.Size,
					Files: FileInfoList{
						fil[i-1],
						fil[i],
					}})
				ofSize++

			case ofSize > 0:
				// Third of equal size found.
				// Add current to the result list.
				res[len(res)-1].Files = append(res[len(res)-1].Files, fil[i])
				ofSize++

			}
		} else {
			// New size found. Save it.
			ofSize = 0
			size = fi.Size
		}
	}
	return res
}

func equalHashes(fil FileInfoList) []EqualSize {
	var res = make([]EqualSize, 0)

	var hstr string
	var ofSize int

	for i, fi := range fil {
		if len(hstr) > 0 && hstr == fi.Hash {
			// Same hash found.
			// log.Printf("Found equal, i=%d, size=%d, name=%s\n", i, size, fi.Name)
			switch {
			case ofSize == 0:
				// Second of equal size found.
				// Add current and previous to result list.
				// Add new equal size.
				res = append(res, EqualSize{
					Size: fi.Size,
					Hash: fi.Hash,
					Files: FileInfoList{
						fil[i-1],
						fil[i],
					}})
				ofSize++

			case ofSize > 0:
				// Third of equal size found.
				// Add current to the result list.
				res[len(res)-1].Files = append(res[len(res)-1].Files, fil[i])
				ofSize++

			}
		} else {
			// New size found. Save it.
			ofSize = 0
			hstr = fi.Hash
		}
	}
	return res
}

// EqualHash returns a slice with files that have an equal size and hash.
func EqualHashes(fil FileInfoList) []EqualSize {
	es := SortSizeHashes(fil)

	return equalHashes(es)
}
