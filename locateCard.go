package main

// **do-this** : is the tag I have used to denote sections that I have yet to 'complete'

import (
	"fmt"
	"regexp"
)

// Convenience global, used in two functions here, and in the calling func
var foundElement *charSetStruct

// This func works intelligently with Hiragana or Romaji in all Exercises, but only via the '??' Directive ...
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

	if isAlphanumeric { // Then we probably have a Romaji string to locate help on
		// Iterate through the array to find the element with the desired Romaji 'value' in card.KeyR
		for _, card := range fileOfCards {
			if card.Romaji == targetString { // card.KeyR is the Romaji-characters string
				foundElement = &card // foundElement is a global
				break
			}
		}
		if foundElement != nil { // Providing that we found something ...
			fmt.Printf("%s", colorRed)
			fmt.Println("Romaji Help on:", foundElement.Romaji)
			fmt.Printf("%s", colorReset)
			/*
				Hint1h
				Hint2k
				Hint3TT
				HintSansR
			*/
			// We probably have a Romaji string to locate help on, so skip HintSansR
			fmt.Println(foundElement.HiraHint)
			fmt.Println(foundElement.KataHint)
			fmt.Println(foundElement.TT_Hint)
		}
	} else { // else we have a Hiragana char to find help on
		// Iterate through array to find element w the desired Hiragana 'value' in card.Keyh  (see block comment below)
		for _, card := range fileOfCards {
			if card.Hira == targetString { // card.Keyh is the Hiragana character (see block comment below) ***
				foundElement = &card
				break
			}
		}
		if foundElement != nil { // Providing that we found something ...
			fmt.Printf("%s", colorRed)
			fmt.Println("Hiragana Help on:", foundElement.Hira) //  (see block comment below) ***
			fmt.Printf("%s", colorReset)
			/*

				HiraHint   string
				KataHint   string
				TT_Hint    string
				SansR_Hint string
				}
			*/
			// We have a Hiragana char to find help on, so skip Hint2k
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
			if card.Hira == setKeyRequest { // card.Value (card.Keyh) is the field that contains the Hiragana
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

func practice_this_card(aCard_reset_request_as_Romaji string) {
	for _, card := range fileOfCards {
		if card.Romaji == aCard_reset_request_as_Romaji { // card.KeyR is the field that contains the Romaji
			// v v v if we find a 'card' in the range of 'fileOfCards',
			// ... we set the foundElement global var, which is used in reSet_aCard_andThereBy_reSet_thePromptString()
			foundElement = &card // foundElement is a global var and contains all the fields of a card or element
			aCardA = *foundElement
			break
		}
	}
	if foundElement == nil {
		fmt.Println("Element not found in: silentlyLocateCard(setKeyRequest string)")
	}
}
