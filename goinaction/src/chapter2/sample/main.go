package main

import (
	"chapter2/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.Println("In function init.")
}

func main() {
	log.Println("In function main.")

	feeds, _ := search.RetrieveFeeds()
	for _, feed := range feeds {
		log.Println(feed)
	}
}
