package schedule

import (
	"fmt"
	"strings"

	"github.com/user/cronmap/internal/parser"
)

var dayNames = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

// Slot represents a scheduled job at a specific day and hour.
type Slot struct {
	Day     int // 0=Sunday ... 6=Saturday
	Hour    int
	Minutes []int
	Command string
}

// WeeklySchedule maps day -> hour -> list of Slots.
type WeeklySchedule map[int]map[int][]Slot

// Build converts a slice of CronEntries into a WeeklySchedule.
func Build(entries []*parser.CronEntry) (WeeklySchedule, error) {
	schedule := make(WeeklySchedule)

	for _, entry := range entries {
		if entry == nil {
			continue
		}

		days, err := parser.ExpandField(entry.DayOfWeek, 0, 6)
		if err != nil {
			return nil, fmt.Errorf("expanding day-of-week: %w", err)
		}
		hours, err := parser.ExpandField(entry.Hour, 0, 23)
		if err != nil {
			return nil, fmt.Errorf("expanding hour: %w", err)
		}
		minutes, err := parser.ExpandField(entry.Minute, 0, 59)
		if err != nil {
			return nil, fmt.Errorf("expanding minute: %w", err)
		}

		for _, d := range days {
			if schedule[d] == nil {
				schedule[d] = make(map[int][]Slot)
			}
			for _, h := range hours {
				schedule[d][h] = append(schedule[d][h], Slot{
					Day:     d,
					Hour:    h,
					Minutes: minutes,
					Command: entry.Command,
				})
			}
		}
	}

	return schedule, nil
}

// Render produces a human-readable text representation of the weekly schedule.
func Render(ws WeeklySchedule) string {
	var sb strings.Builder

	for day := 0; day <= 6; day++ {
		hours, ok := ws[day]
		if !ok {
			continue
		}
		fmt.Fprintf(&sb, "=== %s ===\n", dayNames[day])
		for hour := 0; hour <= 23; hour++ {
			slots, ok := hours[hour]
			if !ok {
				continue
			}
			for _, slot := range slots {
				minStrs := make([]string, len(slot.Minutes))
				for i, m := range slot.Minutes {
					minStrs[i] = fmt.Sprintf("%02d", m)
				}
				fmt.Fprintf(&sb, "  %02d:[%s]  %s\n", hour, strings.Join(minStrs, ","), slot.Command)
			}
		}
	}

	return sb.String()
}
