package exporter

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/user/cronmap/internal/schedule"
)

// CSVHeader is the header row for CSV output.
const CSVHeader = "day,hour,minute,command"

// ToCSV converts a weekly schedule into CSV format.
// Each slot is rendered as a row with day, hour, minute, and command.
func ToCSV(week schedule.Week) (string, error) {
	var buf bytes.Buffer

	buf.WriteString(CSVHeader)
	buf.WriteByte('\n')

	days := []string{
		"Sunday", "Monday", "Tuesday", "Wednesday",
		"Thursday", "Friday", "Saturday",
	}

	for _, day := range days {
		slots, ok := week[day]
		if !ok {
			continue
		}
		for _, slot := range slots {
			line := fmt.Sprintf("%s,%d,%d,%s\n",
				day,
				slot.Hour,
				slot.Minute,
				escapeCSV(slot.Command),
			)
			buf.WriteString(line)
		}
	}

	return buf.String(), nil
}

// escapeCSV wraps a field in double quotes if it contains a comma, quote, or newline.
func escapeCSV(s string) string {
	if strings.ContainsAny(s, ",\"\n") {
		s = strings.ReplaceAll(s, "\"", "\"\"")
		return "\"" + s + "\""
	}
	return s
}
