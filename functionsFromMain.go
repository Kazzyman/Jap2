package main

import "fmt"

func displayRight_logRight(userInput, promptField, actual_objective_type string) { // ::: - -
	log_right_andUpdateGame(promptField, userInput)
	fmt.Printf("%s", colorGreen)

	if actual_objective_type == "hira" { // actual_objective_type == "Hira"
		fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.HiraHint)
	} else { // else it is Romaji, so:
		if limitedToDifficultDescriptions == true {
			fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
		} else { // Then this correct guess must be a REGULAR Kata
			fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
		}
	}
	fmt.Printf("%s", colorReset)
}

func logRight_zu(userInput, promptField, actual_objective_type string) { // ::: - -
	log_right_andUpdateGame(promptField, userInput)
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
	fmt.Printf("%s", colorReset)
	fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
}

func logRightZu2(userInput, promptField, actual_objective_type, objective string) { // ::: - -
	log_right_andUpdateGame(promptField, userInput)
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
}
