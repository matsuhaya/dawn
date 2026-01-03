package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type Record struct {
	Date    string `json:"date"`
	Time    string `json:"time"`
	Routine string `json:"routine"`
}

type Data struct {
	Records []Record `json:"records"`
}

func getDataPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "dawn", "data.json"), nil
}

func ensureDataDir() error {
	path, err := getDataPath()
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	return os.MkdirAll(dir, 0755)
}

func loadData() (*Data, error) {
	path, err := getDataPath()
	if err != nil {
		return nil, err
	}

	data := &Data{Records: []Record{}}

	file, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return data, nil
	}
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, data); err != nil {
		return nil, err
	}

	return data, nil
}

func saveData(data *Data) error {
	if err := ensureDataDir(); err != nil {
		return err
	}

	path, err := getDataPath()
	if err != nil {
		return err
	}

	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, file, 0644)
}

type TimeStatus int

const (
	TimeOK TimeStatus = iota
	TimeTooEarly
	TimeTooLate
)

func checkTime(t time.Time) TimeStatus {
	hour := t.Hour()
	if hour < 5 {
		return TimeTooEarly
	}
	if hour >= 8 {
		return TimeTooLate
	}
	return TimeOK
}

func hasRecordToday(data *Data, today string) bool {
	for _, r := range data.Records {
		if r.Date == today {
			return true
		}
	}
	return false
}

func addRecord(data *Data, routine string, t time.Time) {
	record := Record{
		Date:    t.Format("2006-01-02"),
		Time:    t.Format("15:04"),
		Routine: routine,
	}
	data.Records = append(data.Records, record)

	sort.Slice(data.Records, func(i, j int) bool {
		if data.Records[i].Date != data.Records[j].Date {
			return data.Records[i].Date > data.Records[j].Date
		}
		return data.Records[i].Time > data.Records[j].Time
	})
}
