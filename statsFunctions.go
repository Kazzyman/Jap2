package main

// **do-this**
import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func if_it_is_a_Directive(in string) (isAlphanumeric bool) {
	findAlphasIn := regexp.MustCompile(`[a-zA-Z]`)
	switch true {
	case findAlphasIn.MatchString(in): // <-- 'in'
		isAlphanumeric = true
	default:
		isAlphanumeric = false
	}
	return isAlphanumeric
}

// LOGGERS:
func log_right(prompt_it_was string) {
	logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(prompt_it_was)
	logHits_in_cyclicArrayHits("Right", prompt_it_was)
}
func log_oops(prompt_it_was, field_it_was, guess string) {
	logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(prompt_it_was)
	logHits_in_cyclicArrayHits("Oops", prompt_it_was)
	logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(prompt_it_was +
		":it was:" + field_it_was + ":but you had guessed:" + guess)
}

//
// 'Reinforce-or-Skip' loggers|Inserters:
func logSkipThisPrompt_inThe_frequencyMapOf_IsFineOnChars(promptToSkip string) {
	frequencyMapOf_IsFineOnChars[promptToSkip]++
}
func logReinforceThisPrompt_inThe_frequencyMapOf_need_workOn(promptToWorkMoreOn string) {
	frequencyMapOf_need_workOn[promptToWorkMoreOn]++
}

// Universal hits logger|Inserter:
// Used in exercises 1, 2, 3, '4', 6, 7
func logHits_in_cyclicArrayHits(RightOrOops, JChar string) {
	cyclicArrayHits.InsertRightOrOops(RightOrOops)
	cyclicArrayHits.InsertChar(JChar)
}

//
// A special Universal logger|Inserter: so we can drill the user more on chars he has missed
//
// Used in exercises 1, 2, 3, '4', 6, 7
func logJcharsGottenWrong_in_cyclicArrayOfTheJcharsGottenWrong(Jchar string) {
	cyclicArrayOfTheJcharsGottenWrong.InsertCharsWrong(Jchar)
}

// Directives:
func hits() {
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
	for str, freq := range frequencyMapRightOrOops {
		if str == "Right" {
			fmt.Printf(colorGreen)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorGreen)
			fmt.Printf(" %d\n", freq)
			fmt.Printf(colorReset)
		} else if str == "Oops" { // it is Oops
			fmt.Printf(colorRed)
			fmt.Printf(" %s ", str)
			fmt.Printf(colorCyan)
			fmt.Printf("Frequency:")
			fmt.Printf(colorRed)
			fmt.Printf(" %d\n", freq)
			fmt.Printf(colorReset)
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

func log_to_JapLog_file_inception_time(selectedExercise string) {
	currentTime := time.Now()
	if selectedExercise == "Romaji_Prompt" { // 1
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 1 'Romaji_Prompt' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExercise == "Romaji_w_Kata_Prompt" { // 2
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 2 'Romaji+Kata_Prompt' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExercise == "Respond_w_Hira_or_Romaji_to_kataPrompt_3" { // 3, '4'
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 3 or 4 'Kata_Prompt-Respond-w-Hira|Romaji' occured at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExercise == "drillLines" { // 5
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 5 'Drill Lines' occured at: %s \n",
			currentTime.Format("01-02-2006 15:04:05 Monday"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExercise == "Mixed_prompts" { // 6
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 6 'Hira_prompt' occurred at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExercise == "Most_Difficult" { // 7
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 7 'Most_Difficult' occurred at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
		//
		//
	} else if selectedExercise == "Sequential_Kata" { // 8
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 8 'Sequential_Kata' occurred at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	} else if selectedExercise == "Sequential_Hira" { // 9
		fileHandleBig, err := os.OpenFile("JapLog.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600) // append to file
		check(err)                                                                                 // ... gets a file handle to JapLog.txt
		_, err2 := fmt.Fprintf(fileHandleBig, "\nInception of e xercise 9 'Sequential_Hira' occurred at: %s \n",
			currentTime.Format("15:04:05 on Monday 01-02-2006"))
		check(err2)
		_ = fileHandleBig.Close()
	}
}
