package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

// todo Bug::: stcr still has an issue. Though it does not crash the app if it is run immediately after an stc ...
// ... it is actually recoverable if it is run after an stc.
var thisIsOurFirstRodeo bool

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time in nanoseconds.
	kata_roma = true
	gottenHonestly = true // default is to assume honesty, I guess
	// theGameIsRunning = true
	guessedLastCardCorrectlySoGetFreshOne = true // This will get us an initial card.
	guessLevelCounter = 1

	fmt.Println()
	countSLOC() // Determine and display Source Lines Of Code.

	display_start_menu_etc()

	fmt.Println("What is your first name? (one word)")
	_, _ = fmt.Scan(&nameOfPlayer)
	List_of_game_types()
	fmt.Println(colorRed + "Select a game" + colorReset)
	_, _ = fmt.Scan(&type_of_game)
	switch type_of_game {
	case "1": // hko
		gameDuration = 2*len(fileOfCardsHiraKata) + (2 * len(fileOfCardsEasyKanji)) + len(fileOfCardsKanjiHard)
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = true
		limitedToHiraPrompts = true
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "2": // konly
		gameDuration = 105
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "3": // honly
		gameDuration = 105
		kata_hira = false
		kata_roma = false
		limitedToHiraPrompts = true
		limitedToKataPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "4": // ronly
		gameDuration = 105
		kata_hira = false
		kata_roma = false
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "5": // khSimplex
		gameDuration = 55
		kata_roma = false
		kata_hira = false
		limitedToKataPromptsAndSimplexHiraObj = true
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "6": // donly
		gameDuration = 20
		kata_hira = false
		kata_roma = false
		limitedToDifficultDescriptions = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = true
		limitedToSpelling = false
	case "7": // kh
		gameDuration = 106
		kata_roma = false
		kata_hira = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "8": // rhSimplex
		gameDuration = 55
		kata_hira = false
		kata_roma = false
		limitedToRomaPromptsAndSimplexHiraObj = true
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "9": // kr
		gameDuration = 105
		kata_hira = false
		kata_roma = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	default: // "0" // mix ::: NOT SURE ABOUT ALL THESE FALSE SETTINGS FOR THE MIX CASE ??
		gameDuration = 206
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	}
	display_limited_gaming_dir_list()
	now_using_game_duration_set_by_game_type = true
	the_game_begins()
	begin_Kana_practice()
}

