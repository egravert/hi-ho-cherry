package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	matches = 1000000
	players = 4
)

func main() {

	// initialize game state
	g := Game{r: rand.New(rand.NewSource(time.Now().UnixNano()))}
	for i := 0; i < players; i++ {
		g.p = append(g.p, new(Player))
	}

	// main game loop
	for i := 1; i <= matches; i++ {
		g.Play()
		g.Reset()
		if i <= 10 || i == 100 || i == 1000 || i == 10000 || i == 100000 || i == 1000000 {
			fmt.Printf("%d, %5.2f, %5.2f, %5.2f, %5.2f\n",
				i,
				float32(g.p[0].Wins)/float32(i)*100,
				float32(g.p[1].Wins)/float32(i)*100,
				float32(g.p[2].Wins)/float32(i)*100,
				float32(g.p[3].Wins)/float32(i)*100)
		}
	}
}
