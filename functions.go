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
	cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
	cyclicArrayHits = CyclicArrayHits{}
	cyclicArrayPulls = CyclicArrayPulls{}
	// Also, flush (clear) the maps
	hiraHitMap = make(map[string]CardInfo)
	frequencyMapOf_IsFineOnChars = make(map[string]int)
	frequencyMapOf_need_workOn = make(map[string]int)
	pulledButNotUsedMap = make(map[string]int)
	total_prompts = 0
	gameCorrectOnThirdAttemptAccumulator = 0
	game_loop_counter = 0
	guessLevelCounter = 1
	theGameIsRunning = true // ::: this flag is the only thing that "starts" a game
	// gameCorrectOnFirstAttemptAccumulator = 1   // ::: here it is/was not able to process the last guess prior to game ending.
	// gameCorrectOnSecondAttemptAccumulator = -1 // ::: kluge !!
	// ::: if the first query of a game is gotten right on the first attempt, it is logged as a 2nd and not a 1st
	gameCorrectOnFirstAttemptAccumulator = 0
	gameCorrectOnSecondAttemptAccumulator = 0
	gameFailedOnThirdAttemptAccumulator = 0

	//
	// ::: file-writing and time-stamping is all that follows ==========================================
	currentTime := time.Now()
	TimeOfStartFromInceptionOfGame = time.Now()

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	// unneeded section
	var kind string
	kind = "game"
	if kata_roma {
		kind = "kata_roma"
	}
	if kata_hira {
		kind = "kata_hira"
	}

	_, err1 := fmt.Fprintf(fileHandle,
		"\n\n-- A %s/%s, %s began at: %s",
		actual_prompt_char_type, actual_objective_type, kind, currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
}
func the_game_ends(suppressAllMessages, do_not_reset_gaming_flag, suppressPointsReporting bool) { // ::: - -
	if do_not_reset_gaming_flag {
		// do nada
	} else {
		theGameIsRunning = false
	}
	now_using_game_duration_set_by_game_type = false
	// gameDuration = 0
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

	total_seconds := (minutes * 60) + (seconds)
	firstAtemptAcumF := float32(gameCorrectOnFirstAttemptAccumulator - 1)
	secondAtemptAcumF := float32(gameCorrectOnSecondAttemptAccumulator)
	thirdAtemptAcumF := float32(gameCorrectOnThirdAttemptAccumulator)
	failedOnThirdAttF := float32(gameFailedOnThirdAttemptAccumulator)
	totalSecondsF := float32(total_seconds)
	points2print = (firstAtemptAcumF - (secondAtemptAcumF / 4) - (thirdAtemptAcumF / 2) -
		(failedOnThirdAttF * 2)) / totalSecondsF * 100

	// Create the formatted string
	TotalRun := fmt.Sprintf("%02d:%02d", minutes, seconds)

	// ::: ------------------------------------------------------------------- v v v v
	// Process_users_input_as_a_guess() // ::: trying this ---- and these ------- v v v v
	if guessLevelCounter == 2 {
		if gottenHonestly {
			gameCorrectOnFirstAttemptAccumulator++
			gottenHonestly = false
		}
	} else if guessLevelCounter == 3 {
		gameCorrectOnSecondAttemptAccumulator++
	} else { // ::: never/rarely executes ????
		// gameCorrectOnThirdAttemptAccumulator++
		// fmt.Printf("guessLevelCounter is:%d, where it maybe-should-be-4-? \n", guessLevelCounter)
		// ... then ... the guessLevelCounter was 4?.
		// gameCorrectOnThirdAttemptAccumulator++
		// ::: fail/error accumulator gets incremented (or at least it gets displayed [below]) during the oops message
	} // - - - - - - - - - - - - - - - - - - - - - - - - - ::: -----------------------------------------------------------------
	if suppressAllMessages {
		// do dada
	} else {
		// suppressPointsReporting = true // this is not actually set as true anywhere in the project // todo: fix that
		fmt.Printf("\nThe status of suppressPointsReporting is %t", suppressPointsReporting)
		if gameCorrectOnFirstAttemptAccumulator > 0 && gameCorrectOnSecondAttemptAccumulator > 0 && gameCorrectOnThirdAttemptAccumulator > 0 && gameFailedOnThirdAttemptAccumulator == 0 { // ::: done
			fmt.Println(colorRed)
			fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d%s right on your second try,\n"+
				"... and you got %s%d right on your third try. \n\n", TotalRun, colorReset, gameCorrectOnFirstAttemptAccumulator-1, colorRed, colorReset, gameCorrectOnSecondAttemptAccumulator,
				colorRed, colorReset, gameCorrectOnThirdAttemptAccumulator)
			if gameDuration == 2*len(fileOfCardsHiraKata)+(2*len(fileOfCardsEasyKanji))+len(fileOfCardsKanjiHard) && !suppressPointsReporting {
				// todo : also calculate game scores for games of other lengths
				fmt.Printf("Points = %f\n", (firstAtemptAcumF-(secondAtemptAcumF/4)-(thirdAtemptAcumF/2)-(failedOnThirdAttF*2))/totalSecondsF*100)
			}
		} else if gameCorrectOnFirstAttemptAccumulator > 0 && gameCorrectOnSecondAttemptAccumulator > 0 && gameCorrectOnThirdAttemptAccumulator == 0 && gameFailedOnThirdAttemptAccumulator == 0 { // ::: done
			fmt.Println(colorRed)
			fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d right on your second try. \n\n", TotalRun, colorReset, gameCorrectOnFirstAttemptAccumulator-1,
				colorRed, colorReset, gameCorrectOnSecondAttemptAccumulator)
			if gameDuration == 2*len(fileOfCardsHiraKata)+(2*len(fileOfCardsEasyKanji))+len(fileOfCardsKanjiHard) && !suppressPointsReporting {
				// todo : also calculate game scores for games of other lengths
				fmt.Printf("Points = %f\n", (firstAtemptAcumF-(secondAtemptAcumF/4)-(thirdAtemptAcumF/2)-(failedOnThirdAttF*2))/totalSecondsF*100)
			}
		} else if gameCorrectOnFirstAttemptAccumulator > 0 && gameCorrectOnSecondAttemptAccumulator == 0 && gameCorrectOnThirdAttemptAccumulator == 0 && gameFailedOnThirdAttemptAccumulator == 0 { // ::: done
			fmt.Println(colorRed)
			fmt.Printf("\nYour Game run-time was:%s,  Gongratulations! you got %s%d correct on your first try. \n\n", TotalRun, colorReset, gameCorrectOnFirstAttemptAccumulator-1)
			if gameDuration == 2*len(fileOfCardsHiraKata)+(2*len(fileOfCardsEasyKanji))+len(fileOfCardsKanjiHard) && !suppressPointsReporting {
				// todo : also calculate game scores for games of other lengths
				fmt.Printf("Points = %f\n", (firstAtemptAcumF-(secondAtemptAcumF/4)-(thirdAtemptAcumF/2)-(failedOnThirdAttF*2))/totalSecondsF*100)
			}
		} else {
			fmt.Println(colorRed)
			fmt.Printf("\nYour Game run-time was:%s,  you got %s%d%s correct on your first try,  %s%d%s right on your second try,\n"+
				"... and you got %s%d%s right on your third try, and were unable to answer correctly without a hint "+
				"%s%d times. \n\n", TotalRun, colorReset, gameCorrectOnFirstAttemptAccumulator-1, colorRed, colorReset, gameCorrectOnSecondAttemptAccumulator,
				colorRed, colorReset, gameCorrectOnThirdAttemptAccumulator, colorRed, colorReset, gameFailedOnThirdAttemptAccumulator)
			if gameDuration == 2*len(fileOfCardsHiraKata)+(2*len(fileOfCardsEasyKanji))+len(fileOfCardsKanjiHard) && !suppressPointsReporting {
				// todo : also calculate game scores for games of other lengths
				fmt.Printf("Points = %f\n", (firstAtemptAcumF-(secondAtemptAcumF/4)-(thirdAtemptAcumF/2)-(failedOnThirdAttF*2))/totalSecondsF*100)
			}
		}

		fmt.Printf(" --- View the text log file for further comments re this or previous game stats --- \n\n")

	}

	// End timer and report elapsed time and other stats to a file.
	_, err1 := fmt.Fprintf(fileHandle,
		"\n-- A game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), game_loop_counter-1)
	check_error(err1)

	_, err3 := fmt.Fprintf(fileHandle, "%s's results were as follows: Right on first attempt:%d, on 2nd attempt:%d, 3rd attempt:%d, even a hint was ineffective:%d, %d/%d\n",
		nameOfPlayer, gameCorrectOnFirstAttemptAccumulator-1,
		gameCorrectOnSecondAttemptAccumulator, gameCorrectOnThirdAttemptAccumulator,
		gameFailedOnThirdAttemptAccumulator, game_loop_counter-1, gameDuration)
	check_error(err3)

	_, err2 := fmt.Fprintf(fileHandle,
		"The Elapsed time of the game was: %s \n",
		TotalRun)
	check_error(err2)
	// todo : also calculate game scores for games of other lengths
	//	if gameDuration == 2*len(fileOfCardsHiraKata)+(2*len(fileOfCardsEasyKanji)) && !suppressPointsReporting { // Calculate and print the Point total to the log file only if a full game has been played
	// If it was a full length game, |or| if we are overriding that test:
	if gameDuration == 2*len(fileOfCardsHiraKata)+(2*len(fileOfCardsEasyKanji))+len(fileOfCardsKanjiHard) || !suppressPointsReporting {
		// fmt.Printf("Points: %f", points2print)
		_, err3 := fmt.Fprintf(fileHandle,
			"Points: %f \n\n\n",
			points2print)
		check_error(err3)
	}
}

