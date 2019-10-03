package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	AppVersion = "0.0.1"
)

var (
	argVersion  = flag.Bool("version", false, "バージョンを出力.")
	argSeparate = flag.String("separate", ",", "入力の区切り文字を指定 (デフォルトは , カンマ区切り).")
)

func Choice(words string, sp string) string {
	splited := strings.Split(words, sp)
	rand.Seed(time.Now().Unix())
	result := splited[rand.Intn(len(splited))]
	return result
}

func main() {
	flag.Parse()

	if *argVersion {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	var sp string
	if *argSeparate != "" {
		sp = *argSeparate
	}

	var word string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdin := bufio.NewScanner(os.Stdin)
		stdin.Scan()
		words := stdin.Text()
		word = Choice(words, sp)
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err)
		}
	}
	fmt.Printf(word)
}
