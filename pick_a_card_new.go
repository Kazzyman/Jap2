package main

import (
	"math/rand"
)

// The objective of this function is to randomly query the user, but more on some characters than others
// 4 exercises: todo plus two optional ones
//   Hira prompting, Romaji objective ::: 1

//   Kata prompting, Romaji objective ::: 2, todo plus optional extended kata
//   Kata prompting, Hira objective ::: 3, todo plus optional extended kata

//   Romaji prompting, Hira objective ::: 4
//
// 3 frequency categories (or decks) : rare, medium, often ::: 4 sets meaning 12 decks in all (plus 2) , so 14 instances
//
//    1. fileOfCardsKanjiHard, 2. fileOfCardsHiraKata, and 3. fileOfCardsEasyKanji
//   Rarely drilled (e.g., ya, yu, and yo constructs)
//   Moderate in frequency (chars of average difficulty)
//   Often drilled (difficult or potentially confusing chars)

// This one is the default func: it randomly "calls" one of the above four functions and also sets whichDeck (abandoned)
func randomize_over_all_decks() { // ::: 6 - -
	randIndexRare := rand.Intn(len(fileOfCardsKanjiHard))
	randIndexMedium := rand.Intn(len(fileOfCardsHiraKata))
	randIndexOften := rand.Intn(len(fileOfCardsEasyKanji))
	randIndexExtendedKata := rand.Intn(len(fileOfCardsE))

	var randomized_selector = 0
	if include_Extended_kata_deck {
		randomized_selector = rand.Intn(32) // If including the extended deck 0-31 = 32
	} else {
		randomized_selector = rand.Intn(30) // 30 random numbers, 0->29
	}

	// 2,3,4 = 9
	// ::: Hira prompting, Romaji objective: 1
	if randomized_selector == 0 || randomized_selector == 13 {
		aCard = fileOfCardsKanjiHard[randIndexRare]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 1 || randomized_selector == 2 || randomized_selector == 16 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 3 || randomized_selector == 4 || randomized_selector == 5 || randomized_selector == 17 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Hira
		actual_prompt_char_type = "hira"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}

	// 1,3,4,1 = 9 = 18
	//
	// ::: Kata prompting, Romaji objective: (plus an option for including Extended Kata): 2
	if randomized_selector == 6 { // 13
		aCard = fileOfCardsKanjiHard[randIndexRare]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 7 || randomized_selector == 8 || randomized_selector == 14 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	if randomized_selector == 9 || randomized_selector == 10 || randomized_selector == 11 || randomized_selector == 15 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}
	// ::: Optional todo[の] extended Kata deck, via dir: exko
	if randomized_selector == 30 {
		aCard = fileOfCardsE[randIndexExtendedKata]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}

	// 1,2,3,1 = 7 = 25
	//
	// ::: Kata prompting, Hira objective: 3
	if randomized_selector == 24 {
		aCard = fileOfCardsKanjiHard[randIndexRare]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 25 || randomized_selector == 26 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 27 || randomized_selector == 28 || randomized_selector == 29 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	// ::: Optional todo[の] extended Kata deck, via dir: exko
	if randomized_selector == 31 {
		aCard = fileOfCardsE[randIndexExtendedKata]
		actual_prompt_char = aCard.Kata
		actual_prompt_char_type = "kata"
		actual_objective = aCard.Romaji
		actual_objective_type = "roma"
	}

	// 2,2,3 = 7 = 32
	//
	// ::: Romaji prompting, Hira objective: 4
	if randomized_selector == 18 || randomized_selector == 12 {
		aCard = fileOfCardsKanjiHard[randIndexRare]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 19 || randomized_selector == 20 {
		aCard = fileOfCardsHiraKata[randIndexMedium]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	if randomized_selector == 21 || randomized_selector == 22 || randomized_selector == 23 {
		aCard = fileOfCardsEasyKanji[randIndexOften]
		actual_prompt_char = aCard.Romaji
		actual_prompt_char_type = "roma"
		actual_objective = aCard.Hira
		actual_objective_type = "hira"
	}
	return
}
