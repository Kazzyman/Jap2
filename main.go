package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	gameOn = false
	game = "off"
	display_List_of_Directives()
	for {
		if gameOn {
			// game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		if game_loop_counter > game_duration {
			game_off()
		}
		new_prompt, objective, objective_kind := pick_RandomCard_Assign_fields() // This line is done after each ^^Right!
		begin(new_prompt, objective, objective_kind)
	}
}

func begin(promptField, objective, objective_kind string) { // May be a Hira, Kata, or Romaji prompt  - -
	if game_loop_counter > game_duration {
		game_off()
	}
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature
	for {
		// These prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
		if objective_kind == "Romaji" {
			in = promptForRomajiWithDir(promptField) // Get user's input, from a randomly selected prompt
		} else if objective_kind == "Extended_Romaji" {
			in = promptForRomajiWithDirE(promptField) // A special prompt for Extended Kata, if|when deployed
		} else if objective_kind == "Hira" {
			in = promptForHiraWithDir(promptField)
		}

		DetectedDirective := false
		DetectedDirective = testForDirective(in) // Sets DetectedDirective true if a "Directive" was detected
		if DetectedDirective {
			if in == "set" { // respond_to_UserSuppliedDirective(in, new_objective_kind) will want to return values is "set" is switched on
				promptField, objective, objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
			} else {
				respond_to_UserSuppliedDirective(in, objective_kind)
			}
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
	// These prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	if objective_kind == "Romaji" {
		in = promptForRomajiWithDir(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Extended_Romaji" {
		in = promptForRomajiWithDirE(promptField) // A special prompt for Extended Kata, if|when deployed
	} else if objective_kind == "Hira" {
		in = promptForHiraWithDir(promptField)
	}

	DetectedDirective := false
	DetectedDirective = testForDirective(in)
	if DetectedDirective {
		if in == "set" { // See prior comments
			promptField, objective, objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
		} else {
			respond_to_UserSuppliedDirective(in, objective_kind)
		}
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

func rightOrOops(in, promptField, objective, objective_kind string, skipOops bool) { // - -
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()
			// These prompts, deployed by new_objective_kind, take new_prompt
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Refer to the previous comments re the following mirrored section:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				if in == "set" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, new_objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, new_objective_kind)
				}
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
			// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
			// These prompts, deployed by new_objective_kind, take new_prompt
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Refer to the previous comments re the following mirrored section:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				if in == "set" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, new_objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, new_objective_kind)
				}
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else { // it is not a Right, and is therefor an Oops, no new aCard has been fetched etc. Not even checkMemory has been run
			if skipOops {
				// Then do nothing
			} else {
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				fmt.Printf("      　^^Oops! ")
			}
			tryAgain(promptField, objective, objective_kind) // passing the old original values
		}
	}
}

func tryAgain(promptField, objective, objective_kind string) { // - -
	fmt.Printf("Try again \n")
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative)
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	if objective_kind == "Romaji" {
		in = promptForRomaji(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Extended_Romaji" {
		in = promptForRomajiE(promptField) // A special prompt for Extended Kata, if|when deployed
	} else if objective_kind == "Hira" {
		in = promptForHira(promptField)
	}
	// **** Note here ^ ^ ^ the missing "WithDir" suffix to "promptForHira" as Directives are currently inoperative

	// ...
	// Note the lack of a Directive handling section which normally follows prompting, ergo currently inoperative
	//
	var thisCaseOfAnInHasAlreadyBeenProcessedAbove = false
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
			log_right(promptField)
			fmt.Printf("%s", colorGreen)
			fmt.Printf("        ^^Right! ")
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()
			// These prompts, deployed by new_objective_kind, take new_prompt
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Refer to the previous comments re the following mirrored section:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				if in == "set" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, new_objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, new_objective_kind)
				}
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
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()
			// These prompts, deployed by new_objective_kind, take new_prompt
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Refer to the previous comments re the following mirrored section:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				if in == "set" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, new_objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, new_objective_kind)
				}
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
	var in string // var declaration needed as a ":=" would not work within the conditional ~ "in" not in signature
	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative)
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	if objective_kind == "Romaji" {
		in = promptForRomaji(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Extended_Romaji" {
		in = promptForRomajiE(promptField) // A special prompt for Extended Kata, if|when deployed
	} else if objective_kind == "Hira" {
		in = promptForHira(promptField)
	}

	// **** Note here ^ ^ ^ the missing "WithDir" suffix to "promptForHira" as Directives are currently inoperative

	// ...
	// Note the lack of a Directive handling section which normally follows prompting, ergo currently inoperative
	//
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
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()
			// These prompts, deployed by new_objective_kind, take new_prompt
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Refer to the previous comments re the following mirrored section:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				if in == "set" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, new_objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, new_objective_kind)
				}
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
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()
			// These prompts, deployed by new_objective_kind, take new_prompt
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Refer to the previous comments re the following mirrored section:
			DetectedDirective := false
			DetectedDirective = testForDirective(in)
			if DetectedDirective {
				if in == "set" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, new_objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, new_objective_kind)
				}
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
