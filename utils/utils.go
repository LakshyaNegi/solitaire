package utils

import (
	"log"
	"os"
)

func Log(str string) {
	f, err := os.Create("game.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
}
