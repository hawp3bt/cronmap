package exporter

import (
	"strings"
	"testing"

	"github.com/example/cronmap/internal/schedule"
)

func makeWeekHTML() schedule.Week {
	return schedule.Week{
		"Monday": {
			{Hour: 9, Minute: 0, Command: "/usr/bin/backup"},
			{Hour: 17, Minute: 30, Command: "/usr/bin/report"},
		},
		"Wednesday": {
			{Hour: 3, Minute: 15, Command: "echo hello <world>"},
		},
	}
}

func TestToHTML_ContainsDoctype(t *testing.T) {
	out := ToHTML(makeWeekHTML())
	if !strings.Contains(out, "<!DOCTYPE html>") {
		t.Error("expected HTML doctype declaration")
	}
}

func TestToHTML_ContainsTableHeader(t *testing.T) {
	out := ToHTML(makeWeekHTML())
	for _, col := range []string{"Day", "Time", "Command"} {
		if !strings.Contains(out, "<th>"+col+"</th>") {
			t.Errorf("expected table header %q", col)
		}
	}
}

func TestToHTML_ContainsDayName(t *testing.T) {
	out := ToHTML(makeWeekHTML())
	if !strings.Contains(out, "Monday") {
		t.Error("expected day name Monday in output")
	}
}

func TestToHTML_ContainsCommand(t *testing.T) {
	out := ToHTML(makeWeekHTML())
	if !strings.Contains(out, "/usr/bin/backup") {
		t.Error("expected command /usr/bin/backup in output")
	}
}

func TestToHTML_ContainsTime(t *testing.T) {
	out := ToHTML(makeWeekHTML())
	if !strings.Contains(out, "09:00") {
		t.Error("expected time 09:00 in output")
	}
	if !strings.Contains(out, "17:30") {
		t.Error("expected time 17:30 in output")
	}
}

func TestToHTML_EscapesHTMLInCommand(t *testing.T) {
	out := ToHTML(makeWeekHTML())
	// Raw '<' should be escaped to &lt; in HTML output
	if strings.Contains(out, "echo hello <world>") {
		t.Error("expected HTML special chars to be escaped in command")
	}
	if !strings.Contains(out, "&lt;world&gt;") {
		t.Error("expected &lt;world&gt; in escaped output")
	}
}

func TestToHTML_EmptyWeek(t *testing.T) {
	out := ToHTML(schedule.Week{})
	if !strings.Contains(out, "<table>") {
		t.Error("expected table element even for empty week")
	}
	if strings.Contains(out, "<td>") {
		t.Error("expected no data rows for empty week")
	}
}
