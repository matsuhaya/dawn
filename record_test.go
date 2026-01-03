package main

import (
	"testing"
	"time"
)

func TestCheckTime(t *testing.T) {
	tests := []struct {
		name     string
		hour     int
		minute   int
		expected TimeStatus
	}{
		{"4:59 - too early", 4, 59, TimeTooEarly},
		{"5:00 - OK", 5, 0, TimeOK},
		{"6:30 - OK", 6, 30, TimeOK},
		{"7:59 - OK", 7, 59, TimeOK},
		{"8:00 - too late", 8, 0, TimeTooLate},
		{"12:00 - too late", 12, 0, TimeTooLate},
		{"0:00 - too early", 0, 0, TimeTooEarly},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTime := time.Date(2024, 1, 15, tt.hour, tt.minute, 0, 0, time.Local)
			result := checkTime(testTime)
			if result != tt.expected {
				t.Errorf("checkTime(%02d:%02d) = %v, want %v", tt.hour, tt.minute, result, tt.expected)
			}
		})
	}
}

func TestHasRecordToday(t *testing.T) {
	data := &Data{
		Records: []Record{
			{Date: "2024-01-15", Time: "06:30", Routine: "読書"},
			{Date: "2024-01-14", Time: "05:45", Routine: "ランニング"},
		},
	}

	tests := []struct {
		name     string
		date     string
		expected bool
	}{
		{"record exists", "2024-01-15", true},
		{"record exists 2", "2024-01-14", true},
		{"no record", "2024-01-13", false},
		{"no record future", "2024-01-16", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasRecordToday(data, tt.date)
			if result != tt.expected {
				t.Errorf("hasRecordToday(%s) = %v, want %v", tt.date, result, tt.expected)
			}
		})
	}
}

func TestAddRecord(t *testing.T) {
	data := &Data{Records: []Record{}}
	testTime := time.Date(2024, 1, 15, 6, 30, 0, 0, time.Local)

	addRecord(data, "読書", testTime)

	if len(data.Records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(data.Records))
	}

	r := data.Records[0]
	if r.Date != "2024-01-15" {
		t.Errorf("Date = %s, want 2024-01-15", r.Date)
	}
	if r.Time != "06:30" {
		t.Errorf("Time = %s, want 06:30", r.Time)
	}
	if r.Routine != "読書" {
		t.Errorf("Routine = %s, want 読書", r.Routine)
	}
}

func TestAddRecordSortsDescending(t *testing.T) {
	data := &Data{Records: []Record{}}

	addRecord(data, "1日目", time.Date(2024, 1, 13, 6, 0, 0, 0, time.Local))
	addRecord(data, "3日目", time.Date(2024, 1, 15, 6, 0, 0, 0, time.Local))
	addRecord(data, "2日目", time.Date(2024, 1, 14, 6, 0, 0, 0, time.Local))

	if len(data.Records) != 3 {
		t.Fatalf("expected 3 records, got %d", len(data.Records))
	}

	expectedDates := []string{"2024-01-15", "2024-01-14", "2024-01-13"}
	for i, expected := range expectedDates {
		if data.Records[i].Date != expected {
			t.Errorf("Records[%d].Date = %s, want %s", i, data.Records[i].Date, expected)
		}
	}
}
