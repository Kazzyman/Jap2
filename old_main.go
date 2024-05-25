package main

import (
	"fmt"
	"math/rand"
	"time"
)

// todo: review bool tests used for conditional execution of blocks.

func oldmain() {
	fmt.Println()
	countSLOC()                      // Determine and display Source Lines Of Code.
	rand.Seed(time.Now().UnixNano()) // seed the random number generator with the current time in nanoseconds.
	gameOn = false
	game = "off"
	display_List_of_Directives()
	oldmainLoop("junk", "junk", "junk") // Junk arguments are required, will not be used.
}

// mainLoop(new_prompt, objective, objective_kind string)
// Pick a new card -- OR, don't pick a new one, as for the case of:  if cameFrom_stcR_NOTstc {...} else {...}
func oldmainLoop(new_prompt, objective, objective_kind string) { // Only using the vars in signature to create needed local vars. - -
	// mainLoop(...) is endless, excepting os.Exit(1) via Directive "q".
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
			oldbegin(new_prompt, objective, objective_kind) // todo]  NOTICE: the use of the new_prompt variant vs the promptField variant.
			// These vars ^ ^ ^ would not have existed if we had not created them via this function's signature. This is all well and good,
			// ... because cameFrom_stcR_NOTstc will be false initially, and these local vars will have retained their values obtained
			// ... below v v v in the "else" clause: exactly as they had been after the last call to pick_RandomCard_Assign_fields().
		} else {
			// In this NORMAL case, we want to clobber new_prompt as we pick the next fresh random card.
			new_prompt, objective, objective_kind = pick_RandomCard_Assign_fields()
			// todo]  Note: a new line similar to this ^ ^ ^ is done after each correct guess.
			oldbegin(new_prompt, objective, objective_kind) // Notice: the use of the new_prompt variant vs the promptField variant.
			// ^ ^ ^ Since begin() is called here, from within this for loop, we will continue here after begin() finishes. And,
			// ... as mentioned, "these local vars will have retained their values".  todo]  Hopefully a feature and not a bug.
		}
	}
}

func oldbegin(promptField, objective, objective_kind string) { // promptField could be a Hira, Kata, or a Romaji-type prompt.  - -
	if game_loop_counter > game_duration {
		game_off()
	}
	var userInput string // This var declaration was needed since ":=" would not work within the conditional because "in" not in signature.

	for {

		// todo] Note: 1:2 prompts (deployed via objective_kind) all use the promptField variant (rather than the new_prompt variant).
		// Note: objective_kind may have been randomized by pick_RandomCard_Assign_fields() :: may NOT YET be "limited" to a particular type.

		// 1 of 2 identical prompts.
		userInput = promptFirstCase(promptField, objective_kind)

		// 1 of 2 somewhat-similar constructs that use promptField etc. -- Instead of new_prompt, new_objective etc.
		aDirectiveWasDetected := false
		aDirectiveWasDetected = detectDirective(userInput)
		if aDirectiveWasDetected { // Respond to directive just one of two ways.
			if userInput == "stc" || userInput == "stcr" {
				promptField, objective, objective_kind = respond_to_UserSuppliedDirective(userInput)
			} else {
				// Note: when no stc or stcr was run we do not update promptField, objective, and objective_kind. (it's a naked call).
				respond_to_UserSuppliedDirective(userInput)
			}
			continue // Re-prompt after "Directive" handling, with the same/original promptField  todo]  EXCEPT after an stc directive.
		} else {
			// Passing (below) a recursion flag of false, means we WILL execute rightOrOops(...)
			evaluateUsersGuess(userInput, promptField, objective, objective_kind, false, false)
			// Note: all bool flags passed in the above call are put to work throughout the called func.
			break // Having finished with all potential guessing, return to the mainLoop() func.
		}
	} // Returns to mainLoop to potentially randomly-select a fresh new Card.
}

