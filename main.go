package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	// UUIDの正規表現パターンを定義
	uuidRegex := regexp.MustCompile(`[a-fA-F0-9]{8}-([a-fA-F0-9]{4}-){3}[a-fA-F0-9]{12}`)

	// 標準入力からデータを読み込むためのスキャナを作成
	reader := bufio.NewReader(os.Stdin)

	for {
		// 一行ずつ読み込む
		line, err := reader.ReadString('\n')

		// EOFに到達した場合はループを抜ける
		if err == io.EOF {
			// 最後の行にも置換処理を適用
			line = uuidRegex.ReplaceAllString(line, "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
			fmt.Print(line)
			break
		}

		// エラーが発生した場合はエラーメッセージを表示して終了
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to read from stdin: %v\n", err)
			os.Exit(1)
		}

		// UUIDを置換
		line = uuidRegex.ReplaceAllString(line, "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")

		// 結果を出力
		fmt.Print(line)
	}
}
