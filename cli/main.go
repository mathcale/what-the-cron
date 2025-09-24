package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: what-the-cron <cron-expression>")
		os.Exit(1)
	}

	expr := strings.Join(os.Args[1:], " ")

	human, err := humanizeCron(expr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	next, err := nextRun(expr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(human)
	fmt.Printf("Next execution: %s\n", next.Format("2006-01-02 15:04:05"))
}

func nextRun(expr string) (time.Time, error) {
	schedule, err := cron.ParseStandard(expr)
	if err != nil {
		return time.Time{}, err
	}

	return schedule.Next(time.Now()), nil
}

func humanizeCron(expr string) (string, error) {
	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return "", fmt.Errorf("invalid cron expression: expected 5 fields")
	}

	minute := describeField("minute", fields[0])
	hour := describeField("hour", fields[1])
	day := describeField("day-of-month", fields[2])
	month := describeField("month", fields[3])
	weekday := describeWeekday(fields[4])

	var timePart string
	if minute != "" && hour != "" {
		timePart = "At " + minute + " past " + hour
	} else if minute != "" {
		timePart = "At " + minute
	} else if hour != "" {
		timePart = "At " + hour
	}

	location := []string{}
	if day != "" {
		location = append(location, day)
	}

	if month != "" {
		location = append(location, month)
	}

	if weekday != "" {
		location = append(location, weekday)
	}

	result := timePart

	if len(location) > 0 {
		result += " " + strings.Join(location, " ")
	}

	return result, nil
}

func describeField(name, expr string) string {
	if expr == "*" {
		return ""
	}

	if strings.HasPrefix(expr, "*/") {
		n, err := strconv.Atoi(expr[2:])
		if err != nil {
			return ""
		}
		return fmt.Sprintf("every %d%s %s", n, ordinal(n), name)
	}

	if strings.Contains(expr, "-") {
		parts := strings.Split(expr, "-")
		if len(parts) == 2 {
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			return fmt.Sprintf("%s from %d through %d", name, a, b)
		}
	}

	if n, err := strconv.Atoi(expr); err == nil {
		return fmt.Sprintf("%s %d", name, n)
	}

	return ""
}

func describeWeekday(expr string) string {
	if expr == "*" {
		return ""
	}

	if strings.HasPrefix(expr, "*/") {
		n, err := strconv.Atoi(expr[2:])
		if err != nil {
			return ""
		}

		return fmt.Sprintf("every %d day-of-week", n)
	}

	if strings.Contains(expr, "-") {
		parts := strings.Split(expr, "-")

		if len(parts) == 2 {
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			dayA := weekdayName(a)
			dayB := weekdayName(b)
			return fmt.Sprintf("on every day-of-week from %s through %s", dayA, dayB)
		}
	}

	if n, err := strconv.Atoi(expr); err == nil {
		return fmt.Sprintf("on %s", weekdayName(n))
	}

	return ""
}

func weekdayName(n int) string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if n >= 0 && n <= 6 {
		return days[n]
	}

	return ""
}

func ordinal(n int) string {
	if n%10 == 1 && n%100 != 11 {
		return "st"
	}

	if n%10 == 2 && n%100 != 12 {
		return "nd"
	}

	if n%10 == 3 && n%100 != 13 {
		return "rd"
	}

	return "th"
}
