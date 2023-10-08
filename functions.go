package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

// DIRECTIVES : --------------------------------------------------------------------------------------------
//
func handle_doubleQuestMark_directive(objective_kind string) { //                 - -
	var Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn string
	//
	fmt.Printf("\n  -- Type either a Hiragana or Romaji prompt you need help with:> ")
	_, _ = fmt.Scan(&Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn)
	//
	locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn, objective_kind)
	fmt.Println("")
}

// Handles the Directive 'set'
func reSet_aCard_andThereBy_reSet_thePromptString() (prompt, objective, objective_kind string) { //     - -
	var theHiraganaOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Hiragana to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)

	// Determine if the user has entered a valid Hiragana char (instead of, accidentally, an alpha char or string)
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(theHiraganaOfCardToSilentlyLocate):
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}
	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'
	if isAlphanumeric == true {
		fmt.Println("Are you in alphanumeric input mode?")
		fmt.Printf("... if so, change it to Hiragana (or I mignt die)\n")
		fmt.Printf("%s", colorRed) //
		fmt.Printf(" cautiously ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("re-enter your selection, in Hiragana mode :> ")
		fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)
		// May yet send an Alpha string to the next func, which will itself deal with it elegantly
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		// new_prompt, new_objective, new_objective_kind
		prompt = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		prompt = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		fmt.Println("")
	}
	return prompt, objective, objective_kind
}

// end of DIRECTIVES -----------------------------------------------------------------------------------

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { //      - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}

func testForDirective(in string) (result bool) {
	if in == "set" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "menu" ||
		in == "reset" ||
		in == "stat" ||
		in == "st" ||
		in == "dir" ||
		in == "notes" ||
		in == "quit" ||
		in == "q" ||
		in == "exit" ||
		in == "ex" ||
		in == "stats" ||
		in == "rm" ||
		in == "gameon" ||
		in == "gameoff" ||
		in == "gamed" {
		// Then:
		result = true
	}
	return result
}

var game string
var gameOn bool
var startBeforeCall = time.Now()
var TimeOfStartFromTop = time.Now()

func game_on() (game string) { // - -
	game = "on"
	gameOn = true
	fmt.Println("The game is on")

	startBeforeCall = time.Now()
	currentTime := time.Now()
	TimeOfStartFromTop = time.Now()

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game began at: %s \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
	return game
}
func game_off() (game string) { // - -
	game = "off"
	gameOn = false
	game_duration = 1000

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	// calculate elapsed time
	t_s2 := time.Now()
	elapsed := t_s2.Sub(TimeOfStartFromTop)

	// cast time durations to a String type for Fprintf "formatted print"
	TotalRun := elapsed.String()

	fmt.Printf("\nRun time was: %s, gameOn is: %t \n\n", TotalRun, gameOn)

	// End timer and report elapsed time
	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), total_prompts)
	check_error(err1)
	_, err2 := fmt.Fprintf(fileHandle,
		"\n Elapsed time of game was: %s \n",
		TotalRun)
	check_error(err2)
	return game
}

var game_loop_counter int

func respond_to_UserSuppliedDirective(in, objective_kind string) (prompt, objective, kind string) {
	var count int
	switch in {
	case "gamed":
		fmt.Println("Enter a number for how many prompts there will be in the game")
		_, _ = fmt.Scan(&count)
		game_duration = count - 2
	case "gameon":
		// game_loop_counter ++
		game_on()
	case "gameoff":
		game_off()
	case "reset":
		// Flush (clear) the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// Also, flush (clear) the maps
		total_prompts = 0
		//
		//goland:noinspection ALL
		fmt.Println("\nArrays and maps flushed:\n")
		fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
		fmt.Println("    cyclicArrayHits")
		fmt.Println("    frequencyMapOf_IsFineOnChars")
		//goland:noinspection ALL
		fmt.Println("    frequencyMapOf_need_workOn\n")
		fmt.Println("  And, all Game values have also been reset")
	case "quit":
		os.Exit(1)
	case "q":
		os.Exit(1)
	case "exit":
		os.Exit(1)
	case "ex":
		os.Exit(1)
	case "??": // Directives follow:
		handle_doubleQuestMark_directive(objective_kind)
	case "?":
		fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
	case "set":
		prompt, objective, kind = reSet_aCard_andThereBy_reSet_thePromptString()
	case "stat":
		hits()
	case "st":
		hits()
	case "stats":
		hits()
	case "notes":
		//goland:noinspection ALL  **do-this**
		fmt.Println("\nIn the traditional Hepburn romanization system, the sound じ in hiragana is romanized as \"ji\" \n" +
			"and the katakana ジ is also romanized as \"ji\" \n\n" +
			"However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as\n" +
			" \"zi\" instead of \"ji\"\n\n" +
			"The sound gi:ぎ in hiragana is romanized as \"gi\" and the katakana ギ is also romanized as \"gi\"\n")
		//goland:noinspection ALL  **do-this**
		fmt.Println("゜is called \"handakuten\" 半濁点 translates to \"half-voicing mark\" or \"semi-voiced mark\"\n" +
			"゛is called \"dakuten\" 濁点 meaning 'voiced mark' or 'voicing mark'")
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		display_Listing_of_Directives()
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
	default:
		// fmt.Println("Directive not found") // Does not work because only existent cases are passed to the switch
	}
	return prompt, objective, kind
}

var whichDeck int

func pick_RandomCard_Assign_fieldsT() (promptField, objective, objective_kind string) { // - -
	randIndex := rand.Intn(len(fileOfCardsMostDifficult))

	aCard = fileOfCardsMostDifficult[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Hira
	objective = aCard.Romaji
	objective_kind = "Romaji"
	whichDeck = 3
	return promptField, objective, objective_kind
}

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	randIndex := rand.Intn(len(fileOfCards))
	randIndexS := rand.Intn(len(fileOfCardsS))
	randIndexD := rand.Intn(len(fileOfCardsMostDifficult))

	randomFileOfCards := rand.Intn(12)

	// Hira prompting, Romaji objective:
	if randomFileOfCards == 0 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 1
	}
	if randomFileOfCards == 1 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 2
	}
	if randomFileOfCards == 2 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 3
	}

	// Kata prompting, Romaji objective:
	if randomFileOfCards == 3 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 1
	}
	if randomFileOfCards == 4 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 2
	}
	if randomFileOfCards == 5 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 3
	}

	// Romaji prompting, Hira objective:
	if randomFileOfCards == 6 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 1

	}
	if randomFileOfCards == 7 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 2
	}
	if randomFileOfCards == 8 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 3
	}

	// Kata prompting, Hira objective:
	if randomFileOfCards == 9 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 1
	}
	if randomFileOfCards == 10 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 2
	}
	if randomFileOfCards == 11 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 3
	}
	return promptField, objective, objective_kind
}
