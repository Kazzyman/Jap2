package main

import (
	"fmt"
	"os"
	"regexp"
)

/*
It was a very bad idea from the beginning to detect a user's Directive by testing the input for the presence of Alpha
characters. Because, in the case of a Hira or Kata prompt, Alpha chars in the input do not correspond to the input being
a Directive. It would have been better to have every prompt func check for a string such as "dir" or "help" and, in such a case,
pass control to the Directive handler func which would then prompt with a list of available Directives.
*/

// DIRECTIVES : --------------------------------------------------------------------------------------------
//
func handle_doubleQuestMark_directive() { //                 - -
	// Works with Hiragana or Romaji user input in all Exercises
	var hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom string // the value to find in locateCard.go
	//
	fmt.Printf("\n  -- Type either the Hiragana char or the Romaji that you need help with:> ")
	_, _ = fmt.Scan(&hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom)
	//
	locateCardAndDisplayHelpFieldsContainedInIt(hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom)
	fmt.Println("")
}

// Handles the Directive 'set'
func reSet_aCard_andThereBy_reSet_thePromptString() { //     - -
	var theHiraganaOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Hiragana to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt & \"aCard\\.fields\" :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)

	// Determine if the user has entered a valid Hiragana char (instead of, accidentally, an alpha char or string)
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(theHiraganaOfCardToSilentlyLocate):
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}
	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'
	if isAlphanumeric == true {
		fmt.Println("Are you in alphanumeric input mode?")
		fmt.Printf("... if so, change it to Hiragana (or I mignt die)\n")
		fmt.Printf("%s", colorRed) //
		fmt.Printf(" cautiously ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("re-enter your selection, in Hiragana mode :> ")
		fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&theHiraganaOfCardToSilentlyLocate)
		// May yet send an Alpha string to the next func, which will itself deal with it elegantly
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCardA = *foundElement                                // Set the global var-object 'aCard'
		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCardA = *foundElement                                // Set the global var-object 'aCard'
		fmt.Println("")
	}
}

//
func read_map_of_fineOn() { //     - -
	if len(frequencyMapOf_IsFineOnChars) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe FineOn Map is empty\n")
		fmt.Printf(colorReset)
	}
	for s, f := range frequencyMapOf_IsFineOnChars {
		fmt.Printf(" --- From MapOf_IsFineOn: string is:")
		fmt.Printf(colorCyan)
		fmt.Printf("%s", s)
		fmt.Printf(colorReset)
		fmt.Printf(", freq is:")
		fmt.Printf(colorRed)
		fmt.Printf("%d", f)
		fmt.Printf(colorReset)
		fmt.Printf(" ---\n")
	}
	fmt.Println("")
}

//
func read_map_of_needWorkOn() { //     - -
	if len(frequencyMapOf_need_workOn) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe need_workOn map is empty\n")
		fmt.Printf(colorReset)
	}
	for s, f := range frequencyMapOf_need_workOn {
		fmt.Printf(" --- From frequencyMapOf_need_workOn: string is:")
		fmt.Printf(colorCyan)
		fmt.Printf("%s", s)
		fmt.Printf(colorReset)
		fmt.Printf(", freq is:")
		fmt.Printf(colorRed)
		fmt.Printf("%d", f)
		fmt.Printf(colorReset)
		fmt.Printf(" ---\n")
	}
	fmt.Println("")
}

// end of DIRECTIVES -----------------------------------------------------------------------------------

/*
.
.
.
*/
// Globals:
// .
// A map v v v v v v v is used to store correct guesses, and is only written to after a correct guess
var frequencyMapOf_IsFineOnChars = make(map[string]int) // create the map, dir 'read_map_of_fineOn' reads it - -
var frequencyMapOf_need_workOn = make(map[string]int)   // - -

// .
// Parse the map to see if the new random card matches any members of the map
// Each time we get a new random card we check the map to see if it has been done correctly 3 or more times ...
// ... if it has, then we pick another card, and check it
// else we break
func check_it_for_fine_on() { //          - -
	for s, f := range frequencyMapOf_IsFineOnChars {
		if s == aCardA.Romaji { // if it is in the map we need to check the freq
			if f >= 2 { // if the freq is 3+ we need another card
				// read_map_of_fineOn() // we show the map
				// read_map_of_needWorkOn() // We show the map
				// fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR)
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                 // ... gets a file handle to JapLog.txt
				_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyH:%s freq:%d \n",
					s, aCardA.Hira, f)
				check(err2)
				pick_RandomCard_Assign_fields() // We get that new card ...
				// fmt.Println(" ... so here is a new one ... \n")
				check_it_for_fine_on() // ... and we check THAT new card with a recursive call
				_ = fileHandleBig.Close()
			} else { // else the card had a freq less than 3, so ...
				break //  ... we exit the loop and the func -- we will use this card
			}
		} else { // else the latest card which was randomly selected was not found YET in the map ...
			continue // ... so we continue the loop to finish reading the map
		}
	}
}

