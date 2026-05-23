package stats_test

import (
	"testing"

	"github.com/example/cronmap/internal/parser"
	"github.com/example/cronmap/internal/stats"
)

func makeEntry(days, hours []int, cmd string) *parser.Entry {
	return &parser.Entry{
		DaysOfWeek: days,
		Hours:      hours,
		Minutes:    []int{0},
		Command:    cmd,
	}
}

func TestCompute_TotalEntries(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry([]int{1}, []int{9}, "cmd1"),
		makeEntry([]int{2}, []int{10}, "cmd2"),
	}
	s := stats.Compute(entries)
	if s.TotalEntries != 2 {
		t.Errorf("expected 2 total entries, got %d", s.TotalEntries)
	}
}

func TestCompute_UniqueCommands(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry([]int{1}, []int{9}, "backup"),
		makeEntry([]int{2}, []int{10}, "backup"),
		makeEntry([]int{3}, []int{11}, "cleanup"),
	}
	s := stats.Compute(entries)
	if s.UniqueCommands != 2 {
		t.Errorf("expected 2 unique commands, got %d", s.UniqueCommands)
	}
}

func TestCompute_BusiestDay(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry([]int{1}, []int{9}, "a"),
		makeEntry([]int{1}, []int{10}, "b"),
		makeEntry([]int{3}, []int{11}, "c"),
	}
	s := stats.Compute(entries)
	if s.BusiestDay != "Monday" {
		t.Errorf("expected Monday as busiest day, got %s", s.BusiestDay)
	}
}

func TestCompute_BusiestHour(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry([]int{1}, []int{8}, "a"),
		makeEntry([]int{2}, []int{8}, "b"),
		makeEntry([]int{3}, []int{12}, "c"),
	}
	s := stats.Compute(entries)
	if s.BusiestHour != 8 {
		t.Errorf("expected hour 8 as busiest, got %d", s.BusiestHour)
	}
}

func TestCompute_NilSkipped(t *testing.T) {
	entries := []*parser.Entry{nil, makeEntry([]int{0}, []int{6}, "cmd")}
	s := stats.Compute(entries)
	if s.TotalEntries != 2 {
		t.Errorf("expected 2 total (including nil slot), got %d", s.TotalEntries)
	}
	if s.UniqueCommands != 1 {
		t.Errorf("expected 1 unique command, got %d", s.UniqueCommands)
	}
}

func TestCompute_EmptyEntries(t *testing.T) {
	s := stats.Compute(nil)
	if s.TotalEntries != 0 || s.UniqueCommands != 0 {
		t.Error("expected zero stats for empty input")
	}
}
