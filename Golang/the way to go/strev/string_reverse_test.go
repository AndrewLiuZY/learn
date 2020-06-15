package strev

import "testing"

func TestReverse(t *testing.T) {
	if Reverse("abcde") != "edcba" {
		t.Log("Must be edcba")
		t.Fail()
	}
}
