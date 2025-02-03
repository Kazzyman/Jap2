package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
)

// todo Bug::: stcr still has an issue. Though it does not crash the app if it us run immediately after an stc ...
// ... it is actually recoverable if it is run after an stc.
var thisIsOurFirstRodeo bool

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time in nanoseconds.
	theGameIsRunning = true
	kata_roma = true
	guessLevelCounter = 1
	// limitedToKataPrompts = true
	// limitedToHiraPrompts = true
	gottenHonestly = true
	fmt.Println()
	countSLOC()                          // Determine and display Source Lines Of Code.
	gotLastCardRightSoGetFreshOne = true // This will get us an initial card.

	display_start_menu_etc()
	guessLevelCounter = 1

	fmt.Println("Welcome to the game. Dir options: off/goff, stc, stcr, q, dirg")
	fmt.Println("What is your first name? (one word)")
	_, _ = fmt.Scan(&nameOfPlayer)
	fmt.Println("Enter a number (102) for how many prompts there will be in the game")
	_, _ = fmt.Scan(&game_duration_set_by_user)
	List_of_game_types()
	fmt.Println(colorRed + "Enter the type of game" + colorReset)
	_, _ = fmt.Scan(&type_of_game)
	switch type_of_game {
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
	}
	display_limited_gaming_dir_list()
	now_using_game_duration_set_by_user = true
	the_game_begins()
	begin_Kana_practice()
}

func begin_Kana_practice() { // ::: - -
	for {
		if gotLastCardRightSoGetFreshOne {
			pick_RandomCard_Assign_fields() // This has within it lots of code to assure novel and fresh cards.
			guessLevelCounter = 1
		}

		// Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
		prompt_the_user_for_input() // Each time we do this we increment the guessLevelCounter.
		// Obtain users input.
		_, _ = fmt.Scan(&usersSubmission)

		if theGameIsRunning == true {
			// Skip looking to start a game if one is already running.
		} else if usersSubmission == "game" {
			guessLevelCounter = 1
			fmt.Println("Welcome to the game. Dir options: off/goff, stc, stcr, q, dirg")
			fmt.Println("What is your first name? (one word)")
			_, _ = fmt.Scan(&nameOfPlayer)
			fmt.Println("Enter a number (102) for how many prompts there will be in the game")
			_, _ = fmt.Scan(&game_duration_set_by_user)
			List_of_game_types()
			fmt.Println(colorRed + "Enter the type of game" + colorReset)
			_, _ = fmt.Scan(&type_of_game)
			switch type_of_game {
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
			}
			display_limited_gaming_dir_list()
			now_using_game_duration_set_by_user = true
			the_game_begins()
		}

		if theGameIsRunning != true {
			// skip looking to dirg since it only makes sense when a game is running.
		} else if usersSubmission == "dirg" {
			display_limited_gaming_dir_list()
		}

		if usersSubmission == "stc" {
			reSet_via_a_hira_aCard_andThereBy_reSet_thePromptString()
		} else if usersSubmission == "stcr" {
			reSet_aCard_via_a_romaji_andThereBy_reSet_thePromptString()
		}

		// During gaming, disallow checking for Directives other than q, and goff.
		if theGameIsRunning {
			if usersSubmission == "q" {
				os.Exit(1)
			}
			if usersSubmission == "off" || usersSubmission == "goff" {
				the_game_ends()
			}
		} else {
			// If user's input is a Directive, handle it.
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
			gotLastCardRightSoGetFreshOne = true
			submission_already_processed_above = true // ::: because we are doing it now.

			logRight_zu(usersSubmission, actual_prompt_char, actual_objective_type)
		} else {
			gotLastCardRightSoGetFreshOne = false
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
				gotLastCardRightSoGetFreshOne = true
				submission_already_processed_above = true // ::: because we are doing it now.

				logRightZu2(usersSubmission, actual_prompt_char, actual_objective_type, actual_objective)
			} else {
				gotLastCardRightSoGetFreshOne = false
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
			gotLastCardRightSoGetFreshOne = true
			displayRight_logRight(usersSubmission, actual_prompt_char, actual_objective_type)
		} else {
			gotLastCardRightSoGetFreshOne = false
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
		failedOnThirdAttemptAccumulator++
	}
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was) // used to be the only instance of this func being called
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

func display_start_menu_etc() { // ::: - -
	display_List_of_Directives()
}
