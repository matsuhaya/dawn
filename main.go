package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		recordRoutine("")
		return
	}

	switch args[0] {
	case "log":
		showLog()
	case "history":
		showHistory()
	case "edit":
		editData()
	case "help":
		printHelp()
	default:
		routine := strings.Join(args, " ")
		recordRoutine(routine)
	}
}

func recordRoutine(routine string) {
	now := time.Now()

	status := checkTime(now)
	switch status {
	case TimeTooEarly:
		fmt.Println(MsgTooEarly)
		return
	case TimeTooLate:
		fmt.Println(MsgTooLate)
		return
	}

	data, err := loadData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	today := now.Format("2006-01-02")
	if hasRecordToday(data, today) {
		fmt.Println(MsgAlreadyDone)
		return
	}

	if routine == "" {
		routine = promptRoutine()
		if routine == "" {
			return
		}
	}

	addRecord(data, routine, now)

	if err := saveData(data); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(MsgSuccess)
}

func promptRoutine() string {
	fmt.Print(MsgPrompt)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(input)
}

func showLog() {
	data, err := loadData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	stats := calculateStreaks(data)
	printStreakStats(stats)
}

func showHistory() {
	data, err := loadData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	printHistory(data)
}

func editData() {
	path, err := getDataPath()
	if err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	if err := ensureDataDir(); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		data := &Data{Records: []Record{}}
		if err := saveData(data); err != nil {
			fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
			os.Exit(1)
		}
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	cmd := exec.Command("sh", "-c", editor+" "+path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "エラー: %v\n", err)
		os.Exit(1)
	}
}
