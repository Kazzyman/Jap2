package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
expected_response_type is "romaji_char_as_users_guess" and it should also be "romaji"
*/
func main() {
	fmt.Println()
	countSLOC()                      // Determine and display Source Lines Of Code.
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time in nanoseconds.
	gameOn = false
	game = "off"
	guess_attempt_count = 0 // Obviously.
	display_start_menu_etc()
	Kana_practice()

	// todo: determine and set a type var after scanning the usersInput
	typeOfUsersInput = "romaji"
	usersInput = "ronly"
	categorize_and_process_users_input()
}

func Kana_practice() { // ::: done ::: - -
	for {
		if returning_from_handling_a_directive {
			returning_from_handling_a_directive = false
			// Bypass those elements contained within the following else clause.
			// Note: the dir handler: handle_directive() will have already done each of those things, if needed;
			// including providing for consecutive directives -- therefore nothing is left to do upon "returning".
		} else {
			pick_RandomCard_Assign_fields() // The state of display type and response expected have been set in globals ::: done
			/*
				aCard.Hira, aCard.Kata, aCard.Romaji ::: done

				expected_response_type = "hira_char_as_users_guess"
				expected_response_type = "romaji_char_as_users_guess" ::: done



				displayed_prompt_type = "hira"
				displayed_prompt_type = "romaji"
				displayed_prompt_type = "kata" ::: done
			*/
			// Prompt according to guess_attempt_count and type of displayed prompt and type of requested response.
			prompt_the_user_for_input() // ::: done

			// Obtain users input. ::: done
			_, _ = fmt.Scan(&usersInput)

			// Is users input a Directive, etc. ? Categorize and process users input accordingly.
			categorize_and_process_users_input()
		}

		// Loop, will now get a fresh card and the process continues (exit Directive is "q").
	}
}

func categorize_and_process_users_input() { // ::: - -
	// Is users input alpha?

	// If users input is alpha, is it on the list of directives?

	// If users input is a Directive, handle it.
	handle_directive()

	// If users input is not a Directive, does its type match expected_response_type ?
	if typeOfUsersInput == objective_kind {
		// Process users input as a guess if usersInput type matches objective type.
		Process_users_input_as_a_guess()
	} else {
		// Display a message informing the user that he should change his input method.
		fmt.Println(colorRed)
		fmt.Println("Please change your input method to match the char type that was requested:)")
	}
}

func Process_users_input_as_a_guess() { // ::: - -
	// todo]  Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	if usersInput == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true // todo: is this bool setting well-placed? Compare others.
		if usersInput == "zu" {
			logRight_zu(usersInput, promptField, objective_kind)
		} else {
			log_oops(promptField, objective, usersInput)
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true { // todo: is this bool test even needed??
		// todo]  Note: that "else", here, means that objective != "zu"  ... but the objective may yet be ず or づ.

		if objective == "ず" || objective == "づ" {
			if usersInput == "づ" || usersInput == "ず" {
				logRightZu2(usersInput, promptField, objective_kind, objective)
			} else {
				log_oops(promptField, objective, usersInput)
			}
		}

		// todo]  Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
		if usersInput == objective {
			logRight(usersInput, promptField, objective_kind)

		} else {
			log_oops(promptField, objective, usersInput)
		}
	}
}

func handle_directive() { // ::: - -
	// returning_from_handling_a_directive = true
	// 
	for {
		detectDirective(usersInput)
		if aDirectiveWasDetected {
			aDirectiveWasDetected = false
			if usersInput == "stc" || usersInput == "stcr" {
				respond_to_UserSupplied_stc_Directive(usersInput)
			} else {
				respond_to_UserSuppliedDirective(usersInput)
			}
			// Re-prompt.
			prompt_the_user_for_input()
			_, _ = fmt.Scan(&usersInput)
			detectDirective(usersInput)
			if aDirectiveWasDetected {
				handle_directive()
			} else {
				returning_from_handling_a_directive = true
				break
			}
		}
	}
}

func prompt_the_user_for_input() { // ::: - -
	/*
		expected_response_type = "hira_char_as_users_guess"
		expected_response_type = "romaji_char_as_users_guess" ::: done

		displayed_prompt_type = "hira"  ::: all
		displayed_prompt_type = "romaji"
		displayed_prompt_type = "kata"

		roma-hira ::: done
		hira-roma ::: done

		kata-hira ::: done
		kata-roma ::: done
	*/
	if guess_attempt_count == 0 {
		guess_attempt_count++
		display_start_menu_etc()
	}

	if guess_attempt_count == 1 {
		guess_attempt_count++
		if displayed_prompt_type == "romaji" && expected_response_type == "hira_char_as_users_guess" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected, or '?' for help with: %s \n", aCard.Romaji)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else if displayed_prompt_type == "hira" && expected_response_type == "romaji_char_as_users_guess" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected, or '?' for help with: %s \n", aCard.Hira)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else if displayed_prompt_type == "kata" && expected_response_type == "hira_char_as_users_guess" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected, or '?' for help with: %s \n", aCard.Kata)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else if displayed_prompt_type == "kata" && expected_response_type == "romaji_char_as_users_guess" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected, or '?' for help with: %s \n", aCard.Kata)
			fmt.Printf(" or, type '??' for help with something else ... \n")
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)
		} else {
			fmt.Printf("Missing one or more elements of prompting style "+
				"displayed_prompt_type is %s, and expected_response_type is %s\n", displayed_prompt_type, expected_response_type)
		}

	} else if guess_attempt_count == 2 {
		guess_attempt_count++
		if displayed_prompt_type == "romaji" && expected_response_type == "hira_char_as_users_guess" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else if displayed_prompt_type == "hira" && expected_response_type == "romaji_char_as_users_guess" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else if displayed_prompt_type == "kata" && expected_response_type == "hira_char_as_users_guess" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else if displayed_prompt_type == "kata" && expected_response_type == "romaji_char_as_users_guess" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must try to guess\n Here:> ")
		} else {
			fmt.Printf("Missing one or more elements of prompting style "+
				"displayed_prompt_type is %s, and expected_response_type is %s\n", displayed_prompt_type, expected_response_type)
		}

	} else if guess_attempt_count == 3 {
		guess_attempt_count++
		if displayed_prompt_type == "romaji" && expected_response_type == "hira_char_as_users_guess" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else if displayed_prompt_type == "hira" && expected_response_type == "romaji_char_as_users_guess" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else if displayed_prompt_type == "kata" && expected_response_type == "hira_char_as_users_guess" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else if displayed_prompt_type == "kata" && expected_response_type == "romaji_char_as_users_guess" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")
		} else {
			fmt.Printf("Missing one or more elements of prompting style "+
				"displayed_prompt_type is %s, and expected_response_type is %s\n", displayed_prompt_type, expected_response_type)
		}

	} else if guess_attempt_count == 4 {
		display_failure_of_final_guess_message_etc(usersInput)
	} else {
		fmt.Printf("The value of guess_attempt_count is out of range, it is %d \n", guess_attempt_count)
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
