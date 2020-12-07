package main

import (
	"testing"
)

func TestFindValidPasswordsCount(t *testing.T) {
	got := FindValidPasswordsCount([]PasswordPolicyPassword{
		{1, 3, "a", "abcde"},
		{1, 3, "b", "cdefg"},
		{2, 9, "c", "ccccccccc"},
	})
	expected := 2

	if got != expected {
		t.Error("Fail")
	}
}