func begin_Kana_practice() { // ::: - -
	for {
		if guessedLastCardCorrectlySoGetFreshOne {
			pick_RandomCard_Assign_fields() // This has within it lots of code to assure novel and fresh cards.
			guessLevelCounter = 1
		}

		// Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
		prompt_the_user_for_input() // Each time we do this we increment the guessLevelCounter.
		// Obtain users input.
		_, _ = fmt.Scan(&usersSubmission)

		if theGameIsRunning {
			// Skip looking to start a game if one is already running
			// During gaming, disallow checking for irrelevant Directives, and substitute game versions of certain directives

			switch usersSubmission {
			case "abt":
				about_app()
				// During gameplay, directive handling must not be counted as guesses, hence the decrementing of guessLevelCounter
				guessLevelCounter--
			case "nts":
				notes_on_kana()
				// During gameplay, directive handling must not be counted as guesses, hence the decrementing of guessLevelCounter
				guessLevelCounter--
			case "q":
				the_game_ends(true)
				// During gameplay, directive handling must not be counted as guesses, hence the decrementing of guessLevelCounter
				guessLevelCounter--
			case "rm":
				read_map_of_fineOn()
				read_map_of_needWorkOn()
				read_pulledButNotUsedMap()
				// read_pulls_not_used_array()
				// During gameplay, directive handling must not be counted as guesses, hence the decrementing of guessLevelCounter
				guessLevelCounter--
			case "st":
				st_stats_function()
				// During gameplay, directive handling must not be counted as guesses, hence the decrementing of guessLevelCounter
				guessLevelCounter--
			case "game":
				List_of_game_types()
				fmt.Printf("You are currently playing game number %s\n", type_of_game)
				// During gameplay, directive handling must not be counted as guesses, hence the decrementing of guessLevelCounter
				guessLevelCounter--
			case "dir":
				display_limited_gaming_dir_list()
				// During gameplay, directive handling must not be counted as guesses, hence the decrementing of guessLevelCounter
				guessLevelCounter--
			}

		} else if usersSubmission == "game" {
			reset_all_data(true)
			guessLevelCounter = 1
			fmt.Println("What is your first name? (one word)")
			_, _ = fmt.Scan(&nameOfPlayer)
			List_of_game_types()
			fmt.Println(colorRed + "Select a game" + colorReset)
			_, _ = fmt.Scan(&type_of_game)
			switch type_of_game {
			case "1": // hko
				gameDuration = 2*len(fileOfCardsHiraKata) + (2 * len(fileOfCardsEasyKanji)) + len(fileOfCardsKanjiHard)
				kata_hira = false
				kata_roma = false
				limitedToKataPrompts = true
				limitedToHiraPrompts = true
				limitedToRomaPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			case "2": // konly
				gameDuration = 105
				kata_hira = false
				kata_roma = false
				limitedToKataPrompts = true
				limitedToHiraPrompts = false
				limitedToRomaPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			case "3": // honly
				gameDuration = 105
				kata_hira = false
				kata_roma = false
				limitedToHiraPrompts = true
				limitedToKataPrompts = false
				limitedToRomaPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			case "4": // ronly
				gameDuration = 105
				kata_hira = false
				kata_roma = false
				limitedToRomaPrompts = true
				limitedToKataPrompts = false
				limitedToHiraPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			case "5": // khSimplex
				gameDuration = 55
				kata_roma = false
				kata_hira = false
				limitedToKataPromptsAndSimplexHiraObj = true
				limitedToKataPrompts = true
				limitedToHiraPrompts = false
				limitedToRomaPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			case "6": // donly
				gameDuration = 20
				kata_hira = false
				kata_roma = false
				limitedToDifficultDescriptions = true
				limitedToKataPrompts = false
				limitedToHiraPrompts = false
				limitedToRomaPrompts = true
				limitedToSpelling = false
			case "7": // kh
				gameDuration = 106
				kata_roma = false
				kata_hira = true
				limitedToKataPrompts = false
				limitedToHiraPrompts = false
				limitedToRomaPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			case "8": // rhSimplex
				gameDuration = 55
				kata_hira = false
				kata_roma = false
				limitedToRomaPromptsAndSimplexHiraObj = true
				limitedToRomaPrompts = true
				limitedToKataPrompts = false
				limitedToHiraPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			case "9": // kr
				gameDuration = 105
				kata_hira = false
				kata_roma = true
				limitedToKataPrompts = false
				limitedToHiraPrompts = false
				limitedToRomaPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			default: // "0" // mix ::: NOT SURE ABOUT ALL THESE FALSE SETTINGS FOR THE MIX CASE ??
				gameDuration = 206
				kata_hira = false
				kata_roma = false
				limitedToKataPrompts = false
				limitedToHiraPrompts = false
				limitedToRomaPrompts = false
				limitedToDifficultDescriptions = false
				limitedToSpelling = false
			}
			display_limited_gaming_dir_list()
			now_using_game_duration_set_by_game_type = true
			the_game_begins()

		} else {
			// If user's input is a Directive, handle it. ::: ONLY if a game IS NOT running !!!
			// todo 1 : if a game is running the following handler need not address the case of a "dir" ???
			// todo 2 : ... but since this is complex due to potential recursions I will leave it as it is.
			frontEnd_Possible_Recursive_DirHandler() // ::: this contains the only other prompt
		}
		// Process what can now be assumed to be an actual guess.
		priorToProcessingUsersSubmission_check_IfTypeEnteredRightly()
	}
}

/*
.
.
*/

