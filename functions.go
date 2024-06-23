package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { // ::: - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}

/*
.
*/
// DIRECTIVES : --------------------------------------------------------------------------------------------
// func handle_doubleQuestMark_directive(actual_objective_type string) { //        - -
func handle_doubleQuestMark_directive() { // ::: - -
	var Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn string
	//
	fmt.Printf("\n  -- Type either a Hiragana or Romaji prompt you need help with:> ")
	_, _ = fmt.Scan(&Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn)
	//
	// locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn, actual_objective_type)
	locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn)
	fmt.Println("")
}

/*
.
*/
// Handles the Directive ::: 'stc' <--------------------------
func reSet_via_a_hira_aCard_andThereBy_reSet_thePromptString() { // ::: - -
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

		if limitedToKataPrompts { // ::: we always know the card. So, if we also know one more thing, e.g., the type of the prompt. We then know other things.
			// and, we always know the type of the prompt, ALWAYS!!!! Unfortunately, there are three (not two) prompt types.
			// If there were only two prompt types we could know the object by elimination.
			actual_prompt_char = aCard.Kata
		} else if limitedToRomaPrompts {
			actual_prompt_char = aCard.Romaji
		} else if limitedToHiraPrompts {
			actual_prompt_char = aCard.Hira
		} else if limitedToDifficultDescriptions {
			actual_prompt_char = aCard.Kata
		} else {
			// ::: don't even think about it!
		}

		// by the current objective type
		if actual_objective_type == "roma" {
			actual_objective = aCard.Romaji
		} else { // ::: else it is "hira"
			actual_objective = aCard.Hira
		}
		// ::: or, said differently :
		if actual_objective_type == "hira" {
			actual_objective = aCard.Hira
		} else { // ::: else it is "roma"
			actual_objective = aCard.Romaji
		}

		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'

		// ::: said yet another way :
		// by the current objective type
		if actual_objective_type == "roma" {
			actual_objective = aCard.Romaji
		}
		if actual_objective_type == "hira" {
			actual_objective = aCard.Hira
		}
		fmt.Println("")
	}
}

