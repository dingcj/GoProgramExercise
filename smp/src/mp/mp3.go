package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat    int
	porcess int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music ", source)

	p.porcess = 0

	for p.porcess < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.porcess += 10
	}

	fmt.Println("\nFinished playing ", source)
}
