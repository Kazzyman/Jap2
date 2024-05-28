package main

import (
	"fmt"
	"math/rand"
)

var pulledButNotUsedMap = make(map[string]int)

func pick_RandomCard_Assign_fields() { // ::: - -
	// There are 4 different exercises: four different picking sections each with its own prompting.
	// Set (acquire) one of the 4 combinations of actual_prompt_char, actual_objective, actual_objective_type
	// This section only does the initial pick, and MAY be duplicated below in the for loop.
	if limitedToKataPrompts && limitedToHiraPrompts {

		// "Randomize" - or, at least, alternate between kata and hira prompting. // count starts as 1 and is a global var
		for {
			if count == 1 {
				actual_prompt_char_type = "kata"
				kata_prompting_romaji_objective() // 1
				// then fall to count++ and break
			} else {
				actual_prompt_char_type = "hira"
				hira_prompting_romaji_objective() // 2
				count = 1                         // do the other form with the next card
				break
			}
			count++
			break // the next time that a card is fetched we will use the alternate formulation.
		}

	} else if limitedToKataPrompts {
		actual_prompt_char_type = "kata"
		kata_prompting_romaji_objective() // 1
	} else if limitedToHiraPrompts {
		actual_prompt_char_type = "hira"
		hira_prompting_romaji_objective() // 2
	} else if limitedToRomaPrompts {
		actual_prompt_char_type = "roma"
		roma_prompting_hira_objective() // 3
	} else if limitedToDifficultKata {
		actual_prompt_char_type = "kata"
	} else {
		randomize_over_all_decks()
	}
	// Now that we have a newly-acquired actual_prompt_char, check to see if we have it stored in the cyclicArrayPulls slice ...
	// and, if it is so stored, obtain a replacement, and then look again through the entire slice. Repeat the entire
	// process as many times as may be required to finally obtain a value of actual_prompt_char which is novel according to said slice.
	var maxTries int
	for maxTries < 99999 {
		found := false
		for _, lastPull := range cyclicArrayPulls.pulls {
			if lastPull == actual_prompt_char {
				// We also wish to store these duplicates in the map, to keep a tally of such events -- accessible via the rm Dir
				pulledButNotUsedMap[actual_prompt_char]++ // The '++' increments the int value associated with actual_prompt_char
				// fmt.Printf("We've seen the pseudo-random char before; lastPull: %s and actual_prompt_char: %s\n", lastPull, actual_prompt_char)
				found = true
				if limitedToKataPrompts && limitedToHiraPrompts {

					// randomize - alternate

					actual_prompt_char_type = "kata"
					kata_prompting_romaji_objective() // 1

					actual_prompt_char_type = "hira"
					hira_prompting_romaji_objective() // 2

				} else if limitedToKataPrompts {
					actual_prompt_char_type = "kata"
					kata_prompting_romaji_objective() // 1
				} else if limitedToHiraPrompts {
					actual_prompt_char_type = "hira"
					hira_prompting_romaji_objective() // 2
				} else if limitedToRomaPrompts {
					actual_prompt_char_type = "roma"
					roma_prompting_hira_objective() // 3
				} else if limitedToDifficultKata {
					actual_prompt_char_type = "kata"
				} else {
					randomize_over_all_decks()
				}
				break // Exit the inner loop, having a new and potentially novel actual_prompt_char in hand
			}
		}
		maxTries++
		if !found {
			// If actual_prompt_char is not found in the slice ...
			break // All of our work here being done! We hereby exit the outer 99999 for loop
		} else {
			// A match WAS found (and a new actual_prompt_char value has therefore been obtained; so, we need to restart the entire process)
			continue // explicitly continue, i.e., restart range with the replacement value of actual_prompt_char
		}
	}

	// Exited the loops having obtained a newly-acquired, and verified as novel, value of actual_prompt_char, so, ...
	// ... store that newly-acquired actual_prompt_char in an array to be used as a memory of already-seen-and-used chars
	cyclicArrayPulls.InsertKChar(actual_prompt_char) // Only the chars that are actually used are put into the array
	// Also, store that newly-acquired actual_prompt_char in pulledButNotUsedMap
	pulledButNotUsedMap[actual_prompt_char]++ // The '++' increments the int value associated with actual_prompt_char
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

	if actual_prompt_char_type == "hira" {
		actual_objective = aCard.Romaji // true by process of elimination
	}
	if actual_prompt_char_type == "romaji" {
		actual_objective = aCard.Hira // true by process of elimination
	}
	if actual_prompt_char_type == "kata" { // todo ::: objective could be hira or roma
		actual_objective = aCard.Romaji // ::: nailing this down to just roma solves the issue.
	}
}

// The 4 functions called above:

// 4 exercise categories :
//   Hira prompting, Romaji objective
//   Kata prompting, Romaji objective
//      Kata prompting, Hira objective, Missing, OK! ::: more than just OK, it is logically required to be just one of the two.
//   Romaji prompting, Hira objective

func hira_prompting_romaji_objective() { // ::: - -
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
}

func kata_prompting_romaji_objective() { // ::: - -
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
}

// ::: Kata prompting, Hira objective:
// func kata_prompting_hira_objective() // ::: is logically not possible.

func roma_prompting_hira_objective() { // ::: - -
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsRare[randIndexRare] // Randomly pick a 'card' from a 'deck' and store it in a global var
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsMedium[randIndexMedium] // Note the capital S on fileOfCardsS
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsOften[randIndexOften] // Note the capital S on fileOfCardsS
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
}

//

func Difficult_kata_prompting_romaji_objective() { // ::: - -
	randIndexMK := rand.Intn(len(dataMostDiff)) // dataK.go
	aCard = dataMostDiff[randIndexMK]
	// Kata prompting, Romaji objective:
	actual_prompt_char = aCard.Kata
	actual_prompt_char_type = "kata"
	actual_objective = aCard.Romaji
	actual_objective_type = "roma"
}
