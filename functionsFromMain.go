package main

import "fmt"

func promptFirstCase(promptField, actual_objective_type string) (userInput string) {
	if actual_objective_type == "Romaji" {
		userInput = promptForRomajiWithDir(promptField)
	} else if actual_objective_type == "Extended_Romaji" {
		userInput = promptForRomajiWithDirE(promptField) // A special prompt for Extended Kata, if|when deployed.
	} else if actual_objective_type == "Hira" {
		userInput = promptForHiraWithDir(promptField)
	}
	return userInput
}

func promptSecondCase(new_prompt, actual_objective_type string) (userInput string) {
	if actual_objective_type == "Romaji" {
		userInput = promptForRomajiWithDir(new_prompt)
	} else if actual_objective_type == "Extended_Romaji" {
		userInput = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed.
	} else if actual_objective_type == "Hira" {
		userInput = promptForHiraWithDir(new_prompt)
	}
	return userInput
}

func olddirectiveHandler(userInput, inew_prompt, inew_objective, iactual_objective_type string) (new_prompt, new_objective, actual_objective_type string) {
	aDirectiveWasDetected := false
	aDirectiveWasDetected = detectDirective(userInput)
	if aDirectiveWasDetected {
		if userInput == "stc" || userInput == "stcr" {
			new_prompt, new_objective, actual_objective_type = respond_to_UserSuppliedDirective(userInput)
		} else {
			respond_to_UserSuppliedDirective(userInput)
		}
		// evaluateUsersGuess(userInput, inew_prompt, inew_objective, iactual_objective_type, true, false)
	} else {
		// evaluateUsersGuess(userInput, inew_prompt, inew_objective, iactual_objective_type, true, true)
	}
	return new_prompt, new_objective, actual_objective_type
}

func logRight_zu(userInput, promptField, actual_objective_type string) {
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

func logRight(userInput, promptField, actual_objective_type string) {
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

func logOopsLoser(usersInput string) {
	log_oops(aCard.Hira, aCard.Romaji, usersInput)
	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
}

func logRightZu2(userInput, promptField, actual_objective_type, objective string) {
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
