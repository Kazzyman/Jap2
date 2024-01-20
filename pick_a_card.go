package main

import (
	"fmt"
	"math/rand"
)

var pulledButNotUsedMap = make(map[string]int)

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	if limitedToKataPrompts {
		promptField, objective, objective_kind = kata_prompting_romaji_objective()
	} else if limitedToHiraPrompts {
		promptField, objective, objective_kind = hira_prompting_romaji_objective()
	} else if limitedToRomaPrompts {
		promptField, objective, objective_kind = roma_prompting_hira_objective()
	} else {
		promptField, objective, objective_kind = randomize_over_all_decks()
	}
	// Now that we have a newly-acquired promptField, check to see if we have it stored in the cyclicArrayPulls slice ...
	// and, if it is so stored, obtain a replacement, and then look again through the entire slice. Repeat the entire
	// process as many times as may be required to finally obtain a value of promptField which is novel according to said slice.
	var maxTries int
	for maxTries < 99999 {
		found := false
		for _, lastPull := range cyclicArrayPulls.pulls {
			if lastPull == promptField {
				// We also wish to store these duplicates in the map, to keep a tally of such events -- accessible via the rm Dir
				pulledButNotUsedMap[promptField]++ // The '++' increments the int value associated with promptField
				// fmt.Printf("We've seen the pseudo-random char before; lastPull: %s and promptField: %s\n", lastPull, promptField)
				found = true
				if limitedToKataPrompts {
					promptField, objective, objective_kind = kata_prompting_romaji_objective()
				} else if limitedToHiraPrompts {
					promptField, objective, objective_kind = hira_prompting_romaji_objective()
				} else if limitedToRomaPrompts {
					promptField, objective, objective_kind = roma_prompting_hira_objective()
				} else {
					promptField, objective, objective_kind = randomize_over_all_decks()
				}
				break // Exit the inner loop, having a new and potentially novel promptField in hand
			}
		}
		maxTries++
		if !found {
			// If promptField is not found in the slice ...
			break // All of our work here being done! We hereby exit the outer naked for loop
		} else {
			// A match WAS found (and a new promptField value has therefore been obtained; so, we need to restart the entire process)
			continue // explicitly continue, i.e., restart range with the replacement value of promptField
		}
	}

	// Exited the loops having obtained a newly-acquired, and verified as novel, value of promptField, so, ...
	// ... store that newly-acquired promptField in an array to be used as a memory of already-seen-and-used chars
	cyclicArrayPulls.InsertKChar(promptField) // Only the chars that are actually used are put into the array
	// Also, store that newly-acquired promptField in pulledButNotUsedMap
	pulledButNotUsedMap[promptField]++ // The '++' increments the int value associated with promptField
	if maxTries == 99999 {
		maxTries = 0
		cyclicArrayPulls = CyclicArrayPulls{}
		read_pulledButNotUsedMap()
		newHits()
		read_map_of_needWorkOn()
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// Also, flush (clear) the maps
		total_prompts = 0
		// kanjiHitMap = make(map[string]CardInfo)
		frequencyMapOf_IsFineOnChars = make(map[string]int)
		frequencyMapOf_need_workOn = make(map[string]int)
		fmt.Println("You have finished all the cards, the decks will all be restarted")
	}
	return // promptField, objective, objective_kind
}

func kata_prompting_romaji_objective() (promptField, objective, objective_kind string) {

	// This is for a future project.
	// randIndexMK := rand.Intn(len(fileOfCardsKataPromptsOnly)) // dataK.go
	// aCard = fileOfCardsKataPromptsOnly[randIndexMK]

	// But for now ...
	randIndex := rand.Intn(len(fileOfCards)) // data.go
	aCard = fileOfCards[randIndex]           // Randomly pick a 'card' from a 'deck' and store it in a global var

	// Kata prompting, Romaji objective:
	promptField = aCard.Kata
	objective = aCard.Romaji
	objective_kind = "Romaji"

	return
}

func hira_prompting_romaji_objective() (promptField, objective, objective_kind string) {

	randIndex := rand.Intn(len(fileOfCards)) // data.go
	aCard = fileOfCards[randIndex]           // Randomly pick a 'card' from a 'deck' and store it in a global var

	// Hira prompting, Romaji objective:
	promptField = aCard.Hira
	objective = aCard.Romaji
	objective_kind = "Romaji"

	return
}

func roma_prompting_hira_objective() (promptField, objective, objective_kind string) {

	randIndex := rand.Intn(len(fileOfCards)) // data.go
	aCard = fileOfCards[randIndex]           // Randomly pick a 'card' from a 'deck' and store it in a global var

	// Romaji prompting, Hira objective:
	promptField = aCard.Romaji
	objective = aCard.Hira
	objective_kind = "Hira"

	return
}

func randomize_over_all_decks() (promptField, objective, objective_kind string) {
	randIndex := rand.Intn(len(fileOfCards))               // data.go
	randIndexS := rand.Intn(len(fileOfCardsS))             // dataS.go
	randIndexD := rand.Intn(len(fileOfCardsMostDifficult)) // dataD.go
	randIndexE := rand.Intn(len(fileOfCardsE))             // optional dataExt.go

	if include_Extended_kata_deck {
		randomFileOfCards = rand.Intn(13)
	} else {
		randomFileOfCards = rand.Intn(12)
	}

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

	// Kata prompting, Romaji objective: (plus an option for including Extended Kata)
	if include_Extended_kata_deck { //              ^ ^                 v v v ^ ^ ^
		if randomFileOfCards == 12 {
			aCard = fileOfCardsE[randIndexE] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kata
			objective = aCard.Romaji
			objective_kind = "Extended_Romaji" // Used to set a special prompt for Extended Kata
			whichDeck = 4
		}
	}

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
	return
}
