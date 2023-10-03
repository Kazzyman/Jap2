package main

// **do-this** menu

import (
	"fmt"
	"os"
	"time"
	// "math/rand"
)

// 'Directive Menu' shared by all exercises; displays only at inception or in response to the 'dir' Directive
func display_Listing_of_Directives_allExercisesHave_inCommon() { // (unique)     - -
	fmt.Println("View source code at https://github.com/Kazzyman/Japanese")
	fmt.Println("    (using US or Alpha-Numeric input mode):")
	fmt.Println("        Enter 'menu' to return to the the main menu ")
	fmt.Println("        Enter 'dir' to redisplay this menu of available directives")
	fmt.Println("        Enter 'notes' for some background on Romaji conventions")
	fmt.Println("        Enter '?' for context-sensitive help ")
	fmt.Println("        Enter '??' for help on a particular Hiragana char")
	fmt.Println("        Enter 'set' to set a new specified prompt & \"key\" ")
	fmt.Println("        Enter 'stat' to view what you have done so far in the current session")
	fmt.Println("        Enter 'reset' to reset (flush or clear) the stats logs")
	fmt.Println("        Enter 'rm' to read_map_of_fineOn")
	// fmt.Println("        Enter 'stack' to prime or stack the frequencyMapOf_IsFineOnChars map")
	fmt.Println("        Enter 'exit' or 'quit' to terminate the app")
	//goland:noinspection ALL
	fmt.Println("\n")
}

func mm_list() {
	fmt.Printf("  Main Menu: \n\n" +
		"  Enter '1' to practice recognizing Romaji: and typing Hiragana (simple, quite useful)\n" +
		"  Enter '2' for recognizing Romaji-Katakana pairs: typing Hiragana (somewhat easy, useful)\n" +
		"  Enter '3' to practice recognizing Katakana, type either Hiragana or Romaji (very versatile)\n" +
		"  Enter '5' to practice typing drill lines\n" +
		"  Enter '6' mixed Hira and Kata prompts: answer with Romaji\n" +
		"  Enter '7' to practice the Most-Difficult kata, type either Hiragana or Romaji\n" +
		"  Enter '8' to practice Sequential Kata, type either Hiragana or Romaji\n" +
		"  Enter '9' to practice Sequential Hira, type either Hiragana or Romaji\n" +
		"  Enter '10' to practice recognizing Difficult Romaji: and typing Difficult Hiragana\n" +
		"  Enter 'exit' to quit\n\n\n")
}

//goland:noinspection ALL
func body_of_Romaji_instructions() { //                                                       - -
	fmt.Println("\n\n\n")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type the Hiragana corresponding to the Romaji prompt\n")
	display_Listing_of_Directives_allExercisesHave_inCommon()
	fmt.Println("Using Hiragana-input-mode, Type the Hiragana corresponding to the Romaji prompt: \n")
}

//goland:noinspection ALL
func body_of_Romaji_plus_Kata_instructions() { //                                             - -
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Exercise 2")
	fmt.Println("Practicing touch-typing (TT) Hiragana in response to Romaji-Katakana prompts:\n")
	fmt.Println("Using Hiragana-input-mode on your sys, Type a Hiragana corresponding to the Romaji-Kata prompt\n")
	display_Listing_of_Directives_allExercisesHave_inCommon() // The func is located at the end of this file
	fmt.Println("Using Hiragana-input-mode, Type a Hiragana corresponding to the Romaji-Katakana prompt: \n")
}

//goland:noinspection ALL
func body_of_KataExerciseInstructions() { //                                                  - -
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Practicing recognizing Katakana chars: Using Hiragana-input-mode on your system ...")
	fmt.Println("... enter the Hiragana chars that correspond to the Katakana.\n")
	fmt.Println("Or, alternatively, type the Romaji that corresponds to the Katakana\n")
	display_Listing_of_Directives_allExercisesHave_inCommon()
	fmt.Println("Type either a Hiragana or Romaji that corresponds to the Katakana prompt: \n")
}

