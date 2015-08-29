package model

import (
	"testing"
)

func TestOne(t *testing.T) {
	fi := &FileInfo{}
	_ = fi
}

func TestFileInfoString(t *testing.T) {
	var s = "testname"
	fi := &FileInfo{
		Name: s,
	}
	if s != fi.String() {
		t.Errorf("Exptected %s, got %s", s, fi.String())
	}
}

func TestFileInfoListString(t *testing.T) {
	var s = "testname"
	fi := &FileInfo{
		Name: s,
	}

	fil := FileInfoList{
		fi,
	}
	if len(fil) != 1 {
		t.Errorf("Exptected %d, got %d", 1, len(fil))
	}
}

func TestFileInfoListNilString(t *testing.T) {
	var s = ""
	var fi FileInfoList

	s = fi.String()

	if s != "" {
		t.Errorf("Exptected %s, got %s", "", s)
	}
}

func TestFileInfoListAppendString(t *testing.T) {
	var s = "testname"
	fi := &FileInfo{
		Name: s,
	}

	var fil FileInfoList
	fil = append(fil, fi)

	if len(fil) != 1 {
		t.Errorf("Exptected %d, got %d", 1, len(fil))
	}
}