func evaluateUsersGuess(userInput, promptField, objective, objective_kind string, recursion, recall bool) { // - -
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
		rightOrOops(userInput, promptField, objective, objective_kind) // This func may, intern, call: tryAgain() ... and lastTry()
	} // ^ ^ ^ ^ does NOT return directly back here.

	if recall {
		if gameOn {
			game_loop_counter++
		} else {
			game_loop_counter = 0
		} // todo]  Note: Executes ONLY if recursion is True and recall is also True!
		rightOrOops(userInput, promptField, objective, objective_kind) // This func may, intern, call: tryAgain() ... and lastTry()
		// ^ ^ ^ ^ does NOT return directly back here.
	} else {
		// Do nothing.
	}
	// ^ ^ ^ If evaluateUsersGuess() has been called after handling a "Directive" then rightOrOops() is omitted entirely.

	// 2 of 2 identical prompts.
	userInput = promptFirstCase(promptField, objective_kind)

	// 2 of 2 somewhat-similar constructs that use promptField etc. -- Instead of new_prompt, new_objective etc.
	aDirectiveWasDetected := false
	aDirectiveWasDetected = detectDirective(userInput)
	if aDirectiveWasDetected { // Respond to directive just one of two ways.
		if userInput == "stc" || userInput == "stcr" {
			promptField, objective, objective_kind = respond_to_UserSuppliedDirective(userInput)
		} else {
			// Note: when no stc or stcr was run we do not update promptField, objective, and objective_kind. (it's a naked call).
			respond_to_UserSuppliedDirective(userInput)
		}

		// todo: consider rewriting the following documentation.
		/*
			Recursively ...
			after responding to a Directive, prompt via recursion, from the same/original promptField??
			With recursion true & recall false, rightOrOops() will NOT be done with this recursion,
			...                                                                        t, f, t Skips rightOrOops
		*/
		evaluateUsersGuess(userInput, promptField, objective, objective_kind, true, false) //
	} else {
		/*
			Do a normal, i.e., unconditional recursion with skipOops set to false
			via recall==true & skipOops==false [recursion false or recall true, means do rightOrOops]
		*/
		evaluateUsersGuess(userInput, promptField, objective, objective_kind, true, true) // t,t,f Does rightOrOops
	}
}

func rightOrOops(userInput, promptField, objective, objective_kind string) { // - -
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false // todo]  Note: the first use of this flag !!

	// todo]  Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true // todo: is this bool setting well-placed? Compare others.
		if userInput == "zu" {
			logRight_zu(userInput, promptField, objective_kind)
			// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// 1 of 9 identical calls.
			userInput = promptSecondCase(new_prompt, new_objective_kind)
			new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)

		} else {
			/*
				if skipOops {
								// Do nothing
							} else {
			*/
			log_oops(promptField, objective, userInput)
			fmt.Printf("%s", colorRed)
			oldtryAgain(promptField, objective, objective_kind) // passing the old original values
			// }
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true { // todo: is this bool test even needed??
		// todo]  Note: that "else", here, means that objective != "zu"  ... but the objective may yet be ず or づ.

		if objective == "ず" || objective == "づ" {
			if userInput == "づ" || userInput == "ず" {
				logRightZu2(userInput, promptField, objective_kind, objective)
				// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
				new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field

				// 2 of 9 identical calls.
				userInput = promptSecondCase(new_prompt, new_objective_kind)
				new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)

			} else {
				/*
					if skipOops {
						// Then do nothing
					} else {

				*/
				log_oops(promptField, objective, userInput)
				fmt.Printf("%s", colorRed)
				oldtryAgain(promptField, objective, objective_kind) // passing the old original values
				// }
				// todo: only in this second, and the third, occurrence of this structure did we find tryAgain() below the else block.
				// tryAgain(promptField, objective, objective_kind)
			}
		}

		// todo]  Now; having dispensed with the odd case of the zu clan -- we return you to our usual programming :)
		if userInput == objective {
			logRight(userInput, promptField, objective_kind)
			// Since this was a correct guess, next we obtain new values in-preparation of "returning" to the caller.
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// 3 of 9 identical calls.
			userInput = promptSecondCase(new_prompt, new_objective_kind)
			new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)

		} else {
			/*
				if skipOops {
								// Do nothing
							} else {
			*/
			log_oops(promptField, objective, userInput)
			fmt.Printf("%s", colorRed)
			oldtryAgain(promptField, objective, objective_kind)
			// }
			// todo: is it OK||correct that we try-again irrespective of skipOops ?? As WAS the case, two lines down.
			// todo: only in the second, and in this third, occurrence of this structure did we find tryAgain() below the else block.
			// tryAgain(promptField, objective, objective_kind)
		}
	}
} // end of rightOrOops function.