func priorToProcessingUsersSubmission_check_IfTypeEnteredRightly() { // ::: - -
	if actual_objective_type == "roma" {
		// ::: Determine if the user has entered a valid Romaji char instead of, accidentally, a Hiragana char or string. If so, advise user.
		var isAlphanumeric bool
		findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)

		switch true {
		case findAlphasIn.MatchString(usersSubmission): // if this returns true
			isAlphanumeric = true
		default:
			isAlphanumeric = false
		}

		if isAlphanumeric {
			Process_users_input_as_a_guess()
		} else {
			if theGameIsRunning && usersSubmission == "?" || theGameIsRunning && usersSubmission == "??" {
				fmt.Println("Sorry, but not help is allowed during game play")

				// RE-Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
				guessLevelCounter--
				prompt_the_user_for_input() // Each time we do this we increment the guessLevelCounter.
				// Obtain users input.
				_, _ = fmt.Scan(&usersSubmission)
				priorToProcessingUsersSubmission_check_IfTypeEnteredRightly()
			} else {
				// Display a message informing the user that he should change his input method.
				fmt.Println(colorRed)
				fmt.Println("Please change your input method to match the char type that was requested:)")
				fmt.Printf("Requested type being: %s\n", actual_objective_type)
				fmt.Println(colorReset)

				// RE-Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
				prompt_the_user_for_input() // Each time we do this we increment the guessLevelCounter.
				// Obtain users input.
				_, _ = fmt.Scan(&usersSubmission)
				priorToProcessingUsersSubmission_check_IfTypeEnteredRightly()
			}

		}

	} else if actual_objective_type == "hira" {
		// ::: Determine if the user has entered a valid Hiragana char instead of, accidentally, an alpha char or string. If so, advise user.
		var isAlphanumeric bool
		findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)

		switch true {
		case findAlphasIn.MatchString(usersSubmission): // if this returns true
			isAlphanumeric = true
		default:
			isAlphanumeric = false
		}

		if isAlphanumeric == false { // as will be the case with a Hiragana char. :: if type_of_usersSubmission == actual_objective_type ...
			Process_users_input_as_a_guess()
		} else {
			// Display a message informing the user that he should change his input method.
			fmt.Println(colorRed)
			fmt.Println("Please change your input method to match the char type that was requested:)")
			fmt.Printf("Requested type being: %s\n", actual_objective_type)
			fmt.Println(colorReset)
			// RE-Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
			prompt_the_user_for_input() // Each time we do this we increment the guessLevelCounter.
			// Obtain users input.
			_, _ = fmt.Scan(&usersSubmission)
			priorToProcessingUsersSubmission_check_IfTypeEnteredRightly()
		}
	}
}

/*
.
*/

