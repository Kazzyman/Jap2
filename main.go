package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// todo Bug::: stcr still has an issue. Though it does not crash the app if it us run immediately after an stc ...
// ... it is actually recoverable if it is run after an stc.

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with the current time in nanoseconds.
	gameOn = false
	game = "off"
	guessLevelCounter = 1
	limitedToHiraPrompts = true
	user_guessed_prior_card_rightly = false
	fmt.Println()
	countSLOC()                     // Determine and display Source Lines Of Code.
	pick_RandomCard_Assign_fields() // Pick the first card.
	display_start_menu_etc()
	begin_Kana_practice()
}

func begin_Kana_practice() { // ::: - -
	for { // ::: Application loop.
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		if game_loop_counter > game_duration {
			game_off()
		}
		// ::: debug off fmt..Printf("guessLevelCounter: %d, top of Kana_practice(), inside non_standard_origin...\n", guessLevelCounter)
		// ::: debug off fmt..Printf("at top of Kana_practice()\n")

		if user_guessed_prior_card_rightly {
			pick_RandomCard_Assign_fields()
			guessLevelCounter = 1
		}

		// ::: debug off fmt..Printf("guessLevelCounter: %d, before first prompt in Kana_practice()\n", guessLevelCounter)

		// Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
		prompt_the_user_for_input()

		// ::: debug off fmt..Printf("guessLevelCounter: %d, after first prompt in Kana_practice()\n", guessLevelCounter)

		// Obtain users input.
		_, _ = fmt.Scan(&usersSubmission)

		// Categorize and process users input accordingly, e.g., is users input a Directive, etc.?
		// todo : remove : categorize_and_process_users_input() -- and the above comment.

		// If users input is a Directive, handle it.
		potentially_handle_a_directive()

		// ::: debug off fmt.Printf("guessLevelCounter: %d, after HD call in \"cate\" \n", guessLevelCounter)

		processUsersSubmissionIfTypeEnteredRightly()

		// ::: debug off fmt..Printf("guessLevelCounter: %d, after call to categorize_and... at bottom of Kana_practice()\n", guessLevelCounter)
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

			// ::: debug off fmt.Printf("guessLevelCounter: %d, after call to Process_users_input_as_a_guess()
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

			// ::: debug off fmt.Printf("guessLevelCounter: %d, after call to Process_users_input_as_a_guess()
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
	if usersSubmission == "zu" {
		submission_already_processed_above = true // todo: is this bool setting well-placed? Compare others.
		if usersSubmission == "zu" {
			logRight_zu(usersSubmission, actual_prompt_char, actual_objective_type)
		} else {
			log_oops(actual_prompt_char, actual_objective, usersSubmission)
		}
	} else if submission_already_processed_above != true { // todo: is this bool test even needed??
		// todo]  Note: that "else", here, means that actual_objective != "zu"  ... but the actual_objective may yet be ず or づ.

		if actual_objective == "ず" || actual_objective == "づ" {
			if usersSubmission == "づ" || usersSubmission == "ず" {
				logRightZu2(usersSubmission, actual_prompt_char, actual_objective_type, actual_objective)
			} else {
				log_oops(actual_prompt_char, actual_objective, usersSubmission)
			}
		}
	}

	// todo]  Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
	// ::: debug off fmt.Printf("guess is %s, objective is %s\n", usersSubmission, actual_objective)

	if usersSubmission == actual_objective {
		from_non_standard_origin_DirHandler = false
		user_guessed_prior_card_rightly = true
		// ::: debug off fmt..Printf("guess is %s, objective is %s\n", usersSubmission, actual_objective)
		logRight(usersSubmission, actual_prompt_char, actual_objective_type)
		// ::: debug off fmt..Println("nearing last else in Process_users_input_as_a_guess()")
	} else {
		user_guessed_prior_card_rightly = false
		from_non_standard_origin_DirHandler = false
		if guessLevelCounter >= 4 {
			guessLevelCounter = 1
			logOopsLoser(usersSubmission)
		} else {
			user_guessed_prior_card_rightly = false
			from_non_standard_origin_DirHandler = false
			// ::: debug off fmt..Println("at: log_oops  in Process_users_input_as_a_guess()　::")
			// ::: debug off fmt..Printf("usersSubmission was: %s, actual_objective was: %s\n", usersSubmission, actual_objective)
		}
	}
}

func potentially_handle_a_directive() { // ::: - -
	detectDirective(usersSubmission)
	if its_a_directive {
		its_a_directive = false
		from_non_standard_origin_DirHandler = true

		respond_to_UserSupplied_Directive(usersSubmission)

		prompt_the_user_for_input()
		guessLevelCounter = 1 // Directives normally only get done at level 1 ::: ??

		_, _ = fmt.Scan(&usersSubmission)
		detectDirective(usersSubmission)

		if its_a_directive {
			its_a_directive = false
			potentially_handle_a_directive() // ::: Recursion
		} else {
			from_non_standard_origin_DirHandler = true
			return
		}
	}
}

func display_failure_of_final_guess_message_etc(userInput string) { // ::: - -
	log_oops(aCard.Hira, aCard.Romaji, userInput)
	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
}

func display_start_menu_etc() { // ::: - -
	display_List_of_Directives()
}
