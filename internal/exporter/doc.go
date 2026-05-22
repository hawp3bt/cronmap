// Package exporter provides multiple output format serializers for a weekly
// cron schedule. Supported formats include:
//
//   - JSON  (ToJSON)  — structured JSON array of day/slot objects
//   - iCal  (ToICal)  — RFC 5545 VCALENDAR stream for calendar import
//   - CSV   (ToCSV)   — comma-separated values with header row
//
// All exporters accept a schedule.Week map and return the serialized string
// representation along with any encoding error.
package exporter
