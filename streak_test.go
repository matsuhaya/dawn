package main

import (
	"testing"
	"time"
)

func TestCalculateStreaksEmpty(t *testing.T) {
	data := &Data{Records: []Record{}}
	stats := calculateStreaks(data)

	if stats.Current != 0 {
		t.Errorf("Current = %d, want 0", stats.Current)
	}
	if stats.Longest != 0 {
		t.Errorf("Longest = %d, want 0", stats.Longest)
	}
	if stats.Sum != 0 {
		t.Errorf("Sum = %d, want 0", stats.Sum)
	}
}

func TestGetUniqueDates(t *testing.T) {
	data := &Data{
		Records: []Record{
			{Date: "2024-01-15", Time: "06:00", Routine: "読書"},
			{Date: "2024-01-15", Time: "06:30", Routine: "瞑想"},
			{Date: "2024-01-14", Time: "06:00", Routine: "ランニング"},
		},
	}

	dates := getUniqueDates(data)

	if len(dates) != 2 {
		t.Errorf("expected 2 unique dates, got %d", len(dates))
	}
}

func TestCalculateLongestStreak(t *testing.T) {
	tests := []struct {
		name     string
		dates    []string
		expected int
	}{
		{
			name:     "empty",
			dates:    []string{},
			expected: 0,
		},
		{
			name:     "single day",
			dates:    []string{"2024-01-15"},
			expected: 1,
		},
		{
			name:     "3 consecutive days",
			dates:    []string{"2024-01-13", "2024-01-14", "2024-01-15"},
			expected: 3,
		},
		{
			name:     "gap in middle",
			dates:    []string{"2024-01-10", "2024-01-12", "2024-01-13"},
			expected: 2,
		},
		{
			name:     "multiple streaks",
			dates:    []string{"2024-01-01", "2024-01-02", "2024-01-10", "2024-01-11", "2024-01-12"},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateLongestStreak(tt.dates)
			if result != tt.expected {
				t.Errorf("calculateLongestStreak() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestCalculateCurrentStreak(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	twoDaysAgo := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	threeDaysAgo := time.Now().AddDate(0, 0, -3).Format("2006-01-02")

	tests := []struct {
		name     string
		dates    []string
		expected int
	}{
		{
			name:     "empty",
			dates:    []string{},
			expected: 0,
		},
		{
			name:     "today only",
			dates:    []string{today},
			expected: 1,
		},
		{
			name:     "yesterday only",
			dates:    []string{yesterday},
			expected: 1,
		},
		{
			name:     "today and yesterday",
			dates:    []string{today, yesterday},
			expected: 2,
		},
		{
			name:     "3 days including today",
			dates:    []string{today, yesterday, twoDaysAgo},
			expected: 3,
		},
		{
			name:     "streak from yesterday",
			dates:    []string{yesterday, twoDaysAgo, threeDaysAgo},
			expected: 3,
		},
		{
			name:     "old date only - no current streak",
			dates:    []string{twoDaysAgo},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateCurrentStreak(tt.dates)
			if result != tt.expected {
				t.Errorf("calculateCurrentStreak() = %d, want %d", result, tt.expected)
			}
		})
	}
}

func TestStreakSum(t *testing.T) {
	data := &Data{
		Records: []Record{
			{Date: "2024-01-15", Time: "06:00", Routine: "読書"},
			{Date: "2024-01-14", Time: "06:00", Routine: "ランニング"},
			{Date: "2024-01-10", Time: "06:00", Routine: "英語"},
		},
	}

	stats := calculateStreaks(data)

	if stats.Sum != 3 {
		t.Errorf("Sum = %d, want 3", stats.Sum)
	}
}
