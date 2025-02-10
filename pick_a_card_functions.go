package main

import (
	"fmt"
	"math/rand"
)

/* ::: these are the options available to the user:
todo, since these only have to do with prompting, we will eventually, want to add two options to limit objective type to romaji or hira
as it is, objective type is always "randomized" for the ko cases of hko
	'honly' Use only Hira prompting ::: 1
	'konly' Use only Kata prompting, could be either hira or romaji objective ::: 2, or 3
	'hko' Use BOTH Hira & Kata prompting (prompt with either, but no romaji prompting) ::: 1, 2, or 3
	'ronly' Use only Romaji prompting ::: 4
	'donly' Difficult descriptive prompting (romaji objective) ::: 5
*/
// ::: The sixth "option" is the default of randomizing over all decks except the difficult one (6)

/*
There are 5 exercises.
That is, there are 4 classes plus one special case.

The four classes are:
1. Hira prompting that requires a Romaji response.

2. Kata prompting that requires a Romaji response.
3. Kata prompting that requires a Hira response.

4. Romaji prompting that requires a Hira response.

the one special case is:
5. Difficult-descriptive-only prompting, requiring a Romaji response.

So, herein we will determine which of those 5 is in progress; and deploy the correct function.
*/

// This global map is used in the next function and also in func read_pulledButNotUsedMap() {
var pulledButNotUsedMap = make(map[string]int) // Global map, this, declared and made right hear.

var hiraOrRomajiObjective = 1 // just like count
// var kataPromptsObjective = "oneOfTwoPotentialOptions" // hira, or romaji

