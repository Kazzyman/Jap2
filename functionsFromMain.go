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
	fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
	fmt.Printf("%s", colorReset)
	log_right_andUpdateGame(promptField, userInput)

}

func logRightZu2(userInput, promptField, actual_objective_type, objective string) { // ::: - -
	// log_right_andUpdateGame(promptField, userInput)
	fmt.Printf("%s", colorGreen)
	if actual_objective_type == "hira" {
		fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.HiraHint)
		fmt.Printf("... it could have been either ず or %s since they have the same sound!\n", objective)
	} else { // it can be assumed to have been of roma type
		if limitedToDifficultDescriptions == true {
			fmt.Printf("      　%s %s   - \n", aCard.Hira, aCard.SansR_Hint)
			fmt.Printf("... it could have been either ず or %s since they have the same sound!\n", objective)
		} else {
			fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
			fmt.Printf("... it could have been either ず or %s since they have the same sound!\n", objective)
		}
	}
	fmt.Printf("%s", colorReset)
	log_right_andUpdateGame(promptField, userInput)

}
