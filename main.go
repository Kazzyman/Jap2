package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// var in string

func pick_RandomCard_Assign_aCard() (prompt_fieldOfCard string) { // - -
	randIndex := rand.Intn(len(fileOfCards))
	randIndexS := rand.Intn(len(fileOfCardsS))
	randIndexD := rand.Intn(len(fileOfCardsMostDifficult))

	// ToDo: done: randomize between fileOfCards, fileOfCardsS, and fileOfCardsMostDifficult
	// ToDo: remove "prompt_fieldOfCard = aCard.Hira" lines after completing main todo
	randomFileOfCards := rand.Intn(3)
	// fmt.Printf("randomFileOfCards is %d \n", randomFileOfCards)
	// randomFileOfCards = 1
	if randomFileOfCards == 0 {
		aCard = fileOfCards[randIndex]  // randomly pick a 'card' from the 'deck' and store it in a global var
		prompt_fieldOfCard = aCard.Hira // prompt_fieldOfCard is also the prompt_string
	}
	if randomFileOfCards == 1 {
		aCard = fileOfCardsS[randIndexS] // randomly pick a 'card' from the 'deck' and store it in a global var
		prompt_fieldOfCard = aCard.Hira  // prompt_fieldOfCard is also the prompt_string
	}
	if randomFileOfCards == 2 {
		aCard = fileOfCardsMostDifficult[randIndexD] // randomly pick a 'card' from the 'deck' and store it in a global var
		prompt_fieldOfCard = aCard.Hira              // prompt_fieldOfCard is also the prompt_string
	}
	return prompt_fieldOfCard // prompt_fieldOfCard is also the prompt_string
}

func main() {
	rand.Seed(time.Now().UnixNano())
	// for {
	newPrompt_fieldOfCard := pick_RandomCard_Assign_aCard() // fieldOfCard is also the prompt_string
	// ToDo: randomize a prompt type or field from said Card, i.e., Hira, Kata, or Romaji
	begin(newPrompt_fieldOfCard) // fieldOfCard is also the prompt_string
	// }
}
func begin(prompt_fieldOfCard string) { // May be a Hira, Kata, or Romaji prompt  - -
	// fmt.Printf("We begin with FOC: %s \n", fieldOfCard)
	/*

		isThis_a_cardWeNeedMoreWorkOn := 0 // now and then we will consider working on a char the user has had trouble with
		// /
		// /
		// --- Here begins the main loop, which encapsulates the entirety of this func ---
		for {

			isThis_a_cardWeNeedMoreWorkOn++


			// in = prompt()


		outOfForLoop: // this loop is only needed because we want to handel the case of successive directives (tricky)
			// ESPECIALLY, this labeled 'for-loop' is used to handel successive directives WITHOUT DOUBLE PROCESSING of guesses
			for {
				if in == "set" || in == "?" || // <-- If it IS a Directive
					in == "??" || in == "menu" ||
					in == "reset" || in == "stat" ||
					in == "dir" || in == "notes" ||
					in == "quit" || in == "exit" ||
					in == "stats" || in == "rm" ||
					in == "stack" { // <-- If it IS a Directive: v v v
					// v v v v v   HANDEL A DIRECTIVE    v v v v v v v v v v v v v v   v v v v v v v
					branchOnUserSelectedDirectiveIfGiven(usersGuessOrOptionDirective, selectedExercise) // <-- Do Directive
					// v v v v v   RE-PROMPT   v v v v v v
					// -- Re-prompt following the execution of a Directive ---------------------------------------- v v v v
					// ****************************************************************************************************

					// in = prompt()

					// v v v v v  Do not process directives from the above re-prompting   v v v v v
					// which is to say, do not execute a meat func unless usersGuessOrOptionDirective is not a Directive
					if in != "set" &&
						in != "?" &&
						in != "??" &&
						in != "menu" &&
						in != "reset" &&
						in != "stat" &&
						in != "stats" &&
						in != "quit" &&
						in != "notes" &&
						in != "rm" &&
						in != "stack" &&
						in != "exit" &&
						in != "dir" {
				//  v ^ v ^ At this point we know that the usersGuessOrOptionDirective is not a Dir and should be considered to ba a guess,
						// ... AND we will, therefore, need to leave the loop after processing the user's guess
						// Based on the Selected Exercise, process the user's guess (determine if it is correct, etc.)

						// meatOf_Mixed_HiraKataExercise(usersGuessOrOptionDirective, true)

						break outOfForLoop
					} else { // It must be a successive Directive, so we continue to process it from the top of the loop,
						continue outOfForLoop // reiterate at labeled loop: CHECK THE RE-PROMPT FOR A SUCCESSIVE DIRECTIVE
					}
				} else {

					// meatOf_Mixed_HiraKataExercise(usersGuessOrOptionDirective, true)

					// It is probably a valid guess, AND we need to leave the loop after processing it
					break outOfForLoop
				}
			} // End of loop used to handel successive directives WITHOUT DOUBLE PROCESSING of guesses
		} // End of outer loop


	*/
	for {
		in := prompt(prompt_fieldOfCard) // fieldOfCard is also the prompt_string
		DetectedDirective := false
		DetectedDirective = testForDirective(in) // Sets true if a Directive was detected
		if DetectedDirective {
			respond_to_UserSuppliedDirective(in)
			continue // ... After Directive handling, re-prompt with the same fieldOfCard
		} else {
			// recursion false or recall true, means do rightOrOops
			evaluateUsersGuess(in, prompt_fieldOfCard, false, false, false) // do rightOrOops
			break                                                           // ... Having finished with all potential guessing, return to main ...
		}
	} // ... Returns to main to randomly-select the next fieldOfCard
	main()
}

