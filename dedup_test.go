package dedup

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanDirectory(t *testing.T) {
	var wf = func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() {
			t.Logf("%s%s%#v", path, "===================", info)

		}
		return nil
	}

	filepath.Walk(".", wf)
}
