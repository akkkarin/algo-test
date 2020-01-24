package interview

import "testing"

func TestLinkedListStepReverse(t *testing.T) {
	l := gen(10)
	//l.Print()

	l.Reverse()
	l.Print()
}
