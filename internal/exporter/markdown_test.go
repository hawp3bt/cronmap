package exporter

import (
	"strings"
	"testing"

	"github.com/user/cronmap/internal/schedule"
)

func makeWeekMD() schedule.Week {
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

func TestToMarkdown_ContainsHeader(t *testing.T) {
	out := ToMarkdown(makeWeekMD())
	if !strings.Contains(out, "# Weekly Cron Schedule") {
		t.Error("expected top-level header in markdown output")
	}
}

func TestToMarkdown_ContainsDaySection(t *testing.T) {
	out := ToMarkdown(makeWeekMD())
	if !strings.Contains(out, "## Monday") {
		t.Error("expected Monday section header")
	}
	if !strings.Contains(out, "## Friday") {
		t.Error("expected Friday section header")
	}
}

func TestToMarkdown_ContainsCommand(t *testing.T) {
	out := ToMarkdown(makeWeekMD())
	if !strings.Contains(out, "/usr/bin/backup") {
		t.Error("expected command in markdown output")
	}
}

func TestToMarkdown_ContainsTime(t *testing.T) {
	out := ToMarkdown(makeWeekMD())
	if !strings.Contains(out, "09:00") {
		t.Error("expected formatted time 09:00 in output")
	}
	if !strings.Contains(out, "17:30") {
		t.Error("expected formatted time 17:30 in output")
	}
}

func TestToMarkdown_TableSeparator(t *testing.T) {
	out := ToMarkdown(makeWeekMD())
	if !strings.Contains(out, "|-------|---------|")	{
		t.Error("expected markdown table separator row")
	}
}

func TestToMarkdown_EscapesPipe(t *testing.T) {
	week := schedule.Week{
		"Monday": []schedule.Slot{
			{Hour: 1, Minute: 0, Command: "echo foo | bar"},
		},
	}
	out := ToMarkdown(week)
	if !strings.Contains(out, `echo foo \| bar`) {
		t.Errorf("expected pipe to be escaped in markdown, got:\n%s", out)
	}
}

func TestToMarkdown_EmptyWeek(t *testing.T) {
	out := ToMarkdown(schedule.Week{})
	if !strings.Contains(out, "# Weekly Cron Schedule") {
		t.Error("expected header even for empty week")
	}
	if strings.Contains(out, "##") {
		t.Error("expected no day sections for empty week")
	}
}
