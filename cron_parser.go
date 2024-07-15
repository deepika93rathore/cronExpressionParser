package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func expandField(field string, min int, max int) []int {
	var result []int
	if field == "*" {
		for i := min; i <= max; i++ {
			result = append(result, i)
		}
	} else if strings.Contains(field, ",") {
		parts := strings.Split(field, ",")
		for _, part := range parts {
			result = append(result, expandField(part, min, max)...)
		}
	} else if strings.Contains(field, "-") {
		parts := strings.Split(field, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
	} else if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		stepField := parts[0]
		step, _ := strconv.Atoi(parts[1])
		var start int
		if stepField == "*" {
			start = min
		} else {
			start, _ = strconv.Atoi(stepField)
		}
		for i := start; i <= max; i += step {
			result = append(result, i)
		}
	} else {
		value, _ := strconv.Atoi(field)
		result = append(result, value)
	}
	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cron_parser \"<cron_expression>\"")
		return
	}

	cronExpression := os.Args[1]
	fields := strings.Fields(cronExpression)

	if len(fields) != 6 {
		fmt.Println("Error: Invalid cron expression")
		return
	}

	minuteField := fields[0]
	hourField := fields[1]
	dayOfMonthField := fields[2]
	monthField := fields[3]
	dayOfWeekField := fields[4]
	command := fields[5]

	minutes := expandField(minuteField, 0, 59)
	hours := expandField(hourField, 0, 23)
	daysOfMonth := expandField(dayOfMonthField, 1, 31)
	months := expandField(monthField, 1, 12)
	daysOfWeek := expandField(dayOfWeekField, 0, 6)

	fmt.Printf("%-14s %s\n", "minute", joinInts(minutes))
	fmt.Printf("%-14s %s\n", "hour", joinInts(hours))
	fmt.Printf("%-14s %s\n", "day of month", joinInts(daysOfMonth))
	fmt.Printf("%-14s %s\n", "month", joinInts(months))
	fmt.Printf("%-14s %s\n", "day of week", joinInts(daysOfWeek))
	fmt.Printf("%-14s %s\n", "command", command)
}

func joinInts(ints []int) string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(v)
	}
	return strings.Join(strs, " ")
}