func evaluateUsersGuess(in, prompt_fieldOfCard string, recursion, recall, skippOops bool) { // fieldOfCard is also the prompt_string  --
	fmt.Printf("top of evaluateUsersGuess\n")
	// Compare in to fieldOfCard with an if-else [Right! - Oops!] logging etc.
	// Allow for three guesses, third failed guess gives hints
	// Preface second and third prompt with "Try again", and "Last try", respectively
	// After an interim guess, call: prompt(prompt string) (in string)

	// After responding to a Directive, do a rightOrOops() for the current/original fieldOfCard
	if recursion {
		// do not
	} else {
		rightOrOops(in, skippOops, aCard) // This func may call: tryAgain() ... which may call: lastTry()
	}
	if recall {
		rightOrOops(in, skippOops, aCard) // This func may call: tryAgain() ... which may call: lastTry()
	} else {
		// do not
	}
	// recursion false or recall true, means that we do a rightOrOops
	/*
		// ToDo: this if and its recursion bool parameter should probably be omitted; because main also does a random pick
		if recursion == false { // After responding to a Directive, from the same/original fieldOfCard
			fieldOfCard = pick_RandomCard_Assign_aCard()
		}

	*/
	// In any case, prompt
	in = prompt(prompt_fieldOfCard) // fieldOfCard is also the prompt_string
	DetectedDirective := false
	DetectedDirective = testForDirective(in)
	if DetectedDirective {
		respond_to_UserSuppliedDirective(in)
		// Recursively ...
		// ... after responding to a Directive, prompt via recursion, from the same/original fieldOfCard
		evaluateUsersGuess(in, prompt_fieldOfCard, true, false, true) // skip: recursion false or recall true, means do rightOrOops
	}
	// else, go back to main for a random card etc. ????
	fmt.Printf("we are passing in as: %s, and fieldOfCard as: %s \n", in, prompt_fieldOfCard)
	evaluateUsersGuess(in, prompt_fieldOfCard, true, true, false) // do: recursion false or recall true, means do rightOrOops
	// maybe do a main() // ??? but let's explore rightOrOops() etc. first ???
}

