package main

// This is for a future project.
// This dataset was intended as an alternate to the set in data.go used in kata_prompting_romaji_objective()

var fileOfCardsKataPromptsOnly = []charSetStruct{
	{"ヌ", "ぬ", "nu",
		" nu:ぬ:ヌ   Compare:  me:メ  &  na:ナ ",
		" nu flew ヌ -- till it hit a ceiling ",
		" TT: R-ring<--- to the '1' char, at least two lines cross, nu flew ヌ -- hit a ceiling ",
		" ?:ぬ:ヌ   ?-flew ヌ -- till it hit a ceiling  TT: R-ring<--- to the '1' char"},

	{"ネ", "ね", "ne",
		" ne:ね:ネ  Kata looks like a hoe that got nailed ネ ",
		" something got nailed-down by that spike at the left ",
		" TT: R-ring<v to the ',<' chars  ...  Yea, to me it is a hoe:ほ that got nailed ネ ",
		" ?:ね:ネ  Kata looks like a hoe that got nailed ネ  TT: R-ring<v to the ',<' chars"},

	{"メ", "め", "me",
		" me:メ:め   Compare: nu:ヌ",
		" me",
		" TT: right-pinky-slide down to the '?' char",
		" メ:め messy-lesser ver of nuddle:ぬ　 TT: R-pinky-slide down to the '?' char"},

	{"リ", "り", "ri",
		" ri:リ:り  actually, very very similar, actually!",
		" ri:リ:り Longer on the Right is Ri",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},

	{"イ", "い", "i",
		" i:イ:い  the Kata イ looks more like a hiragana te:て, but still two mostly-vertical lines: イ",
		" i:い:イ  for the Kata ... maybe shift the two lines of the hiragana い to forms an 'I'  i:イ",
		" TT: left-middle< to the 'E' char  い:イ  Longer on the left is i, Compare: ri:リ:り",
		" い:イ at least it is still two mostly-vertical lines   TT: L-middle< to the 'E' char"},

	{"ア", "あ", "a",
		" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
		" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
		" TT: left-ring-> to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
		" Compare:  me:め  あ  nu:ぬ  ne:ね   あ-Fuck mae-te!  TT: L-middle<- to the '3' char"},

	{"ワ", "わ", "wa",
		" wa:ワ:わ  Water-fall, pissing in the wind (making わしこ) ",
		" wa",
		" TT: right-ring to the '0' char",
		" ワ:わ  aqua-fall   TT: right-ring to the '0' char"},

	{"ヲ", "を", "wo",
		" wo:ヲ:を   Compare: fu:フ　wa:ワ　u:ウ　a:ア　te:テ:て　 ",
		" wo:ヲ, at least it looks something like wa:ワ, though shifted",
		" TT: right-ring up to the '0' char (shifted)",
		" ヲ:を   Compare: wa:ワ   TT: right-ring up to the '0' char (shifted)"},

	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},

	{"モ", "も", "mo",
		" mo:モ:も ",
		" mo",
		" TT: right-middle<v to the 'M' char",
		" モ:も   TT: right-middle<v to the 'M' char,  m is for ?"},

	{"フ", "ふ", "fu",
		" fu:ふ:フ  ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" fu: think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" ふ:フ  TT: left-ring to the '2' char,  it is the big mountain in Japan "},

	{"ナ", "な", "na",
		" na:な:ナ  Kata ナ is just like the left side of the Hira な -- な:ナ ",
		" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  ( め ぬ )",
		" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
		" な:ナ  TT: index< to the 'U' char, Hiragana t-ies a knot, see the two chars thar "},

	{"ツ", "つ", "tsu",
		" tsu:つ:ツ  see water crashing on the she shore",
		" tsu",
		" TT: left-pinky on 'Z' char",
		" ?:つ:ツ see water crashing on the she shore  TT: left-pinky on 'Z' char"},

	{"テ", "て", "te",
		" te:て:テ  Looks like a T",
		" te:て:テ ,  Compare:  chi:チ & te:テ:て ",
		" TT: left-ring< to the 'W' char, te:テ:て  went the wrong way, and gained a flat hat ",
		" ?:て:テ ,  Compare:  chi:チ & テ:て  TT: left-ring< to the 'W' char"},

	{"ソ", "そ", "so",
		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},

	{"シ", "し", "shi",
		" shi:し:シ  'she' sleeps & snores゛,  Compare: tsu:ツ & so:ソ　to the more-laid-back shi:シ ",
		" shi:し:シ  し looks like a sheep hook,  Compare: ツシ  tsu-shi ",
		" TT: on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" し:シ  TT: on the 'D' char   Compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し"},

	{"オ", "お", "o",
		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" お:オ, on-the-go', cloak trailing behind   TT: index--> to the '6' char"},

	{"ン", "ん", "n",
		" n:ン:ん  looks like a cursive n, and sounds like it too",
		" n:ン:ん  ン looks like 'no' has a nose",
		" TT: index to the 'Y' char",
		" ン:ん   TT: index to the 'Y' char, sounds the way it looks"},

	{"ア", "あ", "a",
		" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
		" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
		" TT: left-ring-> to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
		" Compare:  me:め  あ  nu:ぬ  ne:ね   あ-Fuck mae-te!  TT: L-middle<- to the '3' char"},

	{"ウ", "う", "u",
		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ooh:lays-atop tsu:つ  u:ウ is more-angular with a tick for its top line ",
		" TT: L-middle> to the '4' char  う : u it's the middle-Finger, or wa:ワ with a mohawk ' ",
		" う:ウ ,  TT: left-middle> to the '4' char, う:ウ not ra:ら | wa:ワ with a mohawk"},
	//
}
