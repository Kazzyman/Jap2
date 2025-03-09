package main

// ::: var fileOfCardsKanjiHard = []charSetStruct{ // no specified length makes it a slice vs ...
var fileOfCardsKanjiHard = [2]charSetStruct{ // array (a array is of fixed size, a slice is more flexible).

	// ::: cards where kata = hira

	//
	{"歌手", "歌手", "singer",
		"singer, 歌手 , hiraMissing",
		"KataHint",
		"TT_Hint",
		"SansR_Hint"},

	{"教師", "教師", "teacherAcademic",
		"teacherAcademic, 教師 , きょうし",
		"KataHint",
		"TT_Hint",
		"SansR_Hint"},
	//
	/*
		// cards where kata and hira fields differ
			{"僕", "ぼく", "male-I",
				"僕",
				"male-I",
				"僕は男です (I am a male:ぼく)",
				"I, myself (boku, mainly used by males)"},

			{"俺", "おれ", "I-male-informal",
				"I-male-informal",
				"俺:おれ",
				"俺は男です (I am a male[between friends])",
				"I, myself (mainly used by males [informal/casual])"},

	*/
}
