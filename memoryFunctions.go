package main

import (
	"fmt"
	"os"
)

// Parse the maps to see if the new random card matches any members of the map, and return new card fields if warranted
func checkMemory(promptField, objective, objective_kind string) (new_promptField, new_objective, new_objective_kind string) {
	new_promptField, new_objective, new_objective_kind = check_it_for_fine_on(promptField, objective, objective_kind)
	checkOops_map(new_promptField, new_objective, new_objective_kind) // Potentially calls: pick_RandomCard_Assign_fields() via step 2 of ...
	// ... of check_it_for_needing_more_practice [will that work ?????????]

	// this made no difference because, more often than not, check_it_for_fine_on() returns copies of what it was fed
	// But, now I am trying adding back the original version that requires objective_kind as an argument
	// checkOops_map(promptField, objective, objective_kind) // Potentially calls: pick_RandomCard_Assign_fields() via step 2 of ...
	// ... of check_it_for_needing_more_practice [will that work ?????????]
	return new_promptField, new_objective, new_objective_kind
}

/*
Parse the fine_on map to see if the new random card has already been done correctly 3 or more times
if it has, then pick another random card, and recursively check IT
*/
func check_it_for_fine_on(promptField, objective, objective_kind string) (new_promptField, new_objective, new_objective_kind string) { // - -
	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // Append to file
	check(err)
	for s, f := range frequencyMapOf_IsFineOnChars { // Parse the map, s is the string, f is the frequency
		if s == objective { // If it is in the map, check the frequency
			if f >= 2 { // if the freq is 3+ we need another random card because objective has already been mastered
				// read_map_of_fineOn() // we show the map
				// read_map_of_needWorkOn() // We show the map
				// fmt.Printf("\n You were correct on: %s twice or more ... \n", aCard.KeyR) // Needs fixing: "aCard.KeyR"
				// Log to a file that this action was taken ...
				_, err1 := fmt.Fprintf(fileHandle, " "+
					"check_it_for_fine_on: o_Objctv:%s ws_fnd in MapOf_FinOn, s:%s freq:%d o_pf:%s o_obKnd:%s\n",
					objective, s, f, promptField, objective_kind)
				check(err1)
				//
				// Get a new random card ***********
				new_promptField, new_objective, new_objective_kind = pick_RandomCard_Assign_fields()
				// fmt.Println(" ... so here is a new one ... \n")
				_, err2 := fmt.Fprintf(fileHandle, " "+
					"check_it_for_fine_on: has delt a new card, new_Objctv:%s new_pFld:%s new_obKnd:%s \n",
					new_objective, new_promptField, new_objective_kind)
				check(err2)
				// check_it_for_fine_on(promptField, objective, objective_kind) // Check THAT new random card with a recursive call
				break //
			} else { // else the card had a freq of less than 2, so return the original fields
				new_promptField = promptField
				new_objective = objective
				new_objective_kind = objective_kind
				break //  exit the loop & func -- continue in using the original card
			}
		} else { // else the latest card which was randomly selected was not found YET in the map ...
			continue // ... so we continue the loop to finish reading the map
		}
	}
	if new_promptField == "" {
		new_promptField = promptField
		new_objective = objective
		new_objective_kind = objective_kind
	}
	_ = fileHandle.Close()
	return new_promptField, new_objective, new_objective_kind
}

