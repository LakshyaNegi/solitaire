package game

import (
	"fmt"
	"log"
	"time"

	"github.com/mattn/go-tty"
)

func (g *game) Start() {
	g.IsRunning = true
	go g.input()

	for g.IsRunning {
		g.draw()
	}
}

func (g *game) input() {
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

		g.executeMove()
	}
}

func (g *game) clearInput(t *time.Timer, d time.Duration) {
	for {
		<-t.C
		g.InputBuf.Reset()
		t.Reset(d)
	}
}
