package main

import "time"

/*
	All global vars NOT located in the constants.go file :
 	... Or/and, NOT located in the objectsAndMethods.go file ::

	I am going to use globals extensively

    aCard.Kata
	aCard.Hira
	aCard.Romaji
*/
// ::: not needed: var prompt_field_to_display string
// Will just be either aCard.Hira, aCard.Kata, or aCard.Romaji
var expected_response_type string // Will be either "hira_char_as_users_guess" or "romaji_char_as_users_guess"
var displayed_prompt_type string  // "hira" or "romaji" or "kata"
var guess_attempt_count int       // Will be either 1, 2, or 3
var objective string
var promptField string
var objective_kind string
var typeOfUsersInput string

// prompt functions will be of 3 types ONLY and will ONLY display the prompt.
// obtain_and_process_user_input() will be a separate function that will be responsible to collect, categorize, test, and also display the result of all processing
// ... via consulting guess_attempt_count state (above). This will eliminate a lot of duplication and confusion.

var usersInput string
var userInput_category string // "is_a_directive_that_modifies_A_card", "is_a_misc_directive", "appears_to_be_a_guess"

var aDirectiveWasDetected bool
var returning_from_handling_a_directive bool

var hiraAcard string
var romajiAcard string

var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
var cameFrom_stcR_NOTstc bool

// thisCaseOfAnInHasAlreadyBeenProcessedAbove is being flagged as
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
