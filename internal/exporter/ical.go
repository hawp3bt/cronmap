package exporter

import (
	"fmt"
	"strings"
	"time"

	"github.com/user/cronmap/internal/schedule"
)

const icalHeader = `BEGIN:VCALENDAR
VERSION:2.0
PRODID:-//cronmap//EN
CALSCALE:GREGORIAN
`

const icalFooter = `END:VCALENDAR
`

// ToICal converts a weekly schedule into an iCalendar (RFC 5545) string.
// Each slot is emitted as a VEVENT anchored to the next occurrence of that
// weekday starting from the provided reference Monday.
func ToICal(week schedule.Week, refMonday time.Time) (string, error) {
	if week == nil {
		return "", fmt.Errorf("ical: nil week provided")
	}

	var sb strings.Builder
	sb.WriteString(icalHeader)

	for dayIndex, slots := range week {
		day := refMonday.AddDate(0, 0, dayIndex)
		for _, slot := range slots {
			if slot == nil {
				continue
			}
			for _, hour := range slot.Hours {
				start := time.Date(
					day.Year(), day.Month(), day.Day(),
					hour, 0, 0, 0, time.UTC,
				)
				end := start.Add(time.Minute)

				sb.WriteString("BEGIN:VEVENT\n")
				sb.WriteString(fmt.Sprintf("UID:%s-%d-%d@cronmap\n", slot.Command, dayIndex, hour))
				sb.WriteString(fmt.Sprintf("DTSTART:%s\n", start.Format("20060102T150405Z")))
				sb.WriteString(fmt.Sprintf("DTEND:%s\n", end.Format("20060102T150405Z")))
				sb.WriteString(fmt.Sprintf("SUMMARY:%s\n", escapeICal(slot.Command)))
				sb.WriteString("END:VEVENT\n")
			}
		}
	}

	sb.WriteString(icalFooter)
	return sb.String(), nil
}

// escapeICal performs minimal escaping required by RFC 5545.
func escapeICal(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, ";", `\;`)
	s = strings.ReplaceAll(s, ",", `\,`)
	s = strings.ReplaceAll(s, "\n", `\n`)
	return s
}
