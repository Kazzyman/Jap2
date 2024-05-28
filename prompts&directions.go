package main

import (
	"fmt"
)

func prompt_the_user_for_input() { // ::: - -

	if guessLevelCounter == 1 {
		guessLevelCounter++
		game_loop_counter++
		if game_loop_counter > game_duration {
			game_off()
		}
		if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			if limitedToDifficultKata {
				fmt.Printf(" Hiragana?, or 'dir' - Limited To Difficult Kata\n")
			} else if limitedToRomaPrompts {
				fmt.Printf(" Hiragana?, or 'dir' - Limited To Romaji Prompts\n")
			} else if limitedToKataPrompts && limitedToHiraPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Includes both Kata and Hira Prompts\n")
			} else if limitedToHiraPrompts {
				fmt.Printf(" Hiragana?, or 'dir' - Limited To Hira Prompts\n")
			} else if limitedToKataPrompts {
				fmt.Printf(" Hiragana?, or 'dir' - Limited To Kata Prompts\n")
			} else {
				fmt.Printf(" Hiragana?, or 'dir' - standard mix\n")
			}
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)

		} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			if limitedToDifficultKata {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Difficult Kata\n")
			} else if limitedToRomaPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Romaji Prompts\n")
			} else if limitedToKataPrompts && limitedToHiraPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Includes both Kata and Hira Prompts\n")
			} else if limitedToHiraPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Hira Prompts\n")
			} else if limitedToKataPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Kata Prompts\n")
			} else {
				fmt.Printf(" Romaji?, or 'dir' - standard mix\n")
			}
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)

			/*
				} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" { // ::: not a possible thang
			*/

		} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			if limitedToDifficultKata {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Difficult Kata\n")
			} else if limitedToRomaPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Romaji Prompts\n")
			} else if limitedToKataPrompts && limitedToHiraPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Includes both Kata and Hira Prompts\n")
			} else if limitedToHiraPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Hira Prompts\n")
			} else if limitedToKataPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Kata Prompts\n")
			} else {
				fmt.Printf(" Romaji?, or 'dir' - standard mix\n")
			}
			fmt.Printf(" Here:> ")
			fmt.Printf("%s", colorReset)

		} else {
			fmt.Printf("Missing one or more elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
		}

	} else if guessLevelCounter == 2 {
		no_interim_error_flag = false
		guessLevelCounter++
		// correct, single_faults, double_faults, errors
		single_faults++
		if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" Guess again, or '?' for clue\n Here:> ")

		} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" Guess again, or '?' for clue\n Here:> ")

		} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" Guess again, or '?' for clue\n Here:> ")
		} else {
			fmt.Printf("Missing one or more elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
		}

	} else if guessLevelCounter == 3 {
		// correct, single_faults, double_faults, errors
		double_faults++
		single_faults--
		if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Romaji)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Hiragana input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")

		} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Hira)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")

			/*
				} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" { // ::: not a possible thang
			*/

		} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Kata)
			fmt.Printf("%s", colorCyan)
			fmt.Printf(" Romaji input-mode expected," + colorReset +
				" you must guess, just one more time\n Here:> ")

		} else {
			fmt.Printf("Missing one or more elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
		}
		//
		guessLevelCounter++
		//
	} else if guessLevelCounter > 3 {
		display_failure_of_final_guess_message_etc(usersSubmission)
		guessLevelCounter = 1
		begin_Kana_practice()
	} else if guessLevelCounter >= 4 || guessLevelCounter <= -1 {
		fmt.Printf("The value of guessLevelCounter is out of range, it is %d \n", guessLevelCounter)
	}
}

/*
.
*/

// ::: Special prompts for Extended Kata, if|when deployed **********************
// ... Standard: used when NOT soliciting second, or final, guesses ***
func promptForRomajiWithDirE(prompt string) { // ::: - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" E-Romaji input-mode expected, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
}

// ... Special: used when soliciting second, or final, guesses ***
func promptForRomajiE(prompt string) { // ::: - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" E-Romaji input-mode expected," + colorReset +
		" you must guess\n Here:> ")
	fmt.Printf("%s", colorReset)
} // ::: ***********************************************************************

/*
-
-
-
*/

