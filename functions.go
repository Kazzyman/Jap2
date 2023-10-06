package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
)

// DIRECTIVES : --------------------------------------------------------------------------------------------
//
func handle_doubleQuestMark_directive() { //                 - -
	var hiraganaCharOrRomajiAssociatedWithStructElementToDisplayHelpFieldsFrom string
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

// end of DIRECTIVES -----------------------------------------------------------------------------------

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
		display_Listing_of_Directives()
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

var whichDeck int

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
		whichDeck = 1
	}
	if randomFileOfCards == 1 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 2
	}
	if randomFileOfCards == 2 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 3
	}

	// Kata prompting, Romaji objective:
	if randomFileOfCards == 3 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 1
	}
	if randomFileOfCards == 4 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 2
	}
	if randomFileOfCards == 5 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Romaji
		objective_kind = "Romaji"
		whichDeck = 3
	}

	// Romaji prompting, Hira objective:
	if randomFileOfCards == 6 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 1

	}
	if randomFileOfCards == 7 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 2
	}
	if randomFileOfCards == 8 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 3
	}

	// Kata prompting, Hira objective:
	if randomFileOfCards == 9 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 1
	}
	if randomFileOfCards == 10 {
		aCard = fileOfCardsS[randIndexS] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 2
	}
	if randomFileOfCards == 11 {
		aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
		objective = aCard.Hira
		objective_kind = "Hira"
		whichDeck = 3
	}
	return promptField, objective, objective_kind
}
