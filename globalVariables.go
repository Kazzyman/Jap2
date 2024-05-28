package main

import "time"

// count var used only to "Randomize" - or, at least, to alternate between kata and hira prompting when picking a new card.
var count = 1

var fileExplored int // Used only in countSLOC() and the associated about_app()

var actual_objective string      // aCard.Hira  aCard.Romaji
var actual_objective_type string // "hira",  "roma"

var actual_prompt_char string      //  aCard.Hira  aCard.Romaji  aCard.Kata
var actual_prompt_char_type string //  "hira",  "roma",  "kata"

var usersSubmission string

var guessLevelCounter = 0 // Used to determine the level of (or the format of) prompts.

var its_a_directive = false // A flag used to record the fact that we determined that the usersSubmission was a Directive.

var non_standard_origin_stcR bool // Used in silentlyLocateCard to signal that we've come by-way-of the strc Directive.

// user_guessed_prior_card_rightly, used w/ another var thisIsOurFirstRodeo (declared in main.go) to conditionally pick a new card.
var user_guessed_prior_card_rightly bool // , i.e., if user_guessed_prior_card_rightly || thisIsOurFirstRodeo { ...

var submission_already_processed_above bool // Used to avoid the Double-processing of usersSubmission s.

var include_Extended_kata_deck = false

var limitedToKataPrompts bool // These are used to control (limit) which field of the cards will be used as the prompt char.
var limitedToHiraPrompts bool
var limitedToRomaPrompts bool
var limitedToDifficultKata bool

var total_prompts int // Used in various statistical reporting functions.

// Game-Feature-control-and-tallying-vars-Section: ***************************************
var correctOnFirstAttempt int // Used along with no_interim_error_flag to log the obvious.
var single_faults int
var double_faults int
var errors int
var game_loop_counter int
var game_duration = 15
var no_interim_error_flag bool // Used to signal that the user has gotten this card correctly on his first attempt.
var gameOn bool

// ***************************************************************************************

var TimeOfStartFromInceptionOfGame = time.Now()

var whichDeck int // Currently not actually being used. Under review. 

var foundElement *charSetStruct // Used in silentlyLocateCard for the stc and stcr Directives. 

// Global Maps:
//
// Used in : func read_map_of_fineOn()
// ... and in : func logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(promptToSkip string)
var frequencyMapOf_IsFineOnChars = make(map[string]int) // - -
// Used in : func read_map_of_needWorkOn()
// ... and in : func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string)
var frequencyMapOf_need_workOn = make(map[string]int) // - -

// Create a kanjiHitMap map. A map of keys (kanji chars) and/to their associated history data, i.e., the struct: CardInfo
var hiraHitMap = make(map[string]CardInfo)

// The above is a map of the things below, which is keyed by a kanji char string
// The type of card which will be included in the above map
type CardInfo struct { // Technically, this is a constant rather than a global. My other types are found in constants.go
	UsersGuess            []string
	FirstMeaningOnRecord  string
	SecondMeaningOnRecord string
	CorrectGuessCount     int
}