/*
.
*/
// Handles the Directive ::: 'stcr' <--------------------------
func reSet_aCard_via_a_romaji_andThereBy_reSet_thePromptString() { // ::: - -
	var theRomajiOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Romaji to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theRomajiOfCardToSilentlyLocate)

	// Determine if the user has entered a valid Hiragana char (instead of, accidentally, an alpha char or string)
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(theRomajiOfCardToSilentlyLocate):
		isAlphanumeric = false // ::: true in stc/hira version false here *********
	default:
		isAlphanumeric = true // ::: false in stc/hira version, true here *********************************
	}
	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'
	if isAlphanumeric == true { // ::: true in stc/hira version (and here too) **************************
		fmt.Println("Are you in Japanese input mode?")
		fmt.Printf("... if so, change it to US/romaji (or I mignt die)\n")
		fmt.Printf("%s", colorRed) //
		fmt.Printf(" cautiously ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("re-enter your selection, in US/romaji mode :> ")
		fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&theRomajiOfCardToSilentlyLocate)
		// May yet send a hira string to the next func, which will itself deal with it elegantly -- we hope.
		silentlyLocateCardr(theRomajiOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                // Set the global var-object 'aCard'

		if limitedToKataPrompts { // ::: we always know the card. So, if we also know one more thing, e.g., the type of the prompt. We then know other things.
			// and, we always know the type of the prompt, ALWAYS!!!! Unfortunately, there are three (not two) prompt types.
			// If there were only two prompt types we could know the object by elimination.
			actual_prompt_char = aCard.Kata
		} else if limitedToRomaPrompts {
			actual_prompt_char = aCard.Romaji
		} else if limitedToHiraPrompts {
			actual_prompt_char = aCard.Hira
		} else if limitedToDifficultDescriptions {
			actual_prompt_char = aCard.Kata
		} else {
			// ::: don't even think about it!
		}

		// by the current objective type
		if actual_objective_type == "roma" {
			actual_objective = aCard.Romaji
		} else { // ::: else it is "hira"
			actual_objective = aCard.Hira
		}
		// ::: or, said differently :
		if actual_objective_type == "hira" {
			actual_objective = aCard.Hira
		} else { // ::: else it is "roma"
			actual_objective = aCard.Romaji
		}

		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCardr(theRomajiOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                // Set the global var-object 'aCard'

		// by the current objective type
		// ::: said yet another way :
		if actual_objective_type == "roma" {
			actual_objective = aCard.Romaji
		}
		if actual_objective_type == "hira" {
			actual_objective = aCard.Hira
		}
		fmt.Println("")
	}
}

/*
.
*/

func the_game_begins() { // ::: - -
	theGameIsRunning = true // ::: this flag is the only thing that "starts" a game
	game_loop_counter = 0
	// correctOnFirstAttemptAccumulator = 1   // ::: here it is/was not able to process the last guess prior to game ending.
	// correctOnSecondAttemptAccumulator = -1 // ::: kluge !!
	// ::: if the first query of a game is gotten right on the first attempt, it is logged as a 2nd and not a 1st
	correctOnFirstAttemptAccumulator = 0
	correctOnSecondAttemptAccumulator = 0
	failedOnThirdAttemptAccumulator = 0

	//
	// ::: file-writing and time-stamping is all that follows ==========================================
	currentTime := time.Now()
	TimeOfStartFromInceptionOfGame = time.Now()

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	_, err1 := fmt.Fprintf(fileHandle,
		"\n\n-- A game began at: %s",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
}
func the_game_ends() { // ::: - -
	theGameIsRunning = false
	now_using_game_duration_set_by_user = false
	// game_duration_set_by_user = 0
	// game_loop_counter = 0

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	// calculate elapsed time
	t_s2 := time.Now()
	elapsed := t_s2.Sub(TimeOfStartFromInceptionOfGame)

	// Format the elapsed time to display minutes and whole seconds
	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60

	// Create the formatted string
	TotalRun := fmt.Sprintf("%02d:%02d", minutes, seconds)

	// ::: ------------------------------------------------------------------- v v v v
	Process_users_input_as_a_guess() // ::: trying this ---- and these ------- v v v v
	if guessLevelCounter == 2 {
		if gottenHonestly {
			correctOnFirstAttemptAccumulator++
			gottenHonestly = false
		}
	} else if guessLevelCounter == 3 {
		correctOnSecondAttemptAccumulator++
	} else { // :::  this else never runs, so guessLevelCounter is either 2 or 3, but never 1 or 4 ? ---------------------------
		// correctOnThirdAttemptAccumulator++ // ::: why had I commented-out this line ??? if the else never is done, it could not have mattered
		fmt.Printf("guessLevelCounter is:%d, where it maybe-should-be-4-? \n", guessLevelCounter) // ::: never executes ????
		// ... then ... the guessLevelCounter was 4?.
		// correctOnThirdAttemptAccumulator++ // ::: why had I commented-out this line ???
		// ::: fail/error accumulator gets incremented (or at least it gets displayed [below]) during the oops message
	} // - - - - - - - - - - - - - - - - - - - - - - - - - ::: -----------------------------------------------------------------

	if correctOnFirstAttemptAccumulator > 0 && correctOnSecondAttemptAccumulator > 0 && correctOnThirdAttemptAccumulator > 0 && failedOnThirdAttemptAccumulator == 0 { // ::: done
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d%s right on your second try,\n"+
			"... and you got %s%d right on your third try. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator, colorRed, colorReset, correctOnSecondAttemptAccumulator,
			colorRed, colorReset, correctOnThirdAttemptAccumulator)
	} else if correctOnFirstAttemptAccumulator > 0 && correctOnSecondAttemptAccumulator > 0 && correctOnThirdAttemptAccumulator == 0 && failedOnThirdAttemptAccumulator == 0 { // ::: done
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d right on your second try. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator,
			colorRed, colorReset, correctOnSecondAttemptAccumulator)
	} else if correctOnFirstAttemptAccumulator > 0 && correctOnSecondAttemptAccumulator == 0 && correctOnThirdAttemptAccumulator == 0 && failedOnThirdAttemptAccumulator == 0 { // ::: done
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  Gongratulations! you got %s%d correct on your first try. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator)
	} else {
		fmt.Println(colorRed)
		fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d%s right on your second try,\n"+
			"... and you got %s%d%s right on your third try, and were unable to answer correctly without a hint "+
			"%s%d times. \n\n", TotalRun, colorReset, correctOnFirstAttemptAccumulator, colorRed, colorReset, correctOnSecondAttemptAccumulator,
			colorRed, colorReset, correctOnThirdAttemptAccumulator, colorRed, colorReset, failedOnThirdAttemptAccumulator)
	}

	// End timer and report elapsed time and other stats to a file.
	_, err1 := fmt.Fprintf(fileHandle,
		"\n-- A game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), game_loop_counter)
	check_error(err1)

	_, err3 := fmt.Fprintf(fileHandle, "%s's results were as follows: Right on first attempt:%d, on 2nd attempt:%d, 3rd attempt:%d, even a hint was ineffective:%d, %d/%d\n",
		nameOfPlayer, correctOnFirstAttemptAccumulator,
		correctOnSecondAttemptAccumulator, correctOnThirdAttemptAccumulator,
		failedOnThirdAttemptAccumulator, game_loop_counter, game_duration_set_by_user)
	check_error(err3)

	_, err2 := fmt.Fprintf(fileHandle,
		"The Elapsed time of the game was: %s \n\n\n",
		TotalRun)
	check_error(err2)
}