//goland:noinspection ALL
func body_of_instructions_for_Romaji_responces_only() {
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Printf("力いまのレーまじ\n\n" +
		"ひらがなのプロンプトに答える形でローマ字のタッチタイピングの練習をしましょう。\n\n" +
		"ひらがなのプロンプトに対応するローマ字をタイプしてください。\n\n" +
		"(米語入力モードまたはアルファベット/数字入力モードを使用):" +
		"\n\nメインメニューに戻るには「menu」と入力してください。\n\n" +
		"利用可能なディレクティブのメニューを再表示するには「dir」と入力してください。\n\n" +
		"ローマ字の規則に関する背景情報を見るには「notes」と入力してください。 \n\n" +
		"コンテキストセンシティブヘルプの場合は「?」と入力してください。\n\n" +
		"特定のひらがな文字のヘルプの場合は「??」と入力してください。\n\n" +
		"プロンプトと「キー」をリセットするには、「set」と入力してください。\n\n" +
		"現在のセッションでこれまでに行ったことを表示するには、「stat」と入力してください。\n\n" +
		"ヒットログをリセットするには、「reset」と入力してください。\n\n" +
		"アプリを終了するには、「exit」または「quit」と入力してください。\n\n" +
		"\n" +
		"ここでは、以下の点に注意して翻訳しました。\n\n" +
		"「プロンプト」は「prompt」と訳しました。\n" +
		"「キー」は「keys」と訳しました。\n" +
		"「アプリ」は「app」と訳しました。\n" +
		"「?」と「??」は、それぞれ「context-sensitive help」と「help with a specific hiragana character」と訳しました。\n" +
		"「ヒットログ」は「hit log」と訳しました。\n" +
		"また、以下の点についても、注意事項として追加しました。\n\n" +
		"ひらがなを入力する際は、米語入力モードまたはアルファベット/数字入力モードを使用してください。\n" +
		"プロンプトと「キー」をリセットするには、「set」と入力してください。\n" +
		"現在のセッションでこれまでに行ったことを表示するには、「stat」と入力してください。\n" +
		"ヒットログをリセットするには、「reset」と入力してください。\n" +
		"\n\n")
	fmt.Printf("Let's practice typing Romanization by responding to hiragana prompts.\n\n" +
		"ひらがなを「hiragana」と、ローマ字を「Romanization」と訳しました。\n" +
		"Type the Romanization corresponding to the hiragana prompt.\n\n" +
		"(Use US keyboard input mode or alphabet/number input mode):\n\n" +
		"To return to the main menu, type \"menu\".\n\n" +
		"To display the menu of available directives again, type \"dir\".\n\n" +
		"To see background information on Romanization rules, type \"notes\".\n\n" +
		"For context-sensitive help, type \"?\".\n\n" +
		"For help with a specific hiragana character, type \"??\", e.g. \"??あ\".\n\n" +
		"To reset the prompt and \"keys\", type \"set\".\n\n" +
		"To display what you have done so far in the current session, type \"stat\".\n\n" +
		"To reset the hit log, type \"reset\".\n\n" +
		"To exit the app, type \"exit\" or \"quit\".")
	//
	fmt.Println("Exercises 6, 8, and 9")
	fmt.Println("Practicing touch-typing (TT) Romaji in response to mixed prompts:\n")
	fmt.Println("Type the Romaji corresponding to the prompt\n")
	display_Listing_of_Directives_allExercisesHave_inCommon()
	fmt.Println("Type the Romaji corresponding to the prompt: \n")
}

//goland:noinspection ALL
func body_of_Difficult_instructions() { //                                                        - -
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Exercise 7")
	fmt.Println("Practicing recognizing Katakana chars: Using Hiragana-input-mode on your system ...")
	fmt.Println("... enter the Hiragana chars that correspond to the Katakana.\n")
	fmt.Println("Or, alternatively, type the Romaji that corresponds to the Katakana\n")
	display_Listing_of_Directives_allExercisesHave_inCommon() // The func is located at the end of this file
	fmt.Println("Type either the Hiragana or Romaji that corresponds to the Katakana prompt: \n")
}

// This func is executed each time 'menu' is given as a Directive by the user during any Exercise
// Things to do after an Exercise, and before beginning another Exercise
func do_betweenMainMenuSelectionsTTE() {
	currentTime := time.Now()

	fileHandleBig, err := os.OpenFile("Jap2Log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	_, err2 := fmt.Fprintf(fileHandleBig,
		"\nJap2 handled a trans, occured at: %s \n",
		currentTime.Format("15:04:05 on Monday 01-02-2006"))
	check(err2)
	_ = fileHandleBig.Close()

}
