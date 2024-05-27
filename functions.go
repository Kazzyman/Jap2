package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

// Creates a func named check_error which takes one parameter "e" of type error
func check_error(e error) { // ::: - -
	if e != nil {
		panic(e) // use panic() to display error code
	}
}

// DIRECTIVES : --------------------------------------------------------------------------------------------
// func handle_doubleQuestMark_directive(actual_objective_type string) { //        - -
func handle_doubleQuestMark_directive() { // ::: - -
	var Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn string
	//
	fmt.Printf("\n  -- Type either a Hiragana or Romaji prompt you need help with:> ")
	_, _ = fmt.Scan(&Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn)
	//
	// locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn, actual_objective_type)
	locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn)
	fmt.Println("")
}

// Handles the Directive ::: 'stc'
func reSet_via_a_hira_aCard_andThereBy_reSet_thePromptString() { // ::: - -
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
		// ::: we now are wanting to reset all of the variables associated with the fields in our new card ??
		// ::: ... but that may be a difficult task: we will need to know how those fields will be used. MAYBE WE SHOULD JUST SET THEM FROM aCard.x just-before they are needed ???????????????????????????****************
		//                                                                                           :::                    ^^^  always!!!!!!!!!!!
		// But; if we could do THAT we never needed the named vars in the first place!!!!!!!! But if we take context into account maybe we can know what each of the aCard.x fields are being used for??????????????????????
		// ::: in other words, can you always know from context the type of var you are looking for ??  ----  Nah! I found a resolution I can live with (no such thing as kata prompting with hira objective).
		/*
							we would need to figure out what the objective (hira or roma) should be based on the current objective type (actual_objective_type) but how would that even be knowable?????
							actual_prompt_char = aCard.Hira ::: is this one indeterminable? No, it can never be Kata. It will be Hira in the case of stc and it will be Romaji in the case of stcr !!!!?????
							actual_objective = ::: done
							actual_objective_type ::: was already set and knowable

						here, in stc, we ask for the か (hira) card to be set. What should it prompt us with? There ARE 3 reasonable options (roma-kata, and yes, hira). Albeit, if we are under limitation, i.e., we are not doing a mix ...
						... then we would probably like to and expect to see the prompt char type to match the current limitation on prompting. But we cannot say what the objective will be.
							imagine the new card is the ka card and we know that the objective is aCard.Romaji because we found that the actual_objective_type is roma
							... however, the prompt field could end up being either the aCard.Hira or the aCard.Kata, ::: or could it ??????? Well, not if we follow the imposed limitation. Easy, but the mix could be tricky.
			::: some of this has been solved by eliminating the possibility of kata prompting with hira objective.
		*/
		if limitedToKataPrompts { // ::: we always know the card. So, if we also know one more thing, e.g., the type of the prompt. We then know other things.
			// and, we always know the type of the prompt, ALWAYS!!!! Unfortunately, there are three (not two) prompt types. If there were only two prompt types we could know the object by elimination. We need one more piece of info.
			// : the "response" type. That is, the (one of two) possible types of objective type. Where then, is actual_objective_type set??
			actual_prompt_char = aCard.Kata
			// but the actual_objective type is undetermined ?? the objective type (hira or roma) is undetermined ??
		} else if limitedToRomaPrompts {
			actual_prompt_char = aCard.Romaji
		} else if limitedToHiraPrompts {
			// if you have been getting hira prompts, and you run stc and supply a hira char ... you will get that hira char as the prompt
			actual_prompt_char = aCard.Hira
		} else if limitedToDifficultKata {
			actual_prompt_char = aCard.Kata
		} else {
			// tricky ??????
		}

		/*
			If, here in stc, we ask that か be the new aCard; we are saying that we want ka:か:(or kata ka) to be the new prompt. And, if we can determine the type
		*/
		// by the current objective type
		if actual_objective_type == "roma" {
			actual_objective = aCard.Romaji
			// so, the actual_prompt_char could be either aCard.Hira or aCard.Kata, but it should never be aCard.Romaji
		} else { // ::: else it is "hira"
			actual_objective = aCard.Hira
		}
		// ::: or, said differently :
		if actual_objective_type == "hira" {
			actual_objective = aCard.Hira
		} else { // ::: else it is "roma"
			actual_objective = aCard.Romaji
		}

		fmt.Println("")
	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theHiraganaOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                                 // Set the global var-object 'aCard'
		// ::: we now would like to reset all of the variables associated with the fields in our new card ??
		// ::: ... but that may be a challenging task because there is no way to know how those fields will be used.
		/*
			we need to figure out what the objective (hira or roma) should be based on the current objective type
			actual_prompt_char = aCard.Hira ::: is this one indeterminable? No, it can never be Kata. It will be Hira in the case of stc and it will be Romaji in the case of stcr !!!!
			actual_objective = ::: done
			actual_objective_type ::: was already set and knowable
		*/
		// by the current objective type
		// ::: said yet another way :
		if actual_objective_type == "roma" {
			actual_objective = aCard.Romaji
			// so, the actual_prompt_char could be either aCard.Hira or aCard.Kata, but it should never be aCard.Romaji
		}
		if actual_objective_type == "hira" {
			actual_objective = aCard.Hira
			// so, the actual_prompt_char could be either aCard.Romaji or aCard.Kata, but it should never be aCard.Hira
		}
		fmt.Println("")
	}
}

