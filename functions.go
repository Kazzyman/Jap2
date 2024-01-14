package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { //    - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}

// DIRECTIVES : --------------------------------------------------------------------------------------------
//
func handle_doubleQuestMark_directive(objective_kind string) { //        - -
	var Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn string
	//
	fmt.Printf("\n  -- Type either a Hiragana or Romaji prompt you need help with:> ")
	_, _ = fmt.Scan(&Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn)
	//
	locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn, objective_kind)
	fmt.Println("")
}

// Handles the Directive 'stc'
func reSet_aCard_andThereBy_reSet_thePromptString() (prompt, objective, objective_kind string) { //  - -
	var theHiraganaOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Hiragana to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt :> ")
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
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		// new_prompt, new_objective, new_objective_kind
		prompt = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		prompt = aCard.Hira
		objective = aCard.Romaji
		objective_kind = "Romaji"
		fmt.Println("")
	}
	return prompt, objective, objective_kind
}

func testForDirective(in string) (result bool) { // - -
	/*
	 'nts' for some background on Romaji conventions
	 'dir' redisplay this menu of available Directives
	 'gdc' set the Duration Counter for a Game session
	 'bgs' or 'goff' Begin or end a Game Session
	 '?' context-sensitive help on the current character
	 '??' for help on a particular Hiragana character
	 'st' Statistics
	 'abt' for trivia about this app
	 'rs' to reset (flush or clear) all stats logs etc.
	 'rm' Read the current contents of the Maps
	 'stc' (Set-Card) force the use of a specific card
	 'exko' load the Extended Kata deck
	 'exkf' un-load the Extended Kata deck
	*/
	if in == "stc" ||
		in == "?" || // <-- If it IS a directive
		in == "??" ||
		in == "rs" ||
		in == "st" ||
		in == "dir" ||
		in == "nts" ||
		in == "q" ||
		in == "rm" ||
		in == "bgs" ||
		in == "goff" ||
		in == "abt" ||
		in == "gdc" ||
		in == "exko" ||
		in == "exkf" {
		// Then:
		result = true
	}
	return result
}

func game_on() (game string) { // - -
	game = "on"
	gameOn = true
	fmt.Println("The game is on")

	startBeforeCall = time.Now()
	currentTime := time.Now()
	TimeOfStartFromTop = time.Now()

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game began at: %s \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check_error(err1)
	return game
}
func game_off() (game string) { // - -
	game = "off"
	gameOn = false
	game_duration = 998

	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check_error(err)

	currentTime := time.Now()

	// calculate elapsed time
	t_s2 := time.Now()
	elapsed := t_s2.Sub(TimeOfStartFromTop)

	// cast time durations to a String type for Fprintf "formatted print"
	TotalRun := elapsed.String()

	fmt.Printf("\nRun time was: %s, gameOn is: %t \n\n", TotalRun, gameOn)

	// End timer and report elapsed time
	_, err1 := fmt.Fprintf(fileHandle,
		"\n The game ended at: %s  Total prompts was: %d \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"), total_prompts)
	check_error(err1)
	_, err2 := fmt.Fprintf(fileHandle,
		"\n Elapsed time of game was: %s \n",
		TotalRun)
	check_error(err2)
	return game
}

func read_pulledButNotUsedMap() {
	if len(pulledButNotUsedMap) == 0 {
		fmt.Printf(colorRed)
		fmt.Printf("\nThe seenMap is empty\n")
		fmt.Printf(colorReset)
	}
	for s, f := range pulledButNotUsedMap {
		if s != "" {
			if s != "primedK0" {
				if f > 1 {
					fmt.Printf(" --- From pulledButNotUsedMap: string is:")
					fmt.Printf(colorCyan)
					fmt.Printf("%s", s)
					fmt.Printf(colorReset)
					fmt.Printf(", freq is:")
					fmt.Printf(colorRed)
					fmt.Printf("%d", f)
					fmt.Printf(colorReset)
					fmt.Printf("\n")
				}
			}
		}
	}
	fmt.Println("")
	/*
			var indexIntoArray int
			for indexIntoArray < len(cyclicArrayPulls.pulls) {
				if cyclicArrayPulls.pulls[indexIntoArray] != "" {
					if cyclicArrayPulls.pulls[indexIntoArray] != "primedK0" {
						fmt.Printf("Char stored in cyclicArrayPulls: %s \n", cyclicArrayPulls.pulls[indexIntoArray])
					}
				}
				indexIntoArray++
			}
		//
			//
			//

	*/
}
func read_pulls_not_used_array() {
	for i, lastPull := range cyclicArrayPulls.pulls {
		if i < len(cyclicArrayPulls.pulls) {
			if lastPull != "" {
				fmt.Printf("array: %s\n", lastPull)
			}
		} else {
			break
		}
	}
}
func about_app() {
	fmt.Printf("\n" +
		"This app consists of the following files and lines of code:\n\n" +
		"1701 constants.go ~ 35\n" +
		"343 functions.go ~ 300\n" +
		"157 locateCard.go ~ 80\n" +
		"357 main.go ~ 320\n" +
		"127 memoryFunctions.go ~ 45\n" +
		"133 objectsAndMethods.go ~ 40\n" +
		"82 prompts&directions.go ~ 75\n" +
		"145 statsFunctions.go ~ 125 \n\n" +
		"3045 lines of code (SOC) \n" +
		"or, about 1,020 functional lines of code \n\n")
}
func reset_all_data() {
	// Flush (clear) the old stats and hits arrays
	cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
	cyclicArrayHits = CyclicArrayHits{}
	cyclicArrayPulls = CyclicArrayPulls{}
	// Also, flush (clear) the maps
	hiraHitMap = make(map[string]CardInfo)
	frequencyMapOf_IsFineOnChars = make(map[string]int)
	frequencyMapOf_need_workOn = make(map[string]int)
	pulledButNotUsedMap = make(map[string]int)
	total_prompts = 0
	game = "off"
	gameOn = false
	game_duration = 998
	game_loop_counter = 0
	total_prompts = 0
	//
	//goland:noinspection ALL
	fmt.Println("\nArrays and maps flushed:\n")
	fmt.Println("    cyclicArrayOfTheJcharsGottenWrong")
	fmt.Println("    cyclicArrayHits")
	fmt.Println("    cyclicArrayPulls and hiraHitMap and pulledButNotUsedMap")
	fmt.Println("    frequencyMapOf_IsFineOnChars")
	//goland:noinspection ALL
	fmt.Println("    frequencyMapOf_need_workOn\n")
	fmt.Println("  And, all Game values have also been reset")
}