func rightOrOops(in string, skippOops bool, aCard charSetStruct) {
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// *******************************************************************************************
	// Handle special prompt cases prior to doing the normal "if in == " processing, i.e.,
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************

	if aCard.Romaji == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:

			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.Hira)
			logHits_in_cyclicArrayHits("Right", aCard.Hira)

		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			// The following logger is commented-out because "in" was prob a typo
			// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
			logHits_in_cyclicArrayHits("Oops", aCard.Hira)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.Hira +
				":it was:" + aCard.Romaji + ":but you had guessed:" + in)
			// ToDo: call tryAgain() ... which may call: lastTry()
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if 'in' == aCard.KeyR"
		// ********************************************************
		if in == aCard.Romaji { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:

			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.Hira)
			logHits_in_cyclicArrayHits("Right", aCard.Hira)

			new_fieldOfCard := pick_RandomCard_Assign_aCard() // gets a new card etc
			in = prompt(new_fieldOfCard)                      // gets a new in, having prompted with the new field
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, new_fieldOfCard, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, new_fieldOfCard, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}

		} else {
			if skippOops {
				// do nothing
			} else {
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　^^Oops! ")
				// log the miss:
				// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH) // Because, it was prob a typo
				logHits_in_cyclicArrayHits("Oops", aCard.Hira)
				logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.Hira +
					":it was:" + aCard.Romaji + ":but you had guessed:" + in)
			}
			// ToDo: call tryAgain() ... which may call: lastTry()
			tryAgain(aCard.Hira)
		}
	}
}

func tryAgain(prompt_fieldOfCard string) {
	fmt.Printf("Try again \n")
	in := prompt(prompt_fieldOfCard)
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// *******************************************************************************************
	// Handle special prompt cases prior to doing the normal "if in == " processing, i.e.,
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************

	if aCard.Romaji == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:

			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.Hira)
			logHits_in_cyclicArrayHits("Right", aCard.Hira)

		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			// The following logger is commented-out because "in" was prob a typo
			// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
			logHits_in_cyclicArrayHits("Oops", aCard.Hira)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.Hira +
				":it was:" + aCard.Romaji + ":but you had guessed:" + in)
			// ToDo: call tryAgain() ... which may call: lastTry()
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if 'in' == aCard.KeyR"
		// ********************************************************
		if in == aCard.Romaji { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:

			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.Hira)
			logHits_in_cyclicArrayHits("Right", aCard.Hira)

			newPrompt_fieldOfCard := pick_RandomCard_Assign_aCard() // gets a new card etc
			in = prompt(newPrompt_fieldOfCard)                      // gets a new in, having prompted with the new field
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, newPrompt_fieldOfCard, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, newPrompt_fieldOfCard, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}

			/*
				fieldOfCard2 := pick_RandomCard_Assign_aCard() // gets a new card etc.
				in = prompt(fieldOfCard2)                      // gets a new in, having prompted with the new field
				evaluateUsersGuess(in, fieldOfCard2, true)     // runs this func again

			*/

		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops Again! ")
			// log the miss:
			// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH) // Because, it was prob a typo
			logHits_in_cyclicArrayHits("Oops", aCard.Hira)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.Hira +
				":it was:" + aCard.Romaji + ":but you had guessed:" + in)
			// ToDo: call: lastTry()
			lastTry(prompt_fieldOfCard)
		}
	}
}

func lastTry(prompt_fieldOfCard string) {
	fmt.Printf("Last Try! \n")
	in := prompt(prompt_fieldOfCard)
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove bool
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	// *******************************************************************************************
	// Handle special prompt cases prior to doing the normal "if in == " processing, i.e.,
	// One 'very'-special key handler to emphasize the sameness of two variants of the zu sound
	// *******************************************************************************************

	if aCard.Romaji == "zu" { // yes, at this point fieldOfCard is actually aCard.Romaji
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// log the hit:

			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.Hira)
			logHits_in_cyclicArrayHits("Right", aCard.Hira)

		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
			// log the miss:
			// The following logger is commented-out because "in" was prob a typo
			// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH)
			logHits_in_cyclicArrayHits("Oops", aCard.Hira)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.Hira +
				":it was:" + aCard.Romaji + ":but you had guessed:" + in)
			// ToDo: call tryAgain() ... which may call: lastTry()
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// ********************************************************
		// 'else', no special cases were found, so we process normal cases of "if 'in' == aCard.KeyR"
		// ********************************************************
		if in == aCard.Romaji { // if 'in' is the appropriate Romaji to match the hira or kata prompt
			fmt.Printf("%s", colorGreen)
			fmt.Printf("     　^^Right! \n") // intentional '\n'
			fmt.Printf("%s", colorReset)
			// log the hit:

			logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(aCard.Hira)
			logHits_in_cyclicArrayHits("Right", aCard.Hira)

			newPrompt_fieldOfCard := pick_RandomCard_Assign_aCard() // gets a new card etc
			in = prompt(newPrompt_fieldOfCard)                      // gets a new in, having prompted with the new field
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				// Recursively ...
				// ... after responding to a Directive, do rightOrOops() from the same/original fieldOfCard
				evaluateUsersGuess(in, newPrompt_fieldOfCard, true, false, true) // skip: recursion false or recall true, means do rightOrOops
			} else {
				evaluateUsersGuess(in, newPrompt_fieldOfCard, true, true, false) // do: recursion false or recall true, means do rightOrOops
			}
			/*
				fieldOfCard3 := pick_RandomCard_Assign_aCard() // gets a new card etc.
				in = prompt(fieldOfCard3)                      // gets a new in, having prompted with the new field
				evaluateUsersGuess(in, fieldOfCard3, true)     // runs this func again

			*/

		} else {
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^That was your last try, Oops! ")
			// log the miss:
			// logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(aCard.KeyH) // Because, it was prob a typo
			logHits_in_cyclicArrayHits("Oops", aCard.Hira)
			logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(aCard.Hira +
				":it was:" + aCard.Romaji + ":but you had guessed:" + in)
			fmt.Printf("Here is your hint \n")
		}
	}
	main()
}

