package filter_test

import (
	"testing"

	"github.com/user/cronmap/internal/filter"
	"github.com/user/cronmap/internal/parser"
)

func makeEntry(cmd string, days, hours []int) *parser.Entry {
	return &parser.Entry{
		Command:   cmd,
		DayOfWeek: days,
		Hour:      hours,
		Minute:    []int{0},
	}
}

func TestByDay_Match(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry("backup", []int{0, 6}, []int{2}),
		makeEntry("report", []int{1, 2, 3}, []int{9}),
	}
	got := filter.ByDay(entries, 1)
	if len(got) != 1 || got[0].Command != "report" {
		t.Fatalf("expected [report], got %v", got)
	}
}

func TestByDay_NoMatch(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry("task", []int{3}, []int{10}),
	}
	got := filter.ByDay(entries, 5)
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}

func TestByDay_NilSkipped(t *testing.T) {
	entries := []*parser.Entry{nil, makeEntry("job", []int{2}, []int{8})}
	got := filter.ByDay(entries, 2)
	if len(got) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(got))
	}
}

func TestByHour_Match(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry("midnight", []int{0}, []int{0}),
		makeEntry("noon", []int{0}, []int{12}),
	}
	got := filter.ByHour(entries, 12)
	if len(got) != 1 || got[0].Command != "noon" {
		t.Fatalf("expected [noon], got %v", got)
	}
}

func TestByCommand_CaseInsensitive(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry("/usr/bin/Backup.sh", []int{0}, []int{1}),
		makeEntry("/usr/bin/report.sh", []int{0}, []int{2}),
	}
	got := filter.ByCommand(entries, "backup")
	if len(got) != 1 || got[0].Command != "/usr/bin/Backup.sh" {
		t.Fatalf("expected [Backup.sh], got %v", got)
	}
}

func TestByCommand_EmptySubstr(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry("a", []int{0}, []int{0}),
		makeEntry("b", []int{1}, []int{1}),
	}
	got := filter.ByCommand(entries, "")
	if len(got) != 2 {
		t.Fatalf("expected all entries, got %d", len(got))
	}
}

func TestApply_CombinedFilters(t *testing.T) {
	entries := []*parser.Entry{
		makeEntry("/scripts/backup.sh", []int{0}, []int{2}),
		makeEntry("/scripts/backup.sh", []int{1}, []int{2}),
		makeEntry("/scripts/cleanup.sh", []int{0}, []int{2}),
	}
	opts := filter.Options{DayOfWeek: 0, Hour: 2, Command: "backup"}
	got := filter.Apply(entries, opts)
	if len(got) != 1 || got[0].Command != "/scripts/backup.sh" || got[0].DayOfWeek[0] != 0 {
		t.Fatalf("unexpected result: %v", got)
	}
}

func TestDefaultOptions(t *testing.T) {
	opts := filter.DefaultOptions()
	if opts.DayOfWeek != -1 || opts.Hour != -1 || opts.Command != "" {
		t.Fatalf("unexpected defaults: %+v", opts)
	}
}
