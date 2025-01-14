package main

import (
	"fmt"
	"strings"
)

func recordGuess(kana, users_Guess, Meaning_on_record, second_meaning string) {
	// Check if the kanji exists in the map, not in the CardInfo struct, but in the map: kanjiHitMap
	cardInfo, exists := hiraHitMap[kana] // Use kanji as key to kanjiHitMap map
	// fmt.Printf("cardInfo is:%v, and exists is:%v \n", cardInfo, exists)
	// cardInfo is a struct object of type CardInfo. [kanji] is FROM the signature of this func, and will be a kanji string which is the key to the map
	if !exists {
		// If a card by the name[kanji] doesn't already exist, create a new blank card of type CardInfo
		cardInfo = CardInfo{} // A new blank card which we will append to (or add to) below
	}

	// Append the user's guess [1] and the meaning that we have on record [2] to either the new blank card or the existing card
	// ... oh, and also increment the count [3]
	cardInfo.UsersGuess = append(cardInfo.UsersGuess, users_Guess) // [1]
	if cardInfo.FirstMeaningOnRecord == "" {
		cardInfo.FirstMeaningOnRecord = Meaning_on_record // [2]
		if cardInfo.SecondMeaningOnRecord == "" {
			cardInfo.SecondMeaningOnRecord = second_meaning
		}
	}
	cardInfo.CorrectGuessCount++ // [3]

	// Update the map with the new or updated cardInfo
	hiraHitMap[kana] = cardInfo
}

// LOGGERS:
func log_right_andUpdateGame(prompt_it_was, in string) { // - -
	recordGuess(prompt_it_was, in, aCard.Kata, aCard.Hira)
	if guessLevelCounter < 3 {
		logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(prompt_it_was)
	}
	logHits_in_cyclicArrayHits("Right", prompt_it_was)

	if theGameIsRunning {
		game_loop_counter++
		if game_loop_counter > game_duration_set_by_user {
			the_game_ends()
		}

		if weHadFailed_And_OnlyGotThisRightBecauseOfTheClue {
			// Then that fail has already been logged and we need to skip all logging.
			weHadFailed_And_OnlyGotThisRightBecauseOfTheClue = false
		} else {
			if guessLevelCounter == 2 {
				if gottenHonestly { // todo] do not accumulate if after an "error" or hint
					correctOnFirstAttemptAccumulator++ // ::: 1st
					gottenHonestly = false
				}
			} else if guessLevelCounter == 3 {
				correctOnSecondAttemptAccumulator++ // ::: 2nd
			} else {
				// ... then ... the guessLevelCounter was 4.
				// fmt.Printf("here in log right, guessLevelCounter is:%d, and it should be 4\n", guessLevelCounter)
				correctOnThirdAttemptAccumulator++ // ::: 3rd
			}
			// ::: The other accumulator++  thang : failedOnThirdAttemptAccumulator ]todo[ ... gets handled in log_oops()
		}
	}

	// Only get a fresh card if been here.
	gotLastCardRightSoGetFreshOne = true

}

// 'Reinforce-or-Skip' loggers|Inserters:
func logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(promptToSkip string) { // - -
	frequencyMapOf_IsFineOnChars[promptToSkip]++
}
func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string) { // - -
	frequencyMapOf_need_workOn[promptToWorkMoreOn]++
}

// Universal hits logger|Inserter:
func logHits_in_cyclicArrayHits(RightOrOops, JChar string) { // - -
	cyclicArrayHits.InsertRightOrOops(RightOrOops)
	cyclicArrayHits.InsertChar(JChar)
}

// A special Universal logger|Inserter: so we could drill the user more on chars he has missed
func logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(Jchar string) { // - -
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(Jchar)
}

