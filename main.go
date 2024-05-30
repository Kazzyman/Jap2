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
	theGameIsRunning = false
	guessLevelCounter = 1
	limitedToHiraPrompts = true
	gottenHonestly = true
	fmt.Println()
	countSLOC()                          // Determine and display Source Lines Of Code.
	gotLastCardRightSoGetFreshOne = true // This will get us an initial card.

	display_start_menu_etc()
	begin_Kana_practice()
}

func begin_Kana_practice() { // ::: - -
	for {
		if gotLastCardRightSoGetFreshOne {
			pick_RandomCard_Assign_fields()
			guessLevelCounter = 1
		}

		// Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
		prompt_the_user_for_input() // each time we do this we increment the guessLevelCounter.
		// Obtain users input.
		_, _ = fmt.Scan(&usersSubmission)

		// Do you want to play a game?
		if usersSubmission == "game" {
			guessLevelCounter = 1
			fmt.Println("Welcome to the game. You may end the game early via a submission of off or goff")
			fmt.Println("What is your name?")
			_, _ = fmt.Scan(&nameOfPlayer)
			fmt.Println("Enter a number for how many prompts there will be in the game")
			_, _ = fmt.Scan(&game_duration_set_by_user)
			now_using_game_duration_set_by_user = true
			the_game_begins()
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
*/

func priorToProcessingUsersSubmission_check_IfTypeEnteredRightly() { // ::: - -
	if actual_objective_type == "roma" {
		// ::: Determine if the user has entered a valid Romaji char instead of, accidentally, a Hiragana char or string. If so, advise user.
		var isAlphanumeric bool
		findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)

		switch true {
		case findAlphasIn.MatchString(usersSubmission): // if this returns true
			isAlphanumeric = true // then do this
		default:
			isAlphanumeric = false // else do this.
		}
		/*
		   .
		*/
		switch isAlphanumeric {
		case isAlphanumeric: // if isAlphanumeric was set true above :: if type_of_usersSubmission == actual_objective_type ...
			Process_users_input_as_a_guess() // then do this ...
		default:
			// ... else display a message informing the user that he should change his input method.
			fmt.Println(colorRed)
			fmt.Println("Please change your input method to match the char type that was requested:)")
			fmt.Printf("Requested type being: %s\n", actual_objective_type)
			fmt.Println(colorReset)
		}

	} else if actual_objective_type == "hira" {
		// ::: Determine if the user has entered a valid Hiragana char instead of, accidentally, an alpha char or string. If so, advise user.
		var isAlphanumeric bool
		findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
		switch true {
		case findAlphasIn.MatchString(usersSubmission):
			isAlphanumeric = true
		default:
			isAlphanumeric = false
		}
		if isAlphanumeric == false { // as will be the case with a Hiragana char. :: if type_of_usersSubmission == actual_objective_type ...
			Process_users_input_as_a_guess() // then do this ...
		} else {
			// Display a message informing the user that he should change his input method.
			fmt.Println(colorRed)
			fmt.Println("Please change your input method to match the char type that was requested:)")
			fmt.Printf("Requested type being: %s\n", actual_objective_type)
			fmt.Println(colorReset)
		}
	}
}

/*
.
*/

func Process_users_input_as_a_guess() { // ::: - -
	// ::: Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	if actual_objective == "zu" {
		if usersSubmission == "zu" {

			gotLastCardRightSoGetFreshOne = true
			submission_already_processed_above = true // ::: because we are doing it now.

			logRight_zu(usersSubmission, actual_prompt_char, actual_objective_type)
		} else {
			gotLastCardRightSoGetFreshOne = false
			// logging etc. of Oops is being done in log_oops() by way of prompt_the_user_for_input()
		}
	} else {
		// todo]  Note: that "else", here, means that actual_objective != "zu"  ... but the actual_objective may yet be ず or づ.
		if actual_objective == "ず" || actual_objective == "づ" {
			if usersSubmission == "づ" || usersSubmission == "ず" {

				gotLastCardRightSoGetFreshOne = true
				submission_already_processed_above = true // ::: because we are doing it now.

				logRightZu2(usersSubmission, actual_prompt_char, actual_objective_type, actual_objective)
			} else {
				gotLastCardRightSoGetFreshOne = false
				// logging etc. of Oops is being done in log_oops() by way of prompt_the_user_for_input()
			}
		}
	} // If the objective was any form of zu,  it has been processed above.
	/*
	   ===
	   ===
	*/
	// ::: Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
	if submission_already_processed_above == true { // If it was handled as a zu case above ...
		// ... then there is nothing to do, else
	} else {
		if usersSubmission == actual_objective { // ::: Right
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
	log_oops_andUpdateGame(aCard.Hira, aCard.Romaji, userInput)
	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
}
func log_oops_andUpdateGame(prompt_it_was, field_it_was, guess string) { // - -
	if theGameIsRunning {
		failedOnThirdAttemptAccumulator++
	}
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was)
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

func display_start_menu_etc() { // ::: - -
	display_List_of_Directives()
}
