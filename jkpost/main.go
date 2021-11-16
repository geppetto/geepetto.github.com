package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("insert title")
		os.Exit(1)
	}
	t := os.Args[1]
	now := time.Now()
	fn := now.Format("2006-01-02-" + t + ".md")
	c := "---\n" +
		"layout: post\n" +
		"title: \"" + t + "\"\n" +
		"date: " + now.Format("2006-01-02 15:04:05") + " +0900\n" +
		"---\n" +
		"\n"
	fmt.Println(c)
	fmt.Println(fn)

	f, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fmt.Fprintf(f, c)
}
