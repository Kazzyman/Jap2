package main

import (
	"fmt"
)

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

// Only used in check_it_for_needing_more_practice() // ... abandoned
func find_this_card_to_practice(promptField_found_in_map, currently_known_typeOf_objective string) { // - -
	if whichDeck == 1 {
		if currently_known_typeOf_objective == "Hira" {
			for _, card := range fileOfCards {
				if card.Hira == promptField_found_in_map {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Hira type card not found in: fileOfCards")
			}
		} else if currently_known_typeOf_objective == "Romaji" {
			for _, card := range fileOfCards {
				if card.Hira == promptField_found_in_map {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Romaji type card not found in: fileOfCards")
			}
		}
	} else if whichDeck == 2 {
		if currently_known_typeOf_objective == "Hira" {
			for _, card := range fileOfCardsS {
				if card.Hira == promptField_found_in_map {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Hira type card not found in: fileOfCardsS")
			}
		} else if currently_known_typeOf_objective == "Romaji" {
			for _, card := range fileOfCardsS {
				if card.Hira == promptField_found_in_map {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Romaji type card not found in: fileOfCardsS")
			}
		}
	} else if whichDeck == 3 {
		if currently_known_typeOf_objective == "Hira" {
			for _, card := range fileOfCardsMostDifficult {
				if card.Hira == promptField_found_in_map {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Hira type card not found in: fileOfCardsMostDifficult")
			}
		} else if currently_known_typeOf_objective == "Romaji" {
			for _, card := range fileOfCardsMostDifficult {
				if card.Hira == promptField_found_in_map {
					fmt.Printf("I found your card from the deck, prompt was %s \n", promptField_found_in_map)
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Romaji type card not found in: fileOfCardsMostDifficult")
			}
		}
	}
}
