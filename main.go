package main

import (
	"fmt"
	"math/rand"
	"time"
)

// todo bug: actual_objective_type is not being set
/*
aCard.Kata
aCard.Hira
aCard.Romaji

actual_objective ::         aCard.Hira  aCard.Romaji
actual_objective_type ::    "hira",  "roma"

actual_prompt_char ::       aCard.Hira  aCard.Romaji  aCard.Kata
actual_prompt_char_type ::  "hira",  "roma",  "kata"

usersSubmission  ::         scanned from keyboard
type_of_usersSubmission ::  "hira",  "roma"                 // rarely used.
usersSubmissionMode ::           "hira",  "roma"

guessLevelCounter ::        0-4  ??
its_a_directive ::  true,  false
non_standard_origin_DirHandler = false

user_guessed_prior_card_rightly bool  // may need this ?
submission_already_processed_above bool // ? will I need this ?

var skip_for_reasons bool
*/
func main() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time in nanoseconds.
	gameOn = false
	game = "off"
	guessLevelCounter = 1 // Obviously.

	// todo ::: determine and set type_of_usersSubmission at appropriate times
	// type_of_usersSubmission = "romaji"

	// todo ::: actual_objective_type or equivalent also needed
	// actual_prompt_char_type

	// usersSubmission = "ronly"
	user_guessed_prior_card_rightly = false
	// categorize_and_process_users_input()
	fmt.Println()
	countSLOC() // Determine and display Source Lines Of Code.
	// pick_RandomCard_Assign_fields()
	randomize_over_all_decks()
	display_start_menu_etc()
	Kana_practice()
}