func oldtryAgain(promptField, objective, objective_kind string) { // - -
	fmt.Printf("       Try again, two more attempts remain \n%s", colorReset)
	var userInput string // var declaration needed as a ":=" would not work within the conditional because "in" not in signature.

	// **** Now that we are trying again, after a failed guess, Directives are currently inoperative ...
	// ... deployed by objective_kind, and using promptField (rather than the new_prompt variant).

	// A unique prompt construct.
	if objective_kind == "Romaji" {
		userInput = promptForRomaji1(promptField) // todo]  Note the suffix of 1 on promptForRomaji
	} else if objective_kind == "Extended_Romaji" {
		userInput = promptForRomajiE(promptField) // A special prompt for Extended Kata, if|when deployed.
	} else if objective_kind == "Hira" {
		userInput = promptForHira1(promptField)
	}

	// todo: is this bool setting well-placed? Compare others.
	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false // todo: I hope this is the right one? false and all?

	// todo]  Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true // todo: is this bool setting well-placed? Compare others.
		if userInput == "zu" {
			logRight_zu(userInput, promptField, objective_kind)
			// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// 4 of 9 identical calls.
			userInput = promptSecondCase(new_prompt, new_objective_kind)
			new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)

		} else {
			/*
				if skipOops {
								// Do nothing
							} else {
			*/ // Here is what we do if objective is zu AND if in != "zu"
			log_oops(promptField, objective, userInput)
			fmt.Printf("%s", colorRed)
			// fmt.Printf("       Try again, two more attempts remain \n %s", colorReset) // print in red, then reset to white
			// original (new) ver: tryAgain(promptField, objective, objective_kind)
			oldlastTry(promptField, objective, objective_kind) // passing the old original values
			// }
		}
	} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true { // todo: is this bool test even needed??
		// todo]  Note: that "else", here, means that objective != "zu"  ... but the objective may yet be ず or づ.

		if objective == "ず" || objective == "づ" {
			if userInput == "づ" || userInput == "ず" {
				logRightZu2(userInput, promptField, objective_kind, objective)
				// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
				new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields() // Gets a new card and extract the new prompt field

				// 5 of 9 identical calls.
				userInput = promptSecondCase(new_prompt, new_objective_kind)
				new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)

			} else {
				log_oops(promptField, objective, userInput)
				fmt.Printf("%s", colorRed)
				oldlastTry(promptField, objective, objective_kind)
			}

		} else if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true { // todo: is this bool test even needed??

			if userInput == objective {
				logRight(userInput, promptField, objective_kind)
				new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

				// 6 of 9 identical calls.
				userInput = promptSecondCase(new_prompt, new_objective_kind)
				new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)

			} else {
				log_oops(promptField, objective, userInput)
				fmt.Printf("%s", colorRed)
				oldlastTry(promptField, objective, objective_kind)
			}
		}
	}
} // end of tryAgain function.

func oldlastTry(promptField, objective, objective_kind string) { // - -
	fmt.Printf("       Last Try! \n%s", colorReset)
	var userInput string // var declaration needed as a ":=" would not work within the conditional ~ "in" not in signature.

	// A unique prompt construct.
	if objective_kind == "Romaji" {
		userInput = promptForRomaji2(promptField) // todo]  Note the 2 suffix on promptForRomaji
	} else if objective_kind == "Extended_Romaji" {
		userInput = promptForRomajiE(promptField) // A special prompt for Extended Kata, if|when deployed.
	} else if objective_kind == "Hira" {
		userInput = promptForHira2(promptField)
	}

	thisCaseOfAnInHasAlreadyBeenProcessedAbove = false // todo: is this bool setting well-placed? Compare others.

	// todo]  Firstly, we'll do some special processing to address the strange case of zu, ず, and づ.

	// todo: no bool test ?? It is done below for the if userInput == objective {  ???  ---
	if objective == "zu" {
		thisCaseOfAnInHasAlreadyBeenProcessedAbove = true // todo: is this bool setting well-placed? Compare others.
		if userInput == "zu" {
			logRight_zu(userInput, promptField, objective_kind)
			// Since this was "^^Right!", next we obtain new values in-preparation of "returning" to caller
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// 7 of 9 identical calls.
			userInput = promptSecondCase(new_prompt, new_objective_kind)
			new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)
		}
	} else if objective == "ず" || objective == "づ" {
		if userInput == "づ" || userInput == "ず" {
			thisCaseOfAnInHasAlreadyBeenProcessedAbove = true
			logRightZu2(userInput, promptField, objective_kind, objective)
		} else {
			logOopsLoser(userInput)
			return
		}
		fmt.Printf("%s", colorReset)
		new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

		// 8 of 9 identical calls.
		userInput = promptSecondCase(new_prompt, new_objective_kind)
		new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)
	}

	// todo: is it OK||correct that this "if" stands alone outside of the "else", and has its own bool test ?? Not so in prior funcs.
	if thisCaseOfAnInHasAlreadyBeenProcessedAbove != true { // todo: is this bool test even needed??
		if userInput == objective {
			logRight(userInput, promptField, objective_kind)
			// Since this was a correct guess, next we obtain new values in-preparation of "returning" to the caller.
			new_prompt, new_objective, new_objective_kind := pick_RandomCard_Assign_fields()

			// 9 of 9 identical calls.
			userInput = promptSecondCase(new_prompt, new_objective_kind)
			new_prompt, new_objective, new_objective_kind = directiveHandler(userInput, new_prompt, new_objective, new_objective_kind)

		} else {
			logOopsLoser(userInput)
		}
	}
}
