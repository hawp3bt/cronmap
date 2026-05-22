package exporter_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/example/cronmap/internal/exporter"
	"github.com/example/cronmap/internal/schedule"
)

func makeWeek() map[string][]schedule.Slot {
	return map[string][]schedule.Slot{
		"Monday": {
			{Hour: 9, Minute: 0, Command: "/bin/backup"},
			{Hour: 17, Minute: 30, Command: "/bin/report"},
		},
		"Friday": {
			{Hour: 12, Minute: 15, Command: "/bin/cleanup"},
		},
	}
}

func TestToJSON_ValidOutput(t *testing.T) {
	var buf bytes.Buffer
	if err := exporter.ToJSON(&buf, makeWeek()); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if buf.Len() == 0 {
		t.Fatal("expected non-empty output")
	}
}

func TestToJSON_WellFormedJSON(t *testing.T) {
	var buf bytes.Buffer
	_ = exporter.ToJSON(&buf, makeWeek())

	var out exporter.WeeklyJSON
	if err := json.Unmarshal(buf.Bytes(), &out); err != nil {
		t.Fatalf("output is not valid JSON: %v", err)
	}
}

func TestToJSON_SlotCount(t *testing.T) {
	var buf bytes.Buffer
	_ = exporter.ToJSON(&buf, makeWeek())

	var out exporter.WeeklyJSON
	_ = json.Unmarshal(buf.Bytes(), &out)

	if len(out.Schedule) != 3 {
		t.Fatalf("expected 3 slots, got %d", len(out.Schedule))
	}
}

func TestToJSON_DayOrder(t *testing.T) {
	var buf bytes.Buffer
	_ = exporter.ToJSON(&buf, makeWeek())

	var out exporter.WeeklyJSON
	_ = json.Unmarshal(buf.Bytes(), &out)

	// Monday comes before Friday in day order
	if out.Schedule[0].Day != "Monday" {
		t.Errorf("expected first slot day Monday, got %s", out.Schedule[0].Day)
	}
	if out.Schedule[2].Day != "Friday" {
		t.Errorf("expected last slot day Friday, got %s", out.Schedule[2].Day)
	}
}

func TestToJSON_EmptySchedule(t *testing.T) {
	var buf bytes.Buffer
	if err := exporter.ToJSON(&buf, map[string][]schedule.Slot{}); err != nil {
		t.Fatalf("unexpected error on empty schedule: %v", err)
	}
	var out exporter.WeeklyJSON
	_ = json.Unmarshal(buf.Bytes(), &out)
	if len(out.Schedule) != 0 {
		t.Errorf("expected 0 slots for empty input, got %d", len(out.Schedule))
	}
}