func Kana_practice() { // ::: done ::: - -
	for {
		if non_standard_origin_DirHandler { // ::: May not actually need this ??, could "user_guessed_prior_card_rightly" suffice??
			non_standard_origin_DirHandler = false
			fmt.Printf("guessLevelCounter: %d, top of Kana_practice(), inside non_standard_origin...\n", guessLevelCounter)
			// fmt.Printf("at top of Kana_practice()\n")
			// Bypass those elements contained within the following else clause.
			// Note: the dir handler: handle_directive() will have already done each of those things, if needed;
			// including providing for consecutive directives -- therefore nothing is left to do upon "returning".
			// guessLevelCounter = 1 // ::: as a debug test
		} else {
			if user_guessed_prior_card_rightly {
				// pick_RandomCard_Assign_fields()
				randomize_over_all_decks()
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
			fmt.Printf("guessLevelCounter: %d, before first prompt in Kana_practice()\n", guessLevelCounter)
			// prompt_the_user_for_input(true) // ::: done // todo, lets try true
			prompt_the_user_for_input() // ::: try this
			fmt.Printf("guessLevelCounter: %d, after first prompt in Kana_practice()\n", guessLevelCounter)

			// Obtain users input. ::: done
			_, _ = fmt.Scan(&usersSubmission)

			// Is users input a Directive, etc. ? Categorize and process users input accordingly.
			categorize_and_process_users_input()
			fmt.Printf("guessLevelCounter: %d, after call to categorize_and... at bottom of Kana_practice()\n", guessLevelCounter)
		}
		// Loop, will now get a fresh card and the process continues (exit Directive is "q").
	}
}

func categorize_and_process_users_input() { // ::: - -
	// Is users input alpha?

	// If users input is alpha, is it on the list of directives?

	// If users input is a Directive, handle it.
	fmt.Printf("guessLevelCounter: %d, prior to HD call in \"cate\" \n", guessLevelCounter)

	potentially_handle_a_directive()

	fmt.Printf("guessLevelCounter: %d, after HD call in \"cate\" \n", guessLevelCounter)

	// fmt.Println("just past: handle_directive() in: categorize_and_process_users_input()  ")
	// If users input is not a Directive, does its type match actual_objective_type ?

	// if type_of_usersSubmission == actual_objective_type {
	// Process users input as a guess if usersSubmission type matches actual_objective type.
	fmt.Printf("entering Process_users_input_as_a_guess()\n")
	Process_users_input_as_a_guess()
	fmt.Printf("guessLevelCounter: %d, after call to Process_users_input_as_a_guess() in CAPUI \n", guessLevelCounter)
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

	// ::: if we dont stop it, execution continues top of Kana_practice()
	// todo]  Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
	fmt.Printf("guess is %s, objective is %s\n", usersSubmission, actual_objective)

	if usersSubmission == actual_objective {
		non_standard_origin_DirHandler = false
		user_guessed_prior_card_rightly = true
		fmt.Printf("guess is %s, objective is %s\n", usersSubmission, actual_objective)
		logRight(usersSubmission, actual_prompt_char, actual_objective_type)
		fmt.Println("nearing last else in Process_users_input_as_a_guess()")
	} else {
		user_guessed_prior_card_rightly = false
		non_standard_origin_DirHandler = false
		if guessLevelCounter >= 4 {
			guessLevelCounter = 1
			logOopsLoser(usersSubmission)
			/*
				log_oops(aCard.Hira, aCard.Romaji, usersSubmission)
					fmt.Printf("%s", colorRed)
					fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
					fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
			*/

		} else {
			// log_oops(actual_prompt_char, actual_objective, usersSubmission)
			user_guessed_prior_card_rightly = false
			non_standard_origin_DirHandler = false
			fmt.Println("at: log_oops  in Process_users_input_as_a_guess()　::")
			fmt.Printf("usersSubmission was: %s, actual_objective was: %s\n", usersSubmission, actual_objective)
		}
	}
}

func potentially_handle_a_directive() { // ::: - -

	detectDirective(usersSubmission)
	if its_a_directive {
		its_a_directive = false
		non_standard_origin_DirHandler = true
		if usersSubmission == "stc" || usersSubmission == "stcr" {
			respond_to_UserSupplied_stc_Directive(usersSubmission)
		} else {
			respond_to_UserSuppliedDirective(usersSubmission)
		}

		// Re-prompt.
		// unlessComingFromDirHandler := true // todo ::: fix this
		// prompt_the_user_for_input(unlessComingFromDirHandler)
		prompt_the_user_for_input()
		guessLevelCounter = 1 // Directives normally only get done at level 1 ::: ??

		_, _ = fmt.Scan(&usersSubmission)
		detectDirective(usersSubmission)

		if its_a_directive {
			its_a_directive = false
			// guessLevelCounter = 1
			potentially_handle_a_directive() // ::: Recursion
		} else {
			non_standard_origin_DirHandler = true
			return
		}
	}
}

// func prompt_the_user_for_input(unlessComingFromDirHandler bool) { // ::: - -
func prompt_the_user_for_input() { // ::: - -
	/*
			actual_objective_type = "hira"
			actual_objective_type = "roma" ::: done

			actual_prompt_char_type = "hira"  ::: all
			actual_prompt_char_type = "romaji"
			actual_prompt_char_type = "kata"

			roma-hira ::: done
			hira-roma ::: done

			kata-hira ::: done
			kata-roma ::: done

		if unlessComingFromDirHandler {
			// skip prompt.
		} else if guessLevelCounter == 0 {
			guessLevelCounter++
			display_start_menu_etc()
		}

	*/

	if guessLevelCounter == 1 {
		guessLevelCounter++
		if actual_prompt_char_type == "romaji" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected, or '?' for help with: %s \n", aCard.Romaji)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected, or '?' for help with: %s \n", aCard.Hira)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected, or '?' for help with: %s \n", aCard.Kata)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected, or '?' for help with: %s \n", aCard.Kata)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else {
			fmt.Printf("Missing one or more elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)

		}

	} else if guessLevelCounter == 2 {
		guessLevelCounter++
		if actual_prompt_char_type == "romaji" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else {
			fmt.Printf("Missing one or more elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)

		}

	} else if guessLevelCounter == 3 {
		if actual_prompt_char_type == "romaji" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else {
			fmt.Printf("Missing one or more elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
		}
		// guessLevelCounter++
		//
	} else if guessLevelCounter > 3 {
		display_failure_of_final_guess_message_etc(usersSubmission)
		guessLevelCounter = 1
		Kana_practice()
	} else if guessLevelCounter >= 4 || guessLevelCounter <= -1 {
		fmt.Printf("The value of guessLevelCounter is out of range, it is %d \n", guessLevelCounter)
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
