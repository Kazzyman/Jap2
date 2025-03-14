package main

import (
	"fmt"
	"os"
)

func prompt_the_user_for_input() { // ::: - -

	if guessLevelCounter == 1 { // ::: --------- first prompt -------- 1 1 1 1 1 1 ------------------
		// fmt.Printf("\nGuess level counter is : %d\n", guessLevelCounter)
		guessLevelCounter++

		// Roma_Hira ::: the only way to have Roma prompting is via limitedToRomaPrompts or standard mix - ?
		if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Romaji)

			fmt.Printf("%s", colorCyan)
			if limitedToRomaPrompts {
				if limitedToRomaPromptsAndSimplexHiraObj {
					fmt.Printf(" Hiragana?, or 'dir' - Limited To Romaji Prompts w/Simplex obj") // could be, could b
				} else {
					fmt.Printf(" Hiragana?, or 'dir' - Limited To Romaji Prompts") // could be, could be
				}
			} else {
				fmt.Printf(" Hiragana?, or 'dir' - standard mix") // We should print "standard mix" only if we have proved that we are not here by random?
			}

			if theGameIsRunning { // more shit is appended when a game is running
				fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d%s\n",
					colorReset, nameOfPlayer, colorRed, colorReset, gameCorrectOnFirstAttemptAccumulator,
					colorRed, colorReset, gameCorrectOnSecondAttemptAccumulator, colorRed, colorReset, gameCorrectOnThirdAttemptAccumulator,
					colorRed, colorReset, gameFailedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, gameDuration, colorReset)
			} else {
				fmt.Println()
			}

			if theGameIsRunning { // Determine which "prompt" string to deploy
				fmt.Printf("%s Game:> %s", colorGreen, colorReset)
			} else {
				fmt.Printf(" Hira:> ")
			}
			fmt.Printf("%s", colorReset)

			// Hira_Roma
		} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Hira)

			fmt.Printf("%s", colorCyan)
			if limitedToKataPrompts && limitedToHiraPrompts { // ::: either or
				fmt.Printf(" Romaji?, or 'dir' - Includes Kata&Hira Prompts, in this case hira")
			} else if limitedToHiraPrompts { // if we had been limited to Kata prompts only we could not have landed here in this Hira_Roma block
				fmt.Printf(" Romaji?, or 'dir' - Limited To Hira Prompts")
			} else {
				// the only possible scenarios for Hira_Roma were/are mixed prompting or limited to Hira or Kata ???
				fmt.Printf(" Romaji?, or 'dir' - standard mix")
			}

			if theGameIsRunning {
				fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d%s\n",
					colorReset, nameOfPlayer, colorRed, colorReset, gameCorrectOnFirstAttemptAccumulator,
					colorRed, colorReset, gameCorrectOnSecondAttemptAccumulator, colorRed, colorReset, gameCorrectOnThirdAttemptAccumulator,
					colorRed, colorReset, gameFailedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, gameDuration, colorReset)
			} else {
				fmt.Println()
			}

			if theGameIsRunning {
				fmt.Printf("%s Game:> %s", colorGreen, colorReset)
			} else {
				fmt.Printf(" Roma:> ")
			}
			fmt.Printf("%s", colorReset)

			// Kata_Hira ::: the only way to have Kata prompting is via limitedToKataPrompts or standard mix
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" {
			fmt.Printf("%s", aCard.Kata)

			fmt.Printf("%s", colorCyan)
			if limitedToKataPrompts && limitedToHiraPrompts {
				fmt.Printf(" Hiragana?, or 'dir' - Includes Kata&Hira Prompts, in this case kata")
			} else if limitedToKataPrompts {
				if limitedToKataPromptsAndSimplexHiraObj {
					fmt.Printf(" Hiragana?, or 'dir' - Limited To Kata Prompts w/Simplex Hira")
				} else {
					fmt.Printf(" Hiragana?, or 'dir' - Limited To Kata Prompts")
				}
			} else if kata_hira {
				fmt.Printf(" Hiragana?, or 'dir' - Limited To Kata Prompts w/hira obj")
			} else {
				fmt.Printf(" Hiragana?, or 'dir' - standard mix")
			}

			if theGameIsRunning {
				fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d%s\n",
					colorReset, nameOfPlayer, colorRed, colorReset, gameCorrectOnFirstAttemptAccumulator,
					colorRed, colorReset, gameCorrectOnSecondAttemptAccumulator, colorRed, colorReset, gameCorrectOnThirdAttemptAccumulator,
					colorRed, colorReset, gameFailedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, gameDuration, colorReset)
			} else {
				fmt.Println()
			}

			if theGameIsRunning {
				fmt.Printf("%s Game:> %s", colorGreen, colorReset)
			} else {
				fmt.Printf(" Hira:> ")
			}
			fmt.Printf("%s", colorReset)

			// Kata_Roma
		} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
			fmt.Printf("%s", aCard.Kata)

			fmt.Printf("%s", colorCyan)
			if limitedToKataPrompts && limitedToHiraPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Includes Kata&Hira Prompts, in this case kata")
			} else if limitedToKataPrompts {
				fmt.Printf(" Romaji?, or 'dir' - Limited To Kata Prompts")
			} else if kata_roma {
				fmt.Printf(" Romaji input-mode expected, or 'dir' - Limited To Kata Prompts w/roma obj")
			} else if limitedToDifficultDescriptions { // ::: yes this goes here
				fmt.Printf(" Romaji?, or 'dir' - Limited to Difficult Descriptions")
			} else {
				fmt.Printf(" Romaji?, or 'dir' - standard mix")
			}

			if theGameIsRunning {
				fmt.Printf(", %s%s is playing: %s1st:%s%d%s, 2nd:%s%d%s, 3rd:%s%d%s, fails:%s%d, %s%d/%d%s\n",
					colorReset, nameOfPlayer, colorRed, colorReset, gameCorrectOnFirstAttemptAccumulator,
					colorRed, colorReset, gameCorrectOnSecondAttemptAccumulator, colorRed, colorReset, gameCorrectOnThirdAttemptAccumulator,
					colorRed, colorReset, gameFailedOnThirdAttemptAccumulator, colorCyan, game_loop_counter, gameDuration, colorReset)
			} else {
				fmt.Println()
			}

			if theGameIsRunning {
				fmt.Printf("%s Game:> %s", colorGreen, colorReset)
			} else {
				fmt.Printf(" Roma:> ")
			}
			fmt.Printf("%s", colorReset)
			//
		} else {
			fmt.Printf("Miss-matched elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
		}

		gottenHonestly = true
		/*
			;
		*/

	} else if guessLevelCounter == 2 { // ::: --------------- Guess again ------------------ 2 2 2 2 2 2 2 2 2 2 2 -------------------
		guessLevelCounter++
		if theGameIsRunning {
			if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
				fmt.Printf("%s", aCard.Romaji)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorRed +
					" Guess again\n " + colorYellow + "Game:> " + colorReset)
			} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Hira)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorRed +
					" Guess again\n " + colorYellow + "Game:> " + colorReset)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorRed +
					" Guess again\n " + colorYellow + "Game:> " + colorReset)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" { // was a bug fix ???
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorRed +
					" Guess again\n " + colorYellow + "Game:> " + colorReset)
			} else {
				// ::: got this message in error ??
				fmt.Printf("Miss-matched elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
			}
		} else {
			if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
				fmt.Printf("%s", aCard.Romaji)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorRed +
					" Guess again," + colorReset + " or '?' for clue\n " + colorYellow + "Hira:> " + colorReset)
			} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Hira)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorRed +
					" Guess again, " + colorReset + "or '?' for clue\n " + colorYellow + "Roma:> " + colorReset)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorRed +
					" Guess again, " + colorReset + "or '?' for clue\n " + colorYellow + "Roma:> " + colorReset)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" { // was a bug fix ???
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorRed +
					" Guess again," + colorReset + " or '?' for clue\n " + colorYellow + "Hira:> " + colorReset)
			} else {
				// ::: got this message in error ??
				fmt.Printf("Miss-matched elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
			}
		}
		/*.
		.

		*/
	} else if guessLevelCounter == 3 { // ::: -------------- you must guess, just one more time -----3 3 3 3 3 3 3 3 3 3 --------------------------------------

		guessLevelCounter++
		if theGameIsRunning {
			if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
				fmt.Printf("%s", aCard.Romaji)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Game:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Romaji)
			} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Hira)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Game:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Hira)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Game:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Kata)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" { // was a bug fix ???
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Game:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Kata)
			} else {
				fmt.Printf("Miss-matched elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
			}
		} else {
			if actual_prompt_char_type == "roma" && actual_objective_type == "hira" {
				fmt.Printf("%s", aCard.Romaji)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Hira:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Romaji)
			} else if actual_prompt_char_type == "hira" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Hira)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Roma:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Hira)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "roma" {
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Romaji input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Roma:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Kata)
			} else if actual_prompt_char_type == "kata" && actual_objective_type == "hira" { // was a bug fix ???
				fmt.Printf("%s", aCard.Kata)
				fmt.Printf("%s", colorCyan)
				fmt.Printf(" Hiragana input-mode expected," + colorReset +
					" you must guess, " + colorRed + "just one more time\n Hira:> " + colorReset)
				logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.Kata)
			} else {
				fmt.Printf("Miss-matched elements of prompting style actual_prompt_char_type is %s, and actual_objective_type is %s\n", actual_prompt_char_type, actual_objective_type)
			}
		}
		//
		gottenHonestly = false
		/*
			;
		*/
	} else if guessLevelCounter == 4 { // ::: ----- ^^Oops! That was your last try looser. Here's a clue, just for you: --------  > 3  > 3  > 3  > 3  > 3  > 3  ---- i.e. 4  ----
		display_failure_of_final_guess_message_etc(usersSubmission) // ::: ^^^^
		weHadFailed_And_OnlyGotThisRightBecauseOfTheClue = true
		guessLevelCounter = 1 // ::: was 1, why 1 ??, prob because it is set to 1 twice in main
		begin_Kana_practice()
	} else if guessLevelCounter >= 5 || guessLevelCounter <= -1 {
		fmt.Printf("The value of guessLevelCounter is out of range, it is %d \n", guessLevelCounter)
	} else {
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
func display_limited_gaming_dir_list() {
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"' End this game early")
	fmt.Printf("        Enter '" + colorGreen +
		"game" + colorReset +
		"' display the list of game types & the current game\n")
	fmt.Printf("        Enter '" + colorGreen +
		"nts" + colorReset +
		"' display some notes on Kana\n")
	fmt.Printf("        Enter '" + colorGreen +
		"abt" + colorReset +
		"' display facts about the code\n")
	fmt.Printf("        Enter '" + colorGreen +
		"rm" + colorReset +
		"' perform a read of the maps\n")
	fmt.Printf("        Enter '" + colorGreen +
		"st" + colorReset +
		"' show the stats\n")
}

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
		"spell" + colorReset +
		"' English spelling bee")
	fmt.Println("        Enter '" + colorGreen +
		"hko" + colorReset +
		"' Use Kata & Hira prompting (Only Roma objectives)")
	fmt.Println("        Enter '" + colorGreen +
		"konly" + colorReset +
		"' Use only Kata prompting (mix Hira & Roma objectives)")
	fmt.Println("        Enter '" + colorGreen +
		"honly" + colorReset +
		"' Use only Hira prompting (so only Roma responses)")

	fmt.Println("        Enter '" + colorGreen +
		"ronly" + colorReset +
		"' Use only Romaji prompting (so only Hira responses)")
	fmt.Println("        Enter '" + colorGreen +
		"rhSimplex" + colorReset +
		"' Use Romaji (fix to kata_hira?) prompting & Simplex Hira obj")

	fmt.Println("        Enter '" + colorGreen +
		"donly" + colorReset +
		"' Difficult descriptive prompting")

	fmt.Println("        Enter '" + colorGreen +
		"kh" + colorReset +
		"' Use only kata_hira prompt_response")
	fmt.Println("        Enter '" + colorGreen +
		"khSimplex" + colorReset +
		"' Use kata prompts w/Simplex Hira")

	fmt.Println("        Enter '" + colorGreen +
		"kr" + colorReset +
		"' Use only kata_roma prompt_response")

	fmt.Println("        Enter '" + colorGreen +
		"mix" + colorReset +
		"' Use mixed prompts & objectives")

	fmt.Println("        Enter '" + colorGreen +
		"help" + colorReset +
		"' For instructions on how to use this app")
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
	fmt.Println("        Enter '" + colorGreen +
		"game" + colorReset +
		"', to begin a session and log stats to a file")
}

