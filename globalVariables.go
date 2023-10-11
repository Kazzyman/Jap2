package main

import "time"

// All global vars NOT located in the constants.go file :
// ... Or/and, NOT located in the objectsAndMethods.go file ::

var include_Extended_kata_deck = false

var randomFileOfCards int

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
//
// Used in : func read_map_of_needWorkOn()
// ... and in : func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string)
var frequencyMapOf_need_workOn = make(map[string]int) // - -
