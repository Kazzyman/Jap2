package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// pick returns: promptField, objective, objective_kind
func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	randIndex := rand.Intn(len(fileOfCards))
	randIndexS := rand.Intn(len(fileOfCardsS))
	randIndexD := rand.Intn(len(fileOfCardsMostDifficult))

	randomFileOfCards := rand.Intn(12)

	// Hira prompting, Romaji objective:
	if randomFileOfCards == 0 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
	}
	if randomFileOfCards == 1 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
	}
	if randomFileOfCards == 2 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
	}

	// Kata prompting, Romaji objective:
	if randomFileOfCards == 3 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
	}
	if randomFileOfCards == 4 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
	}
	if randomFileOfCards == 5 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
	}

	// Romaji prompting, Hira objective:
	if randomFileOfCards == 6 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
	}
	if randomFileOfCards == 7 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
	}
	if randomFileOfCards == 8 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
	}

	// Kata prompting, Hira objective:
	if randomFileOfCards == 9 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
	}
	if randomFileOfCards == 10 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
	}
	if randomFileOfCards == 11 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
	}
	return promptField, objective, objective_kind
}

var isThis_a_cardWeNeedMoreWorkOn int

func checkOops_map(promptField, objective, objective_kind string) {
	// Go get a card we've missed before, instead of the random one
	// isThis_a_cardWeNeedMoreWorkOn := 0 // Now & then we will consider working on a char the user has had trouble with
	isThis_a_cardWeNeedMoreWorkOn++
	// If we have gone a while without augmenting the random picks with cards we have prev missed ...
	if isThis_a_cardWeNeedMoreWorkOn > 2 {
		// Log to a file that this action was taken
		fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // Append to file
		check(err)
		_, err2 := fmt.Fprintf(fileHandle, "\nChecked card:%s in the frequencyMapOf_need_workOn after:%d cycles\n",
			promptField, isThis_a_cardWeNeedMoreWorkOn)
		check(err2)
		_ = fileHandle.Close() // Close the file, discard any resulting errors returned
		isThis_a_cardWeNeedMoreWorkOn = 0
		check_it_for_needing_more_practice(promptField, objective, objective_kind)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	newPrompt_field, objective, objective_kind := pick_RandomCard_Assign_fields()
	checkOops_map(newPrompt_field, objective, objective_kind)
	begin(newPrompt_field, objective, objective_kind)
}
func begin(promptField, objective, objective_kind string) { // May be a Hira, Kata, or Romaji prompt  - -
	// In any case:
	// check_it_for_fine_on6() // Attempt to check for both Hira & Romaji ?? (probably meant: both Hira & Kata)
	check_it_for_fine_on() // Checks for Romaji only??
	var in string
	for {
		if objective_kind == "Romaji" {
			in = promptForRomaji(promptField) // Get user's input, from a randomly selected prompt
		} else if objective_kind == "Hira" {
			in = promptForHira(promptField) // Get user's input, from a randomly selected prompt
		}
		DetectedDirective := false
		DetectedDirective = testForDirective(in) // Sets true if a Directive was detected
		if DetectedDirective {
			respond_to_UserSuppliedDirective(in)
			continue // ... After Directive handling, re-prompt with the same promptField
		} else {
			// Recursion false, or recall true, means do rightOrOops()
			evaluateUsersGuess(in, promptField, objective, objective_kind, false, false, false) // f,f,f Does rightOrOops
			break                                                                               // ... Having finished with all potential guessing, return to main ...
		}
	} // ... Returns to main() to randomly-select the next fieldOfCard
	main()
}

func evaluateUsersGuess(in, promptField, objective, objective_kind string, recursion, recall, skippOops bool) { // - -
	if recursion {
		// do nothing
	} else {
		rightOrOops(in, promptField, objective, objective_kind, skippOops) // This func may call: tryAgain() ... which may call: lastTry()
	}
	if recall {
		rightOrOops(in, promptField, objective, objective_kind, skippOops) // This func may call: tryAgain() ... which may call: lastTry()
	} else {
		// do nothing
	}
	// In any case, prompt
	if objective_kind == "Romaji" {
		in = promptForRomaji(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Hira" {
		in = promptForHira(promptField) // Get user's input, from a randomly selected prompt
	}
	DetectedDirective := false
	DetectedDirective = testForDirective(in)
	if DetectedDirective {
		respond_to_UserSuppliedDirective(in)
		// Recursively ...
		// ... after responding to a Directive, prompt via recursion, from the same/original promptField
		// skip: recursion false or recall true, means do rightOrOops
		evaluateUsersGuess(in, promptField, objective, objective_kind, true, false, true) // t,f,t Skips rightOrOops
	}
	// else, go back to main for a random card etc. ????
	// do: recursion false or recall true, means do rightOrOops
	evaluateUsersGuess(in, promptField, objective, objective_kind, true, true, false) // t,t,f Does rightOrOops
	// Maybe do a main() here??? But let's explore rightOrOops() etc. first???
}

// pick returns: promptField, objective, objective_kind
func rightOrOops(in, promptField, objective, objective_kind string, skippOops bool) {
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			checkOops_map(new_prompt, new_objective, new_objective_kind)
			if new_objective_kind == "Romaji" {
				in = promptForRomaji(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHira(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		if in == objective { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			checkOops_map(new_prompt, new_objective, new_objective_kind)
			if new_objective_kind == "Romaji" {
				in = promptForRomaji(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHira(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}

		} else {
			if skippOops {
				// do nothing
			} else {
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　^^Oops! ")
			}
			tryAgain(promptField, objective, objective_kind)
		}
	}
}

// pick returns: promptField, objective, objective_kind
func tryAgain(promptField, objective, objective_kind string) {
	fmt.Printf("Try again \n")
	var in string
	if objective_kind == "Romaji" {
		in = promptForRomaji(promptField) // Gets a new in, having prompted with the new field
	} else if objective_kind == "Hira" {
		in = promptForHira(promptField) // Gets a new in, having prompted with the new field
	}

	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			checkOops_map(new_prompt, new_objective, new_objective_kind)
			if new_objective_kind == "Romaji" {
				in = promptForRomaji(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHira(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}
		} else {
			// pick returns: promptField, objective, objective_kind
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		if in == objective { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			checkOops_map(new_prompt, new_objective, new_objective_kind)
			if new_objective_kind == "Romaji" {
				in = promptForRomaji(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHira(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops Again! ")
			lastTry(promptField, objective, objective_kind)
		}
	}
}

func lastTry(promptField, objective, objective_kind string) {
	fmt.Printf("Last Try! \n")
	var in string
	if objective_kind == "Romaji" {
		in = promptForRomaji(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Hira" {
		in = promptForHira(promptField) // Get user's input, from a randomly selected prompt
	}
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	if objective == "zu" { // yes, at this point fieldOfCard is actually aCard.Romaji
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			checkOops_map(new_prompt, new_objective, new_objective_kind)
			if new_objective_kind == "Romaji" {
				in = promptForRomaji(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHira(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		if in == objective { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			checkOops_map(new_prompt, new_objective, new_objective_kind)
			if new_objective_kind == "Romaji" {
				in = promptForRomaji(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHira(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}
		} else {
			log_oops(aCard.Hira, aCard.Romaji, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^That was your last try, Oops! ")
			fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
		}
	}
	main()
}

func promptForRomaji(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField) // fieldOfCard is also the prompt_string
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected,\n Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

func promptForHira(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField) // fieldOfCard is also the prompt_string
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Hiragana input-mode expected,\n Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}
