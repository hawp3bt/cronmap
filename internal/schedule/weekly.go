package schedule

import (
	"fmt"
	"sort"
	"strings"

	"github.com/example/cronmap/internal/parser"
)

// Slot represents a single scheduled event within a day.
type Slot struct {
	Hour    int
	Minute  int
	Command string
}

// dayNames maps cron weekday numbers (0=Sunday) to names.
var dayNames = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

// Build converts a slice of parsed cron entries into a weekly map.
func Build(entries []*parser.Entry) map[string][]Slot {
	week := make(map[string][]Slot)

	for _, e := range entries {
		if e == nil {
			continue
		}
		for _, dow := range e.DaysOfWeek {
			if dow < 0 || dow > 6 {
				continue
			}
			day := dayNames[dow]
			for _, hour := range e.Hours {
				for _, min := range e.Minutes {
					week[day] = append(week[day], Slot{
						Hour:    hour,
						Minute:  min,
						Command: e.Command,
					})
				}
			}
		}
	}

	for day := range week {
		sort.Slice(week[day], func(i, j int) bool {
			a, b := week[day][i], week[day][j]
			if a.Hour != b.Hour {
				return a.Hour < b.Hour
			}
			return a.Minute < b.Minute
		})
	}

	return week
}

// Render produces a plain-text weekly schedule view.
func Render(week map[string][]Slot) string {
	var sb strings.Builder
	for _, day := range dayNames {
		slots, ok := week[day]
		if !ok {
			continue
		}
		sb.WriteString(day + ":\n")
		for _, s := range slots {
			sb.WriteString(fmt.Sprintf("  %02d:%02d  %s\n", s.Hour, s.Minute, s.Command))
		}
	}
	return sb.String()
}
