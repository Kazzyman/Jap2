package main

// ::: var fileOfCardsKanjiHard = []charSetStruct{ // no specified length makes it a slice vs ...
var fileOfCardsKanjiHard = [2]charSetStruct{ // array (a array is of fixed size, a slice is more flexible).

	// ::: cards where kata = hira; moreover, neither is as implied, rather both are Kanji

	//
	{"歌手", "歌手", "singer",
		"vocal entertainer",
		"vocal entertainer",
		"vocal entertainer",
		"vocal entertainer"},

	{"教師", "教師", "teacherAcademic",
		"T A",
		"T A",
		"T A",
		"T A"},
	//
	/*
		// cards where kata and hira fields differ (each contains what it is supposed to contain)
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
