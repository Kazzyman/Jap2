package main

import "fmt"

func promptFirstCase(objective_kind, promptField string) (in string) {
	if objective_kind == "Romaji" {
		in = promptForRomajiWithDir(promptField)
	} else if objective_kind == "Extended_Romaji" {
		in = promptForRomajiWithDirE(promptField) // A special prompt for Extended Kata, if|when deployed.
	} else if objective_kind == "Hira" {
		in = promptForHiraWithDir(promptField)
	}
	return in
}

func promptSecondCase(new_objective_kind, new_prompt string) (in string) {
	if new_objective_kind == "Romaji" {
		in = promptForRomajiWithDir(new_prompt)
	} else if new_objective_kind == "Extended_Romaji" {
		in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed.
	} else if new_objective_kind == "Hira" {
		in = promptForHiraWithDir(new_prompt)
	}
	return in
}

func directiveHandler(in, inew_prompt, iobjective_kind, inew_objective, inew_objective_kind string) (new_prompt, new_objective, new_objective_kind string) {
	aDirectiveWasDetected := false
	aDirectiveWasDetected = detectDirective(in)
	if aDirectiveWasDetected {
		if in == "stc" || in == "stcr" {
			new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, iobjective_kind)
		} else {
			respond_to_UserSuppliedDirective(in, iobjective_kind)
		}
		evaluateUsersGuess(in, inew_prompt, inew_objective, inew_objective_kind, true, false, true)
	} else {
		evaluateUsersGuess(in, inew_prompt, inew_objective, inew_objective_kind, true, true, false)
	}
	return new_prompt, new_objective, new_objective_kind
}

func logRight_zu(in, promptField, objective_kind string) {
	log_right(promptField, in)
	fmt.Printf("%s", colorGreen)
	if objective_kind == "Hira" {
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

func logRight(in, promptField, objective_kind string) {
	// logRight(in, promptField, objective_kind)
	log_right(promptField, in)
	fmt.Printf("%s", colorGreen)
	if objective_kind == "Hira" {
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

func logOopsLoser(in string) {
	log_oops(aCard.Hira, aCard.Romaji, in)
	fmt.Printf("%s", colorRed)
	fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
	fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
}

func logRightZu2(in, promptField, objective_kind, objective string) {
	log_right(promptField, in)
	fmt.Printf("%s", colorGreen)
	if objective_kind == "Hira" {
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
