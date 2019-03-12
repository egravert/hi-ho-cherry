package main

import (
	"math/rand"
)

type Game struct {
	p []*Player
	r *rand.Rand
}

func (g Game) Reset() {
	for _, p := range g.p {
		p.Reset()
	}
}

func (g Game) Play() {
	for rnd := 1; ; rnd++ {
		for _, p := range g.p {
			p.Spin(g.r)
			if p.Winner() {
				p.Wins++
				return
			}
		}
	}
}
