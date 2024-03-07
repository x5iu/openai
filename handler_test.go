package openai

import (
	"bytes"
	"strings"
	"testing"
)

func TestSetVal(t *testing.T) {
	var (
		dstInt   = int(0402)
		dstBytes = []byte("test")
	)
	if err := setVal(&dstInt, int8(99)); err != nil {
		t.Errorf("setVal: %s", err)
		return
	} else if dstInt != 99 {
		t.Errorf("setVal: %v != %v", dstInt, 99)
		return
	}
	if err := setVal(&dstBytes, "0402"); err != nil {
		t.Errorf("setVal: %s", err)
		return
	} else if !bytes.Equal(dstBytes, []byte("0402")) {
		t.Errorf("setVal: %s != %s", string(dstBytes), "0402")
		return
	}
	if err := setVal(dstInt, 0); err == nil {
		t.Errorf("setVal: expects errors, got nil")
		return
	} else if !strings.Contains(err.Error(), "is not Pointer") {
		t.Errorf("setVal: expects NotPointerError, got => %s", err)
		return
	}
	if err := setVal(&dstInt, "test"); err == nil {
		t.Errorf("setVal: expects errors, got nil")
		return
	} else if !strings.Contains(err.Error(), "is not assignable to") {
		t.Errorf("setVal: expects NotAssignableToError, got => %s", err)
		return
	}
}
