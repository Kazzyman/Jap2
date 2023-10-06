package main

import (
	"fmt"
	"regexp"
)

// Convenience global, used in two functions here, and in the calling func
var foundElement *charSetStruct

// ... Used only in handle_doubleQuestMark_directive()  '(a Directive)'
func locateCardAndDisplayHelpFieldsContainedInIt(targetString string) {
	var isAlphanumeric bool
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	//
	switch true { // The single 'case:' below will be true if targetString is Alpha
	case findAlphasIn.MatchString(targetString):
		// targetString is Alpha, therefore targetString must be a Romaji string
		isAlphanumeric = true
	default: // 'else' the targetString must, instead, be a Hiragana char
		isAlphanumeric = false
	}
	// This should be OK as is (Oct 6, 2023)
	if isAlphanumeric { // Then we probably have a Romaji string to locate help on
		// Iterate through the array to find the element with the desired Romaji
		for _, card := range fileOfCards {
			if card.Romaji == targetString {
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
	} else { // else we have a Hiragana char to find help on
		// Iterate through array to find element w the desired Hiragana
		for _, card := range fileOfCards {
			if card.Hira == targetString {
				foundElement = &card
				break
			}
		}
		if foundElement != nil { // Providing that we found something ...
			fmt.Printf("%s", colorRed)
			fmt.Println("Hiragana Help on:", foundElement.Hira)
			fmt.Printf("%s", colorReset)

			fmt.Println(foundElement.HiraHint)
			fmt.Println(foundElement.TT_Hint)
			fmt.Println(foundElement.SansR_Hint)
		} else {
			fmt.Println("Element not found in: locateCardAndDisplayHelpFieldsContainedInIt(targetString string)")
		}
	}
}

// used ONLY in the 'set' directive to reSet the prompt & "aCard." fields
func silentlyLocateCard(setKeyRequest string) { //  - -

	var isAlphanumeric bool

	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(setKeyRequest):
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}

	if isAlphanumeric == false { // ... then we should be safe to proceed with what will be a Hiragana char
		for _, card := range fileOfCards {
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
