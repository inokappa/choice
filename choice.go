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
	AppVersion = "0.0.2"
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
	var words string

	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
	}
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdin := bufio.NewScanner(os.Stdin)
		var slices []string
		for stdin.Scan() {
			slices = append(slices, stdin.Text())
		}
		if len(slices) == 1 {
			words = strings.Join(slices, sp)
		} else {
			words = strings.Join(slices, ",")
			sp = ","
		}

		word = Choice(words, sp)
	}
	fmt.Printf(word)
}
