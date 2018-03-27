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

	UniqueSliceString(&strs)

	if len(strs) != 2 {
		t.Error(errors.New("UniqueSliceString error"))
	}
	for index, item := range strs {
		if item != except[index] {
			t.Error(errors.New("UniqueSliceString error"))
		}
	}
}
