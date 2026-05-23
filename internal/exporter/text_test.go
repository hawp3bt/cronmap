package exporter

import (
	"strings"
	"testing"

	"github.com/yourorg/cronmap/internal/schedule"
)

func makeWeekText() *schedule.Week {
	return &schedule.Week{
		Days: map[string][]schedule.Slot{
			"Monday": {
				{Hour: 2, Minute: 0, Command: "/usr/bin/backup"},
				{Hour: 6, Minute: 30, Command: "/usr/bin/report"},
			},
			"Wednesday": {
				{Hour: 12, Minute: 15, Command: "/usr/bin/sync"},
			},
		},
	}
}

func TestToText_ContainsHeader(t *testing.T) {
	out := ToText(makeWeekText())
	if !strings.Contains(out, "Day") {
		t.Error("expected header to contain 'Day'")
	}
	if !strings.Contains(out, "Jobs") {
		t.Error("expected header to contain 'Jobs'")
	}
}

func TestToText_ContainsDayName(t *testing.T) {
	out := ToText(makeWeekText())
	if !strings.Contains(out, "Monday") {
		t.Error("expected output to contain 'Monday'")
	}
	if !strings.Contains(out, "Wednesday") {
		t.Error("expected output to contain 'Wednesday'")
	}
}

func TestToText_ContainsCommand(t *testing.T) {
	out := ToText(makeWeekText())
	if !strings.Contains(out, "/usr/bin/backup") {
		t.Error("expected output to contain '/usr/bin/backup'")
	}
}

func TestToText_ContainsTime(t *testing.T) {
	out := ToText(makeWeekText())
	if !strings.Contains(out, "02:00") {
		t.Error("expected output to contain '02:00'")
	}
	if !strings.Contains(out, "06:30") {
		t.Error("expected output to contain '06:30'")
	}
}

func TestToText_NilWeek(t *testing.T) {
	out := ToText(nil)
	if out != "" {
		t.Errorf("expected empty string for nil week, got %q", out)
	}
}

func TestToText_BorderChars(t *testing.T) {
	out := ToText(makeWeekText())
	for _, ch := range []string{"┌", "┐", "└", "┘", "─", "│", "┬", "┴", "├", "┤", "┼"} {
		if !strings.Contains(out, ch) {
			t.Errorf("expected border character %q in output", ch)
		}
	}
}