/*
.
*/
func read_pulledButNotUsedMap() { // ::: - -
	if len(pulledButNotUsedMap) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe seenMap is empty\n")
		fmt.Printf(colorReset)
	}
	for s, f := range pulledButNotUsedMap {
		if s != "" {
			if s != "primedK0" {
				if f > 1 {
					fmt.Printf(" --- From pulledButNotUsedMap: string is:")
					fmt.Printf(colorCyan)
					fmt.Printf("%s", s)
					fmt.Printf(colorReset)
					fmt.Printf(", freq is:")
					fmt.Printf(colorRed)
					fmt.Printf("%d", f)
					fmt.Printf(colorReset)
					fmt.Printf("\n")
				}
			}
		}
	}
	fmt.Println("")
	/*
			var indexIntoArray int
			for indexIntoArray < len(cyclicArrayPulls.pulls) {
				if cyclicArrayPulls.pulls[indexIntoArray] != "" {
					if cyclicArrayPulls.pulls[indexIntoArray] != "primedK0" {
						fmt.Printf("Char stored in cyclicArrayPulls: %s \n", cyclicArrayPulls.pulls[indexIntoArray])
					}
				}
				indexIntoArray++
			}
		//
			//
			//

	*/
}

/*
.

	func read_pulls_not_used_array() {
		for i, lastPull := range cyclicArrayPulls.pulls {
			if i < len(cyclicArrayPulls.pulls) {
				if lastPull != "" {
					fmt.Printf("array: %s\n", lastPull)
				}
			} else {
				break
			}
		}
	}

/*
.
*/
func about_app() { // ::: - -
	fmt.Printf("\nThis app consists of the following lines of code accross %d files:\n\n", fileExplored)
	countSLOC()
	fmt.Printf("As Calculated in real-time by countSLOC(), a custom internal function.\n\n")
}

/*
.
*/
func reset_all_data() { // ::: - -
	// Flush (clear) the old stats and hits arrays
	limitedToKataPrompts = false
	limitedToHiraPrompts = false
	limitedToRomaPrompts = false
	limitedToDifficultDescriptions = false
	include_Extended_kata_deck = false
	cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
	cyclicArrayHits = CyclicArrayHits{}
	cyclicArrayPulls = CyclicArrayPulls{}
	// Also, flush (clear) the maps
	hiraHitMap = make(map[string]CardInfo)
	frequencyMapOf_IsFineOnChars = make(map[string]int)
	frequencyMapOf_need_workOn = make(map[string]int)
	pulledButNotUsedMap = make(map[string]int)
	total_prompts = 0
	correctOnFirstAttemptAccumulator = 0
	correctOnSecondAttemptAccumulator = 0
	correctOnThirdAttemptAccumulator = 0
	failedOnThirdAttemptAccumulator = 0
	game_loop_counter = 0
	// jim = 15
	theGameIsRunning = false

	total_prompts = 0
	//
	//goland:noinspection ALL
	fmt.Println("\nArrays and maps flushed:\n")
	fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
	fmt.Println("    cyclicArrayHits")
	fmt.Println("    cyclicArrayPulls and hiraHitMap and pulledButNotUsedMap")
	fmt.Println("    frequencyMapOf_IsFineOnChars")
	//goland:noinspection ALL
	fmt.Println("    frequencyMapOf_need_workOn\n")
	fmt.Println(colorCyan + "  Limitations re Kata, Hira, and Romaji prompting; as well as all Game values have also been reset\n" + colorReset)
}

/*
.
*/
func notes_on_kana() { // ::: - -
	//goland:noinspection ALL  **do-this**
	fmt.Println("\nIn the traditional Hepburn romanization system, the sound じ in hiragana is romanized as \"ji\" \n" +
		"and the katakana ジ is also romanized as \"ji\" \n\n" +
		"However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as\n" +
		" \"zi\" instead of \"ji\"\n\n" +
		"The sound gi:ぎ in hiragana is romanized as \"gi\" and the katakana ギ is also romanized as \"gi\"\n")
	//goland:noinspection ALL  **do-this**
	fmt.Println("゜is called \"handakuten\" 半濁点 translates to \"half-voicing mark\" or \"semi-voiced mark\"\n" +
		"゛is called \"dakuten\" 濁点 meaning 'voiced mark' or 'voicing mark'")
	fmt.Println("\"digraphs\" is the word that refers to what I have called conjunctions, like ひゅ, for example ")
}

