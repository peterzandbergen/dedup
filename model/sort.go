// sort.go contains the sort methods for the FileInfoList.
// We want to be able to sort on Name and on Size/Name.
package model

import (
	"sort"
	"strings"
)

// Types to wrap FileInfoList in for different sort functions.
type sizeName FileInfoList
type nameSize FileInfoList
type sizeHash FileInfoList

// Check for complete interface implementation.
var (
	_ sort.Interface = sizeName{}
	_ sort.Interface = nameSize{}
	_ sort.Interface = sizeHash{}
)

// SortSizeName returns a FileInfoList sorted on Size and Name. Elements point
// to the elements in the fil parameter.
func SortSizeName(fil FileInfoList) FileInfoList {
	// Copy the list.
	res := make(FileInfoList, len(fil))
	copy(res, fil)

	// Sort the list.
	s := sizeName(res)
	sort.Sort(s)
	return FileInfoList(s)
}

// Len for sizeName
func (s sizeName) Len() int {
	return len(s)
}

func (s sizeName) Less(i, j int) bool {
	fi := s[i]
	fj := s[j]

	if fi.Size < fj.Size {
		return true
	} else if fi.Size > fj.Size {
		return false
	}

	return strings.ToLower(fi.Name) < strings.ToLower(fj.Name)
}

func (s sizeName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortNameSize(fil FileInfoList) FileInfoList {
	// Copy the list.
	res := make(FileInfoList, len(fil))
	copy(res, fil)

	// Sort the list.
	s := nameSize(res)
	sort.Sort(s)
	return FileInfoList(s)
}

// Sorting on name and size.
func (s nameSize) Len() int {
	return len(s)
}

func (s nameSize) Less(i, j int) bool {
	fi := s[i]
	fj := s[j]
	fin := strings.ToLower(fi.Name)
	fjn := strings.ToLower(fj.Name)

	if fin < fjn {
		return true
	} else if fin > fjn {
		return false
	}

	return fi.Size < fj.Size
}

func (s nameSize) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Sorting on size and hash.
func (s sizeHash) Len() int {
	return len(s)
}

func (s sizeHash) Less(i, j int) bool {
	fi := s[i]
	fj := s[j]

	if fi.Size < fj.Size {
		return true
	} else if fi.Size > fj.Size {
		return false
	}

	return strings.ToLower(fi.Hash) < strings.ToLower(fj.Hash)
}

func (s sizeHash) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func SortSizeHashes(fil FileInfoList) FileInfoList {
	// Assumes that the hash has been set.
	// Copy the list.
	res := make(FileInfoList, len(fil))
	copy(res, fil)

	// Sort the list.
	s := sizeHash(res)
	sort.Sort(s)
	return FileInfoList(s)
}
