package main

import "bufio"
import "fmt"
import "library"
import "mp"
import "os"
import "strconv"
import "strings"

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

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("This music ", tokens[1], " does not exist!")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
	Enter following commands to control the player:
	lib list -- View the existing music lib
	lib add <name><artist><source><type> -- Add a new music to the music lib
	lib remove <name> -- Remove the specified music from the lib
	play name -- Play the specified music`)

	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized commands: ", tokens[0])
		}
	}
}
