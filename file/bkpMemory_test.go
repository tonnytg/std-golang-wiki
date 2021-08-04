package bkpmemory_test

import (
	bkpmemory "github.com/tonnytg/std-golang-wiki/file"
	"testing"
)

func TestBkpmemory(t *testing.T) {
	file := "test_bkp.txt"
	// in message set \n before close for best experience"
	msg := "Hello\n"
	err := bkpmemory.BkpMemory(file, msg)
	if err != nil {
		t.Errorf("can't bkp memory %v", err)
	}
}