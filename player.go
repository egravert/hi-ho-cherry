package main

import "math/rand"

const (
	One = iota + 1
	Two
	Three
	Four
	Dog
	Bird
	SpilledBasket
)

type Player struct {
	cherries int
	Wins     int
}

func (p *Player) Spin(r *rand.Rand) {
	spin := r.Intn(7) + 1
	switch spin {
	case One, Two, Three, Four:
		p.Take(spin)
	case Dog, Bird:
		p.PutBack()
	case SpilledBasket:
		p.Reset()
	}
}

func (p *Player) Reset() {
	p.cherries = 0
}

func (p *Player) Winner() bool {
	return p.cherries >= 10
}

func (p *Player) Take(cherries int) {
	p.cherries += cherries
}

func (p *Player) PutBack() {
	if p.cherries >= 2 {
		p.cherries -= 2
	} else {
		p.cherries--
	}
}
