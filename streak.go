package main

import (
	"sort"
	"time"
)

type StreakStats struct {
	Current int
	Longest int
	Sum     int
}

func calculateStreaks(data *Data) StreakStats {
	if len(data.Records) == 0 {
		return StreakStats{}
	}

	dates := getUniqueDates(data)
	if len(dates) == 0 {
		return StreakStats{}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(dates)))

	current := calculateCurrentStreak(dates)
	longest := calculateLongestStreak(dates)
	sum := len(dates)

	return StreakStats{
		Current: current,
		Longest: longest,
		Sum:     sum,
	}
}

func getUniqueDates(data *Data) []string {
	dateSet := make(map[string]bool)
	for _, r := range data.Records {
		dateSet[r.Date] = true
	}

	dates := make([]string, 0, len(dateSet))
	for d := range dateSet {
		dates = append(dates, d)
	}
	return dates
}

func calculateCurrentStreak(dates []string) int {
	if len(dates) == 0 {
		return 0
	}

	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	if dates[0] != today && dates[0] != yesterday {
		return 0
	}

	streak := 1
	for i := 1; i < len(dates); i++ {
		prevDate, _ := time.Parse("2006-01-02", dates[i-1])
		currDate, _ := time.Parse("2006-01-02", dates[i])

		diff := prevDate.Sub(currDate).Hours() / 24
		if diff == 1 {
			streak++
		} else {
			break
		}
	}

	return streak
}

func calculateLongestStreak(dates []string) int {
	if len(dates) == 0 {
		return 0
	}

	sort.Strings(dates)

	longest := 1
	current := 1

	for i := 1; i < len(dates); i++ {
		prevDate, _ := time.Parse("2006-01-02", dates[i-1])
		currDate, _ := time.Parse("2006-01-02", dates[i])

		diff := currDate.Sub(prevDate).Hours() / 24
		if diff == 1 {
			current++
			if current > longest {
				longest = current
			}
		} else {
			current = 1
		}
	}

	return longest
}
