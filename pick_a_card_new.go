package main

import (
	"math/rand"
)

// The objective of this function is to query the user more on some characters than others
// 4 exercise categories :
//   Hira prompting, Romaji objective
//   Kata prompting, Romaji objective
//   Kata prompting, Hira objective // Missing Kata prompting, Hira objective AS A DIRECTIVE, OK!
//   Romaji prompting, Hira objective
//
// 3 frequency categories :
//   Rarely drilled (e.g., ya, yu, and yo constructs) - 4*1=4
//   Moderate in frequency (chars of average difficulty) - 4*2=8
//   Often drilled (difficult or potentially confusing chars) 4*3=12
//                                                         4+8+12=24

// This one is the default func: it randomly "calls" one of the above four functions and also sets whichDeck (abandoned)
func randomize_over_all_decks() { // ::: - -
	randIndexRare := rand.Intn(len(fileOfCardsRare))
	randIndexMedium := rand.Intn(len(fileOfCardsMedium))
	randIndexOften := rand.Intn(len(fileOfCardsOften))
	randIndexExtendedKata := rand.Intn(len(fileOfCardsE))

	var randomized_selector = 0
	if include_Extended_kata_deck {
		randomized_selector = rand.Intn(25) // If including the extended deck 0-24 = 25
	} else {
		randomized_selector = rand.Intn(24) // 24 random numbers, 0->23
	}

	// ::: Hira prompting, Romaji objective:
	if randomized_selector == 0 || randomized_selector == 13 {
		aCard = fileOfCardsRare[randIndexRare]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 1 || randomized_selector == 2 || randomized_selector == 16 {
		aCard = fileOfCardsMedium[randIndexMedium]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 || randomized_selector == 17 {
		aCard = fileOfCardsOften[randIndexOften]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	//
	// ::: Kata prompting, Romaji objective: (plus an option for including Extended Kata)
	if randomized_selector == 6 { // 13
		aCard = fileOfCardsRare[randIndexRare]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 7 || randomized_selector == 8 || randomized_selector == 14 {
		aCard = fileOfCardsMedium[randIndexMedium]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 9 || randomized_selector == 10 || randomized_selector == 11 || randomized_selector == 15 {
		aCard = fileOfCardsOften[randIndexOften]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	// ::: Optional todo[の] extended Kata deck, via dir: exko
	if randomized_selector == 24 {
		aCard = fileOfCardsE[randIndexExtendedKata]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	/*
		//
		// ::: Kata prompting, Hira objective:
		if randomized_selector == 12 {
			aCard = fileOfCardsRare[randIndexRare]
			actual_prompt_char = aCard.Kata
			actual_prompt_char_type = "kata"
			actual_objective = aCard.Hira
			actual_objective_type = "hira"
		}
		if randomized_selector == 13 || randomized_selector == 14 {
			aCard = fileOfCardsMedium[randIndexMedium]
			actual_prompt_char = aCard.Kata
			actual_prompt_char_type = "kata"
			actual_objective = aCard.Hira
			actual_objective_type = "hira"
		}
		if randomized_selector == 15 || randomized_selector == 16 || randomized_selector == 17 {
			aCard = fileOfCardsOften[randIndexOften]
			actual_prompt_char = aCard.Kata
			actual_prompt_char_type = "kata"
			actual_objective = aCard.Hira
			actual_objective_type = "hira"
		}

	*/
	//
	// ::: Romaji prompting, Hira objective:
	if randomized_selector == 18 || randomized_selector == 12 {
		aCard = fileOfCardsRare[randIndexRare]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 19 || randomized_selector == 20 {
		aCard = fileOfCardsMedium[randIndexMedium]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 21 || randomized_selector == 22 || randomized_selector == 23 {
		aCard = fileOfCardsOften[randIndexOften]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	return
}
