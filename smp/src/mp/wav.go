package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat    int
	porcess int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music ", source)

	p.porcess = 0

	for p.porcess < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.porcess += 10
	}

	fmt.Println("\nFinished playing ", source)
}
