package main

// ::: done
import (
	"fmt"
	"math/rand"
)

var pulledButNotUsedMap = make(map[string]int)

func pick_RandomCard_Assign_fields() (promptField, objective, actual_objective_type string) { // - -
	// There are 4 different exercises: four different picking sections each with its own prompting.
	// Set (acquire) one of the 4 combinations of promptField, objective, actual_objective_type
	// This section only does the initial pick, and MAY be duplicated below in the for loop
	limitedToRomaPrompts = true // ::: why did I do this?
	if limitedToKataPrompts {
		actual_prompt_char_type = "kata"
		actual_prompt_char, actual_objective, actual_objective_type = kata_prompting_romaji_objective() // 1
	} else if limitedToHiraPrompts {
		actual_prompt_char_type = "hira"
		actual_prompt_char, actual_objective, actual_objective_type = hira_prompting_romaji_objective() // 2
	} else if limitedToRomaPrompts {
		actual_prompt_char_type = "romaji"
		actual_prompt_char, actual_objective, actual_objective_type = roma_prompting_hira_objective() // 3
	} else if limitedToDifficultKata {
		actual_prompt_char_type = "kata"
		actual_prompt_char, actual_objective, actual_objective_type = Difficult_kata_prompting_romaji_objective() // 4
	} else {
		randomize_over_all_decks() // 1-4
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
					actual_prompt_char_type = "kata"
					actual_prompt_char, actual_objective, actual_objective_type = kata_prompting_romaji_objective() // 1
				} else if limitedToHiraPrompts {
					actual_prompt_char_type = "hira"
					actual_prompt_char, actual_objective, actual_objective_type = hira_prompting_romaji_objective()
				} else if limitedToRomaPrompts {
					actual_prompt_char_type = "romaji"
					actual_prompt_char, actual_objective, actual_objective_type = roma_prompting_hira_objective()
				} else if limitedToDifficultKata {
					actual_prompt_char_type = "kata"
					actual_prompt_char, actual_objective, actual_objective_type = Difficult_kata_prompting_romaji_objective()
				} else {
					randomize_over_all_decks() // 1-4
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
	// we have a bug in that ::: actual_objective_type  --  is not being set to anything
	/*
		if displayed_prompt_type == "hira" {
				objective = aCard.Hira
			}
	*/
	if actual_prompt_char_type == "hira" {
		// actual_objective = aCard.Hira // ::: these have to be wrong
		actual_objective = aCard.Romaji
	}
	if actual_prompt_char_type == "romaji" { // :::
		actual_objective = aCard.Hira
	}
	if actual_prompt_char_type == "kata" { // todo ::: See. prompt tye could be hira or roma
		actual_objective = aCard.Kata
	}

	if actual_prompt_char_type == "kata" {
		actual_objective = aCard.Kata
	}
	if actual_prompt_char_type == "kata" {
		actual_objective = aCard.Kata
	}
	if actual_prompt_char_type == "kata" {
		actual_objective = aCard.Kata
	}

	return // promptField, objective, actual_objective_type
}

// The 4 functions called above:

// 4 exercise categories :
//   Hira prompting, Romaji objective
//   Kata prompting, Romaji objective
//      Kata prompting, Hira objective, Missing, OK!
//   Romaji prompting, Hira objective

func hira_prompting_romaji_objective() (promptField, objective, actual_objective_type string) {
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsRare[randIndexRare]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsMedium[randIndexMedium]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsOften[randIndexOften]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}

	return
}

func kata_prompting_romaji_objective() (promptField, objective, actual_objective_type string) {
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsRare[randIndexRare]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsMedium[randIndexMedium]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsOften[randIndexOften]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	/*
	   actual_objective ::         aCard.Hira  aCard.Romaji
	   actual_objective_type ::    "hira",  "roma"

	   actual_prompt_char ::       aCard.Hira  aCard.Romaji  aCard.Kata
	   actual_prompt_char_type ::  "hira",  "roma",  "kata"
	*/

	return
}

//   Kata prompting, Hira objective

func roma_prompting_hira_objective() (promptField, objective, actual_objective_type string) {
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsRare[randIndexRare] // Randomly pick a 'card' from a 'deck' and store it in a global var
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "romaji"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsMedium[randIndexMedium] // Note the capital S on fileOfCardsS
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "romaji"
		actual_objective = aCard.Hira
		actual_objective_type = "hira" // ::: fix all these to hira or roma
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsOften[randIndexOften] // Note the capital S on fileOfCardsS
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "romaji"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}

	return
}

//
//
//

func Difficult_kata_prompting_romaji_objective() (promptField, objective, actual_objective_type string) {
	randIndexMK := rand.Intn(len(dataMostDiff)) // dataK.go
	aCard = dataMostDiff[randIndexMK]
	// Kata prompting, Romaji objective:
	actual_prompt_char = aCard.Kata
	actual_prompt_char_type = "kata"
	actual_objective = aCard.Romaji
	actual_objective_type = "roma"
	return
}
