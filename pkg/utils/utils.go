package utils

import "fmt"

// Format seconds to a string in the following formats '1d1h'/'1m1s'
//
//nolint:gocritic,nestif
func AgeFormatter(seconds uint64) string {
	age := ""
	days := seconds / 86400
	hours := (seconds / 3600) % 24
	minutes := (seconds / 60) % 60
	seconds = seconds % 60

	if days > 0 && hours > 0 {
		age = fmt.Sprintf("%dd%dh", days, hours)
	} else if days > 0 {
		age = fmt.Sprintf("%dd", days)
	} else if hours > 0 && minutes > 0 {
		age = fmt.Sprintf("%dh%dm", hours, minutes)
	} else if hours > 0 {
		age = fmt.Sprintf("%dh", hours)
	} else if minutes > 0 && seconds > 0 {
		age = fmt.Sprintf("%dm%ds", minutes, seconds)
	} else if minutes > 0 {
		age = fmt.Sprintf("%dm", minutes)
	} else {
		age = fmt.Sprintf("%ds", seconds)
	}

	return age
}