func notes_on_kana() {
	//goland:noinspection ALL  **do-this**
	fmt.Println("\nIn the traditional Hepburn romanization system, the sound じ in hiragana is romanized as \"ji\" \n" +
		"and the katakana ジ is also romanized as \"ji\" \n\n" +
		"However, in some other romanization systems like the Nihon-shiki and Kunrei-shiki, the sound じ is romanized as\n" +
		" \"zi\" instead of \"ji\"\n\n" +
		"The sound gi:ぎ in hiragana is romanized as \"gi\" and the katakana ギ is also romanized as \"gi\"\n")
	//goland:noinspection ALL  **do-this**
	fmt.Println("゜is called \"handakuten\" 半濁点 translates to \"half-voicing mark\" or \"semi-voiced mark\"\n" +
		"゛is called \"dakuten\" 濁点 meaning 'voiced mark' or 'voicing mark'")
	fmt.Println("\"digraphs\" is the word that refers to what I have called conjunctions, like ひゅ, for example ")
}

func respond_to_UserSuppliedDirective(in, objective_kind string) (prompt, objective, kind string) { // - -
	/*
	 'nts' for some background on Romaji conventions
	 'dir' redisplay this menu of available Directives
	 'gdc' set the Duration Counter for a Game session
	 'bgs' or 'goff' Begin or end a Game Session
	 '?' context-sensitive help on the current character
	 '??' for help on a particular Hiragana character
	 'st' Statistics
	 'abt' for trivia about this app
	 'rs' to reset (flush or clear) all stats logs etc.
	 'rm' Read the current contents of the Maps
	 'stc' (Set-Card) force the use of a specific card
	 'exko' load the Extended Kata deck
	 'exkf' un-load the Extended Kata deck
	*/
	var count int
	switch in {
	case "abt":
		about_app()
	case "gdc":
		fmt.Println("Enter a number for how many prompts there will be in the game")
		_, _ = fmt.Scan(&count)
		game_duration = count - 2
	case "bgs":
		// game_loop_counter ++
		game_on()
	case "goff":
		game_off()
	case "rs":
		reset_all_data()
	case "q":
		os.Exit(1)
	case "??": // Directives follow:
		handle_doubleQuestMark_directive(objective_kind)
	case "?":
		fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
	case "stc":
		prompt, objective, kind = reSet_aCard_andThereBy_reSet_thePromptString()
	case "st":
		newHits()
	case "nts":
		notes_on_kana()
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
		read_pulledButNotUsedMap()
		// read_pulls_not_used_array()
	case "exko":
		include_Extended_kata_deck = true
		fmt.Println("Extended Kata deck has been loaded")
	case "exkf":
		include_Extended_kata_deck = false
		fmt.Println("Extended Kata deck has been un-loaded")
	default:
		// fmt.Println("Directive not found") // Does not work because only existent cases are passed to the switch
	}
	return prompt, objective, kind
}

/*
// This func, sans "_Extended" at the end of its name, becomes a testing version of pick_RandomCard_Assign_fields()
// ... for fileOfCardsE : the deck of Extended Kata
func pick_RandomCard_Assign_fields_Extended() (promptField, objective, objective_kind string) { // - -
	randIndexE := rand.Intn(len(fileOfCardsE))
	aCard = fileOfCardsE[randIndexE] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kata
	objective = aCard.Romaji
	objective_kind = "Extended_Romaji" // Used to set a special prompt for Extended Kata
	whichDeck = 4
	return promptField, objective, objective_kind
}
*/

/*
// Drop "_Test", at the end of its name, & this func becomes a short testing version of pick_RandomCard_Assign_fields()
func pick_RandomCard_Assign_fields_Test() (promptField, objective, objective_kind string) { // - -
	randIndex := rand.Intn(len(fileOfCardsMostDifficult))
	aCard = fileOfCardsMostDifficult[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Hira
	objective = aCard.Romaji
	objective_kind = "Romaji"
	whichDeck = 3
	return promptField, objective, objective_kind
}
*/
