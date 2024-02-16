package main

import "time"

// All global vars NOT located in the constants.go file :
// ... Or/and, NOT located in the objectsAndMethods.go file ::

// Create a kanjiHitMap map. A map of keys (kanji chars) and/to their associated history data, i.e., the struct: CardInfo
var hiraHitMap = make(map[string]CardInfo)

// The above is a map of the things below, which is keyed by a kanji char string

// The type of card which will be included in the above map
type CardInfo struct {
	UsersGuess            []string
	FirstMeaningOnRecord  string
	SecondMeaningOnRecord string
	CorrectGuessCount     int
}

const colorYellow = "\033[33m"

var include_Extended_kata_deck = false

var limitedToKataPrompts bool
var limitedToHiraPrompts bool
var limitedToRomaPrompts bool
var limitedToDifficultKata bool

var randomFileOfCards int
var randomArrangementOfCards int

var returning_from_a_wrong_guess = false
var current_deck string
var current_deckA string
var total_prompts int

var game string
var gameOn bool
var startBeforeCall = time.Now()
var TimeOfStartFromTop = time.Now()

var game_loop_counter int

var whichDeck int

var foundElement *charSetStruct

var game_duration = 998

// Global Maps:
//
// Used in : func read_map_of_fineOn()
// ... and in : func logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(promptToSkip string)
var frequencyMapOf_IsFineOnChars = make(map[string]int) // - -
// Used in : func read_map_of_needWorkOn()
// ... and in : func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string)
var frequencyMapOf_need_workOn = make(map[string]int) // - -
