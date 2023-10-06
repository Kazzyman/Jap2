package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Occasionally we want to get a card that has been missed before, instead of the random one
// i.e., now & then we will consider working on a char the user has had trouble with, rather than a random char
var isThis_a_cardWeNeedMoreWorkOn int

func main() {
	rand.Seed(time.Now().UnixNano())
	gameOn = false
	for {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		if game_loop_counter > game_duration {
			game_off()
		}

		new_prompt, objective, objective_kind := pick_RandomCard_Assign_fields() // This line is done after each ^^Right!
		// After each: pick_RandomCard_Assign_fields() [except this ^ ^ one] we will do a checkMemory()
		begin(new_prompt, objective, objective_kind)
	}
}

var game_duration = 1000 // for testing this will be 10, but in future it should be 800
func begin(promptField, objective, objective_kind string) { // May be a Hira, Kata, or Romaji prompt  - -

	if game_loop_counter > game_duration {
		game_off()
	}
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature
	for {
		if objective_kind == "Romaji" {
			in = promptForRomajiWithDir(promptField) // Get user's input, from a randomly selected prompt
		} else if objective_kind == "Hira" {
			in = promptForHiraWithDir(promptField) // Get user's input, from a randomly selected prompt
		}
		DetectedDirective := false
		DetectedDirective = testForDirective(in) // Sets DetectedDirective true if a "Directive" was detected
		if DetectedDirective {
			respond_to_UserSuppliedDirective(in)
			continue // ... After "Directive" handling, re-prompt with the same/original promptField
		} else {
			// Passing recursion false [or recall true], means do rightOrOops()
			// ...                                                                  f,f,f Does rightOrOops
			evaluateUsersGuess(in, promptField, objective, objective_kind, false, false, false)
			break // ... Having finished with all potential guessing, return to main ...
		}
	} // ... Returns to main()'s inner loop; to randomly-select the next fieldOfCard ::

}

func evaluateUsersGuess(in, promptField, objective, objective_kind string, recursion, recall, skipOops bool) { // - -
	if game_loop_counter > game_duration {
		game_off()
	}
	/*
		This next construct is strange! Because, it seems to allow for rightOrOops() to be done "twice" -- but it does not!
		since rightOrOops() sets in motion a chain of events that never returns directly here where it was called ...
		duplication here never occurs. The construct is needed to allow for rightOrOops() to be omitted if
		evaluateUsersGuess() has been called with recursion true and recall false
	*/
	if recursion {
		// If recursion is true, then do nothing
	} else {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		rightOrOops(in, promptField, objective, objective_kind, skipOops) // This func may call: tryAgain() ... which may call: lastTry()
	}
	if recall {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		rightOrOops(in, promptField, objective, objective_kind, skipOops) // This func may call: tryAgain() ... which may call: lastTry()
	} else {
		// If recall is false, then do nothing
	}
	// ^ ^ ^ If evaluateUsersGuess() has been called after handling a "Directive" then rightOrOops() is omitted entirely

	if objective_kind == "Romaji" {
		in = promptForRomajiWithDir(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Hira" {
		in = promptForHiraWithDir(promptField) // Get user's input, from a randomly selected prompt
	}
	DetectedDirective := false
	DetectedDirective = testForDirective(in)
	if DetectedDirective {
		respond_to_UserSuppliedDirective(in)
		/*
			Recursively ...
			after responding to a Directive, prompt via recursion, from the same/original promptField
			With recursion true & recall false, rightOrOops() will NOT be done with this recursion,
			and, with skipOops true [and passed to rightOrOops()] rightOrOops() will not say ^^Oops! after this recursion
			...                                                                        t, f, t Skips rightOrOops
		*/
		evaluateUsersGuess(in, promptField, objective, objective_kind, true, false, true) //
	} else {
		/*
			Do a normal, i.e., unconditional recursion with skipOops set to false
			via recall==true & skipOops==false [recursion false or recall true, means do rightOrOops]
		*/
		evaluateUsersGuess(in, promptField, objective, objective_kind, true, true, false) // t,t,f Does rightOrOops
	}
}

func rightOrOops(in, promptField, objective, objective_kind string, skipOops bool) {
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			if game == "off" {
				promptField, objective, objective_kind = checkMemory(new_prompt, new_objective, new_objective_kind)
			}
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt) // Gets a new in, having prompted with the new field
			}
			// Refer to the previous comments about the following:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			if skipOops {
				// Then do nothing
			} else {
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				fmt.Printf("       ^^Oops! ")
			}
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		if in == objective {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      　^^Right! \n")
			fmt.Printf("%s", colorReset)
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			if game == "off" {
				promptField, objective, objective_kind = checkMemory(new_prompt, new_objective, new_objective_kind)
			}
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt) // Gets a new in, having prompted with the new field
			}
			// Refer to the previous comments about the following:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			if skipOops {
				// Then do nothing
			} else {
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　^^Oops! ")
			}
			tryAgain(promptField, objective, objective_kind)
		}
	}
}

func tryAgain(promptField, objective, objective_kind string) {
	fmt.Printf("Try again \n")
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature
	if objective_kind == "Romaji" {
		in = promptForRomaji(promptField) // Gets a new in, having prompted with the new field
	} else if objective_kind == "Hira" {
		in = promptForHira(promptField) // Gets a new in, having prompted with the new field
	}

	var thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			if game == "off" {
				promptField, objective, objective_kind = checkMemory(new_prompt, new_objective, new_objective_kind)
			}
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt) // Gets a new in, having prompted with the new field
			}
			// Refer to the previous comments about the following:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		if in == objective {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      　^^Right! \n")
			fmt.Printf("%s", colorReset)
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			if game == "off" {
				promptField, objective, objective_kind = checkMemory(new_prompt, new_objective, new_objective_kind)
			}
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt) // Gets a new in, having prompted with the new field
			}
			// Refer to the previous comments about the following:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^Oops Again! ")
			lastTry(promptField, objective, objective_kind)
		}
	}
}

func lastTry(promptField, objective, objective_kind string) { // - -
	fmt.Printf("Last Try! \n")
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature
	if objective_kind == "Romaji" {
		in = promptForRomaji(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Hira" {
		in = promptForHira(promptField) // Get user's input, from a randomly selected prompt
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
			if game == "off" {
				promptField, objective, objective_kind = checkMemory(new_prompt, new_objective, new_objective_kind)
			}
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! ")
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		if in == objective {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("      　^^Right! \n")
			fmt.Printf("%s", colorReset)
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			if game == "off" {
				promptField, objective, objective_kind = checkMemory(new_prompt, new_objective, new_objective_kind)
			}
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Gets a new in, having prompted with the new field
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt) // Gets a new in, having prompted with the new field
			}
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				respond_to_UserSuppliedDirective(in)
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			log_oops(aCard.Hira, aCard.Romaji, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("      　^^That was your last try, Oops! ")
			fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
		}
	}
}
