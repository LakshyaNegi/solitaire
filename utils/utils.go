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

func CharInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
