package main

// update
// const aCard = charSetStruct{} // did not work
/*
There are four files (or decks) of cards:
	fileOfCardsE extended
	fileOfCardsS
	fileOfCardsMostDifficult
	fileOfCards
*/
// All decks will draw cards per this aCard var
var aCard = charSetStruct{}

// The structure of a single 'card' (aCard.) from fileOfCards
type charSetStruct struct {
	Kata   string
	Hira   string
	Romaji string

	HiraHint   string
	KataHint   string
	TT_Hint    string
	SansR_Hint string
}

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