/*
.
*/
func read_pulledButNotUsedMap() { // ::: - -
	if len(pulledButNotUsedMap) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe seenMap (pulledButNotUsedMap) is empty\n")
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
	fmt.Printf("\nThis app consists of the following lines of code across %d files:\n", fileExplored)
	fmt.Println("   main.go, constants.go, elementsOfsloc.go, functions.go, globalVariables.go, \n" +
		"   locateCard.go, memoryFunctions.go, objectsAndMethods.go, pick_a_card_functions.go, \n" +
		"   pick_a_card_random_all.go, prompts&directions.go, statsFunctions.go, functionsFromMain.go\n")
	countSLOC()
	fmt.Printf("As Calculated in real-time by countSLOC(), a custom internal function.\n\n")
}

/*
.
*/
// Flush (clear) the old stats and hits arrays
func reset_all_data(do_not_reset_gaming_flag, suppressPrinting bool) { // ::: - -
	limitedToKataPrompts = false
	limitedToHiraPrompts = false
	limitedToRomaPrompts = false
	limitedToDifficultDescriptions = false
	include_Extended_kata_deck = false
	//
	kata_roma = true // ::: the "default"
	//
	cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
	cyclicArrayHits = CyclicArrayHits{}
	cyclicArrayPulls = CyclicArrayPulls{}
	// Also, flush (clear) the maps
	hiraHitMap = make(map[string]CardInfo)
	frequencyMapOf_IsFineOnChars = make(map[string]int)
	frequencyMapOf_need_workOn = make(map[string]int)
	pulledButNotUsedMap = make(map[string]int)
	total_prompts = 0
	gameCorrectOnFirstAttemptAccumulator = 0
	gameCorrectOnSecondAttemptAccumulator = 0
	gameCorrectOnThirdAttemptAccumulator = 0
	gameFailedOnThirdAttemptAccumulator = 0
	game_loop_counter = 0
	// jim = 15
	if do_not_reset_gaming_flag {
		// do nada
	} else {
		theGameIsRunning = false
	}

	total_prompts = 0
	//
	//goland:noinspection ALL
	if suppressPrinting {
		// do not print
	} else {
		fmt.Println("\nArrays and maps flushed:\n")
		fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
		fmt.Println("    cyclicArrayHits")
		fmt.Println("    cyclicArrayPulls and hiraHitMap and pulledButNotUsedMap")
		fmt.Println("    frequencyMapOf_IsFineOnChars")
		//goland:noinspection ALL
		fmt.Println("    frequencyMapOf_need_workOn\n")
		fmt.Println(colorCyan + "  Limitations re Kata, Hira, and Romaji prompting; as well as all Game values have also been reset\n" + colorReset)
		fmt.Println("... and will take effect after you dispatch the current card: default kata_roma ...")

	}
}

