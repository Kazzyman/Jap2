package main

// ::: done
import (
	"fmt"
	"math/rand"
)

// maybe the strategy here should be to only add what is needed rather than delete things I don't need. ::: done

/*
aCard.Hira, aCard.Kata, aCard.Romaji ::: done

expected_response_type = "hira_char_as_users_guess"
expected_response_type = "romaji_char_as_users_guess" ::: done

displayed_prompt_type = "hira"  ::: all done
displayed_prompt_type = "romaji"
displayed_prompt_type = "kata"
*/

var pulledButNotUsedMap = make(map[string]int)

func pick_RandomCard_Assign_fields() (promptField, objective, objective_kind string) { // - -
	// There are 4 different exercises: four different picking sections each with its own prompting.
	// Set (acquire) one of the 4 combinations of promptField, objective, objective_kind
	// This section only does the initial pick, and MAY be duplicated below in the for loop
	if limitedToKataPrompts {
		displayed_prompt_type = "kata"
		promptField, objective, objective_kind = kata_prompting_romaji_objective() // 1
	} else if limitedToHiraPrompts {
		displayed_prompt_type = "hira"
		promptField, objective, objective_kind = hira_prompting_romaji_objective() // 2
	} else if limitedToRomaPrompts {
		displayed_prompt_type = "romaji"
		promptField, objective, objective_kind = roma_prompting_hira_objective() // 3
	} else if limitedToDifficultKata {
		displayed_prompt_type = "kata"
		promptField, objective, objective_kind = Difficult_kata_prompting_romaji_objective() // 4
	} else {
		promptField, objective, objective_kind = randomize_over_all_decks() // 1-4
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
					displayed_prompt_type = "kata"
					promptField, objective, objective_kind = kata_prompting_romaji_objective() // 1
				} else if limitedToHiraPrompts {
					displayed_prompt_type = "hira"
					promptField, objective, objective_kind = hira_prompting_romaji_objective()
				} else if limitedToRomaPrompts {
					displayed_prompt_type = "romaji"
					promptField, objective, objective_kind = roma_prompting_hira_objective()
				} else if limitedToDifficultKata {
					displayed_prompt_type = "kata"
					promptField, objective, objective_kind = Difficult_kata_prompting_romaji_objective()
				} else {
					promptField, objective, objective_kind = randomize_over_all_decks() // 1-4
				}
				break // Exit the inner loop, having a new and potentially novel promptField in hand
			}
		}
		maxTries++
		if !found {
			// If promptField is not found in the slice ...
			break // All of our work here being done! We hereby exit the outer 99999 for loop
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
	//
	if maxTries == 99999 {
		maxTries = 0
		cyclicArrayPulls = CyclicArrayPulls{}
		// read_pulledButNotUsedMap()
		// newHits()
		read_map_of_needWorkOn()
		cyclicArrayOfTheJcharsGottenWrong = CyclicArrayOfTheJcharsGottenWrong{}
		cyclicArrayHits = CyclicArrayHits{}
		// Also, flush (clear) the maps
		total_prompts = 0
		// kanjiHitMap = make(map[string]CardInfo)
		frequencyMapOf_IsFineOnChars = make(map[string]int)
		frequencyMapOf_need_workOn = make(map[string]int)
		fmt.Println(colorCyan + "You have finished all the cards; repeats will now ensue ... \n" + colorReset)
	}
	if displayed_prompt_type == "hira" {
		objective = aCard.Hira
	}
	if displayed_prompt_type == "romaji" {
		objective = aCard.Romaji
	}
	if displayed_prompt_type == "kata" {
		objective = aCard.Kata
	}

	if displayed_prompt_type == "kata" {
		objective = aCard.Kata
	}
	if displayed_prompt_type == "kata" {
		objective = aCard.Kata
	}
	if displayed_prompt_type == "kata" {
		objective = aCard.Kata
	}

	promptField = objective
	objective_kind = displayed_prompt_type
	if expected_response_type == "hira char as users guess" {
		displayed_prompt_type = "hira"
	} else if expected_response_type == "romaji char as users guess" {
		displayed_prompt_type = "romaji"
	}
	/*
		aCard.Hira, aCard.Kata, aCard.Romaji ::: done

		expected_response_type = "hira char as users guess"
		expected_response_type = "romaji char as users guess" ::: done

		displayed_prompt_type = "hira"
		displayed_prompt_type = "romaji"
		displayed_prompt_type = "kata" ::: done
	*/
	return // promptField, objective, objective_kind
}

