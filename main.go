package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// todo Bug::: stcr still has an issue. Though it does not crash the app if it us run immediately after an stc ...
// ... it is actually recoverable if it is run after an stc.
var thisIsOurFirstRodeo bool

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time in nanoseconds.
	gameOn = false
	guessLevelCounter = 1
	limitedToHiraPrompts = true
	fmt.Println()
	countSLOC()                     // Determine and display Source Lines Of Code.
	pick_RandomCard_Assign_fields() // Pick the first card.

	thisIsOurFirstRodeo = true
	user_guessed_prior_card_rightly = false

	display_start_menu_etc()
	begin_Kana_practice()
}

func begin_Kana_practice() { // ::: - -
	for { // ::: Application loop.

		// ::: debug off fmt.Printf("guessLevelCounter: %d, top of Kana_practice(), inside non_standard_origin...\n", guessLevelCounter)
		// ::: debug off fmt.Printf("at top of Kana_practice()\n")

		if user_guessed_prior_card_rightly || thisIsOurFirstRodeo {

			// ::: debug off fmt.Printf("passed the sniff test for getting a new card \n")

			thisIsOurFirstRodeo = false
			if no_interim_error_flag {
				correctOnFirstAttempt++
			}
			pick_RandomCard_Assign_fields()
			no_interim_error_flag = true

			guessLevelCounter = 1
		}

		// ::: debug off fmt.Printf("guessLevelCounter: %d, before first prompt in Kana_practice()\n", guessLevelCounter)

		// Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
		prompt_the_user_for_input()

		// ::: debug off fmt.Printf("guessLevelCounter: %d, after first prompt in Kana_practice()\n", guessLevelCounter)

		// Obtain users input.
		_, _ = fmt.Scan(&usersSubmission)

		// Categorize and process users input accordingly, e.g., is users input a Directive, etc.?
		// todo : remove : categorize_and_process_users_input() -- and the above comment.

		// If users input is a Directive, handle it.
		potentially_handle_a_directive()

		// ::: debug off fmt.Printf("guessLevelCounter: %d, after HD call in \"cate\" \n", guessLevelCounter)

		processUsersSubmissionIfTypeEnteredRightly()

		// ::: debug off fmt.Printf("guessLevelCounter: %d, after call to categorize_and... at bottom of Kana_practice()\n", guessLevelCounter)
		// Loop, will now get a fresh card and the process continues indefinitely until exit Directive is "q".
	}
}

func categorize_and_process_users_input() { // ::: - -

}

func processUsersSubmissionIfTypeEnteredRightly() { // ::: - -
	if actual_objective_type == "roma" {
		// Determine if the user has entered a valid Romaji char instead of, accidentally, a Hiragana char or string. Advise.
		var isAlphanumeric bool
		findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
		switch true {
		case findAlphasIn.MatchString(usersSubmission):
			isAlphanumeric = true
		default:
			isAlphanumeric = false
		}
		if isAlphanumeric == true {
			// ::: debug off fmt.Println("you entered an alpha")

			// if type_of_usersSubmission == actual_objective_type
			Process_users_input_as_a_guess()

			// ::: debug off fmt.Printf("guessLevelCounter: %d, after call to Process_users_input_as_a_guess()\n", guessLevelCounter)
		} else {
			// Display a message informing the user that he should change his input method.
			fmt.Println(colorRed)
			fmt.Println("Please change your input method to match the char type that was requested:)")
			fmt.Printf("Requested type being: %s\n", actual_objective_type)
			fmt.Println(colorReset)
		}
	}

	if actual_objective_type == "hira" {
		// Determine if the user has entered a valid Hiragana char instead of, accidentally, an alpha char or string. Advise.
		var isAlphanumeric bool
		findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
		switch true {
		case findAlphasIn.MatchString(usersSubmission):
			isAlphanumeric = true
		default:
			isAlphanumeric = false
		}
		if isAlphanumeric == false {
			// ::: debug off fmt.Println("you entered a Hiragana char")

			// if type_of_usersSubmission == actual_objective_type
			Process_users_input_as_a_guess()

			// ::: debug off fmt.Printf("guessLevelCounter: %d, after call to Process_users_input_as_a_guess()\n", guessLevelCounter)
		} else {
			// Display a message informing the user that he should change his input method.
			fmt.Println(colorRed)
			fmt.Println("Please change your input method to match the char type that was requested:)")
			fmt.Printf("Requested type being: %s\n", actual_objective_type)
			fmt.Println(colorReset)
		}
	}
}

