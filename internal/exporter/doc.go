// Package exporter provides multiple output format writers for a cronmap
// weekly schedule. Supported formats include:
//
//   - JSON  (ToJSON)
//   - iCal  (ToICal)
//   - CSV   (ToCSV)
//   - Markdown (ToMarkdown)
//   - Plain text (ToText)
//   - HTML  (ToHTML)
//
// All exporters accept a schedule.Week value and return a formatted string.
package exporter
