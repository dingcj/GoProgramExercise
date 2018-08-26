package main

import "fmt"
import "library"
import "strconv"

func main() {
	fmt.Println("vim-go")
}

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			lib.Add(&library.MusicEntry{strconv.Itoa(id), tokens[2],
				tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			tmp := lib.RemoveByName(tokens[2])
			if tmp == nil {
				fmt.Println(tokens[2], " is not exist!")
			} else {
				fmt.Println(tokens[2], " is deleted success!")
			}
		}
	default:
		fmt.Println("Unrecognized lib command: ", tokens[1])
	}
}
