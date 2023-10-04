package main

import (
	"fmt"
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