// We need a counter on this loop so that we avoid endless looping in the case of all chars becoming fine on
func check_it_for_fine_on6() { //       - -
	avoidEndlessLooping := 0
	for s, f := range frequencyMapOf_IsFineOnChars {
		avoidEndlessLooping++
		if avoidEndlessLooping < 300 {
			if s == aCardA.Hira || s == aCardA.Kata { // if it is in the map we need to check the freq
				if f >= 2 { // if the freq is 3+ we need another card
					// read_map_of_fineOn() // we show the map
					// read_map_of_needWorkOn() // We show the map
					// fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR)
					// Log to a file that this action was taken **do-this**
					if s == aCardA.Hira {
						fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
						check(err)                                                                                 // ... gets a file handle to JapLog.txt
						_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyH:%s freq:%d \n",
							s, aCardA.Hira, f)
						check(err2)
						_ = fileHandleBig.Close()
					} else if s == aCardA.Kata {
						fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
						check(err)                                                                                 // ... gets a file handle to JapLog.txt
						_, err2 := fmt.Fprintf(fileHandleBig, "\n Prompt %s was found in frequencyMapOf_IsFineOnChars, KeyK:%s freq:%d \n",
							s, aCardA.Kata, f)
						check(err2)
						_ = fileHandleBig.Close()
					}
					pick_RandomCard_Assign_fields() // We get that new card ...
					// fmt.Println(" ... so here is a new one ... \n")
					check_it_for_fine_on6() // ... and we check THAT new card with a recursive call
				} else { // else the card had a freq less than 3, so ...
					break //  ... we exit the loop and the func -- we will use this card
				}
			} else { // else the latest card which was randomly selected was not found YET in the map ...
				continue // ... so we continue the loop to finish reading the map
			}
		} else {
			// Flush (clear) the map; and reset the counter
			frequencyMapOf_IsFineOnChars = make(map[string]int)
			avoidEndlessLooping = 0
			fmt.Println("The frequencyMapOf_IsFineOnChars map has been reset")
			break // This should exit the loop
		}
	}
}

// pick returns: promptField, objective, objective_kind
func check_it_for_needing_more_practice(promptField, objective, objective_kind string) { //       - -
	var skip_this_step bool
	skip_this_step = false
	for s, f := range frequencyMapOf_need_workOn {
		if s == objective { // Check if the latest random card is in the need_workOn map, and check the freq ...
			if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we keep it
				// read_map_of_fineOn() // we show the map
				read_map_of_needWorkOn() // We show the map
				fmt.Printf("\n The Random card: %s was missed once or more \n", promptField)
				fmt.Println("... so we will keep it and quiz you on it ... ")
				// Log to a file that this action was taken **do-this**
				fileHandleBig, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
				check(err)                                                                                  // ... gets a file handle to JapLog.txt
				_, err2 := fmt.Fprintf(fileHandleBig, "\nPart1of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
					s, f)
				check(err2)
				skip_this_step = true
				_ = fileHandleBig.Close()
				break //  ... we exit the loop and the func -- we will keep and use this random card, and skip the next loop
			} else { // else the card had a freq less than 2, so ...
				continue // keep looking through the map for another instance that may in there, with a significant freq
			}
		}
	}
	if skip_this_step == false {
		// The latest random was not in the map, but it is time to serve-up something difficult ... so:
		for s, f := range frequencyMapOf_need_workOn {
			if s == objective { // Check if the latest random is in the map, and check the freq ...
				if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we set it as aCard
					// read_map_of_fineOn() // we show the map
					read_map_of_needWorkOn() // We show the map
					fmt.Printf("\n This Random card: %s was missed 1 or more times \n", objective)
					fmt.Println("... so we will test you on it, since it has been a while")
					// Log to a file that this action was taken **do-this**
					fileHandle, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                              // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandle, "\nPart2of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
						s, f)
					check(err2)
					_ = fileHandle.Close()
					practice_this_card(objective) // locate and assign aCard // set it as new aCard
					break                         //  ... we exit the loop and the func -- we will keep and use this random card
				} else { // else the card had a freq less than 2, so ...
					continue // keep looking through the map for another instance that may in there, with a significant freq
				}
			}
		}
	}
}

func stack_the_map() { //             - -
	promptToSkip := "shi"
	for i := 0; i < 6; i++ {
		frequencyMapOf_IsFineOnChars[promptToSkip]++
	}
	fmt.Printf("\nSix occurrences of 'shi' have been added to frequencyMapOf_IsFineOnChars\n\n")
}

// Creates a func named check which takes one parameter "e" of type error
func check(e error) { //      - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
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
		fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
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
