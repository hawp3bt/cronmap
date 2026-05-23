package exporter_test

import (
	"strings"
	"testing"

	"github.com/user/cronmap/internal/exporter"
	"github.com/user/cronmap/internal/schedule"
)

func makeWeekXML() map[string][]schedule.Slot {
	return map[string][]schedule.Slot{
		"Monday": {
			{Hour: 9, Minute: 0, Command: "backup.sh"},
			{Hour: 17, Minute: 30, Command: "report.sh"},
		},
		"Wednesday": {
			{Hour: 12, Minute: 15, Command: "sync,data"},
		},
	}
}

func TestToXML_ContainsXMLHeader(t *testing.T) {
	out, err := exporter.ToXML(makeWeekXML())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.HasPrefix(out, "<?xml") {
		t.Errorf("expected XML declaration, got: %q", out[:min(40, len(out))])
	}
}

func TestToXML_ContainsRootElement(t *testing.T) {
	out, err := exporter.ToXML(makeWeekXML())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "<schedule>") {
		t.Errorf("expected <schedule> root element in output")
	}
}

func TestToXML_ContainsDayName(t *testing.T) {
	out, err := exporter.ToXML(makeWeekXML())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `name="Monday"`) {
		t.Errorf("expected day name attribute in output")
	}
}

func TestToXML_ContainsCommand(t *testing.T) {
	out, err := exporter.ToXML(makeWeekXML())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "backup.sh") {
		t.Errorf("expected command 'backup.sh' in output")
	}
}

func TestToXML_SlotAttributes(t *testing.T) {
	out, err := exporter.ToXML(makeWeekXML())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out, `hour="9"`) {
		t.Errorf("expected hour attribute in slot")
	}
	if !strings.Contains(out, `minute="0"`) {
		t.Errorf("expected minute attribute in slot")
	}
}

func TestToXML_DayOrder(t *testing.T) {
	out, err := exporter.ToXML(makeWeekXML())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	monIdx := strings.Index(out, `name="Monday"`)
	wedIdx := strings.Index(out, `name="Wednesday"`)
	if monIdx == -1 || wedIdx == -1 {
		t.Fatal("expected both Monday and Wednesday in output")
	}
	if monIdx > wedIdx {
		t.Errorf("expected Monday before Wednesday in output")
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
