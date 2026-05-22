// Package filter provides utilities for narrowing down a slice of parsed
// cron entries by day-of-week, hour, or command substring.
//
// Typical usage:
//
//	entries, _ := parser.Parse(input)
//	opts := filter.Options{DayOfWeek: 1, Hour: -1, Command: "backup"}
//	matched := filter.Apply(entries, opts)
//
// Each filter function (ByDay, ByHour, ByCommand) can also be used
// independently for more granular control.
package filter
