package main

import "fmt"

func promptForRomaji(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected,\n Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

func promptForHira(promptField string) (usersGuessOrOptionDirective string) { //  - -
	fmt.Printf("%s", promptField)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Hiragana input-mode expected,\n Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

func promptForHiraWithDir(prompt string) (usersGuessOrOptionDirective string) { //               - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Hiragana input-mode expected, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}
func promptForRomajiWithDir(prompt string) (usersGuessOrOptionDirective string) { //             - -
	fmt.Printf("%s", prompt)
	fmt.Printf("%s", colorCyan)
	fmt.Printf(" Romaji input-mode expected, or '?' for help with: %s \n", prompt)
	fmt.Printf(" or, type '??' for help with something else ... \n")
	fmt.Printf(" Here:> ")
	fmt.Printf("%s", colorReset)
	_, _ = fmt.Scan(&usersGuessOrOptionDirective)
	return usersGuessOrOptionDirective
}

// 'Directive Menu' shared by all exercises; displays only at inception or in response to the 'dir' Directive
func display_Listing_of_Directives() { // (unique)     - -
	fmt.Println("View source code at https://github.com/Kazzyman/Japanese")
	fmt.Println("    (using US or Alpha-Numeric input mode):")
	fmt.Println("        Enter 'dir' to redisplay this menu of available directives")
	fmt.Println("        Enter 'notes' for some background on Romaji conventions")
	fmt.Println("        Enter '?' for context-sensitive help ")
	fmt.Println("        Enter '??' for help on a particular Hiragana char")
	fmt.Println("        Enter 'set' to set a new specified prompt & \"key\" ")
	fmt.Println("        Enter 'stat' to view what you have done so far in the current session")
	fmt.Println("        Enter 'reset' to reset (flush or clear) the stats logs")
	fmt.Println("        Enter 'rm' to read_map_of_fineOn")
	// fmt.Println("        Enter 'stack' to prime or stack the frequencyMapOf_IsFineOnChars map")
	fmt.Println("        Enter 'exit' or 'quit' to terminate the app")
	//goland:noinspection ALL
	fmt.Println("\n")
}