func Process_users_input_as_a_guess() { // ::: - -
	// todo]  Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	if actual_objective == "zu" {
		submission_already_processed_above = true // todo:  because we are doing it now.
		if usersSubmission == "zu" {
			user_guessed_prior_card_rightly = true
			logRight_zu(usersSubmission, actual_prompt_char, actual_objective_type)
		} else {
			user_guessed_prior_card_rightly = false
			display_failure_of_final_guess_message_etc(usersSubmission)
		}
	} else if submission_already_processed_above != true { // todo: is this bool test even needed? Well, no; but it reads well.
		// todo]  Note: that "else", here, means that actual_objective != "zu"  ... but the actual_objective may yet be ず or づ.
		if actual_objective == "ず" || actual_objective == "づ" {
			if usersSubmission == "づ" || usersSubmission == "ず" {
				submission_already_processed_above = true // todo: because we are doing it now.
				user_guessed_prior_card_rightly = true
				logRightZu2(usersSubmission, actual_prompt_char, actual_objective_type, actual_objective)
			} else {
				user_guessed_prior_card_rightly = false
				display_failure_of_final_guess_message_etc(usersSubmission)
			}
		}
	} // If the objective was any form of zu, it has been processed above.

	// todo]  Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
	// ::: debug off fmt.Printf("1 guess is %s, objective is %s\n", usersSubmission, actual_objective)
	if submission_already_processed_above == true { // If it was handled as a zu case above.
		// Do nothing ??
	} else {
		if usersSubmission == actual_objective {
			user_guessed_prior_card_rightly = true
			// ::: debug off fmt.Printf("2 guess is %s, objective is %s\n", usersSubmission, actual_objective)
			logRight(usersSubmission, actual_prompt_char, actual_objective_type)
			// ::: debug off fmt.Println("nearing last else in Process_users_input_as_a_guess()")
		} else {
			user_guessed_prior_card_rightly = false
			if guessLevelCounter >= 4 {
				guessLevelCounter = 1
				display_failure_of_final_guess_message_etc(usersSubmission)
			} else {
				user_guessed_prior_card_rightly = false
				// ::: debug off fmt.Println("at: log_oops  in Process_users_input_as_a_guess()　::")
				// ::: debug off fmt.Printf("usersSubmission was: %s, actual_objective was: %s\n", usersSubmission, actual_objective)
			}
		}
	}
	submission_already_processed_above = false // So, reset it for the next round.
}

func potentially_handle_a_directive() { // ::: - -
	detectDirective(usersSubmission)
	if its_a_directive {
		its_a_directive = false

		respond_to_UserSupplied_Directive(usersSubmission)

		prompt_the_user_for_input() // ::: this is where the issue is with game result counting.
		guessLevelCounter = 1       // Directives normally only get done at level 1 ::: ??

		_, _ = fmt.Scan(&usersSubmission)
		detectDirective(usersSubmission)

		if its_a_directive {
			its_a_directive = false
			potentially_handle_a_directive() // ::: Recursion
		} else {

			return
		}
	}
}

func display_failure_of_final_guess_message_etc(userInput string) { // ::: - -
	errors++
	single_faults--
	double_faults--
	log_oops(aCard.Hira, aCard.Romaji, userInput)
	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
}
func log_oops(prompt_it_was, field_it_was, guess string) { // - -
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was)
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

func display_start_menu_etc() { // ::: - -
	display_List_of_Directives()
}
