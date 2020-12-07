package main

import (
	"testing"
)

func TestFindValidPasswordsSchemeACount(t *testing.T) {
	got := FindValidPasswordsSchemeACount([]PasswordPolicyPassword{
		{1, 3, "a", "abcde"},
		{1, 3, "b", "cdefg"},
		{2, 9, "c", "ccccccccc"},
	})
	expected := 2

	if got != expected {
		t.Error("Fail")
	}
}

func TestFindValidPasswordsSchemeBCount(t *testing.T) {
	got := FindValidPasswordsSchemeBCount([]PasswordPolicyPassword{
		{1, 3, "a", "abcde"},
		{1, 3, "b", "cdefg"},
		{2, 9, "c", "ccccccccc"},
	})
	expected := 1

	if got != expected {
		t.Error("Fail")
	}
}
