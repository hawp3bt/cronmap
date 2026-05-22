package exporter

import (
	"strings"
	"testing"

	"github.com/user/cronmap/internal/schedule"
)

func makeWeekCSV() schedule.Week {
	return schedule.Week{
		"Monday": []schedule.Slot{
			{Hour: 9, Minute: 0, Command: "/usr/bin/backup"},
			{Hour: 17, Minute: 30, Command: "/usr/bin/report"},
		},
		"Friday": []schedule.Slot{
			{Hour: 6, Minute: 15, Command: "/usr/bin/cleanup"},
		},
	}
}

func TestToCSV_ContainsHeader(t *testing.T) {
	week := makeWeekCSV()
	out, err := ToCSV(week)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(out, CSVHeader) {
		t.Errorf("expected output to start with header %q, got: %q", CSVHeader, out[:min(len(out), 40)])
	}
}

func TestToCSV_RowCount(t *testing.T) {
	week := makeWeekCSV()
	out, err := ToCSV(week)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	// 1 header + 3 data rows
	if len(lines) != 4 {
		t.Errorf("expected 4 lines (1 header + 3 data), got %d", len(lines))
	}
}

func TestToCSV_ContainsCommand(t *testing.T) {
	week := makeWeekCSV()
	out, err := ToCSV(week)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "/usr/bin/backup") {
		t.Errorf("expected output to contain command /usr/bin/backup")
	}
}

func TestToCSV_EscapesCommaInCommand(t *testing.T) {
	week := schedule.Week{
		"Monday": []schedule.Slot{
			{Hour: 8, Minute: 0, Command: "echo hello,world"},
		},
	}
	out, err := ToCSV(week)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `"echo hello,world"`) {
		t.Errorf("expected quoted command in output, got:\n%s", out)
	}
}

func TestToCSV_EmptyWeek(t *testing.T) {
	out, err := ToCSV(schedule.Week{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 1 || lines[0] != CSVHeader {
		t.Errorf("expected only header for empty week, got: %q", out)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
