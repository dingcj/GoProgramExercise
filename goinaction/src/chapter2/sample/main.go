package main

import (
	_ "chapter2/sample/matchers"
	"chapter2/sample/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)

	log.Println("Change log output to std out.")
}

func main() {
	search.Run("president")
}
