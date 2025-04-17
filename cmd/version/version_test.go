package main

import (
	"testing"
)

// TestIncrementVersionPatch checks that the patch part increments by 1.
func TestIncrementVersionPatch(t *testing.T) {
	got, err := incrementVersion("1.2.3", "patch")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "1.2.4"
	if got != want {
		t.Errorf("incrementVersion(\"1.2.3\", \"patch\") = %q; want %q", got, want)
	}
}

// TestIncrementVersionMinor checks that the minor part increments and patch resets to 0.
func TestIncrementVersionMinor(t *testing.T) {
	got, err := incrementVersion("1.2.3", "minor")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "1.3.0"
	if got != want {
		t.Errorf("incrementVersion(\"1.2.3\", \"minor\") = %q; want %q", got, want)
	}
}

// TestIncrementVersionMajor checks that the major part increments and minor & patch reset to 0.
func TestIncrementVersionMajor(t *testing.T) {
	got, err := incrementVersion("1.2.3", "major")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "2.0.0"
	if got != want {
		t.Errorf("incrementVersion(\"1.2.3\", \"major\") = %q; want %q", got, want)
	}
}

// TestIncrementVersionInvalidFormat ensures an error is returned for wrong segment count.
func TestIncrementVersionInvalidFormat(t *testing.T) {
	_, err := incrementVersion("1.2", "patch")
	if err == nil {
		t.Error("expected error for invalid version format, got nil")
	}
}

// TestIncrementVersionUnknownPart ensures an error is returned for an unrecognized part.
func TestIncrementVersionUnknownPart(t *testing.T) {
	_, err := incrementVersion("1.2.3", "foobar")
	if err == nil {
		t.Error("expected error for unknown part to increment, got nil")
	}
}

// TestDeterminePartMajor verifies that titles containing [major] (case-insensitive) select "major".
func TestDeterminePartMajor(t *testing.T) {
	title := "Add feature [Major]"
	part := determinePart(title)
	if part != "major" {
		t.Errorf("determinePart(%q) = %q; want \"major\"", title, part)
	}
}

// TestDeterminePartMinor verifies that titles containing [minor] (case-insensitive) select "minor".
func TestDeterminePartMinor(t *testing.T) {
	title := "Fix bug [mInOr]"
	part := determinePart(title)
	if part != "minor" {
		t.Errorf("determinePart(%q) = %q; want \"minor\"", title, part)
	}
}

// TestDeterminePartDefault verifies that titles without [major] or [minor] default to "patch".
func TestDeterminePartDefault(t *testing.T) {
	title := "Update README"
	part := determinePart(title)
	if part != "patch" {
		t.Errorf("determinePart(%q) = %q; want \"patch\"", title, part)
	}
}
