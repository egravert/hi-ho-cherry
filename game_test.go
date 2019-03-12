package main

import (
	"math/rand"
	"testing"
	"time"
)

func Benchmark_MonteCarlo(b *testing.B) {
	g := Game{r: rand.New(rand.NewSource(time.Now().UnixNano()))}
	for i := 0; i < 4; i++ {
		g.p = append(g.p, new(Player))
	}
	for n := 0; n < b.N; n++ {
		g.Play()
		g.Reset()
	}
}
