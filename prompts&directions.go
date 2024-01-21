package main

import (
	"fmt"
)

func List_of_Directives() {
	fmt.Println("\n\nView source code at https://github.com/Kazzyman/Jap2")
	fmt.Println("    Use Alpha-Numeric (US) input-mode on your system to:")
	fmt.Println("        Enter '" + colorGreen +
		"nts" + colorReset +
		"' for some background on Romaji conventions")
	fmt.Println("        Enter '" + colorGreen +
		"dir" + colorReset +
		"' redisplay this menu of available Directives")
	fmt.Println("        Enter '" + colorGreen +
		"gdc" + colorReset +
		"' set the Duration Counter for a Game session ")
	fmt.Println("        Enter '" + colorGreen +
		"bgs" + colorReset +
		"' or " + colorGreen +
		"'goff'" + colorReset +
		" Begin or end a Game Session ")
	fmt.Println("        Enter '" + colorGreen +
		"?" + colorReset +
		"' context-sensitive help on the current character")
	fmt.Println("        Enter '" + colorGreen +
		"??" + colorReset +
		"' for help on a particular Hiragana or Romaji")
	fmt.Println("        Enter '" + colorGreen +
		"st" + colorReset +
		"' Statistics")
	fmt.Println("        Enter '" + colorGreen +
		"abt" + colorReset +
		"' for trivia about this app")
	fmt.Println("        Enter '" + colorGreen +
		"rs" + colorReset +
		"' to reset (flush or clear) all stats logs etc.")
	fmt.Println("        Enter '" + colorGreen +
		"rm" + colorReset +
		"' Read the current contents of the Maps")
	fmt.Println("        Enter '" + colorGreen +
		"stc" + colorReset +
		"' (Set-Card) force the use of a specific card")
	fmt.Println("        Enter '" + colorGreen +
		"exko" + colorReset +
		"' load the Extended Kata deck")
	fmt.Println("        Enter '" + colorGreen +
		"exkf" + colorReset +
		"' un-load the Extended Kata deck")
	fmt.Println("        Enter '" + colorGreen +
		"konly" + colorReset +
		"' Use only Kata prompting and Romaji objective, rs to ReSet")
	fmt.Println("        Enter '" + colorGreen +
		"honly" + colorReset +
		"' Use only Hira prompting, rs to ReSet")
	fmt.Println("        Enter '" + colorGreen +
		"ronly" + colorReset +
		"' Use only Romaji prompting, rs to ReSet")
	fmt.Println("        Enter '" + colorGreen +
		"donly" + colorReset +
		"' Difficult Kata only, rs to ReSet")
	fmt.Println("        Enter '" + colorGreen +
		"q" + colorReset +
		"', (quit) terminate the app")
}

//

// Special prompts for use when soliciting second guesses
func promptForRomaji1(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected," + colorReset +
		" you must try to guess\n Here:> ")
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}
func promptForHira1(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Hiragana input-mode expected," + colorReset +
		" you must try to guess\n Here:> ")
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// Special prompts for use when soliciting final guesses
func promptForRomaji2(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected," + colorReset +
		" you must guess, just one more time\n Here:> ")
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}
func promptForHira2(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Hiragana input-mode expected," + colorReset +
		" you must guess, just one more time\n Here:> ")
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// Standard prompts for use when NOT soliciting second, or final, guesses
func promptForHiraWithDiro(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Hiragana input-mode expected, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}
func promptForRomajiWithDiro(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// Standard prompts for use when NOT soliciting second, or final, guesses
func promptForHiraWithDir(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Hiragana?, or 'dir'\n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}
func promptForRomajiWithDir(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji?, or 'dir'\n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// Special prompts for Extended Kata, if|when deployed **********************
// ... Standard: used when NOT soliciting second, or final, guesses ***
func promptForRomajiWithDirE(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" E-Romaji input-mode expected, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// ... Special: used when soliciting second, or final, guesses ***
func promptForRomajiE(prompt string) (usersGuessOrOptionDirective string) { // - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" E-Romaji input-mode expected," + colorReset +
		" you must guess\n Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
} // ***********************************************************************

// 'Directive Menu' ; displays only at inception
func display_List_of_Directives() { // (unique)     - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Println("\n")
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n\n", total_prompts)
	fmt.Printf("Extended Kata deck is loaded: %t \n\n", include_Extended_kata_deck)
	fmt.Printf("Limited to Kata prompts with romaji objectives: %t \n", limitedToKataPrompts)
	fmt.Printf("Limited to Hira prompts only:                   %t \n", limitedToHiraPrompts)
	fmt.Printf("Limited to Romaji prompts only:                 %t \n", limitedToRomaPrompts)
	fmt.Printf("Limited to Difficult Kata only:                 %t \n\n", limitedToDifficultKata)
}

// 'Directive Menu' ; displays only in response to "Dir" Directive
func re_display_List_of_Directives() { // (unique)     - -
	frequencyMapRightOrOops := make(map[string]int)
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i] // Parse the cyclicArray
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		// ... one such entry per unique string
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Then, parse the map which was created, and loaded, above
	// As is the way with maps, the frequency map has only one entry for Right, & one for Oops
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" { // Finds the one potential entry for Right
			total_prompts = freq // Obviously, total_prompts has been declared as a global, elsewhere
		} else if str == "Oops" { // Finds the one potential entry for Oops
			total_prompts = total_prompts + freq
		}
	}
	List_of_Directives()
	//goland:noinspection ALL
	fmt.Printf("Game counter: %d, Game Duration: %d \n", game_loop_counter, game_duration+2)
	fmt.Printf("Current Prompt Count Total: %d \n\n", total_prompts)
	fmt.Printf("Extended Kata deck is loaded: %t \n\n", include_Extended_kata_deck)
	fmt.Printf("Limited to Kata prompts with romaji objectives: %t \n", limitedToKataPrompts)
	fmt.Printf("Limited to Hira prompts only:                   %t \n", limitedToHiraPrompts)
	fmt.Printf("Limited to Romaji prompts only:                 %t \n", limitedToRomaPrompts)
	fmt.Printf("Limited to Difficult Kata only:                 %t \n\n", limitedToDifficultKata)
	//
}
