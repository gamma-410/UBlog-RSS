package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)

	switch arg {
	case "number":
		numbers()
	case "info":
		info()
	case "article":
		num := flag.Arg(1)
		i, _ := strconv.Atoi(num)

		article(i)

	default:
		fmt.Println("コマンド引数を指定してください...!")
	}
}

func numbers() {
	feed, err := gofeed.NewParser().ParseURL("https://blog.umiirosoft.com/feed/")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("\n -> 現在 ", len(feed.Items), " 件の投稿があります！\n")
}

func info() {
	feed, err := gofeed.NewParser().ParseURL("https://blog.umiirosoft.com/feed/")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("\n", feed.Title)
	fmt.Println("\n -> ", feed.Description)
	fmt.Println(" -> ", feed.Link, "\n")
}

func article(i int) {
	feed, err := gofeed.NewParser().ParseURL("https://blog.umiirosoft.com/feed/")
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("\n", feed.Items[i].Title)
	fmt.Println("\n -> ", feed.Items[i].Link)
	fmt.Println(" -> ", feed.Items[i].PublishedParsed.Format(time.RFC3339), "\n")
}