// The 4 functions called above:

// 4 exercise categories :
//   Hira prompting, Romaji objective
//   Kata prompting, Romaji objective
//      Kata prompting, Hira objective, Missing, OK!
//   Romaji prompting, Hira objective

func hira_prompting_romaji_objective() (promptField, objective, objective_kind string) {
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsRare[randIndexRare]
		promptField = aCard.Hira
		displayed_prompt_type = "hira"
		objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsMedium[randIndexMedium]
		promptField = aCard.Hira
		displayed_prompt_type = "hira"
		objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsOften[randIndexOften]
		promptField = aCard.Hira
		displayed_prompt_type = "hira"
		objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
	}
	/*

				randIndex := rand.Intn(len(fileOfCards)) // data.go
				aCard = fileOfCards[randIndex]           // Randomly pick a 'card' from a 'deck' and store it in a global var
				// Hira prompting, Romaji objective:
				promptField = aCard.Hira
		displayed_prompt_type = "hira"
				objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
				objective_kind = "Romaji"

	*/
	return
}

func kata_prompting_romaji_objective() (promptField, objective, objective_kind string) {
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsRare[randIndexRare]
		promptField = aCard.Kata
		displayed_prompt_type = "kata"
		objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsMedium[randIndexMedium]
		promptField = aCard.Kata
		displayed_prompt_type = "kata"
		objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsOften[randIndexOften]
		promptField = aCard.Kata
		displayed_prompt_type = "kata"
		objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
	}

	/*

				randIndex := rand.Intn(len(fileOfCards)) // data.go
				aCard = fileOfCards[randIndex]           // Randomly pick a 'card' from a 'deck' and store it in a global var
				// Kata prompting, Romaji objective:
				promptField = aCard.Kata
		displayed_prompt_type = "kata"
				objective = aCard.Romaji
		expected_response_type = "romaji_char_as_users_guess"
				objective_kind = "Romaji"

	*/
	return
}

//   Kata prompting, Hira objective

func roma_prompting_hira_objective() (promptField, objective, objective_kind string) {
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsRare[randIndexRare] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
		displayed_prompt_type = "romaji"
		objective = aCard.Hira
		expected_response_type = "hira_char_as_users_guess"
		objective_kind = "Hira"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsMedium[randIndexMedium] // Note the capital S on fileOfCardsS
		promptField = aCard.Romaji
		displayed_prompt_type = "romaji"
		objective = aCard.Hira
		expected_response_type = "hira_char_as_users_guess"
		objective_kind = "Hira"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsOften[randIndexOften] // Note the capital S on fileOfCardsS
		promptField = aCard.Romaji
		displayed_prompt_type = "romaji"
		objective = aCard.Hira
		expected_response_type = "hira_char_as_users_guess"
		objective_kind = "Hira"
	}

	/*
				randIndex := rand.Intn(len(fileOfCards)) // data.go
				aCard = fileOfCards[randIndex]           // Randomly pick a 'card' from a 'deck' and store it in a global var
				// Romaji prompting, Hira objective:
				promptField = aCard.Romaji
		displayed_prompt_type = "romaji"
				objective = aCard.Hira
		expected_response_type = "hira_char_as_users_guess"
				objective_kind = "Hira"

	*/
	return
}

//
//
//

func Difficult_kata_prompting_romaji_objective() (promptField, objective, objective_kind string) {
	randIndexMK := rand.Intn(len(dataMostDiff)) // dataK.go
	aCard = dataMostDiff[randIndexMK]
	// Kata prompting, Romaji objective:
	promptField = aCard.Kata
	displayed_prompt_type = "kata"
	objective = aCard.Romaji
	expected_response_type = "romaji_char_as_users_guess"
	objective_kind = "Romaji"
	return
}