/*
.
.
*/
func detectDirective(in string) { // ::: - -
	if in == "stc" ||
		in == "stcr" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "rs" ||
		in == "st" ||
		in == "dir" ||
		in == "nts" ||
		in == "q" ||
		in == "rm" ||
		in == "abt" ||
		in == "exko" ||
		in == "exkf" ||
		in == "konly" ||
		in == "honly" ||
		in == "ronly" ||
		in == "donly" ||
		in == "hko" ||
		in == "help" {
		// Then:
		its_a_directive = true
	}
	// return result
}

/*
.
*/

func respond_to_UserSupplied_Directive(usersSubmission string) { // ::: - -
	switch usersSubmission {
	// Directives follow:
	// ::: alphabetically (mostly)
	case "?":
		fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
	case "??":
		// handle_doubleQuestMark_directive(actual_objective_type)
		handle_doubleQuestMark_directive()
	case "help":
		helpText()
	case "hko":
		limitedToKataPrompts = true
		limitedToHiraPrompts = true
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
	case "konly":
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
	case "honly":
		limitedToHiraPrompts = true
		limitedToKataPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
	case "ronly":
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultDescriptions = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
	case "donly":
		limitedToDifficultDescriptions = true
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
		/*
			.
		*/
	case "abt":
		about_app()

	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()
	case "exko":
		include_Extended_kata_deck = true
		fmt.Println("Extended Kata deck has been loaded")
	case "exkf":
		include_Extended_kata_deck = false
		fmt.Println("Extended Kata deck has been un-loaded")
	case "nts":
		notes_on_kana()
	case "q":
		os.Exit(1)
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
		read_pulledButNotUsedMap()
		// read_pulls_not_used_array()
	case "rs":
		reset_all_data()
	case "st":
		st_stats_function()
	case "stc":
		reSet_via_a_hira_aCard_andThereBy_reSet_thePromptString()
	case "stcr":
		reSet_aCard_via_a_romaji_andThereBy_reSet_thePromptString()
	default:
		// fmt.Println("Directive not found") // Does not ever happen because only existent cases are passed to the switch.
	}
}

/*
.
*/
func st_stats_function() { // ::: - -
	newHits()
	if !include_Extended_kata_deck {
		fmt.Println("Extended Kata deck is NOT loaded\n")
	} else {
		fmt.Println("Extended Kata deck is loaded\n")
	}

	if limitedToKataPrompts {
		fmt.Printf("Limited to Kata prompts with romaji objectives: %t \n\n", limitedToKataPrompts)
	}
	if limitedToHiraPrompts {
		fmt.Printf("Limited to Hira prompts only: %t \n\n", limitedToHiraPrompts)
	}
	if limitedToRomaPrompts {
		fmt.Printf("Limited to Romaji prompts only: %t \n\n", limitedToRomaPrompts)
	}
	if limitedToDifficultDescriptions {
		fmt.Printf("Limited to Difficult Kata only: %t \n\n", limitedToDifficultDescriptions)
	}
}

/*
.
*/
func helpText() { // ::: - -
	fmt.Println(colorRed)
	fmt.Printf("\nTo use this app you will need either a Japanese keyboard\n or a software work-around for same. \n\n")
	fmt.Println(colorReset)
}

/*
// This func, sans "_Extended" at the end of its name, becomes ::: a testing version of pick_RandomCard_Assign_fields()
// ... for fileOfCardsE : the deck of Extended Kata
func pick_RandomCard_Assign_fields_Extended() (promptField, objective, actual_objective_type string) { // - -
	randIndexE := rand.Intn(len(fileOfCardsE))
	aCard = fileOfCardsE[randIndexE] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kata
	objective = aCard.Romaji
	actual_objective_type = "Extended_Romaji" // Used to set a special prompt for Extended Kata
	whichDeck = 4
	return promptField, objective, actual_objective_type
}
*/
//
//
/*
// Drop "_Test", at the end of its name, & this func becomes ::: a short testing version of pick_RandomCard_Assign_fields()
func pick_RandomCard_Assign_fields_Test() (promptField, objective, actual_objective_type string) { // - -
	randIndex := rand.Intn(len(fileOfCardsMostDifficult))
	aCard = fileOfCardsMostDifficult[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Hira
	objective = aCard.Romaji
	actual_objective_type = "Romaji"
	whichDeck = 3
	return promptField, objective, actual_objective_type
}
*/
