package model

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	fi := &FileInfo{}
	_ = fi
}

func TestFileInfoString(t *testing.T) {
	var s = "[         100] testname"
	fi := &FileInfo{
		Name: "testname",
		Size: 100,
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

func TestMarshalToJson(t *testing.T) {
	var s = "testname"

	root := &ScanRoot{
		Path: "scanRoot",
	}

	fi := &FileInfo{
		Parent:  root,
		Name:    s,
		ModTime: time.Now(),
	}

	root.Files = append(root.Files, fi, fi)

	b, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		t.Errorf("Error marshalling scanroot: %s", err.Error())
	}
	_ = b
	// t.Logf("\n%s", string(b))
}

func TestMarshalToJsonFile(t *testing.T) {
	var s = "testname"

	root := &ScanRoot{
		Path: "scanRoot",
	}

	fi := &FileInfo{
		Parent:  root,
		Name:    s,
		ModTime: time.Now(),
	}

	root.Files = append(root.Files, fi, fi)

	b, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		t.Errorf("Error marshalling scanroot: %s", err.Error())
	}
	// t.Logf("\n%s", string(b))

	if f, err := os.Create("scanModel.json.txt"); err == nil {
		defer f.Close()
		if _, err := f.Write(b); err != nil {
			t.Errorf("Error writing to file: %s", err.Error())
		}
	} else {
		t.Errorf("Error creating scanmodel.json.txt: %s", err.Error())
	}
}