// Directives:
func hits() { // - -
	// Create maps to store the frequency of each relevant string for that map
	frequencyMapRightOrOops := make(map[string]int)
	frequencyMapChar := make(map[string]int) // These, apparently, create a map to associate a unique string with an int
	frequencyMapWrongs := make(map[string]int)

	//
	// Parse the relevant cyclic array to extract the strings and put them into the relevant map:
	//
	// Load the RightOrOops frequency map
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i]
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}
	// Load the char frequency map // As documented in the foregoing 'for' loop
	for i := 0; i < len(cyclicArrayHits.jchar); i++ {
		str := cyclicArrayHits.jchar[i]
		frequencyMapChar[str]++
	}
	// Load the wrongs frequency map // As documented in the foregoing 'for' loop
	for i := 0; i < len(cyclicArrayOfTheJcharsGottenWrong.jchar); i++ {
		str := cyclicArrayOfTheJcharsGottenWrong.jchar[i]
		frequencyMapWrongs[str]++
	}

	// -- PRINT -- the 'Right' and 'Oops' and their frequencies (Right or Oops) (top of printout)
	for str, freq := range frequencyMapRightOrOops { // The map has only one entry for Right, & one for Oops
		if str == "Right" { // Finds the one potential entry for Right
			fmt.Printf(colorGreen)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorGreen)
			fmt.Printf(" %d\n", freq) // Frequency of Right, per the map
			fmt.Printf(colorReset)
			total_prompts = freq
		} else if str == "Oops" { // Finds the one potential entry for Oops
			fmt.Printf(colorRed)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorRed)
			fmt.Printf(" %d\n", freq) // Frequency of Oops, per the map
			fmt.Printf(colorReset)
			total_prompts = total_prompts + freq
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		}
	}
	// Print the unique Chars and their frequencies (continuing the printout above)
	numberOfUniqueCharsHit := 0
	for str, freq := range frequencyMapChar {
		if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			numberOfUniqueCharsHit++
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorReset)
			fmt.Printf(" %d\n", freq)
		}
	}
	fmt.Printf("Number of unique chars: ")
	fmt.Printf(colorPurple)
	fmt.Printf("%d \n\n", numberOfUniqueCharsHit)
	fmt.Printf(colorReset)

	fmt.Printf("Total prompts:")
	fmt.Printf(colorRed)
	fmt.Printf(" %d\n", total_prompts)
	fmt.Printf(colorReset)

	// Print the ones gotten wrong  (continuing the printout above)
	for str, freq := range frequencyMapWrongs {
		if str == "" {
			// it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			// Print "Gotten Wrong:" + 'str' as multicolored (below)
			fmt.Printf("Gotten Wrong:")            // (in color White)
			fieldsOfStr := strings.Split(str, ":") // Print 'str' as multicolored (below)
			//                              // Gotten Wrong: (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[0]) // KataOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[1]) // it was (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[2]) // RomajiOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[3]) // but you had guessed:_ (in color White) _ is a space char
			fmt.Printf(colorRed)
			fmt.Printf("%s ", fieldsOfStr[4]) // the bad guess_ (in color Red) _ is a space char
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:") // Frequency: (in color Cyan)
			fmt.Printf(colorReset)
			fmt.Printf(" %d \n", freq) // 'a number' (in color White)
		}
	}
	fmt.Println("")
}