func Process_users_input_as_a_guess() { // ::: - -
	// ::: Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	// Since actual_objective can only ever be of roma or hira type, those two we now do.
	// ... First the roma:
	if actual_objective == "zu" {
		if usersSubmission == "zu" {
			if guessLevelCounter == 3 {
				fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check_error(err)
				_, err1 := fmt.Fprintf(fileHandle,
					"\nUser may have mistyped==%s:%s:%s, quarter point deduction", aCard.Romaji, aCard.Hira, aCard.Kata) // mistyped is a word?
				check_error(err1)
			} else if guessLevelCounter == 4 {
				fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check_error(err)
				_, err1 := fmt.Fprintf(fileHandle,
					"\nUser had a some difficulty with==%s:%s:%s, half point deduction", aCard.Romaji, aCard.Hira, aCard.Kata)
				check_error(err1)
			}
			guessedLastCardCorrectlySoGetFreshOne = true
			submission_already_processed_above = true // ::: because we are doing it now.

			logRight_zu(usersSubmission, actual_prompt_char, actual_objective_type)
		} else {
			guessedLastCardCorrectlySoGetFreshOne = false
			// logging etc. of Oops is being done in log_oops() by way of prompt_the_user_for_input()
		}
		// ... and Then the hira: (and in the case of zu, there are two hira forms of zu)
	} else {
		// todo]  Note: that "else", here, means that actual_objective != "zu"  ... but the actual_objective may yet be ず or づ.
		if actual_objective == "ず" || actual_objective == "づ" {
			if usersSubmission == "づ" || usersSubmission == "ず" {
				if guessLevelCounter == 3 {
					fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
					check_error(err)
					_, err1 := fmt.Fprintf(fileHandle,
						"\nUser may have mistyped==%s:%s:%s, quarter point deduciton", aCard.Romaji, aCard.Hira, aCard.Kata) // mistyped is a word?
					check_error(err1)
				} else if guessLevelCounter == 4 {
					fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
					check_error(err)
					_, err1 := fmt.Fprintf(fileHandle,
						"\nUser had a some difficulty with==%s:%s:%s, half point deduction", aCard.Romaji, aCard.Hira, aCard.Kata)
					check_error(err1)
				}
				guessedLastCardCorrectlySoGetFreshOne = true
				submission_already_processed_above = true // ::: because we are doing it now.

				logRightZu2(usersSubmission, actual_prompt_char, actual_objective_type, actual_objective)
			} else {
				guessedLastCardCorrectlySoGetFreshOne = false
				// logging etc. of Oops is being done in log_oops() by way of prompt_the_user_for_input()
			}
		}
	} // If the objective was any form of zu,  it has already been processed above.
	/*
	   ===
	   ===
	*/
	// ::: Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
	if submission_already_processed_above == true { // If it was handled as a zu case above ...
		// ... it was a form of zu, and there is nothing more to do, else
	} else {
		// Write to a file:
		if usersSubmission == actual_objective {
			if guessLevelCounter == 3 {
				fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check_error(err)
				_, err1 := fmt.Fprintf(fileHandle,
					"\nUser may have mistyped==%s:%s:%s, quarter point deduction", aCard.Romaji, aCard.Hira, aCard.Kata) // mistyped is a word?
				check_error(err1)
			} else if guessLevelCounter == 4 {
				fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				check_error(err)
				_, err1 := fmt.Fprintf(fileHandle,
					"\nUser had a some difficulty with==%s:%s:%s, half point deduction", aCard.Romaji, aCard.Hira, aCard.Kata)
				check_error(err1)
			}
			// These two lines are all we do when a match occurs.
			guessedLastCardCorrectlySoGetFreshOne = true
			displayRight_logRight(usersSubmission, actual_prompt_char, actual_objective_type)
		} else {
			guessedLastCardCorrectlySoGetFreshOne = false
			// logging etc. of Oops is being done in log_oops() by way of prompt_the_user_for_input()
		}
	}
	submission_already_processed_above = false // So, reset it for the next round.

}

/*
.
*/

func frontEnd_Possible_Recursive_DirHandler() { // ::: - -
	detectDirective(usersSubmission)
	if its_a_directive {
		its_a_directive = false

		respond_to_UserSupplied_Directive(usersSubmission)

		prompt_the_user_for_input() // re-prompt using same card as before.
		_, _ = fmt.Scan(&usersSubmission)

		detectDirective(usersSubmission)
		if its_a_directive {
			its_a_directive = false
			frontEnd_Possible_Recursive_DirHandler() // ::: Recursion
		} else {
			return // Need this to get out of the recursive func.
		}
	}
}

/*
.
*/

func display_failure_of_final_guess_message_etc(userInput string) { // ::: - -
	// log_oops_andUpdateGame(aCard.Hira, aCard.Romaji, userInput) // this needs fixing, as the first string is not always a Hira
	if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
		// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Romaji)
		log_oops_andUpdateGame(aCard.Romaji, aCard.Romaji, userInput) // ::: this may also need fixing, as the second string may not always be Romaji
	} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
		// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Hira)
		log_oops_andUpdateGame(aCard.Hira, aCard.Romaji, userInput) // ::: this may also need fixing, as the second string may not always be Romaji
	} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
		// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Kata)
		log_oops_andUpdateGame(aCard.Kata, aCard.Romaji, userInput) // ::: this may also need fixing, as the second string may not always be Romaji
	} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" {
		// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Kata)
		log_oops_andUpdateGame(aCard.Kata, aCard.Romaji, userInput) // ::: this may also need fixing, as the second string may not always be Romaji
	}

	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)
	_, err1 := fmt.Fprintf(fileHandle,
		"\nUser had a REAL ISSUE with==%s:%s:%s, two point deduction ", aCard.Romaji, aCard.Hira, aCard.Kata)
	check_error(err1)
}
func log_oops_andUpdateGame(prompt_it_was, field_it_was, guess string) { // - -
	if theGameIsRunning {
		gameFailedOnThirdAttemptAccumulator++
	}
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was) // used to be the only instance of this func being called
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

func display_start_menu_etc() { // ::: - -
	display_List_of_Directives()
}
