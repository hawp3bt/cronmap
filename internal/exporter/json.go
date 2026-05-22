package exporter

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/example/cronmap/internal/schedule"
)

// SlotJSON is the JSON-serializable representation of a single schedule slot.
type SlotJSON struct {
	Day     string `json:"day"`
	Hour    int    `json:"hour"`
	Minute  int    `json:"minute"`
	Command string `json:"command"`
}

// WeeklyJSON is the top-level JSON export structure.
type WeeklyJSON struct {
	Schedule []SlotJSON `json:"schedule"`
}

// ToJSON serialises the weekly schedule to JSON and writes it to w.
func ToJSON(w io.Writer, week map[string][]schedule.Slot) error {
	out := WeeklyJSON{}

	dayOrder := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for _, day := range dayOrder {
		slots, ok := week[day]
		if !ok {
			continue
		}
		for _, s := range slots {
			out.Schedule = append(out.Schedule, SlotJSON{
				Day:     day,
				Hour:    s.Hour,
				Minute:  s.Minute,
				Command: s.Command,
			})
		}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(out); err != nil {
		return fmt.Errorf("exporter: json encode failed: %w", err)
	}
	return nil
}
