/*

Package flip (╯°□°）╯︵ʇxǝʇ

Go library used to flip text.

Installation

Just go get the package:

    go get -u github.com/peterhellberg/flip

Usage

A small usage example

    package main

    import (
    	"fmt"
    	"os"
    	"strings"

    	"github.com/peterhellberg/flip"
    )

    func main() {
    	fmt.Println(flip.Table(strings.Join(os.Args[1:], " ")))
    }

*/
package flip

import "strings"

var upsideDownChars = map[string]string{
	"z": "z", "y": "ʎ", "x": "x", "w": "ʍ", "v": "ʌ", "u": "n", "t": "ʇ",
	"s": "s", "r": "ɹ", "q": "b", "p": "d", "o": "o", "n": "u", "m": "ɯ",
	"l": "ן", "k": "ʞ", "j": "ɾ", "i": "ᴉ", "h": "ɥ", "g": "ƃ", "f": "ɟ",
	"e": "ǝ", "d": "p", "c": "ɔ", "b": "q", "a": "ɐ", " ": " ", "-": "-",
	"+": "+", "Z": "Z", "Y": "⅄", "X": "X", "W": "M", "V": "Λ", "U": "∩",
	"T": "┴", "S": "S", "R": "ɹ", "Q": "Q", "P": "Ԁ", "O": "O", "N": "N",
	"M": "W", "L": "˥", "K": "ʞ", "J": "ſ", "I": "I", "H": "H", "G": "פ",
	"F": "Ⅎ", "E": "Ǝ", "D": "p", "C": "Ɔ", "B": "q", "A": "∀", "9": "6",
	"8": "8", "7": "ㄥ", "6": "9", "5": "ϛ", "4": "ㄣ", "3": "Ɛ",
	"2": "ᄅ", "1": "Ɩ", "0": "0",
}

const (
	// DefaultEmoticon is the default emoticon used to flip tables
	DefaultEmoticon = "(╯°□°）╯︵"

	// GopherEmoticon is the gopher emoticon used to do a gopher flip
	GopherEmoticon = "ʕ╯◔ϖ◔ʔ╯︵"

	// AngryEmoticon is an angry emoticon used to angrily flip the table
	AngryEmoticon = "(ノಠ益ಠ)ノ︵"

	// SparklyEmoticon is a very sparkly emoticon used to flip the table
	SparklyEmoticon = "(ﾉ◕ヮ◕)ﾉ*:･ﾟ✧*:･ﾟ✧ "
)

var (
	// Table is the default table flipper func
	Table = Func(DefaultEmoticon)

	// Gopher is the gopher table flipper func
	Gopher = Func(GopherEmoticon)

	// Angry is the angry table flipper func
	Angry = Func(AngryEmoticon)

	// Sparkly is the sparkly table flipper func
	Sparkly = Func(SparklyEmoticon)

	// Flippers is a map of named flipper funcs
	Flippers = map[string]func(string) string{
		"table":   Table,
		"gopher":  Gopher,
		"angry":   Angry,
		"sparkly": Sparkly,
	}
)

// Func returns a flipper func
func Func(s string) func(string) string {
	f := Flipper(s)

	return func(s string) string {
		return f.Flip(s)
	}
}

// Flipper is a string that decorates a flipped string
type Flipper string

// Flip flips and decorates a string
func (f Flipper) Flip(s string) string {
	return string(f) + UpsideDown(s)
}

// UpsideDown turns the given string upside down
func UpsideDown(s string) string {
	ss := strings.Split(s, "")
	ns := ""

	for i := len(ss) - 1; i >= 0; i-- {
		uc := upsideDownChars[ss[i]]

		if uc == "" {
			uc = upsideDownChars[strings.ToLower(ss[i])]
		}

		if uc != "" {
			ns += uc
		}
	}

	return ns
}

// Reverse reverses the given string
func Reverse(s string) string {
	ss := strings.Split(s, "")
	ns := ""

	for i := len(ss) - 1; i >= 0; i-- {
		ns += ss[i]
	}

	return ns
}
