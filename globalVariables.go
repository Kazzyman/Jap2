package main

import "time"

// count var used only to "Randomize" - or, at least, to alternate between kata and hira prompting when picking a new card.
var count = 1
var kata_roma bool
var kata_hira bool

var fileExplored int // Used only in countSLOC() and the associated about_app()

var actual_objective string      // aCard.Hira  aCard.Romaji
var actual_objective_type string // "hira",  "roma"

var actual_prompt_char string      //  aCard.Hira  aCard.Romaji  aCard.Kata
var actual_prompt_char_type string //  "hira",  "roma",  "kata"

var usersSubmission string

var its_a_directive = false // A flag used to record the fact that we determined that the usersSubmission was a Directive.

var non_standard_origin_stcR bool // Used in silentlyLocateCard to signal that we've come by-way-of the strc Directive.

var gotLastCardRightSoGetFreshOne bool

var submission_already_processed_above bool // Used to avoid the Double-processing of usersSubmission s.

var include_Extended_kata_deck = false

var limitedToKataPrompts bool // These are used to control (limit) which field of the cards will be used as the prompt char.
var limitedToHiraPrompts bool
var limitedToRomaPrompts bool
var limitedToDifficultDescriptions bool

var total_prompts int // Used in various statistical reporting functions.

// ::: Game-Feature-control-and-tallying-vars-Section: ---------------------------------------
var correctOnFirstAttemptAccumulator int
var correctOnSecondAttemptAccumulator int
var correctOnThirdAttemptAccumulator int
var failedOnThirdAttemptAccumulator int
var gottenHonestly bool
var nameOfPlayer string
var weHadFailed_And_OnlyGotThisRightBecauseOfTheClue bool

var guessLevelCounter = 1 // Used to determine the level of (or the format of) prompts.
var game_loop_counter int

const game_duration = 15 // Fun with constants and
var game_duration_set_by_user int
var now_using_game_duration_set_by_user bool

var theGameIsRunning bool

// ::: ---------------------------------------------------------------------------------------

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

const (
	jim = game_duration
)

const rick = "\u001B[33m" // Another example, fun with constants
const dick = rick
