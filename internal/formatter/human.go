package formatter

import (
	"fmt"
	"strings"

	"github.com/user/cronmap/internal/schedule"
)

// TimeSlot represents a single scheduled task at a specific time.
type TimeSlot struct {
	Day     string
	Hour    int
	Minute  int
	Command string
}

// HumanReadable converts a weekly schedule into a slice of human-readable time slots.
func HumanReadable(s schedule.WeeklySchedule) []TimeSlot {
	var slots []TimeSlot
	for _, day := range schedule.DayNames {
		entries, ok := s[day]
		if !ok {
			continue
		}
		for _, e := range entries {
			for _, h := range e.Hours {
				for _, m := range e.Minutes {
					slots = append(slots, TimeSlot{
						Day:     day,
						Hour:    h,
						Minute:  m,
						Command: e.Command,
					})
				}
			}
		}
	}
	return slots
}

// FormatSlot returns a human-readable string for a single TimeSlot.
func FormatSlot(ts TimeSlot) string {
	return fmt.Sprintf("%s at %02d:%02d — %s", ts.Day, ts.Hour, ts.Minute, ts.Command)
}

// FormatAll returns all time slots as a newline-joined string.
func FormatAll(slots []TimeSlot) string {
	lines := make([]string, 0, len(slots))
	for _, ts := range slots {
		lines = append(lines, FormatSlot(ts))
	}
	return strings.Join(lines, "\n")
}
