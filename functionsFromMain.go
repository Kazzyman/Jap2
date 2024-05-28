package main

import "fmt"

func logRight_zu(userInput, promptField, actual_objective_type string) { // ::: - -
	log_right(promptField, userInput)
	fmt.Printf("%s", colorGreen)
	if actual_objective_type == "Hira" {
		fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.SansR_Hint)
	} else { // else it is Romaji, and ...
		if limitedToDifficultKata == true {
			fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
		} else { // ... This correct guess must be a REGULAR Kata
			fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
		}
	}
	fmt.Printf("%s", colorReset)
	fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
}

func logRight(userInput, promptField, actual_objective_type string) { // ::: - -
	// logRight(userInput, promptField, actual_objective_type)
	log_right(promptField, userInput)
	fmt.Printf("%s", colorGreen)
	if actual_objective_type == "Hira" {
		fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.SansR_Hint)
	} else { // else it is Romaji, so:
		if limitedToDifficultKata == true {
			fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
		} else { // Then this correct guess must be a REGULAR Kata
			fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
		}
	}
	fmt.Printf("%s", colorReset)
}

func logRightZu2(userInput, promptField, actual_objective_type, objective string) { // ::: - -
	log_right(promptField, userInput)
	fmt.Printf("%s", colorGreen)
	if actual_objective_type == "Hira" {
		fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.SansR_Hint)
		fmt.Printf("... it could have been either ず or %s since they have the same sound!\n", objective)
	} else {
		if limitedToDifficultKata == true {
			fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
			fmt.Printf("... it could have been either ず or %s since they have the same sound!\n", objective)
		} else {
			fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
			fmt.Printf("... it could have been either ず or %s since they have the same sound!\n", objective)
		}
	}
	fmt.Printf("%s", colorReset)
}
