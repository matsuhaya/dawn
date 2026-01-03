package main

import "fmt"

const (
	MsgSuccess     = "🌅 また一日、生き延びたな"
	MsgTooEarly    = "🚬 夜型の人間には用はない"
	MsgTooLate     = "🚬 朝は待ってくれない"
	MsgAlreadyDone = "🚬 欲張るな、1日1回だ"
	MsgPrompt      = "何をする？: "
)

func printStreakStats(stats StreakStats) {
	fmt.Printf("🔥 Current Streak:  %d days\n", stats.Current)
	fmt.Printf("🏆 Longest Streak:  %d days\n", stats.Longest)
	fmt.Printf("📊 Streak Sum:      %d days\n", stats.Sum)
}

func printHistory(data *Data) {
	if len(data.Records) == 0 {
		fmt.Println("記録がありません")
		return
	}

	for _, r := range data.Records {
		fmt.Printf("%s %s  %s\n", r.Date, r.Time, r.Routine)
	}
}

func printHelp() {
	help := `Dawn - ルーチントラッカー

使い方:
  dawn [ルーチン内容]    ルーチンを記録する（5:00〜7:59のみ）
  dawn log              継続状況を表示する
  dawn history          過去の記録を一覧表示する
  dawn edit             JSONファイルを直接編集する
  dawn help             このヘルプを表示する

例:
  dawn "読書"           読書を記録
  dawn                  対話式で記録`

	fmt.Println(help)
}
