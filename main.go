package main

import (
	"fmt"
	"math/rand"
	"time"
)

var RPS = [3]string{"Rock", "Paper", "Scissors"}
var winner = map[string]string{
	"Rock":     "Scissors",
	"Paper":    "Rock",
	"Scissors": "Paper",
}

func pickOptions(out chan<- []string) {
	for i := 0; i < 100; i++ {
		time := time.Now().UnixNano()
		rand.Seed(time)
		random1 := RPS[rand.Intn(3)]
		random2 := RPS[rand.Intn(3)]
		out <- []string{random1, random2}
	}
	close(out)
}

func referee(in <-chan []string, out chan<- string) {
	for pick := range in {
		player1 := pick[0]
		player2 := pick[1]
		if value, ok := winner[player1]; ok && player2 == value {
			out <- fmt.Sprintf("Player 1 wins, %s beats %s", player1, player2)
		} else if value, ok := winner[player2]; ok && player1 == value {
			out <- fmt.Sprintf("Player 2 wins, %s beats %s", player2, player1)
		} else {
			out <- fmt.Sprintf("Tied, player 1 %s player 2 %s", player1, player2)
		}
	}
	close(out)
}

func main() {
	picks := make(chan []string)
	result := make(chan string)
	go pickOptions(picks)
	go referee(picks, result)

	for v := range result {
		fmt.Println(v)
	}
}
