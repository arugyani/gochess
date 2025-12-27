package util

import "testing"

func TestCollapse(t *testing.T) {
	got := Collapse([]uint64{0x00000000000001, 0x0000000000010, 0x0000000000100})
	want := uint64(0x00000000000111)

	if got != want {
		t.Errorf("got %064b, wanted %064b", got, want)
	}
}
