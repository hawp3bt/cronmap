package exporter

import (
	"fmt"
	"strings"

	"github.com/yourorg/cronmap/internal/schedule"
)

// ToText renders a weekly schedule as a plain-text table suitable for
// terminal output or piping into other tools.
//
// Example output:
//
//	┌─────────────┬───────────────────────────────────┐
//	│ Day         │ Jobs                              │
//	├─────────────┼───────────────────────────────────┤
//	│ Monday      │ 02:00  /usr/bin/backup             │
//	│             │ 06:30  /usr/bin/report             │
//	└─────────────┴───────────────────────────────────┘
func ToText(week *schedule.Week) string {
	if week == nil {
		return ""
	}

	const dayWidth = 11
	const jobWidth = 35

	hr := func(left, mid, right, fill string) string {
		return left + strings.Repeat(fill, dayWidth+2) + mid + strings.Repeat(fill, jobWidth+2) + right + "\n"
	}

	var sb strings.Builder

	sb.WriteString(hr("┌", "┬", "┐", "─"))
	sb.WriteString(fmt.Sprintf("│ %-*s │ %-*s │\n", dayWidth, "Day", jobWidth, "Jobs"))
	sb.WriteString(hr("├", "┼", "┤", "─"))

	for _, day := range schedule.DayNames {
		slots, ok := week.Days[day]
		if !ok || len(slots) == 0 {
			continue
		}

		for i, slot := range slots {
			dayLabel := ""
			if i == 0 {
				dayLabel = day
			}
			job := fmt.Sprintf("%02d:%02d  %s", slot.Hour, slot.Minute, slot.Command)
			if len(job) > jobWidth {
				job = job[:jobWidth-1] + "…"
			}
			sb.WriteString(fmt.Sprintf("│ %-*s │ %-*s │\n", dayWidth, dayLabel, jobWidth, job))
		}
	}

	sb.WriteString(hr("└", "┴", "┘", "─"))
	return sb.String()
}
