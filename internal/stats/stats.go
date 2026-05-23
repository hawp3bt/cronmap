// Package stats provides summary statistics over a parsed crontab schedule.
package stats

import (
	"sort"

	"github.com/example/cronmap/internal/parser"
)

// Summary holds aggregate statistics for a set of cron entries.
type Summary struct {
	TotalEntries  int
	UniqueCommands int
	BusiestDay    string
	BusiestHour   int
	EntriesPerDay map[string]int
	EntriesPerHour map[int]int
}

var dayNames = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

// Compute derives a Summary from a slice of parsed cron entries.
func Compute(entries []*parser.Entry) Summary {
	perDay := make(map[string]int)
	perHour := make(map[int]int)
	commandSet := make(map[string]struct{})

	for _, e := range entries {
		if e == nil {
			continue
		}
		commandSet[e.Command] = struct{}{}
		for _, d := range e.DaysOfWeek {
			if d >= 0 && d <= 6 {
				perDay[dayNames[d]]++
			}
		}
		for _, h := range e.Hours {
			perHour[h]++
		}
	}

	return Summary{
		TotalEntries:   len(entries),
		UniqueCommands: len(commandSet),
		BusiestDay:     busiestKey(perDay),
		BusiestHour:    busiestIntKey(perHour),
		EntriesPerDay:  perDay,
		EntriesPerHour: perHour,
	}
}

func busiestKey(m map[string]int) string {
	var best string
	var max int
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if m[k] > max {
			max = m[k]
			best = k
		}
	}
	return best
}

func busiestIntKey(m map[int]int) int {
	var best, max int
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if m[k] > max {
			max = m[k]
			best = k
		}
	}
	return best
}
