package exporter

import (
	"fmt"
	"strings"

	"github.com/user/cronmap/internal/schedule"
)

// ToMarkdown renders a weekly schedule as a Markdown table.
// Each day becomes a section header with a table of time slots and commands.
func ToMarkdown(week schedule.Week) string {
	var sb strings.Builder

	sb.WriteString("# Weekly Cron Schedule\n\n")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for _, day := range days {
		slots, ok := week[day]
		if !ok || len(slots) == 0 {
			continue
		}

		sb.WriteString(fmt.Sprintf("## %s\n\n", day))
		sb.WriteString("| Time  | Command |\n")
		sb.WriteString("|-------|---------|\n")

		for _, slot := range slots {
			time := fmt.Sprintf("%02d:%02d", slot.Hour, slot.Minute)
			command := escapeMarkdown(slot.Command)
			sb.WriteString(fmt.Sprintf("| %s | %s |\n", time, command))
		}

		sb.WriteString("\n")
	}

	return sb.String()
}

// escapeMarkdown escapes pipe characters in Markdown table cells.
func escapeMarkdown(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}
