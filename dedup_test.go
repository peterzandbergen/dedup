package dedup

import (
	"github.com/peterzandbergen/dedup/model"

	"path/filepath"
	"testing"
)

func TestScanDirectory(t *testing.T) {
	var p = `C:\Users\Peter\Documents`
	// var p = `Z:\`
	if s, err := filepath.Abs(p); err == nil {
		p = s
	}

	root := &model.ScanRoot{
		Path: p,
	}

	err := ScanTree(root)

	if err != nil {
		t.Fatalf("err != nil: %s", err.Error())
	}

	if len(root.Files) <= 0 {
		t.Fatal("res is empty")
	}

	fi := root.Files[0]

	if fi == nil {
		t.Fatal("first element in res is nil")
	}

	// t.Logf("Root %s contains %d elements", root.Files, len(root.Files))
	// t.Logf("%s", res[0].Name)
	// t.Log(res.String())

	// t.Log(root.String())
}
