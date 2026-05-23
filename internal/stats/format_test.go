package stats_test

import (
	"strings"
	"testing"

	"github.com/example/cronmap/internal/parser"
	"github.com/example/cronmap/internal/stats"
)

func buildSummary() stats.Summary {
	entries := []*parser.Entry{
		makeEntry([]int{1, 2}, []int{9, 10}, "backup"),
		makeEntry([]int{1}, []int{9}, "cleanup"),
		makeEntry([]int{5}, []int{18}, "report"),
	}
	return stats.Compute(entries)
}

func TestFormatSummary_ContainsHeader(t *testing.T) {
	out := stats.FormatSummary(buildSummary())
	if !strings.Contains(out, "Crontab Statistics") {
		t.Error("expected header in output")
	}
}

func TestFormatSummary_ContainsTotalEntries(t *testing.T) {
	out := stats.FormatSummary(buildSummary())
	if !strings.Contains(out, "Total entries") {
		t.Error("expected total entries line")
	}
}

func TestFormatSummary_ContainsBusiestDay(t *testing.T) {
	out := stats.FormatSummary(buildSummary())
	if !strings.Contains(out, "Monday") {
		t.Error("expected Monday in busiest day output")
	}
}

func TestFormatSummary_ContainsHourSection(t *testing.T) {
	out := stats.FormatSummary(buildSummary())
	if !strings.Contains(out, "Entries per hour") {
		t.Error("expected hour section in output")
	}
}

func TestFormatSummary_ContainsDaySection(t *testing.T) {
	out := stats.FormatSummary(buildSummary())
	if !strings.Contains(out, "Entries per day") {
		t.Error("expected day section in output")
	}
}