// Handles the Directive ::: 'stcr'
func reSet_aCard_via_a_romaji_andThereBy_reSet_thePromptString() { // ::: - -
	var theRomajiOfCardToSilentlyLocate string
	var isAlphanumeric bool

	fmt.Printf("\nEnter a Romaji to")
	fmt.Printf("%s", colorCyan) //
	fmt.Printf(" reSet the prompt :> ")
	fmt.Printf("%s", colorReset) //
	_, _ = fmt.Scan(&theRomajiOfCardToSilentlyLocate)

	// Determine if the user has entered a valid Romaji char (instead of, accidentally, Hira char or string)
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(theRomajiOfCardToSilentlyLocate):
		isAlphanumeric = true
		// ::: debug off fmt.Printf("verify value of switch argument, isAlphanumeric is %t \n", isAlphanumeric)("verify value of switch argument, isAlphanumeric is %t \n", isAlphanumeric)
	default:
		isAlphanumeric = true // todo ????????
	}
	// Tentatively, prepare to Scan for user's input and attempt locating a matching 'aCard'
	if isAlphanumeric != true {
		fmt.Println("Are you in Hira input mode?")
		fmt.Printf("... if so, change it to Romaji (or I mignt die)\n")
		fmt.Printf("%s", colorRed) //
		fmt.Printf(" cautiously ")
		fmt.Printf("%s", colorCyan)
		fmt.Printf("re-enter your selection, in Romaji mode :> ")
		fmt.Printf("%s", colorReset)
		_, _ = fmt.Scan(&theRomajiOfCardToSilentlyLocate)
		// May yet send an Alpha string to the next func, which will itself deal with it elegantly
		silentlyLocateCard(theRomajiOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                               // Set the global var-object 'aCard'
		// new_prompt, new_objective, actual_objective_type
		actual_prompt_char = aCard.Romaji

		/*
				we need to figure out what the objective (hira or roma) should be based on the current objective type
				actual_prompt_char = aCard.Hira ::: is this one indeterminable? No, it can never be Kata. It will be Hira in the case of stc and it will be Romaji in the case of stcr !!!!
				actual_objective = ::: done
				actual_objective_type ::: was already set and knowable

			here, in stcr, we ask for the ka card to be set. What should it prompt us with? There ARE 3 options. If we are under limitation, i.e., we are not doing a mix ...
			... then we would probably like to and expect to see the prompt char type to match the current limitation.
				imagine the new card is the ka card and we know that the objective is aCard.Romaji because we found that the actual_objective_type is roma
				... however, the prompt field could end up being either the aCard.Hira or the aCard.Kata, ::: or could it ???????

		*/

		actual_objective = aCard.Hira
		actual_objective_type = "hira"
		fmt.Println("")
		// todo, must be re-prompting here ???? it is prompting for and testing for the original Romaji/Hiragana

	} else {
		// Confidently, go-looking for user's input: locate matching 'aCard'
		silentlyLocateCard(theRomajiOfCardToSilentlyLocate) // Set the Convenience-global: foundElement
		aCard = *foundElement                               // Set the global var-object 'aCard'
		actual_prompt_char = aCard.Romaji
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
		fmt.Println("")
	}
	// return prompt, objective, actual_objective_type
}

func detectDirective(in string) (result bool) { // ::: - -
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
		 'stc' (Set-Card) force the use of a specific card (Hira input)
		'stcr' (Set-Card) force the use of a specific card (Romaji input)
		 'exko' load the Extended Kata deck
		 'exkf' un-load the Extended Kata deck
		'konly' Use Kata only
		'hko'  Use both Hira and Kata prompting
		"honly" .. Hira
		"ronly" .. Romaji
		"donly" .. Difficult
	*/
	if in == "stc" ||
		in == "stcr" ||
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
		in == "exkf" ||
		in == "konly" ||
		in == "honly" ||
		in == "ronly" ||
		in == "donly" ||
		in == "hko" {
		// Then:
		its_a_directive = true
	}
	return result
}