// Called only from checkOops_map()
func check_it_for_needing_more_practice(promptField, objective, objective_kind string) { //       - -
	var skip_this_step = false
	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
	check(err)
	for s, f := range frequencyMapOf_need_workOn {
		if s == objective { // Check if the latest random card is in the need_workOn map, and check the freq ...
			if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we keep it
				// read_map_of_fineOn() // Show the contents of the map
				// read_map_of_needWorkOn() // Show the contents of the map
				fmt.Printf("\n The Random card: %s was missed once or more \n", promptField)
				fmt.Println("... so we will keep it and quiz you on it ... ")
				// Log to a file that this action was taken
				_, err2 := fmt.Fprintf(fileHandle, " "+
					"checkOopsmap/cifnmp: Step 1 of 2, fndStr:%s in need_wrkOnMap, freq:%d : keeping card\n",
					s, f)
				check(err2)
				skip_this_step = true
				break //  ... we exit the loop and the func -- we will keep and use this random card, and skip the final section
			} else { // else the card had a freq less than 2, so ...
				continue // keep looking through the map for another instance that may in there, with a significant freq
			}
		}
	}
	_ = fileHandle.Close()

	// so, if we have set skip_this_step to true ...
	// ... we should NOT be keeping the randomly-selected card??
	if skip_this_step == false {
		fileHandle2, err2 := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // Append to file
		check(err2)
		// The latest random was not in the map at sig freq, or not in the map at all ??
		// In any case, it is time to serve-up something new and potentially more difficult ... so:
		for s, f := range frequencyMapOf_need_workOn {
			if s == objective { // Check if the latest random is in the map, and check the freq ...
				if f >= 1 { // ... if the freq is 1+ we definitely need more work on this particular card, so we set it as aCard
					// read_map_of_fineOn() // we show the map
					// read_map_of_needWorkOn() // We show the map
					fmt.Printf("\n This Random card: %s was missed 1 or more times \n", objective)
					fmt.Println("... so we will test you on it, since it has been a while")
					// Log to a file that this action was taken **do-this**
					fileHandle, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
					check(err)                                                                              // ... gets a file handle to JapLog.txt
					_, err2 := fmt.Fprintf(fileHandle, "\nPart2of2 found:%s in frequencyMapOf_need_workOn, freq:%d \n",
						s, f)
					check(err2)
					_ = fileHandle.Close()
					practice_this_card(objective, objective_kind) // locate and assign aCard // set it as new aCard
					break                                         //  ... we exit the loop and the func -- we will keep and use this random card
				} else { // else the card had a freq less than 2, so ...
					continue // keep looking through the map for another instance that may in there, with a significant freq
				}
			}
		}
		_ = fileHandle2.Close()
	}
}

// Only used in check_it_for_needing_more_practice()
func practice_this_card(objective, objective_kind string) { // - -
	if whichDeck == 1 {
		if objective_kind == "Hira" {
			for _, card := range fileOfCards {
				if card.Hira == objective {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Hira objective card not found in: fileOfCards")
			}
		} else if objective_kind == "Romaji" {
			for _, card := range fileOfCards {
				if card.Romaji == objective {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Romaji objective card not found in: fileOfCards")
			}
		}
	} else if whichDeck == 2 {
		if objective_kind == "Hira" {
			for _, card := range fileOfCardsS {
				if card.Hira == objective {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Hira objective card not found in: fileOfCardsS")
			}
		} else if objective_kind == "Romaji" {
			for _, card := range fileOfCardsS {
				if card.Romaji == objective {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Romaji objective card not found in: fileOfCardsS")
			}
		}
	} else if whichDeck == 3 {
		if objective_kind == "Hira" {
			for _, card := range fileOfCardsMostDifficult {
				if card.Hira == objective {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Hira objective card not found in: fileOfCardsMostDifficult")
			}
		} else if objective_kind == "Romaji" {
			for _, card := range fileOfCardsMostDifficult {
				if card.Romaji == objective {
					foundElement = &card  // foundElement is a global var and contains all the fields of a card
					aCard = *foundElement // aCard is also a global var
					break
				}
			}
			if foundElement == nil {
				fmt.Println("Romaji objective card not found in: fileOfCardsMostDifficult")
			}
		}
	}
}

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

// Global Maps:
var frequencyMapOf_IsFineOnChars = make(map[string]int) // create the map, dir 'read_map_of_fineOn' reads it - -
var frequencyMapOf_need_workOn = make(map[string]int)   // - -

func checkOops_map(promptField, objective, objective_kind string) { // - -
	fileHandle, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // Append to file
	check(err)
	isThis_a_cardWeNeedMoreWorkOn++
	// If we have gone a while without augmenting the random picks with cards we have previously missed ...
	if isThis_a_cardWeNeedMoreWorkOn > 2 {
		// Log to a file that this action was taken, and ...
		_, err2 := fmt.Fprintf(fileHandle, " "+
			"checkOops_map A: card:%s in Map_workOn after:%d cycles  passing to check_it_for_needing_more_practice \n",
			promptField, isThis_a_cardWeNeedMoreWorkOn)
		check(err2)
		// Then:
		isThis_a_cardWeNeedMoreWorkOn = 0
		check_it_for_needing_more_practice(promptField, objective, objective_kind)
	} else {
		// Retain and use the randomly picked card
		// Log to a file that this action was taken
		_, err3 := fmt.Fprintf(fileHandle, " "+
			"checkOops_map B: card:%s in frequencyMapOf_need_workOn after:%d cycles  retaining random card \n",
			promptField, isThis_a_cardWeNeedMoreWorkOn)
		check(err3)
	}
	_ = fileHandle.Close() // Close the file, discarding any resulting errors returned
}
