package cron

import (
	"fmt"
	"strconv"
	"strings"
)

type CronHumanizer struct{}

func NewCronHumanizer() *CronHumanizer {
	return &CronHumanizer{}
}

func (h *CronHumanizer) Humanize(expr string) (string, error) {
	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return "", fmt.Errorf("invalid cron expression: expected 5 fields")
	}

	minute := h.describeField("minute", fields[0])
	hour := h.describeField("hour", fields[1])
	day := h.describeField("day-of-month", fields[2])
	month := h.describeField("month", fields[3])
	weekday := h.describeWeekday(fields[4])

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

func (h *CronHumanizer) describeField(name, expr string) string {
	if expr == "*" {
		return ""
	}

	if strings.HasPrefix(expr, "*/") {
		n, err := strconv.Atoi(expr[2:])
		if err != nil {
			return ""
		}

		return fmt.Sprintf("every %d%s %s", n, h.ordinal(n), name)
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

func (h *CronHumanizer) describeWeekday(expr string) string {
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
			dayA := h.weekdayName(a)
			dayB := h.weekdayName(b)
			return fmt.Sprintf("on every day-of-week from %s through %s", dayA, dayB)
		}
	}

	if n, err := strconv.Atoi(expr); err == nil {
		return fmt.Sprintf("on %s", h.weekdayName(n))
	}

	return ""
}

func (h *CronHumanizer) weekdayName(n int) string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	if n >= 0 && n <= 6 {
		return days[n]
	}

	return ""
}

func (h *CronHumanizer) ordinal(n int) string {
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