func Display_the_menu_of_game_types() { // ::: - -
	// fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	/*
		fmt.Println("        Enter '" + colorGreen +
			"hko" + colorReset +
			"' Use Kata & Hira prompting (Roma objectives)")
		fmt.Println("        Enter '" + colorGreen +
			"konly" + colorReset +
			"' Use only Kata prompting (mix Hira & Roma objectives)")
		fmt.Println("        Enter '" + colorGreen +
			"honly" + colorReset +
			"' Use only Hira prompting (so only Roma responses)")
		fmt.Println("        Enter '" + colorGreen +
			"ronly" + colorReset +
			"' Use only Romaji prompting (so only Hira responses)")
		fmt.Println("        Enter '" + colorGreen +
			"rhSimplex" + colorReset +
			"' Use Romaji prompting & Simplex Hira obj")
		fmt.Println("        Enter '" + colorGreen +
			"donly" + colorReset +
			"' Difficult descriptive prompting")
		fmt.Println("        Enter '" + colorGreen +
			"kh" + colorReset +
			"' Use only kata_hira prompt_response")
		fmt.Println("        Enter '" + colorGreen +
			"khSimplex" + colorReset +
			"' Use kata prompts w/Simplex Hira")
		fmt.Println("        Enter '" + colorGreen +
			"kr" + colorReset +
			"' Use only kata_roma prompt_response")
		fmt.Println("        Enter '" + colorGreen +
			"mix" + colorReset +
			"' Use mixed prompts & objectives")
	*/
	fmt.Println("        Enter '" + colorGreen +
		"1" + colorReset +
		"' Use Kata & Hira prompting (Roma objectives)")
	fmt.Println("        Enter '" + colorGreen +
		"2" + colorReset +
		"' Use only Kata prompting (mix Hira & Roma objectives)")
	fmt.Println("        Enter '" + colorGreen +
		"3" + colorReset +
		"' Use only Hira prompting (so only Roma responses)")

	fmt.Println("        Enter '" + colorGreen +
		"4" + colorReset +
		"' Use only Romaji prompting (so only Hira responses)")
	fmt.Println("        Enter '" + colorGreen +
		"5" + colorReset +
		"' Use Kata prompting & Simplex Hira obj")

	fmt.Println("        Enter '" + colorGreen +
		"6" + colorReset +
		"' Difficult descriptive prompting")

	fmt.Println("        Enter '" + colorGreen +
		"7" + colorReset +
		"' Use only kata_hira prompt_response")
	fmt.Println("        Enter '" + colorGreen +
		"8" + colorReset +
		"' Use kata prompts w/Simplex Hira")

	fmt.Println("        Enter '" + colorGreen +
		"9" + colorReset +
		"' Use only kata_roma prompt_response")

	fmt.Println("        Enter '" + colorGreen +
		"0" + colorReset +
		"' Use mixed prompts & objectives")

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

	/*
		if now_using_game_duration_set_by_game_type {
			fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, gameDuration)
		} else {
			fmt.Printf("Game counter: %d, Game Duration: %d \n\n", game_loop_counter, jim)
		}
	*/

	/*
		fmt.Printf("Current Prompt Count Total: %d \n\n", total_prompts)

		fmt.Printf("Extended Kata deck is loaded: %t \n\n", include_Extended_kata_deck)

		fmt.Printf("Limited to Kata_hira prompt_response:             %t \n", kata_hira)
		fmt.Printf("Limited to Kata_roma prompt_response:             %t \n", kata_roma)

		fmt.Printf("Limited to Kata prompts:             %t \n", limitedToKataPrompts)
		fmt.Printf("Limited to Hira prompts only:        %t \n", limitedToHiraPrompts)
		fmt.Printf("Limited to Romaji prompts only:      %t \n", limitedToRomaPrompts)
		fmt.Printf("Limited to Difficult Kata only:      %t \n\n", limitedToDifficultDescriptions)

	*/
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
	if theGameIsRunning {
		if now_using_game_duration_set_by_game_type {
			fmt.Println(colorRed)
			fmt.Printf("\nAn active game is in session, the game_loop_counter is:%s%d%s, and the game_duration is set to:%s%d \n\n", colorReset, game_loop_counter, colorRed, colorReset, gameDuration)
			fmt.Println(colorReset)
		} else {
			fmt.Printf("game_loop_counter:%d, game_duration:%d \n", game_loop_counter, jim)
		}
		fmt.Printf("total_prompts:%d,  \n\n", total_prompts)
	} else {
		fmt.Println(colorRed)
		fmt.Printf("\nYou are not currently engaged in a timed game sesion. \n\n")
		fmt.Println(colorReset)
	}

	if !include_Extended_kata_deck {
		fmt.Println("Extended Kata deck is NOT loaded\n")
	} else {
		fmt.Println("Extended Kata deck is loaded\n")
	}

	if limitedToKataPrompts && limitedToHiraPrompts {
		fmt.Printf("Includes Kata&Hira prompts, w/romaji objectives: %t \n\n", limitedToKataPrompts)
	} else if limitedToKataPrompts {
		fmt.Printf("Limited to Kata prompts: %t \n\n", limitedToKataPrompts)
	} else if kata_hira {
		fmt.Printf("Limited to Kata_hira prompt_response: %t \n\n", kata_hira)
	} else if kata_roma {
		fmt.Printf("Limited to Kata_roma prompt_response: %t \n\n", kata_roma)
	} else if limitedToHiraPrompts {
		fmt.Printf("Limited to Hira prompts: %t \n\n", limitedToHiraPrompts)
	} else if limitedToRomaPrompts {
		fmt.Printf("Limited to Romaji prompts: %t \n\n", limitedToRomaPrompts)
	} else if limitedToDifficultDescriptions {
		fmt.Printf("Limited to Difficult descriptive prompts only (Romaji as objective): %t \n\n", limitedToDifficultDescriptions)
	} else {
		fmt.Printf("Standard Mix, says rix\n\n")
	}
}
