package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println()
	countSLOC()                      // Determine and display Source Lines Of Code.
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time in nanoseconds.
	gameOn = false
	game = "off"
	display_List_of_Directives()
	Kana_practice()
}

func Kana_practice() {
	for {

		pick_RandomCard_Assign_fields() // The state of display type and response expected have be set in globals :::
		/*
			aCard.Hira, aCard.Kata, aCard.Romaji ::: done

			expected_response_type = "hira_char_as_users_guess"
			expected_response_type = "romaji_char_as_users_guess" ::: done

			displayed_prompt_type = "hira"
			displayed_prompt_type = "romaji"
			displayed_prompt_type = "kata" ::: all done
		*/
		// Prompt according to guess_attempt_count and type of displayed prompt and type of requested response.
		prompt_the_user_for_input() // ::: done.

		// Obtain users input.
		_, _ = fmt.Scan(&usersInput)

	}
}

func prompt_the_user_for_input() {
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
	if guess_attempt_count == 1 {
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
			fmt.Println("Missing one or more elements of prompting style")
		}

	} else if guess_attempt_count == 2 {
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
			fmt.Println("Missing one or more elements of prompting style")
		}
	} else if guess_attempt_count == 3 {
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
			fmt.Println("Missing one or more elements of prompting style")
		}
	} else if guess_attempt_count == 4 {
		display_failure_of_final_guess_message_etc(usersInput)
	} else {
		fmt.Printf("The value of guess_attempt_count is out of range, it is %d \n", guess_attempt_count)
	}
}

func display_failure_of_final_guess_message_etc(userInput string) {
	log_oops(aCard.Hira, aCard.Romaji, userInput)
	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
}
