package game

import (
	"fmt"
	"log"
	"time"

	"github.com/mattn/go-tty"
)

func (g *Game) Start() {
	g.IsRunning = true
	go g.input()
}

func (g *Game) input() {
	d := 3 * time.Second

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	timer := time.NewTimer(d)

	for {
		go g.clearInput(timer, d)

		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		g.InputBuf.WriteRune(r)

		in := fmt.Sprintf("Input: %s", g.InputBuf.String())

		g.screen.DrawAt(in, 0, 20)

		g.screen.Flush()

		g.ExecuteMove()
	}
}

func (g *Game) clearInput(t *time.Timer, d time.Duration) {
	for {
		<-t.C
		g.InputBuf.Reset()
		t.Reset(d)
	}
}
