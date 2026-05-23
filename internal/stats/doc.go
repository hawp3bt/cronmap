// Package stats computes and formats aggregate statistics from a collection
// of parsed cron entries. It provides insight into schedule density, the
// busiest days and hours, and the number of unique commands scheduled.
//
// Typical usage:
//
//	entries, _ := parser.Parse(input)
//	summary := stats.Compute(entries)
//	fmt.Println(stats.FormatSummary(summary))
package stats