/*
.
*/
func notes_on_kana() { // ::: - -
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
		in == "STC" ||
		in == "stcr" ||
		in == "STCR" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "rs" ||
		in == "RS" ||
		in == "st" ||
		in == "ST" ||
		in == "dir" ||
		in == "DIR" ||
		in == "nts" ||
		in == "q" ||
		in == "rm" ||
		in == "abt" ||
		in == "exko" ||
		in == "exkf" ||
		in == "konly" ||
		in == "honly" ||
		in == "ronly" ||
		in == "rhSimplex" ||
		in == "donly" ||
		in == "hko" ||
		in == "kh" ||
		in == "khSimplex" ||
		in == "kr" ||
		in == "game" ||
		in == "spell" ||
		in == "mix" ||
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
	guessLevelCounter--
	switch usersSubmission {
	// Directives follow:
	// ::: alphabetically (mostly)
	case "?":
		if limitedToDifficultDescriptions {
			fmt.Printf("\n%s\n\n", aCard.SansR_Hint)
		} else {
			fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
		}
	case "??":
		// handle_doubleQuestMark_directive(actual_objective_type)
		handle_doubleQuestMark_directive()
	case "help":
		helpText()

	case "spell":
		kata_roma = true // Because, we prompt from the kata field and solicit a Romaji response
		kata_hira = false
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = true
		fmt.Printf("-- 	Your settings will go into effect after you dispence with the present card ...\n")
	case "kh":
		kata_roma = false
		kata_hira = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
	case "khSimplex":
		kata_roma = false
		kata_hira = false
		limitedToKataPromptsAndSimplexHiraObj = true
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
	case "kr":
		kata_hira = false
		kata_roma = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")

	case "hko":
		gameDuration = 2*len(fileOfCardsHiraKata) + (2 * len(fileOfCardsEasyKanji)) + len(fileOfCardsKanjiHard)
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = true
		limitedToHiraPrompts = true
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
	case "konly":
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
	case "honly":
		kata_hira = false
		kata_roma = false
		limitedToHiraPrompts = true
		limitedToKataPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
	case "ronly":
		kata_hira = false
		kata_roma = false
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
	case "rhSimplex":
		kata_hira = false
		kata_roma = false
		limitedToRomaPromptsAndSimplexHiraObj = true
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
	case "donly":
		kata_hira = false
		kata_roma = false
		limitedToDifficultDescriptions = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = true
		limitedToSpelling = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
		/*
			.
		*/
	case "abt":
		about_app()

	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		if theGameIsRunning {
			display_limited_gaming_dir_list()
		} else {
			re_display_List_of_Directives()
		}
	case "DIR": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		if theGameIsRunning {
			display_limited_gaming_dir_list()
		} else {
			re_display_List_of_Directives()
		}
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
		reset_all_data(false, false)
	case "st":
		st_stats_function()
	case "stc":
		reSet_via_a_hira_aCard_andThereBy_reSet_thePromptString()
	case "stcr":
		reSet_aCard_via_a_romaji_andThereBy_reSet_thePromptString()
	case "game": // ::: this code reachable only via recursion, e.g., immediately following a "dir" directive command.
		reset_all_data(true, true) // I would like to suppress the prints in this func
		fmt.Println("What is your first name? (one word)")
		_, _ = fmt.Scan(&nameOfPlayer)
		List_of_game_types()
		fmt.Println(colorRed + "Enter the type of game" + colorReset)
		_, _ = fmt.Scan(&type_of_game)
		switch type_of_game {
		case "1": // hko Use Kata & Hira prompting (Roma objectives)
			gameDuration = 2*len(fileOfCardsHiraKata) + (2 * len(fileOfCardsEasyKanji)) + len(fileOfCardsKanjiHard)
			kata_hira = false
			kata_roma = false
			limitedToKataPrompts = true
			limitedToHiraPrompts = true
			limitedToRomaPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
		case "2": // konly Use only Kata prompting (mix Hira & Roma objectives)
			gameDuration = 2 * len(fileOfCardsHiraKata)
			kata_hira = false
			kata_roma = false
			limitedToKataPrompts = true
			limitedToHiraPrompts = false
			limitedToRomaPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
		case "3": // honly Use only Hira prompting (so only Roma responses)
			gameDuration = len(fileOfCardsHiraKata)
			kata_hira = false
			kata_roma = false
			limitedToHiraPrompts = true
			limitedToKataPrompts = false
			limitedToRomaPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
		case "4": // ronly Use only Romaji prompting (so only Hira responses)
			gameDuration = len(fileOfCardsHiraKata)
			kata_hira = false
			kata_roma = false
			limitedToRomaPrompts = true
			limitedToKataPrompts = false
			limitedToHiraPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
		case "5": // khSimplex : Use Kata prompting & Simplex Hira obj
			gameDuration = len(fileOfCards_nonCompound)
			kata_roma = false
			kata_hira = false
			limitedToKataPromptsAndSimplexHiraObj = true
			limitedToKataPrompts = true
			limitedToHiraPrompts = false
			limitedToRomaPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
		case "6": // donly : Difficult descriptive prompting
			gameDuration = len(dataMostDiff)
			kata_hira = false
			kata_roma = false
			limitedToDifficultDescriptions = true
			limitedToKataPrompts = false
			limitedToHiraPrompts = false
			limitedToRomaPrompts = true
			limitedToSpelling = false
			fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
		case "7": // kh : Use only kata_hira prompt_response
			gameDuration = 2 * len(fileOfCardsHiraKata)
			kata_roma = false
			kata_hira = true
			limitedToKataPrompts = false
			limitedToHiraPrompts = false
			limitedToRomaPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
		case "8": // rhSimplex : Use kata prompts w/Simplex Hira
			gameDuration = len(fileOfCards_nonCompound)
			kata_hira = false
			kata_roma = false
			limitedToRomaPromptsAndSimplexHiraObj = true
			limitedToRomaPrompts = true
			limitedToKataPrompts = false
			limitedToHiraPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")
		case "9": // kr : Use only kata_roma prompt_response
			gameDuration = len(fileOfCardsHiraKata)
			kata_hira = false
			kata_roma = true
			limitedToKataPrompts = false
			limitedToHiraPrompts = false
			limitedToRomaPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
		default: // "0" // mix ::: NOT SURE ABOUT ALL THESE FALSE SETTINGS FOR THE MIX CASE ??
			gameDuration = 3*len(fileOfCardsHiraKata) + (2 * len(fileOfCardsEasyKanji)) + len(fileOfCardsKanjiHard) // ::: 3 *
			kata_hira = false
			kata_roma = false
			limitedToKataPrompts = false
			limitedToHiraPrompts = false
			limitedToRomaPrompts = false
			limitedToDifficultDescriptions = false
			limitedToSpelling = false
			// ::: nothing limited means prompt with a mix of all three
			fmt.Printf("-- Your settings will go into effect after you dispence with the present card ...\n")
		}
		display_limited_gaming_dir_list()
		now_using_game_duration_set_by_game_type = true
		the_game_begins()
	case "mix":
		/*
			limitedToKataPrompts bool // These are used to control (limit) which field of the cards will be used as the prompt char.
			var limitedToHiraPrompts bool
			var limitedToRomaPrompts bool
			var limitedToRomaPromptsAndSimplexHiraObj bool
			var limitedToKataPromptsAndSimplexHiraObj bool
			var limitedToDifficultDescriptions bool
			var limitedToSpelling
		*/
		limitedToRomaPromptsAndSimplexHiraObj = false
		limitedToKataPromptsAndSimplexHiraObj = false
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		kata_roma = false
		kata_hira = false
		limitedToSpelling = false
		fmt.Printf("-- Your setting will go into effect after you dispence with the present card ...\n")

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

// ::: The following two test functions are obsolete due to use of "whichDeck" int
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
