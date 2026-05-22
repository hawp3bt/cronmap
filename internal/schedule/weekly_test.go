package schedule

import (
	"strings"
	"testing"

	"github.com/user/cronmap/internal/parser"
)

func makeEntry(minute, hour, dom, month, dow, cmd string) *parser.CronEntry {
	return &parser.CronEntry{
		Minute:     minute,
		Hour:       hour,
		DayOfMonth: dom,
		Month:      month,
		DayOfWeek:  dow,
		Command:    cmd,
	}
}

func TestBuild_SingleEntry(t *testing.T) {
	entries := []*parser.CronEntry{
		makeEntry("0", "9", "*", "*", "1", "/bin/morning.sh"),
	}
	ws, err := Build(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	slots, ok := ws[1][9]
	if !ok || len(slots) == 0 {
		t.Fatal("expected slot on Monday at 09:00")
	}
	if slots[0].Command != "/bin/morning.sh" {
		t.Errorf("unexpected command: %q", slots[0].Command)
	}
}

func TestBuild_EveryDay(t *testing.T) {
	entries := []*parser.CronEntry{
		makeEntry("30", "12", "*", "*", "*", "/bin/lunch.sh"),
	}
	ws, err := Build(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ws) != 7 {
		t.Errorf("expected 7 days, got %d", len(ws))
	}
}

func TestBuild_NilEntrySkipped(t *testing.T) {
	entries := []*parser.CronEntry{nil, makeEntry("0", "8", "*", "*", "5", "/bin/friday.sh")}
	ws, err := Build(entries)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := ws[5]; !ok {
		t.Error("expected Friday entry")
	}
}

func TestRender_ContainsDayName(t *testing.T) {
	entries := []*parser.CronEntry{
		makeEntry("0", "10", "*", "*", "3", "/bin/wednesday.sh"),
	}
	ws, _ := Build(entries)
	output := Render(ws)
	if !strings.Contains(output, "Wednesday") {
		t.Error("expected 'Wednesday' in rendered output")
	}
	if !strings.Contains(output, "/bin/wednesday.sh") {
		t.Error("expected command in rendered output")
	}
}
