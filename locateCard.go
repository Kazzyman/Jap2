package main

import (
	"fmt"
	"regexp"
)

// ::: Used only in handle_doubleQuestMark_directive()  '(a Directive)'
// (temporarily deprecated) ********************* v v v v v ------------------------------------------------v v v v v uses actual_objective_type
func locateCardAndDisplayHelpFieldsContainedInIt_deprecated(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn, actual_objective_type string) {
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	// Determine whether a Hira or a Romaji was entered to find help on
	switch true { // The single 'case:' below will be true if Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn is Alpha
	case findAlphasIn.MatchString(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn):
		// Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn is Alpha, therefore must be a Romaji string
		isAlphanumeric = true
	default: // 'else' must, instead, be a Hiragana char we are to find help on
		isAlphanumeric = false
	}
	// * * * * * *
	// Help for a Romaji prompt, should give only the SansR_Hint if actual_objective_type == "Romaji" ****
	if isAlphanumeric { // We probably have a Romaji string to locate help on
		// Iterate through the array to find the element with the desired Romaji
		for _, card := range fileOfCardsAllSequential {
			if card.Romaji == Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn { // It is a Romaji
				foundElement = &card // foundElement is a global
				break
			}
		}
		if foundElement != nil { // Providing that we found something ...
			fmt.Printf("%s", colorRed)
			fmt.Println("Romaji Help on:", foundElement.Romaji)
			fmt.Printf("%s", colorReset)
			if actual_objective_type == "roma" {
				fmt.Println(foundElement.SansR_Hint)
			} else {
				fmt.Println(foundElement.HiraHint)
				fmt.Println(foundElement.KataHint)
				fmt.Println(foundElement.TT_Hint)
			}
		}
		// * * * * * *
		// Help for a Hiragana prompt, should give only the SansR_Hint if actual_objective_type == "Romaji" ****
	} else { // We probably have a Hiragana string to locate help on
		// Iterate through array to find element w the desired Hiragana
		for _, card := range fileOfCardsAllSequential {
			if card.Hira == Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn { // It is a Hira
				foundElement = &card
				break
			}
		}
		if foundElement != nil { // Providing that we found something ...
			fmt.Printf("%s", colorRed)
			fmt.Println("Hiragana Help on:", foundElement.Hira)
			fmt.Printf("%s", colorReset)
			if actual_objective_type == "roma" { // this if-else clause is different/extra ????????????
				fmt.Println(foundElement.SansR_Hint)
			} else {
				fmt.Println(foundElement.HiraHint)
				fmt.Println(foundElement.KataHint)
				fmt.Println(foundElement.TT_Hint)
			}
		} else {
			fmt.Println("Element not found in: locateCardAndDisplayHelpField...")
		}
	}
}

// Used only in handle_doubleQuestMark_directive()  '(a Directive)' ------------
// ...  -------------  differs from the deprecated ver in that does not use actual_objective_type --------- v v v v v
// func locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn, actual_objective_type string) {
func locateCardAndDisplayHelpFieldsContainedInIt(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn string) {
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	// Determine whether a Hira or a Romaji was entered to find help on
	switch true { // The single 'case:' below will be true if Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn is Alpha
	case findAlphasIn.MatchString(Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn):
		// Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn is Alpha, therefore must be a Romaji string
		isAlphanumeric = true
	default: // 'else' must, instead, be a Hiragana char we are to find help on
		isAlphanumeric = false
	}
	// * * * * * *
	// Help for a Romaji prompt
	if isAlphanumeric { // We probably have a Romaji string to locate help on
		// Iterate through the array to find the element with the desired Romaji
		for _, card := range fileOfCardsAllSequential {
			if card.Romaji == Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn { // It is a Romaji
				foundElement = &card // foundElement is a global
				break
			}
		}
		if foundElement != nil { // Providing that we found something ...
			fmt.Printf("%s", colorRed)
			fmt.Println("Romaji Help on:", foundElement.Romaji)
			fmt.Printf("%s", colorReset)
			fmt.Println(foundElement.HiraHint)
			fmt.Println(foundElement.KataHint)
			fmt.Println(foundElement.TT_Hint)
		}
		// * * * * * *
		// Help for a Hiragana prompt
	} else { // We probably have a Hiragana string to locate help on
		// Iterate through array to find element w the desired Hiragana
		for _, card := range fileOfCardsAllSequential {
			if card.Hira == Hira_or_Romaji_input_sameAsPrompt_toFindHelpOn { // It is a Hira
				foundElement = &card
				break
			}
		}
		if foundElement != nil { // Providing that we found something ...
			fmt.Printf("%s", colorRed)
			fmt.Println("Hiragana Help on:", foundElement.Hira)
			fmt.Printf("%s", colorReset)
			fmt.Println(foundElement.HiraHint)
			fmt.Println(foundElement.KataHint)
			fmt.Println(foundElement.TT_Hint)
		} else {
			fmt.Println("Element not found in: locateCardAndDisplayHelpField...")
		}
	}
}

