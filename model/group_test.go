package model

import (
	"testing"
)

var testList = FileInfoList{
	&FileInfo{
		Name: "aaa",
		Size: 200,
	},
	&FileInfo{
		Name: "bbb",
		Size: 200,
	},
}

func TestGroup(t *testing.T) {
	sn := SortSizeName(testList)

	gn := EqualSize(testList)

	t.Logf("%#v", gn)

	t.Log(sn)
}
