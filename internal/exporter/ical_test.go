package exporter

import (
	"strings"
	"testing"
	"time"

	"github.com/user/cronmap/internal/schedule"
)

func refMonday() time.Time {
	// 2024-01-01 is a Monday
	return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
}

func makeWeekICal() schedule.Week {
	w := make(schedule.Week, 7)
	w[0] = []*schedule.Slot{
		{Command: "backup.sh", Hours: []int{2, 14}},
	}
	w[3] = []*schedule.Slot{
		{Command: "report.sh", Hours: []int{9}},
	}
	return w
}

func TestToICal_ContainsHeader(t *testing.T) {
	out, err := ToICal(makeWeekICal(), refMonday())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "BEGIN:VCALENDAR") {
		t.Error("expected BEGIN:VCALENDAR in output")
	}
	if !strings.Contains(out, "END:VCALENDAR") {
		t.Error("expected END:VCALENDAR in output")
	}
}

func TestToICal_EventCount(t *testing.T) {
	out, err := ToICal(makeWeekICal(), refMonday())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	count := strings.Count(out, "BEGIN:VEVENT")
	// 2 hours for Monday + 1 hour for Thursday = 3 events
	if count != 3 {
		t.Errorf("expected 3 VEVENT blocks, got %d", count)
	}
}

func TestToICal_ContainsCommand(t *testing.T) {
	out, err := ToICal(makeWeekICal(), refMonday())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "SUMMARY:backup.sh") {
		t.Error("expected SUMMARY:backup.sh in output")
	}
}

func TestToICal_NilWeek(t *testing.T) {
	_, err := ToICal(nil, refMonday())
	if err == nil {
		t.Error("expected error for nil week, got nil")
	}
}

func TestToICal_EscapesSpecialChars(t *testing.T) {
	w := make(schedule.Week, 7)
	w[0] = []*schedule.Slot{
		{Command: "cmd;with,special", Hours: []int{6}},
	}
	out, err := ToICal(w, refMonday())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if strings.Contains(out, "cmd;with,special") {
		t.Error("expected special characters to be escaped")
	}
	if !strings.Contains(out, `cmd\;with\,special`) {
		t.Error("expected escaped version in output")
	}
}
