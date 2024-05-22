package main

import (
	"fmt"
)

func mainbkup() {
	fmt.Println()
	countSLOC() // Determine and display Source Lines Of Code.
	// rand.Seed(time.Now().UnixNano())  todo: Can we lose this line?
	gameOn = false
	game = "off"
	display_List_of_Directives()
	mainLoop("junk", "junk", "junk") // Junk arguments are required, will not be used.
}

// Pick a new card -- OR, don't pick a new one, as for the case of:  if cameFrom_stcR_NOTstc {...} else {...}
func mainLoopbkup(new_prompt, objective, objective_kind string) { // Only using the vars in signature to create needed local vars. - -
	for {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		}
		if game_loop_counter > game_duration {
			game_off()
		}
		if cameFrom_stcR_NOTstc { // todo: Once this bool value is set true, how and where does it (and should it) get set false?
			begin(new_prompt, objective, objective_kind) // todo]  NOTICE: the use of the new_prompt variant vs the promptField variant.
			// These vars ^ ^ ^ would not have existed if we had not created them via this function's signature. This is all well and good,
			// ... because cameFrom_stcR_NOTstc will be false initially, and these local vars will have retained their values obtained
			// ... below v v v in the "else" clause: exactly as they had been after the last call to pick_RandomCard_Assign_fields().
		} else {
			// In this NORMAL case, we want to clobber new_prompt as we pick the next fresh random card.
			new_prompt, objective, objective_kind = pick_RandomCard_Assign_fields()
			// todo]  Note: a new line similar to this ^ ^ ^ is done after each correct guess.
			begin(new_prompt, objective, objective_kind) // Notice: the use of the new_prompt variant vs the promptField variant.
			// ^ ^ ^ Since begin() is called here, from within this for loop, we will continue here after begin() finishes. And,
			// ... as mentioned, "these local vars will have retained their values".  todo]  Hopefully a feature and not a bug.
		}
	}
}

func beginbkup(promptField, objective, objective_kind string) { // promptField could be a Hira, Kata, or a Romaji-type prompt.  - -
	if game_loop_counter > game_duration {
		game_off()
	}
	var in string // This var declaration was needed since ":=" would not work within the conditional because "in" not in signature.
	for {
		// todo] Note: 1:2 prompts (deployed via objective_kind) all use the promptField variant (rather than the new_prompt variant).
		// Note: objective_kind may have been randomized by pick_RandomCard_Assign_fields() :: may NOT YET be "limited" to a particular type.
		if objective_kind == "Romaji" {
			in = promptForRomajiWithDir(promptField) // Get user's guess|dir.
		} else if objective_kind == "Extended_Romaji" {
			in = promptForRomajiWithDirE(promptField) // Special prompt for Extended Kata, if|when deployed.
		} else if objective_kind == "Hira" {
			in = promptForHiraWithDir(promptField)
		} // todo: Further/better explain the above via comments, re prompt type and objective type (user's input mode).

		// One of two constructs like this that use promptField etc. instead of new_prompt, new_objective, new_objective_kind.
		aDirectiveWasDetected := false
		aDirectiveWasDetected = detectDirective(in)
		if aDirectiveWasDetected { // Respond to directive just one of two ways.
			if in == "stc" || in == "stcr" {
				promptField, objective, objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
			} else {
				// Note: when no stc or stcr was run we do not update promptField, objective, and objective_kind. (it's a naked call).
				respond_to_UserSuppliedDirective(in, objective_kind)
			}
			continue // Re-prompt after "Directive" handling, with the same/original promptField  todo]  EXCEPT after an stc directive.
		} else {
			// Passing (below) a recursion flag of false, means we WILL execute rightOrOops(...)
			evaluateUsersGuess(in, promptField, objective, objective_kind, false, false, false)
			// Note: all bool flags passed in the above call are put to work throughout the called func.
			break // Having finished with all potential guessing, return to the mainLoop() func.
		}
	} // Returns to mainLoop to potentially randomly-select a fresh new Card.
}

