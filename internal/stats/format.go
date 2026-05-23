package stats

import (
	"fmt"
	"sort"
	"strings"
)

// FormatSummary returns a human-readable multi-line string of the Summary.
func FormatSummary(s Summary) string {
	var sb strings.Builder

	sb.WriteString("=== Crontab Statistics ===\n")
	fmt.Fprintf(&sb, "Total entries   : %d\n", s.TotalEntries)
	fmt.Fprintf(&sb, "Unique commands : %d\n", s.UniqueCommands)
	if s.BusiestDay != "" {
		fmt.Fprintf(&sb, "Busiest day     : %s (%d entries)\n", s.BusiestDay, s.EntriesPerDay[s.BusiestDay])
	}
	fmt.Fprintf(&sb, "Busiest hour    : %02d:00 (%d entries)\n", s.BusiestHour, s.EntriesPerHour[s.BusiestHour])

	sb.WriteString("\nEntries per day:\n")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	for _, d := range days {
		if count, ok := s.EntriesPerDay[d]; ok {
			fmt.Fprintf(&sb, "  %-10s : %d\n", d, count)
		}
	}

	sb.WriteString("\nEntries per hour:\n")
	hours := make([]int, 0, len(s.EntriesPerHour))
	for h := range s.EntriesPerHour {
		hours = append(hours, h)
	}
	sort.Ints(hours)
	for _, h := range hours {
		fmt.Fprintf(&sb, "  %02d:00 : %d\n", h, s.EntriesPerHour[h])
	}

	return sb.String()
}
