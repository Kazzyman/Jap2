package main

import (
	"fmt"
	"math/rand"
	"time"
)

// todo Bug::: stcr still has an issue.

func main() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time in nanoseconds.
	gameOn = false
	game = "off"
	guessLevelCounter = 1 // Obviously.

	// todo ::: determine and set type_of_usersSubmission at appropriate times
	// type_of_usersSubmission = "romaji"

	// todo ::: actual_objective_type or equivalent also needed -- what was I thinking here??
	// actual_prompt_char_type

	// usersSubmission = "ronly"
	limitedToHiraPrompts = true // ::: works
	// limitedToKataPrompts = true // ::: but this breaks it
	user_guessed_prior_card_rightly = false
	// categorize_and_process_users_input()
	fmt.Println()
	countSLOC()                     // Determine and display Source Lines Of Code.
	pick_RandomCard_Assign_fields() // this has a bug: Missing one or more elements of prompting style actual_prompt_char_type is romaji, and actual_objective_type is
	// randomize_over_all_decks()
	display_start_menu_etc()
	Kana_practice()
}

func Kana_practice() { // ::: - -
	for {
		if non_standard_origin_DirHandler { // ::: May not actually need this ??, could "user_guessed_prior_card_rightly" suffice??
			non_standard_origin_DirHandler = false
			// ::: debug off fmt..Printf("guessLevelCounter: %d, top of Kana_practice(), inside non_standard_origin...\n", guessLevelCounter)
			// ::: debug off fmt..Printf("at top of Kana_practice()\n")

			// Bypass those elements contained within the following else clause.
			// Note: the dir handler: handle_directive() will have already done each of those things, if needed;
			// including providing for consecutive directives -- therefore nothing is left to do upon "returning".
			// guessLevelCounter = 1 // ::: as a debug test
		} else {
			if user_guessed_prior_card_rightly {
				pick_RandomCard_Assign_fields()
				// randomize_over_all_decks()
				guessLevelCounter = 1
			}
			/*
				aCard.Hira, aCard.Kata, aCard.Romaji ::: done

				actual_objective_type = "hira"
				actual_objective_type = "roma" ::: done

				actual_prompt_char_type = "hira"
				actual_prompt_char_type = "romaji"
				actual_prompt_char_type = "kata" ::: done
			*/
			// Prompt according to guessLevelCounter, type of displayed prompt, and type of requested response.
			// ::: debug off fmt..Printf("guessLevelCounter: %d, before first prompt in Kana_practice()\n", guessLevelCounter)
			// prompt_the_user_for_input(true) // ::: done // todo, lets try true
			prompt_the_user_for_input() // ::: try this
			// ::: debug off fmt..Printf("guessLevelCounter: %d, after first prompt in Kana_practice()\n", guessLevelCounter)

			// Obtain users input. ::: done
			_, _ = fmt.Scan(&usersSubmission)

			// Is users input a Directive, etc. ? Categorize and process users input accordingly.
			categorize_and_process_users_input()
			// ::: debug off fmt..Printf("guessLevelCounter: %d, after call to categorize_and... at bottom of Kana_practice()\n", guessLevelCounter)
		}
		// Loop, will now get a fresh card and the process continues (exit Directive is "q").
	}
}

func categorize_and_process_users_input() { // ::: - -
	// Is users input alpha?

	// If users input is alpha, is it on the list of directives?

	// If users input is a Directive, handle it.
	// ::: debug off fmt.Printf("guessLevelCounter: %d, prior to HD call in \"cate\" \n", guessLevelCounter)

	potentially_handle_a_directive()

	// ::: debug off fmt.Printf("guessLevelCounter: %d, after HD call in \"cate\" \n", guessLevelCounter)

	// ::: debug off fmt.Println("just past: handle_directive() in: categorize_and_process_users_input()  ")
	// If users input is not a Directive, does its type match actual_objective_type ?

	// if type_of_usersSubmission == actual_objective_type {
	// Process users input as a guess if usersSubmission type matches actual_objective type.
	// ::: debug off fmt.Printf("entering Process_users_input_as_a_guess()\n")
	Process_users_input_as_a_guess()
	// ::: debug off fmt.Printf("guessLevelCounter: %d, after call to Process_users_input_as_a_guess() in CAPUI \n", guessLevelCounter)
	/*
		} else {
			// Display a message informing the user that he should change his input method.
			fmt.Println(colorRed)
			fmt.Println("Please change your input method to match the char type that was requested:)")
			fmt.Printf("type_of_usersSubmission: %s, and actual_objective_type: %s\n", type_of_usersSubmission, actual_objective_type)
		}

	*/

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
		non_standard_origin_DirHandler = false
		user_guessed_prior_card_rightly = true
		// ::: debug off fmt..Printf("guess is %s, objective is %s\n", usersSubmission, actual_objective)
		logRight(usersSubmission, actual_prompt_char, actual_objective_type)
		// ::: debug off fmt..Println("nearing last else in Process_users_input_as_a_guess()")
	} else {
		user_guessed_prior_card_rightly = false
		non_standard_origin_DirHandler = false
		if guessLevelCounter >= 4 {
			guessLevelCounter = 1
			logOopsLoser(usersSubmission)
		} else {
			user_guessed_prior_card_rightly = false
			non_standard_origin_DirHandler = false
			// ::: debug off fmt..Println("at: log_oops  in Process_users_input_as_a_guess()　::")
			// ::: debug off fmt..Printf("usersSubmission was: %s, actual_objective was: %s\n", usersSubmission, actual_objective)
		}
	}
}

func potentially_handle_a_directive() { // ::: - -

	detectDirective(usersSubmission)
	if its_a_directive {
		its_a_directive = false
		non_standard_origin_DirHandler = true

		respond_to_UserSupplied_Directive(usersSubmission)

		prompt_the_user_for_input()
		guessLevelCounter = 1 // Directives normally only get done at level 1 ::: ??

		_, _ = fmt.Scan(&usersSubmission)
		detectDirective(usersSubmission)

		if its_a_directive {
			its_a_directive = false
			potentially_handle_a_directive() // ::: Recursion
		} else {
			non_standard_origin_DirHandler = true
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