func evaluateUsersGuessbkup(in, promptField, objective, objective_kind string, recursion, recall, skipOops bool) { // - -
	if game_loop_counter > game_duration {
		game_off()
	}
	/*
		These following constructs are strange! Because, they seem to allow for rightOrOops() to be done "twice" -- but they do not!
		Since rightOrOops() sets in motion a chain of events that never returns directly back to where it was called ...
		... such duplication never occurs. The constructs are needed to allow for rightOrOops() to be conditionally omitted if ...
		... e.g., evaluateUsersGuess() has been called with its recursion flag true and its recall flag false.
	*/

	if recursion {
		// Just proceed to the next "if" clause AND proceed todo the rightOrOops func, JUST ONCE!
	} else {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		} // todo]  Note: Executes when recursion is False!
		rightOrOops(in, promptField, objective, objective_kind, skipOops) // This func may, intern, call: tryAgain() ... and lastTry()
	} // ^ ^ ^ ^ does NOT return directly back here.

	if recall {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		} // todo]  Note: Executes ONLY if recursion is True and recall is also True!
		rightOrOops(in, promptField, objective, objective_kind, skipOops) // This func may, intern, call: tryAgain() ... and lastTry()
		// ^ ^ ^ ^ does NOT return directly back here.
	} else {
		// Do nothing.
	}
	// ^ ^ ^ If evaluateUsersGuess() has been called after handling a "Directive" then rightOrOops() is omitted entirely.

	// todo]  Note: 2:2 prompts (deployed via objective_kind) all use the promptField variant (rather than the new_prompt variant).
	if objective_kind == "Romaji" {
		in = promptForRomajiWithDir(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Extended_Romaji" {
		in = promptForRomajiWithDirE(promptField) // A special prompt for Extended Kata, if|when deployed
	} else if objective_kind == "Hira" {
		in = promptForHiraWithDir(promptField)
	}

	// Two of two constructs like this that use promptField etc. instead of new_prompt, new_objective, new_objective_kind.
	aDirectiveWasDetected := false
	aDirectiveWasDetected = detectDirective(in)
	if aDirectiveWasDetected { // Respond to directive just one of two ways.
		if in == "stc" || in == "stcr" {
			promptField, objective, objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
		} else {
			// Note: when no stc or stcr was run we do not update promptField, objective, and objective_kind. (it's a naked call).
			respond_to_UserSuppliedDirective(in, objective_kind)
		}

		// todo: consider rewriting the following documentation.
		/*
			Recursively ...
			after responding to a Directive, prompt via recursion, from the same/original promptField??
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

// todo I have added skipOops arguments to all the calls to tryAgain()
func rightOrOopsbkup(in, promptField, objective, objective_kind string, skipOops bool) { // - -
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false

	// todo]  Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true // That is, it will be true :)
		if in == "zu" {
			log_right(promptField, in)
			fmt.Printf("%s", colorGreen)
			if objective_kind == "Hira" {
				fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.SansR_Hint)
			} else { // else it is Romaji, and ...
				if limitedToDifficultKata == true {
					fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
				} else { // ... This correct guess must be a REGULAR Kata, so ...
					fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
				}
			}
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// These prompts (1 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt)
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// One of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
			aDirectiveWasDetected := false
			aDirectiveWasDetected = detectDirective(in)
			if aDirectiveWasDetected {
				if in == "stc" || in == "stcr" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, objective_kind)
				}
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			if skipOops {
				// Do nothing
			} else { // Here is what we do if objective is zu AND if in != "zu"
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				// fmt.Printf("       Try again, two more attempts remain \n %s", colorReset) // print in red, then reset to white
				tryAgain(promptField, objective, objective_kind, skipOops) // passing the old original values
			}
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// todo]  Note: that "else", here, means that objective != "zu"  ... but the objective may yet be ず or づ.

		if objective == "ず" || objective == "づ" {
			if in == "づ" || in == "ず" {
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
				// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
				new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field

				// These prompts (2 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
				if new_objective_kind == "Romaji" {
					in = promptForRomajiWithDir(new_prompt)
				} else if new_objective_kind == "Extended_Romaji" {
					in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
				} else if new_objective_kind == "Hira" {
					in = promptForHiraWithDir(new_prompt)
				}

				// Two of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
				aDirectiveWasDetected := false
				aDirectiveWasDetected = detectDirective(in)
				if aDirectiveWasDetected {
					if in == "stc" || in == "stcr" {
						new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
					} else {
						respond_to_UserSuppliedDirective(in, objective_kind)
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
				}
				tryAgain(promptField, objective, objective_kind, skipOops) // passing the old original values
			}
		}

		// todo]  Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
		if in == objective {
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
			// Since this was a correct guess, next we obtain new values in-preparation of "returning" to the caller.
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// These prompts (3 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt)
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Three of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
			aDirectiveWasDetected := false
			aDirectiveWasDetected = detectDirective(in)
			if aDirectiveWasDetected {
				if in == "stc" || in == "stcr" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, objective_kind)
				}
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			if skipOops {
				// Do nothing
			} else {
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
			}
			tryAgain(promptField, objective, objective_kind, skipOops) // Retaining the old original values.
		}
	}
}

// was: func tryAgain(promptField, objective, objective_kind string) { // - - // todo it had no skipOops bool argument.
func tryAgainbkup(promptField, objective, objective_kind string, skipOops bool) { // - -
	// original ver:
	fmt.Printf("       Try again, two more attempts remain \n%s", colorReset)
	var in string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature.

	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative) ...
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant).
	if objective_kind == "Romaji" {
		in = promptForRomaji1(promptField)
	} else if objective_kind == "Extended_Romaji" {
		in = promptForRomajiE(promptField) // A special prompt for Extended Kata, if|when deployed.
	} else if objective_kind == "Hira" {
		in = promptForHira1(promptField)
	}
	// **** Note here ^ ^ ^ the missing "WithDir" suffix to "promptForHira" as Directives are currently inoperative.

	// ...
	// Note the lack of a Directive handling section which normally follows prompting, ergo currently inoperative.
	//
	// end of original ver.

	// new version:
	// todo should or could this go at the top of the func?
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false // todo I hope this is the right one? false and all?

	// todo]  Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true // That is, it will be true :)
		if in == "zu" {
			log_right(promptField, in)
			fmt.Printf("%s", colorGreen)
			if objective_kind == "Hira" {
				fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.SansR_Hint)
			} else { // else it is Romaji, and ...
				if limitedToDifficultKata == true {
					fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
				} else { // ... This correct guess must be a REGULAR Kata, so ...
					fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
				}
			}
			fmt.Printf("%s", colorReset)
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()
			/* old ver:
					// todo: "duplicate the fancier version of zu handling from rightOrOops.
			if objective == "zu" {
				thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
				if in == "zu" {
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
					fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
					new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			*/ // end of old ver.

			// new ver:
			// These prompts (1 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt)
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// /* delete this dir handling section:
			// One of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
			aDirectiveWasDetected := false
			aDirectiveWasDetected = detectDirective(in)
			if aDirectiveWasDetected {
				if in == "stc" || in == "stcr" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, objective_kind)
				}
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}

			// */
		} else {
			if skipOops {
				// Do nothing
			} else { // Here is what we do if objective is zu AND if in != "zu"
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				// fmt.Printf("       Try again, two more attempts remain \n %s", colorReset) // print in red, then reset to white
				// original (new) ver: tryAgain(promptField, objective, objective_kind, skipOops) // passing the old original values
				lastTry(promptField, objective, objective_kind, skipOops) // passing the old original values
			}
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		// todo]  Note: that "else", here, means that objective != "zu"  ... but the objective may yet be ず or づ.

		if objective == "ず" || objective == "づ" {
			if in == "づ" || in == "ず" {
				/* this is how the new ver went:
				} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
				// todo]  Note: that "else", here, means that objective != "zu"  ... but the objective may yet be ず or づ.

				if objective == "ず" || objective == "づ" {
					if in == "づ" || in == "ず" {
				*/
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
				// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
				new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field
				// the above is obviously all new ver.
				/* this is how it went in new ver:
				new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field

				// These prompts (2 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
				if new_objective_kind == "Romaji" {
				*/

				/* new ver comparison of following original ver:
					// These prompts (2 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
					if new_objective_kind == "Romaji" {
						in = promptForRomajiWithDir(new_prompt)
					} else if new_objective_kind == "Extended_Romaji" {
						in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
					} else if new_objective_kind == "Hira" {
						in = promptForHiraWithDir(new_prompt)
					}

					// Two of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
					aDirectiveWasDetected := false
					aDirectiveWasDetected = detectDirective(in)
					if aDirectiveWasDetected {
						if in == "stc" || in == "stcr" {
							new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
						} else {
							respond_to_UserSuppliedDirective(in, objective_kind)
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
					}
					tryAgain(promptField, objective, objective_kind) // passing the old original values

				*/ // that was new ver comparison of following original ver

				/* todo trying omitting these two lines:
					}
				}

				*/

				// todo]  Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)

				// original ver follows and I expect to be keeping it all:
				// These prompts (4 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
				if new_objective_kind == "Romaji" {
					in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
				} else if new_objective_kind == "Extended_Romaji" {
					in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
				} else if new_objective_kind == "Hira" {
					in = promptForHiraWithDir(new_prompt)
				}

				// /*
				// Four of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
				aDirectiveWasDetected := false
				aDirectiveWasDetected = detectDirective(in)
				if aDirectiveWasDetected {
					if in == "stc" || in == "stcr" {
						new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
					} else {
						respond_to_UserSuppliedDirective(in, objective_kind)
					}
					evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
				} else {
					evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
				}

				// */

				/* from new ver for comparison, and it all looks good so keep surrounding.
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
						} else {
							evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
				*/
			} else {

				// todo:
				/* new version option for next few lines:
				if skipOops {
							// Then do nothing
						} else {
							log_oops(promptField, objective, in)
							fmt.Printf("%s", colorRed)
						}
						tryAgain(promptField, objective, objective_kind) // passing the old original values
				*/

				// original ver:
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				lastTry(promptField, objective, objective_kind, skipOops)
			}

		} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {

			// all original ver ^ ^ ^ v v v.

			if in == objective {
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
				new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

				// These prompts (5 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
				if new_objective_kind == "Romaji" {
					in = promptForRomajiWithDir(new_prompt)
				} else if new_objective_kind == "Extended_Romaji" {
					in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
				} else if new_objective_kind == "Hira" {
					in = promptForHiraWithDir(new_prompt)
				}

				// Five of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
				aDirectiveWasDetected := false
				aDirectiveWasDetected = detectDirective(in)
				if aDirectiveWasDetected {
					if in == "stc" || in == "stcr" {
						new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
					} else {
						respond_to_UserSuppliedDirective(in, objective_kind)
					}
					evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
				} else {
					evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
				}
			} else {
				log_oops(promptField, objective, in)
				fmt.Printf("%s", colorRed)
				lastTry(promptField, objective, objective_kind, skipOops)
			}
		}
		// todo try adding a } somewhere around here:
	} // end of try again.
} // todo, here may work?