// ::: used ONLY in the 'stc' directive to reSet the prompt & "aCard." fields
func silentlyLocateCard(setKeyRequest string) { //  - -
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)

	if non_standard_origin_stcR {
		switch true {
		case findAlphasIn.MatchString(setKeyRequest):
			isAlphanumeric = true
		default:
			isAlphanumeric = true
		}

		if isAlphanumeric == true { // ... then we should be safe to proceed with what will be a romaji char
			for _, card := range fileOfCardsAllSequential {
				if card.Romaji == setKeyRequest {
					// v v v if we find a 'card' in the range of 'fileOfCards',
					// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
					foundElement = &card // foundElement is a global var and contains all the fields of a card or element
					// fmt.Printf("the card we got was: %s\n", foundElement)
					// todo : update the prompt and change the prompt type to Romaji ??? I think we got that now
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Element not found in: silentlyLocateCard(setKeyRequest string)")
			}
		} else {
			fmt.Printf(colorRed)
			fmt.Println("\nYou bastard!")
			fmt.Printf("\nYou have killed me with a Hira string instead of a Romaji char\n\n")
			fmt.Printf("the card we got was: %s\n\n", foundElement)

			fmt.Printf(colorReset)
		}
	} else { // we are doing a Hira reset instead of a Romaji reset
		switch true {
		case findAlphasIn.MatchString(setKeyRequest):
			isAlphanumeric = true
		default:
			isAlphanumeric = false
		}

		if isAlphanumeric == false { // ... then we should be safe to proceed with what will be a Hiragana char
			for _, card := range fileOfCardsAllSequential {
				if card.Hira == setKeyRequest {
					// v v v if we find a 'card' in the range of 'fileOfCards',
					// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
					foundElement = &card // foundElement is a global var and contains all the fields of a card or element
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Element not found in: silentlyLocateCard(setKeyRequest string)")
			}
		} else {
			fmt.Printf(colorRed)
			fmt.Println("\nYou bastard!")
			fmt.Printf("\nYou have killed me with an Alpha string instead of a Hiragana\n\n")
			fmt.Printf(colorReset)
		}
	}
}

// ::: used ONLY in the 'stcr' directive to reSet the prompt & "aCard." fields
func silentlyLocateCardr(setKeyRequest string) { //  - -

	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)

	if non_standard_origin_stcR { // ::: This is never true!!  this gets me to the bottom.

		switch true {
		case findAlphasIn.MatchString(setKeyRequest):
			isAlphanumeric = true

		default:
			isAlphanumeric = true
		}

		if isAlphanumeric != true { // ... then we should be safe to proceed with what will be a romaji char

			for _, card := range fileOfCardsAllSequential {

				if card.Romaji == setKeyRequest {
					// v v v if we find a 'card' in the range of 'fileOfCards',
					// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
					foundElement = &card // foundElement is a global var and contains all the fields of a card or element
					fmt.Printf("since isAlphanumeric != true : the card we got was: %s\n", foundElement)
					// todo : update the prompt and change the prompt type to Romaji ??? I think we got that now
					break
				}
			}

			if foundElement == nil {
				fmt.Println("Element not found in: silentlyLocateCard(setKeyRequest string)")
			}
		} else {
			fmt.Printf(colorRed)
			fmt.Println("\nYou bastard!")
			fmt.Printf("\nYou have killed me with a Hira string instead of a Romaji char\n\n")
			fmt.Printf("the card we got was: %s\n\n", foundElement)
			fmt.Printf(colorReset)
		}
	} else { // we are doing a Hira reset instead of a Romaji reset

		switch true {
		case findAlphasIn.MatchString(setKeyRequest):
			isAlphanumeric = false

		default:
			isAlphanumeric = true
		}

		if isAlphanumeric == false { // ... then we should be safe to proceed with what will be a Hiragana char

			for _, card := range fileOfCardsAllSequential {
				if card.Romaji == setKeyRequest {
					// v v v if we find a 'card' in the range of 'fileOfCards',
					// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
					foundElement = &card // foundElement is a global var and contains all the fields of a card or element
					break
				}
			}

			if foundElement == nil {
				fmt.Println("Element not found in: silentlyLocateCard(setKeyRequest string)")
			}
		} else {
			fmt.Printf(colorRed)
			fmt.Println("\nYou bastard!")
			fmt.Printf("\nYou have killed me with an Alpha string instead of a Hiragana\n\n")
			fmt.Printf(colorReset)
		}
	}
}
