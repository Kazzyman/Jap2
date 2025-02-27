package main

import "fmt"

func displayRight_logRight(userInput, promptField, actual_objective_type string) { // ::: - -
	// log_right_andUpdateGame(promptField, userInput)
	fmt.Printf("%s", colorGreen)

	// if objective was of type hira, the prompt could have been either roma or kata
	// ::: hira objective
	if actual_objective_type == "hira" { // actual_objective_type == "Hira"
		if actual_prompt_char_type == "kata" {
			fmt.Printf("      　%s  - %s\n", aCard.Romaji, aCard.HiraHint)
		} else { // type of prompt must have been of Romaji type
			fmt.Printf("      　%s  - %s\n", aCard.Kata, aCard.HiraHint)
		}
	} else { // else the objective type is Romaji, so:
		// ::: roma objective
		if limitedToDifficultDescriptions == true {
			fmt.Printf("      　%s %s  - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
		} else { // if objective was of type roma, the prompt could have been either hira or kata
			if actual_prompt_char_type == "kata" { //
				// ::: kata prompt
				fmt.Printf("      　%s  - %s\n", aCard.Hira, aCard.HiraHint)
			} else { // type of prompt must have been of hira type, and the objective was a roma
				// ::: hira prompt
				fmt.Printf("      　%s  - %s\n", aCard.Kata, aCard.HiraHint)
			}
			// fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint) // todo ? aCard.SansR_Hint
		}
	}
	fmt.Printf("%s", colorReset)
	log_right_andUpdateGame(promptField, userInput)

}

func logRight_zu(userInput, promptField, actual_objective_type string) { // ::: - -
	// log_right_andUpdateGame(promptField, userInput)
	fmt.Printf("%s", colorGreen)
	if actual_objective_type == "hira" {
		fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.HiraHint)
	} else { // else it is Romaji, and ...
		if limitedToDifficultDescriptions == true {
			fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
		} else { // ... This correct guess must be a REGULAR Kata
			fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
		}
	}
	// fmt.Printf("%s", colorReset)
	fmt.Printf("         ... it could have been either ず or づ as they are the same sound: zu\n")
	fmt.Printf("%s", colorReset)
	log_right_andUpdateGame(promptField, userInput)

}

func logRightZu2(userInput, promptField, actual_objective_type, objective string) { // ::: - -
	// log_right_andUpdateGame(promptField, userInput)
	fmt.Printf("%s", colorGreen)
	if actual_objective_type == "hira" {
		fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.HiraHint)
		fmt.Printf("         ... it could have been either ず or %s since they have the same sound!\n", objective)
	} else { // it can be assumed to have been of roma type
		if limitedToDifficultDescriptions == true {
			fmt.Printf("      　%s %s   - \n", aCard.Hira, aCard.SansR_Hint)
			fmt.Printf("         ... it could have been either ず or %s since they have the same sound!\n", objective)
		} else {
			fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
			fmt.Printf("         ... it could have been either ず or %s since they have the same sound!\n", objective)
		}
	}
	fmt.Printf("%s", colorReset)
	log_right_andUpdateGame(promptField, userInput)

}

func beginGame() {
	fmt.Println("What is your first name? (one word)")
	_, _ = fmt.Scan(&nameOfPlayer)
	List_of_game_types()
	fmt.Println(colorRed + "Select a game" + colorReset)
	_, _ = fmt.Scan(&type_of_game)
	if type_of_game != "1" && type_of_game != "2" {
		fmt.Printf("\n%s is not a valid entry\n", type_of_game)
		fmt.Println(colorRed + "Select a game from the list (1-9)" + colorReset)
		_, _ = fmt.Scan(&type_of_game)
	}
	fmt.Println(colorRed + "Enter proposed game duration, or 0 for a complete game & Points to console" + colorReset)
	_, _ = fmt.Scan(&gameDuration)
	if gameDuration > 0 {
		// gameDuration will be set (below) to the combined length of the decks involved
	} else {
		if type_of_game == "1" {
			gameDuration = 2*len(fileOfCardsHiraKata) + (2 * len(fileOfCardsEasyKanji)) + len(fileOfCardsKanjiHard)
		}
	}
	switch type_of_game {
	case "1": // hko Use Kata & Hira prompting (Roma objectives)
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = true
		limitedToHiraPrompts = true
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "2": // konly Use only Kata prompting (mix Hira & Roma objectives)
		gameDuration = 2 * len(fileOfCardsHiraKata)
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "3": // honly Use only Hira prompting (so only Roma responses)
		gameDuration = len(fileOfCardsHiraKata)
		kata_hira = false
		kata_roma = false
		limitedToHiraPrompts = true
		limitedToKataPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "4": // ronly Use only Romaji prompting (so only Hira responses)
		gameDuration = len(fileOfCardsHiraKata)
		kata_hira = false
		kata_roma = false
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "5": // khSimplex : Use Kata prompting & Simplex Hira obj
		gameDuration = len(fileOfCards_nonCompound)
		kata_roma = false
		kata_hira = false
		limitedToKataPromptsAndSimplexHiraObj = true
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "6": // donly : Difficult descriptive prompting
		gameDuration = len(dataMostDiff)
		kata_hira = false
		kata_roma = false
		limitedToDifficultDescriptions = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = true
		limitedToSpelling = false
	case "7": // kh : Use only kata_hira prompt_response
		gameDuration = 2 * len(fileOfCardsHiraKata)
		kata_roma = false
		kata_hira = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "8": // rhSimplex : Use kata prompts w/Simplex Hira
		gameDuration = len(fileOfCards_nonCompound)
		kata_hira = false // ::: should be true ??
		kata_roma = false
		limitedToRomaPromptsAndSimplexHiraObj = true // ::: ?
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	case "9": // kr : Use only kata_roma prompt_response
		gameDuration = len(fileOfCardsHiraKata)
		kata_hira = false
		kata_roma = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
	default: // "0" // mix ::: NOT SURE ABOUT ALL THESE FALSE SETTINGS FOR THE MIX CASE ??
		gameDuration = 3*len(fileOfCardsHiraKata) + (2 * len(fileOfCardsEasyKanji)) + len(fileOfCardsKanjiHard) // ::: 3 *
		kata_hira = false
		kata_roma = false
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultDescriptions = false
		limitedToSpelling = false
		// ::: nothing limited means prompt with a mix of all three
	}
	display_limited_gaming_dir_list()
	now_using_game_duration_set_by_game_type = true
	the_game_begins()
}