/*
// This one is the default func: it randomly "calls" one of the above four functions and also sets whichDeck (abandoned)
func randomize_over_all_decks_original() (promptField, objective, objective_kind string) {
	randIndex := rand.Intn(len(fileOfCards))   // data.go
	randIndexS := rand.Intn(len(fileOfCardsS)) // dataS.go
	// randIndexD := rand.Intn(len(fileOfCardsMostDifficult)) // dataD.go
	// This next dataset can be added to the ones above with Dir: exko (EXtendedKataOn)
	randIndexE := rand.Intn(len(fileOfCardsE)) // optional dataExt.go

	if include_Extended_kata_deck {
		randomFileOfCards = rand.Intn(8) // If including the extended deck 0-8 = 9
	} else {
		randomFileOfCards = rand.Intn(7) // Normal set of three 0-7 = 8
	}

	// Hira prompting, Romaji objective:
	if randomFileOfCards == 0 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Hira
displayed_prompt_type = "hira"
		objective = aCard.Romaji
expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
		whichDeck = 1
	}

	if randomFileOfCards == 1 {
		aCard = fileOfCardsS[randIndexS] // Note the capital S on fileOfCardsS
		promptField = aCard.Hira
displayed_prompt_type = "hira"
		objective = aCard.Romaji
expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
		whichDeck = 2
	}
	//
		if randomFileOfCards == 2 {
			aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Hira
displayed_prompt_type = "hira"
			objective = aCard.Romaji
expected_response_type = "romaji_char_as_users_guess"
			objective_kind = "Romaji"
			whichDeck = 3
		}

	//

	// Kata prompting, Romaji objective: (plus an option for including Extended Kata)
	if include_Extended_kata_deck { //              ^ ^                 v v v ^ ^ ^
		if randomFileOfCards == 8 { // The optional ninth exercise
			aCard = fileOfCardsE[randIndexE] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kata
displayed_prompt_type = "kata"
			objective = aCard.Romaji
expected_response_type = "romaji_char_as_users_guess"
			objective_kind = "Extended_Romaji" // Used to set a special prompt for Extended Kata
			whichDeck = 4
		}
	}
	if randomFileOfCards == 2 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
displayed_prompt_type = "kata"
		objective = aCard.Romaji
expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
		whichDeck = 1
	}

	if randomFileOfCards == 3 {
		aCard = fileOfCardsS[randIndexS] // Note the capital S on fileOfCardsS
		promptField = aCard.Kata
displayed_prompt_type = "kata"
		objective = aCard.Romaji
expected_response_type = "romaji_char_as_users_guess"
		objective_kind = "Romaji"
		whichDeck = 2
	}
	//
		if randomFileOfCards == 4 {
			aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kata
displayed_prompt_type = "kata"
			objective = aCard.Romaji
expected_response_type = "romaji_char_as_users_guess"
			objective_kind = "Romaji"
			whichDeck = 3
		}

	//

	// Romaji prompting, Hira objective:
	if randomFileOfCards == 4 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Romaji
displayed_prompt_type = "romaji"
		objective = aCard.Hira
expected_response_type = "hira_char_as_users_guess"
		objective_kind = "Hira"
		whichDeck = 1
	}

	if randomFileOfCards == 5 {
		aCard = fileOfCardsS[randIndexS] // Note the capital S on fileOfCardsS
		promptField = aCard.Romaji
displayed_prompt_type = "romaji"
		objective = aCard.Hira
expected_response_type = "hira_char_as_users_guess"
		objective_kind = "Hira"
		whichDeck = 2
	}
	//
		if randomFileOfCards == 7 {
			aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Romaji
displayed_prompt_type = "romaji"
			objective = aCard.Hira
expected_response_type = "hira_char_as_users_guess"
			objective_kind = "Hira"
			whichDeck = 3
		}

	//

	// Kata prompting, Hira objective:
	if randomFileOfCards == 6 {
		aCard = fileOfCards[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
		promptField = aCard.Kata
displayed_prompt_type = "kata"
		objective = aCard.Hira
expected_response_type = "hira_char_as_users_guess"
		objective_kind = "Hira"
		whichDeck = 1
	}

	if randomFileOfCards == 7 {
		aCard = fileOfCardsS[randIndexS] // Note the capital S on fileOfCardsS
		promptField = aCard.Kata
displayed_prompt_type = "kata"
		objective = aCard.Hira
expected_response_type = "hira_char_as_users_guess"
		objective_kind = "Hira"
		whichDeck = 2
	}
	//
		if randomFileOfCards == 10 {
			aCard = fileOfCardsMostDifficult[randIndexD] // Randomly pick a 'card' from a 'deck' and store it in a global var
			promptField = aCard.Kata
displayed_prompt_type = "kata"
			objective = aCard.Hira
expected_response_type = "hira_char_as_users_guess"
			objective_kind = "Hira"
			whichDeck = 3
		}

	//
	return
} // end of Old func
*/
