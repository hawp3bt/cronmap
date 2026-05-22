package parser

import (
	"testing"
)

func TestParse_ValidEntry(t *testing.T) {
	entry, err := Parse("*/5 * * * * /usr/bin/backup.sh")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entry == nil {
		t.Fatal("expected non-nil entry")
	}
	if entry.Minute != "*/5" {
		t.Errorf("expected minute '*/5', got %q", entry.Minute)
	}
	if entry.Command != "/usr/bin/backup.sh" {
		t.Errorf("expected command '/usr/bin/backup.sh', got %q", entry.Command)
	}
}

func TestParse_Comment(t *testing.T) {
	entry, err := Parse("# this is a comment")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entry != nil {
		t.Error("expected nil entry for comment line")
	}
}

func TestParse_EmptyLine(t *testing.T) {
	entry, err := Parse("   ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entry != nil {
		t.Error("expected nil entry for empty line")
	}
}

func TestParse_TooFewFields(t *testing.T) {
	_, err := Parse("* * * *")
	if err == nil {
		t.Error("expected error for too few fields")
	}
}

func TestExpandField_Wildcard(t *testing.T) {
	vals, err := ExpandField("*", 0, 6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(vals) != 7 {
		t.Errorf("expected 7 values, got %d", len(vals))
	}
}

func TestExpandField_Step(t *testing.T) {
	vals, err := ExpandField("*/15", 0, 59)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []int{0, 15, 30, 45}
	if len(vals) != len(expected) {
		t.Fatalf("expected %v, got %v", expected, vals)
	}
	for i, v := range expected {
		if vals[i] != v {
			t.Errorf("index %d: expected %d, got %d", i, v, vals[i])
		}
	}
}

func TestExpandField_Range(t *testing.T) {
	vals, err := ExpandField("1-3", 0, 6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(vals) != 3 || vals[0] != 1 || vals[2] != 3 {
		t.Errorf("unexpected range result: %v", vals)
	}
}

func TestExpandField_List(t *testing.T) {
	vals, err := ExpandField("1,3,5", 0, 6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(vals) != 3 || vals[1] != 3 {
		t.Errorf("unexpected list result: %v", vals)
	}
}