func prompt(fieldOfCard string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", fieldOfCard) // fieldOfCard is also the prompt_string
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Type the corresponding Romaji, or '?' for help with: %s \n", fieldOfCard)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

func testForDirective(in string) (result bool) {
	if in == "set" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "menu" ||
		in == "reset" ||
		in == "stat" ||
		in == "dir" ||
		in == "notes" ||
		in == "quit" ||
		in == "exit" ||
		in == "stats" ||
		in == "rm" ||
		in == "stack" {
		// Then:
		result = true
	}
	return result
}

func respond_to_UserSuppliedDirective(in string) {
	switch in {
	case "menu":
		// Flush the old stats and hits arrays
		fmt.Println("menu Directive handled")
	case "reset":
		// Flush (clear) the old stats and hits arrays
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// Also, flush (clear) the maps
		frequencyMapOf_IsFineOnChars = make(map[string]int)
		frequencyMapOf_need_workOn = make(map[string]int)
		//
		//goland:noinspection ALL
		fmt.Println("\nArrays and maps flushed:\n")
		fmt.Println("cyclicArrayOfTheJcharsGottenWrong")
		fmt.Println("cyclicArrayHits")
		fmt.Println("frequencyMapOf_IsFineOnChars")
		//goland:noinspection ALL
		fmt.Println("frequencyMapOf_need_workOn\n")
	case "quit":
		os.Exit(1)
	case "exit":
		os.Exit(1)
	case "??": // Directives follow:
		handle_doubleQuestMark_directive()
	case "?":

		fmt.Printf("\n%s\n%s\n%s\n\n", aCardA.HiraHint, aCardA.KataHint, aCardA.TT_Hint)
		/*
			HiraHint   string
			KataHint   string
			TT_Hint    string
			SansR_Hint string
		*/
	case "set":

		reSet_aCard_andThereBy_reSet_thePromptString()

	case "stat":
		hits()
	case "stats":
		hits()
	case "notes":
		//goland:noinspection ALL  **do-this**
		fmt.Println("\nIn the traditional Hepburn romanization system, the sound じ in hiragana is romanized as \"ji\" \n" +
			"and the katakana ジ is also romanized as \"ji\" \n\n" +
			"However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as\n" +
			" \"zi\" instead of \"ji\"\n\n" +
			"The sound gi:ぎ in hiragana is romanized as \"gi\" and the katakana ギ is also romanized as \"gi\"\n")
		//goland:noinspection ALL  **do-this**
		fmt.Println("゜is called \"handakuten\" 半濁点 translates to \"half-voicing mark\" or \"semi-voiced mark\"\n" +
			"゛is called \"dakuten\" 濁点 meaning 'voiced mark' or 'voicing mark'")
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):

		body_of_instructions_for_Romaji_responces_only()
		fmt.Println("You are doing Exercise 6")

	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
	case "stack":
		// Load six occurrences of 'shi' to the map_of_fineOn
		stack_the_map()
	default:
		// fmt.Println("Directive not found") // Does not work because only existent cases are passed to the switch
	}
}
