package main

var fileOfCardsRare = []charSetStruct{
	//
	// ya, yu, yo's of hi:ひ:ヒ -------------------------------------------------------------------------
	//
	{"ヒャ", "ひゃ", "hya",
		" hya:ひ:ヒ  hi->bi->pi hya from hi which looks like a smiling hi-hi ひひ",
		" TT: left-index > to the 'V' char  ひゃ ヒャ",
		" T: plus R-index ^^<< to the '7' char (shifted)  ひゃ ヒャ",
		" ?:ひ:ヒ   '?'-man   TT: left-index > to the 'V' char"},
	{"ヒュ", "ひゅ", "hyu",
		" hyu:ひ:ヒ  hi->bi->pi hyu from hi which looks like a smiling hi-hi ひひ",
		" TT: left-index > to the 'V' char  ひゅ ヒュ",
		" TT: plus R-index ^^- to the '8' char (shifted)  ひゅ ヒュ",
		" ?:ひ:ヒ   '?'-man   TT: left-index > to the 'V' char"},
	{"ヒョ", "ひょ", "hyo",
		" hyo:ひ:ヒ  hi->bi->pi hyo from hi which looks like a smiling hi-hi ひひ",
		" TT: left-index > to the 'V' char  ひょ ヒョ",
		" TT: plus R-ring ^^<- to the '9' char (shifted)  ひょ ヒョ",
		" ?:ひ:ヒ   '?'-man   TT: left-index > to the 'V' char"},
	//
	// ya, yu, yo's of bi:び:ビ -------------------------------------------------------------------------
	//
	{"ビャ", "びゃ", "bya",
		" bya:ビ:び  hi->bi->pi bya from hi which looks like a smiling hi-hi ひひ",
		" TT: L-index to the 'V' char, plus a dakuten 濁点  びゃ ビャ",
		" TT: plus R-index ^^<< to the '7' char (shifted)  びゃ ビャ",
		" ?:ビ:び   びゃ ビャ   TT: L-index to 'V' char, plus a dakuten"},
	{"ビュ", "びゅ", "byu",
		" byu:ビ:び  hi->bi->pi byu from hi which looks like a smiling hi-hi ひひ",
		" TT: L-index to the 'V' char, plus a dakuten 濁点  びゅ ビュ",
		" TT: plus R-index ^^- to the '8' char (shifted)  びゅ ビュ",
		" ?:ビ:び   びゅ ビュ   TT: L-index to 'V' char, plus a dakuten"},
	{"ビョ", "びょ", "byo",
		" byo:ビ:び  hi->bi->pi byo from hi which looks like a smiling hi-hi ひひ",
		" TT: L-index to the 'V' char, plus a dakuten 濁点  びょ ビョ",
		" TT: plus R-ring ^^<- to the '9' char (shifted)  びょ ビョ",
		" ?:ビ:び   びょ ビョ   TT: L-index to 'V' char, plus a dakuten"},
	//
	// ya, yu, yo's of pi:ぴ:ピ -------------------------------------------------------------------------
	//
	{"ピャ", "ぴゃ", "pya",
		" pya:ピ:ぴ  hi-->pi  becomes with a handakuten゜半濁点  ぴゃ ピャ",
		" TT: L-index to the 'V' char, plus a handakuten゜半濁点  ぴゃ ピャ",
		" TT: plus R-index ^^<< to the '7' char (shifted)  ぴゃ ピャ",
		" ?:ピ:ぴ    ぴゃ ピャ   TT: L-index to the 'V' char, plus a handakuten゜半濁点"},
	{"ピュ", "ぴゅ", "pyu",
		" pyu:ピ:ぴ   hi-->pi  becomes with a handakuten゜半濁点  ぴゅ ピュ",
		" TT: L-index to the 'V' char, plus a handakuten゜半濁点  ぴゅ ピュ",
		" TT: plus R-index ^^- to the '8' char (shifted)  ぴゅ ピュ",
		" ?:ピ:ぴ    ぴゅ ピュ   TT: L-index to the 'V' char, plus a handakuten゜半濁点"},
	{"ピョ", "ぴょ", "pyo",
		" pyo:ピ:ぴ   hi-->pi  becomes with a handakuten゜半濁点  ぴょ ピョ",
		" TT: L-index to the 'V' char, plus a handakuten゜半濁点  ぴょ ピョ",
		" TT: plus R-ring ^^<- to the '9' char (shifted)  ぴょ ピョ",
		" ?:ピ:ぴ    ぴょ ピョ   TT: L-index to the 'V' char, plus a handakuten゜半濁点"},
	//
	// ya, yu, yo's of ki:き ---------------------------------------------------------------------------
	//
	{"キャ", "きゃ", "kya",
		" kya:きゃ is an easy one, キャ has the same top, & the 'ya' is similar",
		" kya:きゃ キャ has the same top and the 'ya' is similar",
		" TT: 'G' char, then index <-- to the 7 char    きゃ,   Compare:  sa:さ ",
		" きゃ is an easy one, キャ has the same top, & the 'ya' is similar"},
	{"キュ", "きゅ", "kyu",
		" kyu:きゅ キュ has the same top",
		" kyu: ゅ:ュ",
		" TT: yu:ゅ index->, to the '8' char, yu looks like:  ユ are #1",
		" Key-ゅ   and:  ユ are #1     キュ has the same top as き "},
	{"キョ", "きょ", "kyo",
		" kyo:きょ:キョ   キョ  has the same top",
		" kyo:キョ ",
		" TT: ょ:ョ middle to the '9' char, triple yo:ょ ヨ",
		" Key-ょ  キョ has the same top as き   hira:ょ is a yoyo , triple yo:ヨ"},
	//
	// ya, yu, yo's of gi:ぎ ---------------------------------------------------------------------------
	//
	{"ギャ", "ぎゃ", "gya",
		" gya  gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し",
		" gya:ギャ ",
		" TT: ya:ゃ index <-- to the 7 char",
		" ぎ:ギ is an easy Kata, ギ has the same top as ki:き   ka,ki,ku,ke,ko-->ga,?,gu,ge,go"},
	{"ギュ", "ぎゅ", "gyu",
		" gyu:ぎゅ:ギュ   ギ has same top as ki:き   ka,(ki),ku,ke,ko-->ga,(gi),gu,ge,go",
		" gyu:ぎゅ:ギュ ",
		" TT: yu:ゅ index->, to the '8' char, yu looks like:  ユ are #1",
		" ぎ:ギ is an easy Kata, ギ has the same top as ki:き    ka,ki,ku,ke,ko-->ga,?,gu,ge,go"},
	{"ギョ", "ぎょ", "gyo",
		" gyo: gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し ",
		" gyo: ぎ:ギ is always the sound gi, and is Seldom ji ( that usually being じ:ji:ジ )",
		" TT: 'G' char, ぎ:ギ, then ょ:ョ:middle to the '9' char, triple yo:ょ ヨ      Compare: ざ",
		" It's sound is always from き ,and NEVER from し or ち   ka,ki,ku,ke,ko--> ga,?,gu,ge,go"},
	//

	// ya, yu, yo's of shi:し:シ ------------------------------------------------------------------------
	//
	{"シャ", "しゃ", "sha",
		" sha:し:シ  'she' sleeps & snores゛,  Compare: tsu:ツ & so:ソ　to the more-laid-back shi:シ ",
		" sha:し:シ  し looks like a sheep hook,  Compare: ツシ  tsu-shi ",
		" TT: on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" し:シ  TT: on the 'D' char   Compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し"},
	{"シュ", "しゅ", "shu",
		" shu:し:シ  'she' sleeps & snores゛,  Compare: tsu:ツ & so:ソ　to the more-laid-back shi:シ ",
		" shu:し:シ  し looks like a sheep hook,  Compare: ツシ  tsu-shi ",
		" TT: on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" し:シ  TT: on the 'D' char   Compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し"},
	{"ショ", "しょ", "sho",
		" sho:し:シ  'she' sleeps & snores゛,  Compare: tsu:ツ & so:ソ　to the more-laid-back shi:シ ",
		" sho:し:シ  し looks like a sheep hook,  Compare: ツシ  tsu-shi ",
		" TT: on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" し:シ  TT: on the 'D' char   Compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し"},
	//
	// ya, yu, yo's of ji:じ:ジ -------------------------------------------------------------------------
	//
	{"ジャ", "じゃ", "ja",
		" ja:ジ:じ    Seldom is it fr chi:ち ,it's the sound ji, NEVER gi (that being ぎ:gi:ギ)",
		" ja, the sound, VERY RARELY is from chi:ち:ヂ ,but NEVER the sound gi (that being ぎ:gi:ギ)",
		" TT: 'D' char, remember, the sound:ji is nearly-always from shi:し with a dakuten 濁点 じ:ジ",
		" ジ:じ  TT: 'D' char, with a dakuten 濁点     it's the sound:gee:jeep"},
	{"ジュ", "じゅ", "ju",
		" ju:ジ:じ    Seldom is it fr chi:ち ,it's the sound ji, NEVER gi (that being ぎ:gi:ギ)",
		" ju, the sound, VERY RARELY is from chi:ち:ヂ ,but NEVER the sound gi (that being ぎ:gi:ギ)",
		" TT: 'D' char, remember, the sound:ji is nearly-always from shi:し with a dakuten 濁点 じ:ジ",
		" ジ:じ  TT: 'D' char, with a dakuten 濁点     it's the sound:gee:jeep"},
	{"ジョ", "じょ", "jo",
		" jo:ジ:じ    Seldom is it fr chi:ち ,it's the sound ji, NEVER gi (that being ぎ:gi:ギ)",
		" jo, the sound, VERY RARELY is from chi:ち:ヂ ,but NEVER the sound gi (that being ぎ:gi:ギ)",
		" TT: 'D' char, remember, the sound:ji is nearly-always from shi:し with a dakuten 濁点 じ:ジ",
		" ジ:じ  TT: 'D' char, with a dakuten 濁点     it's the sound:gee:jeep"},
	//
	// ya, yu, yo's of chi:ち:チ ------------------------------------------------------------------------
	//
	{"チャ", "ちゃ", "cha",
		" cha:ち:チ ,   ち チ ",
		" ...  chi:ち:チ ,  Compare:  chi:チ  te:テ ",
		" TT: left-pinky on the 'A' char ",
		" ?:ち:チ   Kata bares 'some' resemblance to ち   TT: L-pinky on the 'A' char "},
	{"チュ", "ちゅ", "chu",
		" chu:ち:チ ,   ち チ ",
		" ...  chi:ち:チ ,  Compare:  chi:チ  te:テ ",
		" TT: left-pinky on the 'A' char ",
		" ?:ち:チ   Kata bares 'some' resemblance to ち   TT: L-pinky on the 'A' char "},
	{"チョ", "ちょ", "cho",
		" cho:ち:チ ,   ち チ ",
		" ...  chi:ち:チ ,  Compare:  chi:チ  te:テ ",
		" TT: left-pinky on the 'A' char ",
		" ?:ち:チ   Kata bares 'some' resemblance to ち   TT: L-pinky on the 'A' char "},
	//
	// ya, yu, yo's of ni:に:ニ -------------------------------------------------------------------------
	//
	{"ニャ", "にゃ", "nya",
		" nya:に:ニ  I can see a knee-cap in the Hiragana ",
		" nya:に:ニ  ",
		" TT: right-middle< to the 'I' char ",
		" ?:に:ニ  TT: R-middle< to the 'I' char,  I can see a knee-cap in the Hiragana "},
	{"ニュ", "にゅ", "nyu",
		" nyu:に:ニ  I can see a knee-cap in the Hiragana ",
		" nyu:に:ニ  ",
		" TT: right-middle< to the 'I' char ",
		" ?:に:ニ  TT: R-middle< to the 'I' char,  I can see a knee-cap in the Hiragana "},
	{"ニョ", "にょ", "nyo",
		" nyo:に:ニ  I can see a knee-cap in the Hiragana ",
		" nyo:に:ニ  ",
		" TT: right-middle< to the 'I' char ",
		" ?:に:ニ  TT: R-middle< to the 'I' char,  I can see a knee-cap in the Hiragana "},
	//
	// ya, yu, yo's of mi:み:ミ -------------------------------------------------------------------------
	//
	{"ミャ", "みゃ", "mya",
		" mya:ミ:み  mi is three, or phallic 'me' is mi ",
		" mya is three; lines",
		" TT: right-index to the 'N' char",
		" ミ:み   TT: right-index to the 'N' char"},
	{"ミュ", "みゅ", "myu",
		" myu:ミ:み  mi is three, or phallic 'me' is mi ",
		" myu is three; lines",
		" TT: right-index to the 'N' char",
		" ミ:み   TT: right-index to the 'N' char"},
	{"ミョ", "みょ", "myo",
		" myo:ミ:み  mi is three, or phallic 'me' is mi ",
		" myo is three; lines",
		" TT: right-index to the 'N' char",
		" ミ:み   TT: right-index to the 'N' char"},
	//
	// "group" ya yu yo =============================================================================================
	//
	{"ヤ", "や", "ya",
		" ya:ヤ:や  look very similar, which is too rare",
		" ya",
		" TT: right-index to the '7' char, un-shifted",
		" ヤ:や   TT: right-index to the '7' char, un-shifted"},
	{"ユ", "ゆ", "yu",
		" yu:ユ:ゆ  Katakana looks like number 1, you are #1 ",
		" yu",
		" TT: on the '8' char, shiftless when naked  ",
		" ?:ユ:ゆ  TT: on the '8' char, shiftless when naked "},
	{"ヨ", "よ", "yo",
		" yo:ヨ:よ  Looks like a toy",
		" yo",
		" TT: right-middle to the '9' char, un-shifted",
		" ヨ:よ  Looks like a toy  TT: right-middle to the '9' char, un-shifted"},
	//
	// ya, yu, yo's of ri:り:リ ---------------------------------------------------------------------------
	//
	{"リャ", "りゃ", "rya",
		" rya:リ:り  actually, very very similar, actually!",
		" rya",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},
	{"リュ", "りゅ", "ryu",
		" ryu:リ:り  actually, very very similar, actually!",
		" ryu",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},
	{"リョ", "りょ", "ryo",
		" ryo:リ:り  actually, very very similar, actually!",
		" ryo",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},
	//
	//
	//
	{"ソ", "そ", "so",
		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},
	{"ニ", "に", "ni",
		" ni:に:ニ  I can see a knee-cap in the Hiragana ",
		" ni:に:ニ  ",
		" TT: right-middle< to the 'I' char ",
		" ?:に:ニ  TT: R-middle< to the 'I' char,  I can see a knee-cap in the Hiragana "},
	{"ヒ", "ひ", "hi",
		" hi:ひ:ヒ  a smiling mouth doing a hee-hee hi-hi",
		" hi, pronounced hee ",
		" TT: left-index > to the 'V' char",
		" ?:ひ:ヒ   '?'-man   TT: left-index > to the 'V' char"},
	{"ヘ", "へ", "he",
		" he:ヘ:へ  Hira and Kata are pretty-much the same, both hay:he stacks",
		" he",
		" TT: left-pinky way-up to the '^' char",
		" ヘ:へ   TT: left-pinky way-up to the '^' char, clearly stacks of feed"},
	{"ビ", "び", "bi",
		" bi:ビ:び   hi-bi-pi ",
		" bi is from hi which looks like a smiling hi-hi びび",
		" TT: L-index to the 'V' char, plus a dakuten 濁点",
		" ?:ビ:び from hi, looks like smiling hi-hi びび   TT: L-index to 'V' char, plus a dakuten"},
	{"ベ", "べ", "be",
		" be:ベ:べ  same same, , clearly stacks of feed, hay:he stacks, ",
		" be:  he-->be-->pe , dakuten゛濁点, handakuten゜半濁点",
		" TT: left-pinky way-up to the '^' char, plus dakuten 濁点",
		" TT: left-pinky way-up to the '^' char, plus dakuten 濁点   stack of feed, "},
}