func game_on() (game string) { // ::: - -
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
func game_off() (game string) { // ::: - -
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

func read_pulledButNotUsedMap() { // ::: - -
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

func about_app() { // ::: - -
	fmt.Printf("\n" +
		"This app consists of the following files and lines of code:\n\n" +
		"1701 constants.go ~ 35\n" +
		"343 functions.go ~ 300\n" +
		"157 locateCard.go ~ 80\n" +
		"357 main.go ~ 320\n" +
		"127 memoryFunctions.go ~ 45\n" +
		"133 objectsAndMethods.go ~ 40\n" +
		"82 prompts&directions.go ~ 75\n" +
		"145 statsFunctions.go ~ 125 \n\n")
	countSLOC()
}

func reset_all_data() { // ::: - -
	// Flush (clear) the old stats and hits arrays
	limitedToKataPrompts = false
	limitedToHiraPrompts = false
	limitedToRomaPrompts = false
	limitedToDifficultKata = false
	include_Extended_kata_deck = false
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
	fmt.Println(colorCyan + "  Limitations re Kata, Hira, and Romaji prompting; as well as all Game values have also been reset\n" + colorReset)
}

func notes_on_kana() { // ::: - -
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

/*
.
.
*/
func respond_to_UserSupplied_Directive(usersSubmission string) { // ::: - -
	/* ::: as displayed
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
	'stc' (Set-Card) force the use of a specific card (Hira input)
	'stcr' (Set-Card) force the use of a specific card (Roma input)
	'exko' load the Extended Kata deck
	'exkf' un-load the Extended Kata deck
	*/
	// ::: alphabetically (mostly)
	var count int
	switch usersSubmission {
	case "?":
		fmt.Printf("\n%s\n%s\n%s\n\n", aCard.HiraHint, aCard.KataHint, aCard.TT_Hint)
	case "??": // Directives follow:
		// handle_doubleQuestMark_directive(actual_objective_type)
		handle_doubleQuestMark_directive()
	case "hko":
		limitedToKataPrompts = true
		limitedToHiraPrompts = true
		limitedToRomaPrompts = false
		limitedToDifficultKata = false
	case "konly":
		limitedToKataPrompts = true
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultKata = false
	case "honly":
		limitedToHiraPrompts = true
		limitedToKataPrompts = false
		limitedToRomaPrompts = false
		limitedToDifficultKata = false
	case "ronly":
		limitedToRomaPrompts = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToDifficultKata = false
	case "donly":
		limitedToDifficultKata = true
		limitedToKataPrompts = false
		limitedToHiraPrompts = false
		limitedToRomaPrompts = false
		/*
			.
		*/
	case "abt":
		about_app()
	case "bgs":
		// game_loop_counter ++
		game_on()
	case "dir": // reDisplay the DIRECTORY OF DIRECTIVES (and instructions):
		re_display_List_of_Directives()
	case "exko":
		include_Extended_kata_deck = true
		fmt.Println("Extended Kata deck has been loaded")
	case "exkf":
		include_Extended_kata_deck = false
		fmt.Println("Extended Kata deck has been un-loaded")
	case "gdc":
		fmt.Println("Enter a number for how many prompts there will be in the game")
		_, _ = fmt.Scan(&count)
		game_duration = count - 2
	case "goff":
		game_off()
	case "nts":
		notes_on_kana()
	case "q":
		os.Exit(1)
	case "rm":
		read_map_of_fineOn()
		read_map_of_needWorkOn()
		read_pulledButNotUsedMap()
		// read_pulls_not_used_array()
	case "rs":
		reset_all_data()
	case "st":
		st_stats_function()
	case "stc":
		reSet_via_a_hira_aCard_andThereBy_reSet_thePromptString()
	case "stcr":
		reSet_aCard_via_a_romaji_andThereBy_reSet_thePromptString()
	default:
		// fmt.Println("Directive not found") // Does not ever happen because only existent cases are passed to the switch.
	}
}

func st_stats_function() {
	newHits()
	if !include_Extended_kata_deck {
		fmt.Println("Extended Kata deck is NOT loaded\n")
	} else {
		fmt.Println("Extended Kata deck is loaded\n")
	}

	if limitedToKataPrompts {
		fmt.Printf("Limited to Kata prompts with romaji objectives: %t \n\n", limitedToKataPrompts)
	}
	if limitedToHiraPrompts {
		fmt.Printf("Limited to Hira prompts only: %t \n\n", limitedToHiraPrompts)
	}
	if limitedToRomaPrompts {
		fmt.Printf("Limited to Romaji prompts only: %t \n\n", limitedToRomaPrompts)
	}
	if limitedToDifficultKata {
		fmt.Printf("Limited to Difficult Kata only: %t \n\n", limitedToDifficultKata)
	}
}

/*
// This func, sans "_Extended" at the end of its name, becomes ::: a testing version of pick_RandomCard_Assign_fields()
// ... for fileOfCardsE : the deck of Extended Kata
func pick_RandomCard_Assign_fields_Extended() (promptField, objective, actual_objective_type string) { // - -
	randIndexE := rand.Intn(len(fileOfCardsE))
	aCard = fileOfCardsE[randIndexE] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Kata
	objective = aCard.Romaji
	actual_objective_type = "Extended_Romaji" // Used to set a special prompt for Extended Kata
	whichDeck = 4
	return promptField, objective, actual_objective_type
}
*/
//
//
/*
// Drop "_Test", at the end of its name, & this func becomes ::: a short testing version of pick_RandomCard_Assign_fields()
func pick_RandomCard_Assign_fields_Test() (promptField, objective, actual_objective_type string) { // - -
	randIndex := rand.Intn(len(fileOfCardsMostDifficult))
	aCard = fileOfCardsMostDifficult[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	promptField = aCard.Hira
	objective = aCard.Romaji
	actual_objective_type = "Romaji"
	whichDeck = 3
	return promptField, objective, actual_objective_type
}
*/