// this func handles the cases in which limits have been selected by the user
func pick_RandomCard_Assign_fields() { // ::: - -

	// There are 4 different classes of exercises: four different picking sections each with its own prompting.
	// Set (acquire) one of the potential combinations of actual_prompt_char, actual_objective, actual_objective_type
	// This section only does the initial pick, and MAY be duplicated below in the next (larger) for loop.
	if limitedToKataPrompts && limitedToHiraPrompts {
		// true on both means potentially use either
		// "Randomize" - or, at least, alternate between kata and hira prompting. // count starts as 1 and is a global var
		for {
			if count == 1 {
				actual_prompt_char_type = "kata"

				/*
					for {
						if hiraOrRomajiObjective == 1 {
							kata_prompting_romaji_objective() // ::: 2
						} else {
							// ::: todo: if prompting with kata we should "randomly" ask for either a hira or a romaji objective
							kata_prompting_hira_objective() // ::: 3
							hiraOrRomajiObjective = 1
							break
						}
						hiraOrRomajiObjective++
						break
					}
				*/

				kata_prompting_romaji_objective()

				// then fall to count++ and break
			} else {
				actual_prompt_char_type = "hira"
				hira_prompting_romaji_objective() // ::: 1
				count = 1                         // do the other form with the next card
				break
			}
			count++
			break // the next time that a card is fetched we will use the alternate formulation.
		}
	} else if limitedToDifficultDescriptions {
		actual_prompt_char_type = "kata"                   // ::: because the prompt description is in the kata field
		Difficult_descriptive_prompting_romaji_objective() // ::: 5
	} else if limitedToSpelling {
		Spelling_prompting_romaji_objective() // ::: 8
	} else if limitedToKataPrompts {
		actual_prompt_char_type = "kata"
		if limitedToKataPromptsAndSimplexHiraObj {
			actual_objective_type = "hira"
			kata_prompting_Simplex_hira_objective()
		} else {
			// ::: if prompting with kata we should "randomly" ask for either a hira or a romaji objective
			for {
				if hiraOrRomajiObjective == 1 {
					kata_prompting_romaji_objective() // ::: 2
				} else {
					kata_prompting_hira_objective() // ::: 3
					hiraOrRomajiObjective = 1
					break
				}
				hiraOrRomajiObjective++
				break
			}
		}

	} else if limitedToHiraPrompts {
		actual_prompt_char_type = "hira"
		hira_prompting_romaji_objective() // ::: 1
	} else if limitedToRomaPrompts {
		if limitedToRomaPromptsAndSimplexHiraObj {
			actual_prompt_char_type = "roma"
			roma_prompting_w_simplex_hira_objective() // ::: 7
		} else {
			actual_prompt_char_type = "roma"
			roma_prompting_hira_objective() // ::: 4
		}

	} else if kata_roma {
		kata_prompting_romaji_objective() // ::: 2
	} else if kata_hira {
		kata_prompting_hira_objective() // ::: 3
	} else {
		// if no limits have been set via a directive, prompt as standard mix (prompts will be hira, roma, or kata)
		randomize_over_all_decks() // ::: (6) in the other source code file
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

				//
				// copied from above?
				if limitedToKataPrompts && limitedToHiraPrompts {
					// true on both means potentially use either
					// "Randomize" - or, at least, alternate between kata and hira prompting. // count starts as 1 and is a global var
					for {
						if count == 1 {
							actual_prompt_char_type = "kata"

							/*
								for {
									if hiraOrRomajiObjective == 1 {
										kata_prompting_romaji_objective() // ::: 2
									} else {
										// ::: todo: if prompting with kata we should "randomly" ask for either a hira or a romaji objective
										kata_prompting_hira_objective() // ::: 3
										hiraOrRomajiObjective = 1
										break
									}
									hiraOrRomajiObjective++
									break
								}
							*/

							kata_prompting_romaji_objective()
							// then fall to count++ and break
						} else {
							actual_prompt_char_type = "hira"
							hira_prompting_romaji_objective() // ::: 1
							count = 1                         // do the other form with the next card
							break
						}
						count++
						break // the next time that a card is fetched we will use the alternate formulation.
					}
				} else if limitedToDifficultDescriptions {
					actual_prompt_char_type = "kata"                   // ::: because the prompt description is in the kata field
					Difficult_descriptive_prompting_romaji_objective() // ::: 5

				} else if limitedToSpelling {
					actual_prompt_char_type = "kata"      // ::: because the prompt description is in the kata field
					Spelling_prompting_romaji_objective() // :::

				} else if limitedToKataPrompts {
					// todo add simplex
					if limitedToKataPromptsAndSimplexHiraObj {
						kata_prompting_Simplex_hira_objective()
					} else {
						actual_prompt_char_type = "kata"
						// ::: if prompting with kata we should "randomly" ask for either a hira or a romaji objective
						for {
							if hiraOrRomajiObjective == 1 {
								kata_prompting_romaji_objective() // ::: 2
							} else {
								kata_prompting_hira_objective() // ::: 3
								hiraOrRomajiObjective = 1
								break
							}
							hiraOrRomajiObjective++
							break
						}
					}
				} else if limitedToHiraPrompts {
					actual_prompt_char_type = "hira"
					hira_prompting_romaji_objective() // ::: 1
				} else if limitedToRomaPrompts {
					if limitedToRomaPromptsAndSimplexHiraObj { // ::: ???
						roma_prompting_w_simplex_hira_objective()
					} else {
						actual_prompt_char_type = "roma"
						roma_prompting_hira_objective() // ::: 4
					}

				} else if kata_roma {
					kata_prompting_romaji_objective() // ::: 2
				} else if kata_hira {
					kata_prompting_hira_objective() // ::: 3
				} else if limitedToRomaPromptsAndSimplexHiraObj { // ::: ???
					roma_prompting_w_simplex_hira_objective() // ::: 7
				} else {
					// if no limits have been set via a directive, prompt as standard mix (prompts will be hira, roma, or kata)
					randomize_over_all_decks() // ::: (6)
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
		pulledButNotUsedMap = make(map[string]int)
		// fmt.Println(colorReset + "You have finished all the cards; repeats will now ensue ... \n")
		fmt.Printf("%sYou have finished all the cards in your chosen deck; repeats will now ensue ... \n", colorReset)
	}
}

/*
= ::: ------------ END OF FUNCTION ------------------------------------------------------------------------------------
=
=
*/

// ::: 1
func hira_prompting_romaji_objective() { // ::: - -
	randIndexRare := rand.Intn(len(fileOfCardsKanjiHard))
	randIndexMedium := rand.Intn(len(fileOfCardsHiraKata))
	randIndexOften := rand.Intn(len(fileOfCardsEasyKanji))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsKanjiHard[randIndexRare]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
}

// ::: 2
func kata_prompting_romaji_objective() { // ::: - -
	randIndexRare := rand.Intn(len(fileOfCardsKanjiHard))
	randIndexMedium := rand.Intn(len(fileOfCardsHiraKata))
	randIndexOften := rand.Intn(len(fileOfCardsEasyKanji))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsKanjiHard[randIndexRare]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
}

// ::: 3
func kata_prompting_hira_objective() { // ::: - -
	randIndexRare := rand.Intn(len(fileOfCardsKanjiHard))
	randIndexMedium := rand.Intn(len(fileOfCardsHiraKata))
	randIndexOften := rand.Intn(len(fileOfCardsEasyKanji))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsKanjiHard[randIndexRare]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
}

// ::: 32
func kata_prompting_Simplex_hira_objective() { // ::: - -
	randIndex := rand.Intn(len(fileOfCards_nonCompound))
	aCard = fileOfCards_nonCompound[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	actual_prompt_char = aCard.Kata
	actual_prompt_char_type = "kata"
	actual_objective = aCard.Hira
	actual_objective_type = "hira"
}

// ::: 4 - limited to Roma prompting (solicits complex Hira constructs)
func roma_prompting_hira_objective() { // ::: - -
	randIndexRare := rand.Intn(len(fileOfCardsKanjiHard))
	randIndexMedium := rand.Intn(len(fileOfCardsHiraKata))
	randIndexOften := rand.Intn(len(fileOfCardsEasyKanji))

	randomized_selector := rand.Intn(6) // 6 random numbers, 0->5

	if randomized_selector == 0 {
		aCard = fileOfCardsKanjiHard[randIndexRare] // Randomly pick a 'card' from a 'deck' and store it in a global var
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 1 || randomized_selector == 2 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
}

// ::: 5
func Difficult_descriptive_prompting_romaji_objective() { // ::: - -
	randIndexMK := rand.Intn(len(dataMostDiff)) //
	aCard = dataMostDiff[randIndexMK]
	// prompting from a description in the kata field, Romaji objective:
	actual_prompt_char = aCard.Kata  // ::: because the description being used as prompt is stored in the kata field.
	actual_prompt_char_type = "kata" // ::: because the description being used as prompt is stored in the kata field.
	actual_objective = aCard.Romaji
	actual_objective_type = "roma"
}

// ::: 8
func Spelling_prompting_romaji_objective() {
	randIndexMK := rand.Intn(len(dataSpelling)) //
	aCard = dataSpelling[randIndexMK]
	// prompting from a description in the kata field, Romaji objective:
	actual_prompt_char = aCard.Kata  // ::: because the description being used as prompt is stored in the kata field.
	actual_prompt_char_type = "kata" // ::: because the description being used as prompt is stored in the kata field.
	actual_objective = aCard.Romaji
	actual_objective_type = "roma"
}

// ::: 7 - limited to Roma prompting and simplex Hira chars
func roma_prompting_w_simplex_hira_objective() { // ::: - -
	randIndex := rand.Intn(len(fileOfCards_nonCompound))
	aCard = fileOfCards_nonCompound[randIndex] // Randomly pick a 'card' from a 'deck' and store it in a global var
	actual_prompt_char = aCard.Romaji
	actual_prompt_char_type = "roma"
	actual_objective = aCard.Hira
	actual_objective_type = "hira"
}
