package main

/*
There are several files (or decks) of cards:
	They are each contained in files named:
	data_____.go
*/

// The structure of a single 'card' (aCard.) ; mostly from fileOfCards____ ; mostly. :)
type charSetStruct struct {
	Kata   string // Typically katakana, but kanji in special cases like fileOfCardsKanjiHard
	Hira   string // Typically hiragana, but kanji in special cases like fileOfCardsKanjiHard
	Romaji string // sometimes Meaning; ::: but normally roma notation

	HiraHint   string
	KataHint   string
	TT_Hint    string
	SansR_Hint string
}

// All decks will draw cards per this aCard var
var aCard = charSetStruct{}

/*
 We instantiate a series of struct objects as slices of instances of the charSetStruct type
 e.g., fileOfCards being a "slice" of structures of type charSetStruct
 i.e., We are creating a slice named fileOfCards that holds instances of the charSetStruct type

Each element in a slice is initialized using the composite literal syntax whereby
... we are providing values for each field of the charSetStruct struct: i.e.,
... each set of values enclosed in curly braces { ... } represents an instance of the struct
*/

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"
const colorYellow = "\033[33m"
