package card

import (
	"math/rand"
	"time"
)

var SYMS = []string{
	"♥",
	"♦",
	"♠",
	"♣",
}

var VALS = []string{
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"0",
	"V",
	"D",
	"K",
	"T",
}

// returned random int (0...n)
func R(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}