func newHits() {
	// Create maps to store the frequency of each relevant string for that map
	frequencyMapRightOrOops := make(map[string]int)
	frequencyMapWrongs := make(map[string]int)

	//
	// Parse the relevant cyclic array to check_for_match_in_other_fields the strings and put them into the relevant map:
	//
	// Load the RightOrOops frequency map
	for i := 0; i < len(cyclicArrayHits.RightOrOops); i++ {
		str := cyclicArrayHits.RightOrOops[i]
		// Apparently; this loads a string into, and increments the frequency of, that particular string, in the map
		frequencyMapRightOrOops[str]++ // Specifically, the '++' must increment the int value associated with str
	}

	// Load the wrongs frequency map // As documented in the foregoing 'for' loop
	for i := 0; i < len(cyclicArrayOfTheJcharsGottenWrong.jchar); i++ {
		str := cyclicArrayOfTheJcharsGottenWrong.jchar[i]
		frequencyMapWrongs[str]++
	}

	// -- PRINT -- 'Right' and its frequency
	for str, freq := range frequencyMapRightOrOops { // The map has only one entry for Right, & one for Oops
		if str == "Right" { // Finds the one potential entry for Right
			fmt.Printf(colorGreen)
			fmt.Printf("\n%s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorGreen)
			fmt.Printf(" %d\n", freq) // Frequency of Right, per the map
			fmt.Printf(colorReset)
			total_prompts = freq
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		}
	}
	//
	//
	// Print the unique Chars and their frequencies (continuing the printout above)
	numberOfUniqueKanaCharsHit := 0
	for hiraString, cardInfoData := range hiraHitMap {
		/*
			if cardInfoData.CorrectGuessCount == 0 {
				// it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
			} else {

		*/

		numberOfUniqueKanaCharsHit++
		fmt.Printf(" %s ", hiraString)
		fmt.Printf(colorGreen)
		fmt.Printf("Your guesses:%s, ", cardInfoData.UsersGuess)
		fmt.Printf(colorYellow)
		fmt.Printf("Meanings: %s, %s ", cardInfoData.FirstMeaningOnRecord, cardInfoData.SecondMeaningOnRecord)
		fmt.Printf(colorCyan)
		fmt.Printf("Freq:")
		fmt.Printf(colorReset)
		fmt.Printf(" %d\n", cardInfoData.CorrectGuessCount)
		// }
	}
	fmt.Printf("Number of unique chars: ")
	fmt.Printf(colorPurple)
	fmt.Printf("%d \n", numberOfUniqueKanaCharsHit)
	fmt.Printf(colorReset)

	fmt.Printf("Total prompts:")
	fmt.Printf(colorRed)
	fmt.Printf(" %d\n", total_prompts)
	fmt.Printf(colorReset)

	// -- PRINT -- 'Oops' and its frequency
	for str, freq := range frequencyMapRightOrOops { // The map has only one entry for Right, & one for Oops
		if str == "Oops" { // Finds the one potential entry for Oops
			fmt.Printf(colorRed)
			fmt.Printf("%s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorRed)
			fmt.Printf(" %d\n", freq) // Frequency of Oops, per the map
			fmt.Printf(colorReset)
			total_prompts = total_prompts + freq
		} else if str == "" {
			// else, it is an 'empty' position in the map due to empty uninitialized position in the cyclic array
		}
	}
	//
	// Print the ones gotten wrong  (continuing the printout above)
	for str, freq := range frequencyMapWrongs {
		if str == "" {
			// it is an 'empty' position in the map due to empty uninitialized positions in the cyclic array
		} else {
			// Print "Gotten Wrong:" + 'str' as multicolored (below)
			fmt.Printf("Gotten Wrong:")            // (in color White)
			fieldsOfStr := strings.Split(str, ":") // Print 'str' as multicolored (below)
			//                              // Gotten Wrong: (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[0]) // KataOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[1]) // it was (in color White)
			fmt.Printf(colorRed)
			fmt.Printf("%s", fieldsOfStr[2]) // RomajiOrWhatEver (in color Red)
			fmt.Printf(colorReset)
			fmt.Printf(" %s: ", fieldsOfStr[3]) // but you had guessed:_ (in color White) _ is a space char
			fmt.Printf(colorRed)
			fmt.Printf("%s ", fieldsOfStr[4]) // the bad guess_ (in color Red) _ is a space char
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:") // Frequency: (in color Cyan)
			fmt.Printf(colorReset)
			fmt.Printf(" %d \n", freq) // 'a number' (in color White)
		}
	}
	fmt.Println("")
}
