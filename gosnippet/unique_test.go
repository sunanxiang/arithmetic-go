package gosnippet

import (
	"testing"
	"errors"
)

func TestUniqueSliceString(t *testing.T) {
	var strs []string
	except := []string{"test", "mock"}
	for i := 0; i < 10; i++ {
		strs = append(strs, "test")
	}
	for i := 0; i < 10; i++ {
		strs = append(strs, "mock")
	}

	rlt := UniqueSliceString(strs)
	if len(rlt) != 2 {
		t.Error(errors.New("UniqueSliceString error"))
	}
	for index, item := range rlt {
		if item != except[index] {
			t.Error(errors.New("UniqueSliceString error"))
		}
	}
}