func lastTrybkup(promptField, objective, objective_kind string, skipOops bool) { // - -
	fmt.Printf("       Last Try! \n%s", colorReset)
	var in string // var declaration needed as a ":=" would not work within the conditional ~ "in" not in signature.

	// **** Now that we are trying again, after a failed guess, prompts do not solicit Directives:(currently inoperative) ...
	// ... so, these prompts, deployed by objective_kind, take promptField (rather than the new_prompt variant)
	if objective_kind == "Romaji" {
		in = promptForRomaji2(promptField) // Get user's input, from a randomly selected prompt
	} else if objective_kind == "Extended_Romaji" {
		in = promptForRomajiE(promptField) // A special prompt for Extended Kata, if|when deployed
	} else if objective_kind == "Hira" {
		in = promptForHira2(promptField)
	}
	// **** Note here ^ ^ ^ the missing "WithDir" suffix to "promptForHira" as Directives are currently inoperative

	// ...
	// Note the lack of a Directive handling section which normally follows prompting, ergo currently inoperative
	//
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false

	// todo: "duplicate the fancier version of zu handling from rightOrOops.
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
		if in == "zu" {
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
			fmt.Printf("It could have been either ず or づ as they are the same sound: zu\n")
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// These prompts (6 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt)
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Six of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
			aDirectiveWasDetected := false
			aDirectiveWasDetected = detectDirective(in)
			if aDirectiveWasDetected {
				if in == "stc" || in == "stcr" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, objective_kind)
				}
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			log_oops(promptField, objective, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! And I lied about that being your last try, you've got this, try again...\n %s", colorReset)
		}

	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true {
		if in == objective {
			log_right(promptField, in)
			fmt.Printf("%s", colorGreen)
			if objective_kind == "Hira" {
				fmt.Printf("      　%s %s   - %s\n", aCard.Romaji, aCard.Kata, aCard.SansR_Hint)
			} else { // else it is Romaji, so:
				if limitedToDifficultKata == true {
					fmt.Printf("      　%s %s  \n", aCard.Hira, aCard.SansR_Hint)
				} else { // Then this correct guess must be a regular Kata
					fmt.Printf("      　%s %s   - %s\n", aCard.Hira, aCard.Kata, aCard.SansR_Hint)
				}
			}
			fmt.Printf("%s", colorReset)
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// These prompts (7 of 7 like this), deployed by new_objective_kind, take new_prompt etc.
			if new_objective_kind == "Romaji" {
				in = promptForRomajiWithDir(new_prompt) // Get user's input, from a randomly selected prompt
			} else if new_objective_kind == "Extended_Romaji" {
				in = promptForRomajiWithDirE(new_prompt) // A special prompt for Extended Kata, if|when deployed
			} else if new_objective_kind == "Hira" {
				in = promptForHiraWithDir(new_prompt)
			}

			// Seven of seven constructs like this that use new_prompt, new_objective, and new_objective_kind instead of promptField etc.
			aDirectiveWasDetected := false
			aDirectiveWasDetected = detectDirective(in)
			if aDirectiveWasDetected {
				if in == "stc" || in == "stcr" {
					new_prompt, new_objective, new_objective_kind = respond_to_UserSuppliedDirective(in, objective_kind)
				} else {
					respond_to_UserSuppliedDirective(in, objective_kind)
				}
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, false, true)
			} else {
				evaluateUsersGuess(in, new_prompt, new_objective, new_objective_kind, true, true, false)
			}
		} else {
			log_oops(aCard.Hira, aCard.Romaji, in)
			fmt.Printf("%s", colorRed)
			fmt.Printf("     ^^Oops! That was your last try looser. Here's a clue, just for you: ...\n %s", colorReset)
			fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
		}
	}
}