// List_of_Directives ::: is in use, a few steps removed.
func List_of_Directives() { // ::: - -
	fmt.Println("\n\nView source code at https://github.com/Kazzyman/Jap2")
	fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	fmt.Println("        Enter '" + colorGreen +
		"nts" + colorReset +
		"' for some background on Romaji conventions")
	fmt.Println("        Enter '" + colorGreen +
		"dir" + colorReset +
		"' redisplay this menu of available Directives")
	fmt.Println("        Enter '" + colorGreen +
		"gdc" + colorReset +
		"' set the Duration Counter for a Game session ")
	fmt.Println("        Enter '" + colorGreen +
		"bgs" + colorReset +
		"' or " + colorGreen +
		"'goff'" + colorReset +
		" Begin or end a Game Session ")
	fmt.Println("        Enter '" + colorGreen +
		"?" + colorReset +
		"' context-sensitive help on the current character")
	fmt.Println("        Enter '" + colorGreen +
		"??" + colorReset +
		"' for help on a particular Hiragana or Romaji")
	fmt.Println("        Enter '" + colorGreen +
		"st" + colorReset +
		"' Statistics")
	fmt.Println("        Enter '" + colorGreen +
		"abt" + colorReset +
		"' for trivia about this app")
	fmt.Println("        Enter '" + colorGreen +
		"rs" + colorReset +
		"' to reset (flush or clear) all stats logs etc.")
	fmt.Println("        Enter '" + colorGreen +
		"rm" + colorReset +
		"' Read the current contents of the Maps")
	fmt.Println("        Enter '" + colorGreen +
		"stc" + colorReset +
		"' (Set-Card) force the use of a specific card (Hira input)")
	fmt.Println("        Enter '" + colorGreen +
		"stcr" + colorReset +
		"' (Set-Card) force the use of a specific card (Roma input)")
	fmt.Println("        Enter '" + colorGreen +
		"exko" + colorReset +
		"' load the Extended Kata deck")
	fmt.Println("        Enter '" + colorGreen +
		"exkf" + colorReset +
		"' un-load the Extended Kata deck")

	fmt.Println("        Enter '" + colorGreen +
		"hko" + colorReset +
		"' Use BOTH Kata & Hira prompting w/ a Romaji objective, rs to ReSet")
	fmt.Println("        Enter '" + colorGreen +
		"konly" + colorReset +
		"' Use only Kata prompting and Romaji objective, rs to ReSet")
	// missing Kata prompting and Hira objective
	fmt.Println("        Enter '" + colorGreen +
		"honly" + colorReset +
		"' Use only Hira prompting and Romaji objective, rs to ReSet")
	fmt.Println("        Enter '" + colorGreen +
		"ronly" + colorReset +
		"' Use only Romaji prompting and Hira objective , rs to ReSet")

	fmt.Println("        Enter '" + colorGreen +
		"donly" + colorReset +
		"' Difficult Kata only, rs to ReSet")

	fmt.Println("        Enter '" + colorGreen +
		"help" + colorReset +
		"' For instructions on how to use this app")

	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
}

/*
.

*/

// 'Directive Menu' ::: in use : displays only at inception
func display_List_of_Directives() { // (unique) // ::: - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Println("\n")
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration)
	fmt.Printf("Current Prompt Count Total: %d \n\n", total_prompts)
	fmt.Printf("Extended Kata deck is loaded: %t \n\n", include_Extended_kata_deck)
	fmt.Printf("Limited to Kata prompts with Romaji objectives: %t \n", limitedToKataPrompts)
	fmt.Printf("Limited to Hira prompts only:                   %t \n", limitedToHiraPrompts)
	fmt.Printf("Limited to Romaji prompts only:                 %t \n", limitedToRomaPrompts)
	fmt.Printf("Limited to Difficult Kata only:                 %t \n\n", limitedToDifficultKata)
}

/*
..
..

*/

// ::: in current use
// ::: 'Directive Menu' ; displays only in response to "Dir" Directive
func re_display_List_of_Directives() { // (unique) // ::: - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Printf("game_loop_counter:%d, game_duration:%d \n", game_loop_counter, game_duration)
	fmt.Printf("total_prompts:%d,  \n\n", total_prompts)

	if !include_Extended_kata_deck {
		fmt.Println("Extended Kata deck is NOT loaded\n")
	} else {
		fmt.Println("Extended Kata deck is loaded\n")
	}

	if limitedToKataPrompts && limitedToHiraPrompts {
		fmt.Printf("Includes BOTH Kata AND Hira prompts, with romaji objectives: %t \n\n", limitedToKataPrompts)
	}
	if limitedToKataPrompts {
		fmt.Printf("Limited to Kata prompts, with ONLY Romaji objectives: %t \n\n", limitedToKataPrompts)
		// Remembering that the Kata-Hira combination is not logically sound (too difficult to determine hira or roma)
	}
	if limitedToHiraPrompts {
		fmt.Printf("Limited to Hira prompts (Romaji objectives): %t \n\n", limitedToHiraPrompts)
	}
	if limitedToRomaPrompts {
		fmt.Printf("Limited to Romaji prompts (Hira objective): %t \n\n", limitedToRomaPrompts)
	}
	if limitedToDifficultKata {
		fmt.Printf("Limited to Difficult Kata only (may ask for either Hira or Romaji as objective): %t \n\n", limitedToDifficultKata)
	}
}
