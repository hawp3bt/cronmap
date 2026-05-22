package formatter_test

import (
	"strings"
	"testing"

	"github.com/user/cronmap/internal/formatter"
	"github.com/user/cronmap/internal/parser"
	"github.com/user/cronmap/internal/schedule"
)

func makeSchedule(t *testing.T, line string) schedule.WeeklySchedule {
	t.Helper()
	entry, err := parser.Parse(line)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	return schedule.Build([]*parser.Entry{entry})
}

func TestHumanReadable_ReturnsSlots(t *testing.T) {
	// every day at 09:30
	sched := makeSchedule(t, "30 9 * * * /usr/bin/backup")
	slots := formatter.HumanReadable(sched)
	if len(slots) != 7 {
		t.Errorf("expected 7 slots (one per day), got %d", len(slots))
	}
}

func TestHumanReadable_CorrectTime(t *testing.T) {
	sched := makeSchedule(t, "15 14 * * 1 /bin/task")
	slots := formatter.HumanReadable(sched)
	if len(slots) != 1 {
		t.Fatalf("expected 1 slot, got %d", len(slots))
	}
	if slots[0].Hour != 14 || slots[0].Minute != 15 {
		t.Errorf("expected 14:15, got %02d:%02d", slots[0].Hour, slots[0].Minute)
	}
	if slots[0].Command != "/bin/task" {
		t.Errorf("unexpected command: %s", slots[0].Command)
	}
}

func TestFormatSlot_Output(t *testing.T) {
	ts := formatter.TimeSlot{Day: "Monday", Hour: 8, Minute: 5, Command: "/bin/run"}
	out := formatter.FormatSlot(ts)
	if !strings.Contains(out, "Monday") {
		t.Errorf("expected day in output, got: %s", out)
	}
	if !strings.Contains(out, "08:05") {
		t.Errorf("expected time 08:05 in output, got: %s", out)
	}
	if !strings.Contains(out, "/bin/run") {
		t.Errorf("expected command in output, got: %s", out)
	}
}

func TestFormatAll_Newlines(t *testing.T) {
	slots := []formatter.TimeSlot{
		{Day: "Monday", Hour: 1, Minute: 0, Command: "a"},
		{Day: "Tuesday", Hour: 2, Minute: 30, Command: "b"},
	}
	out := formatter.FormatAll(slots)
	lines := strings.Split(out, "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 lines, got %d", len(lines))
	}
}
