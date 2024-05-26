package main

import "time"

/*
aCard.Kata
aCard.Hira
aCard.Romaji

actual_objective ::         aCard.Hira  aCard.Romaji
actual_objective_type ::    "hira",  "roma"

actual_prompt_char ::       aCard.Hira  aCard.Romaji  aCard.Kata
actual_prompt_char_type ::  "hira",  "roma",  "kata"

usersSubmission  ::         scanned from keyboard
type_of_usersSubmission ::  "hira",  "roma"                 // rarely used.
usersInputMode ::           "hira",  "roma"

guessLevelCounter ::        0-4  ??
byWayOfDirectiveHandler ::  true,  false
its_a_directive ::  true,  false
non_standard_origin_DirHandler = false
*/

var actual_objective string      // aCard.Hira  aCard.Romaji
var actual_objective_type string // "hira",  "roma"

var actual_prompt_char string      //  aCard.Hira  aCard.Romaji  aCard.Kata
var actual_prompt_char_type string //  "hira",  "roma",  "kata"

var usersSubmission string
var type_of_usersSubmission string //  "hira",  "roma"                 // rarely used.
var usersInputMode string          //  "hira",  "roma"

var guessLevelCounter = 0
var byWayOfDirectiveHandler = false
var its_a_directive = false
var non_standard_origin_DirHandler = false
var non_standard_origin_stcR bool

var user_guessed_prior_card_rightly bool    // may need this ?
var submission_already_processed_above bool // ? will I need this ?

var skip_for_reasons bool // ? will I need this ?

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

var include_Extended_kata_deck = false

var limitedToKataPrompts bool
var limitedToHiraPrompts bool
var limitedToRomaPrompts bool
var limitedToDifficultKata bool

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
