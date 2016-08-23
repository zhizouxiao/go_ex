package stack_test

import (
	"github.com/zhizouxiao/programming_in_go/1_chapter/stack"
	"testing"
)

func TestStack(t *testing.T) {
	count := 1
	var aStack stack.Stack
	assertTrue(t, aStack.Len() == 0, "expeacted empty Stack", count)

	aStack.Push(1)
	aStack.Push(2)
	aStack.Push("three")
	assertTrue(t, aStack.Len() == 3, "expected nonempty Stack", count)
}

func assertTrue(t *testing.T, condition bool, message string, id int) {
	if !condition {
		t.Error("# %d: %s", id, message)
	}
}
