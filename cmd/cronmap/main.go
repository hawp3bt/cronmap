package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/user/cronmap/internal/formatter"
	"github.com/user/cronmap/internal/parser"
	"github.com/user/cronmap/internal/schedule"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: cronmap <crontab-file>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	var entries []*parser.Entry
	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		entry, err := parser.Parse(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "line %d skipped: %v\n", lineNum, err)
			continue
		}
		if entry != nil {
			entries = append(entries, entry)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		os.Exit(1)
	}

	weekly := schedule.Build(entries)
	slots := formatter.HumanReadable(weekly)
	if len(slots) == 0 {
		fmt.Println("No scheduled tasks found.")
		return
	}
	fmt.Println(formatter.FormatAll(slots))
}
