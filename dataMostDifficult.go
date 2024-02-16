package main

/*
. 7 Most difficult Kata:  type Romaji or Hira
. 10 Most difficult: Romaji: type Hira:  /nu, /ne, /na, /me, /ri, /ra, /a, /u, /o
*/
// When this deck is used, only Romaji objectives are requested
/*
Easy to hard:
  no の
  shi し
  etc
  Kata
  se vs sa

*/
var dataMostDiff = []charSetStruct{
	// All of the cards in this deck can be a problem (confusing?) when running with the standard mix (abandoned)
	//
	// ma is a stripped and lowered ho: ま vs ほ
	{"Just a single vertical w two horiz & a loop at its base", "ま - ", "ma",
		"Compare:  ma:ま  vs  ha:は",
		" ma is a stripped and lowered ho: ま vs ほ",
		" ma is a stripped and lowered ho: ま vs ほ",
		" ma is a stripped and lowered ho: ま vs ほ"},
	//
	// ho is a modified ma: ま becomes ほ
	{"Same as ha but w a cap line", "ほ:は - ", "ho",
		"Compare:  ma:ま  vs  ha:は  vs  ho:ほ",
		" ho is a modified ma: ま becomes ほ",
		" ho is a modified ma: ま becomes ほ",
		" ho is a modified ma: ま becomes ほ"},

	{"Same as ke, but w loop at base", "は:け - ", "ha",
		"ha:は  vs  ke:け",
		" ",
		" ",
		" "},

	{"Same as ha, but no loop at base", "け:は - ", "ke",
		"ha:は  vs  ke:け",
		" ",
		" ",
		" "},

	{"Two vertical lines as Hira, or a T with sloped top", "い:イ - ", "i",
		"i:い  vs  i:イ",
		" ",
		" ",
		" "},

	{"T with sloped top, or Two vertical lines as Hira", "い:イ - ", "i",
		"i:い  vs  i:イ",
		" ",
		" ",
		" "},

	{"T bent left, w two horiz top lines", "て becomes テ - ", "te",
		"テ  vs  て  (both are te)",
		" ",
		" ",
		" "},

	{"ra or u ?  As Kanji it has a vertical stroke", "ra:ら vs u:う - ", "ra",
		"u:うウ  vs  ra:らラ",
		" ",
		" ",
		" "},

	{"This Katakana has a horiz atop a Kata fu", "ra:ラ vs fu:フ - ", "ra",
		"ra:ラ  also: ら as Hira;  Compare: u:う　ra:ら　",
		" ",
		" ",
		" "},

	//
	/*
		// : zu has 2 Hira values: づ, ず
		{"ヅ", "づ", "zu",
			" zu:づ:ヅ ず:ズ , is either ず:ズ or づ:ヅ , as they are both the same sound ",
			" zu: for the Kataす think 'Sue' ズ is running with her hair flowing behind her",
			" TT: L-pinky>v to the 'Z' char, plus dakuten 濁点 ",
			" ?:づ:ヅ ず:ズ  that big wave つ sounds the same, TT: L-pinky>v to the 'Z' char"},

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

		{"ソ", "そ", "so",
			" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
			" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
			" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
			" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},

		{"オ", "お", "o",
			" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
			" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
			" TT: index--> to the '6' char     o:お:オ ",
			" お:オ, on-the-go', cloak trailing behind   TT: index--> to the '6' char"},
		{"ナ", "な", "na",
			" na:な:ナ  Kata ナ is just like the left side of the Hira な -- な:ナ ",
			" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  ( め ぬ )",
			" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
			" な:ナ  TT: index< to the 'U' char, Hiragana t-ies a knot, see the two chars thar "},

		{"ン", "ん", "n",
			" n:ン:ん  looks like a cursive n, and sounds like it too",
			" n:ン:ん  ン looks like 'no' has a nose",
			" TT: index to the 'Y' char",
			" ン:ん   TT: index to the 'Y' char, sounds the way it looks"},

		{"ラ", "ら", "ra",
			" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
			" ra is usually pronounced as la",
			" TT: right-ring up to the 'O' char",
			" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},
		{"ウ", "う", "u",
			" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
			" u:ooh:lays-atop tsu:つ  u:ウ is more-angular with a tick for its top line ",
			" TT: L-middle> to the '4' char  う : u it's the middle-Finger, or wa:ワ with a mohawk ' ",
			" う:ウ ,  TT: left-middle> to the '4' char, う:ウ not ra:ら | wa:ワ with a mohawk"},

		{"エ", "え", "e",
			" e:え:'エ' does have a vague, angular resemblance: eye see it as an, eye; may-bey, may-b-eh e ",
			" e:え ... see 'エ' as a ... an ... eye eh; may-bey, m'ay'-b-eh e   e:え:エ ",
			" TT: left-index> to the '5' char   e:え:エ ",
			" エ:え   TT: L-index> to the '5' char,  エ, eye see it as a ... an, eye ??  eh?:え:エ"},

		{"ク", "く", "ku",
			" ku:く:ク ,  Kata: starting with one angle, they settled for this??  く:ク ",
			" ku:く:ク ,  Compare:  ta:タ  ke:ケ  ku:ク  クケタ,  it's kuku, kookoo I tell you! ",
			" TT: R-index<- to the 'H' char  く:ク ",
			"く:ク ,  Compare:  クケタ ta:タ  ke:ケ  ?:ク   TT: R-index<- to the 'H' char"},

		{"サ", "さ", "sa",
			" sa:さ:サ , if you 'se':せ そサ ,  Compare:  Hira se:せ  to  Kata sa:サ  せサ ",
			" ... , but it looks an awful-lot like Hira se:せ, I'd say   ... just 'Hir'-say ?? ",
			" TT: ring>v to the 'X' char  --  at least Kata 'sa' still goes to the left ... ",
			" さ:サ ,  Compare:  Hira:se:せ  to  Kata:?:サ  せサ se:? TT: ring>v to the 'X' char"},
		{"セ", "せ", "se",
			" se:せ:セ  ... セ is just a bit more angular, as is the way with most Katakana ",
			" se:せ looks like a face saying something ... 'se'ing something ",
			" TT: right-pinky to the 'P' char",
			" ?:せ:セ  TT: R-pinky to the 'P' char,  セ is more angular, as with most Katakana"},

		{"タ", "た", "ta",
			" ta:た:タ  ... it's a ku:くク with a drool クタ... and that's くool I guess ",
			" ta:た:タ ,  Compare:  ku:ク  &  ke:ケ ",
			" TT: left-pinky< to the 'Q' char ",
			" top of タ looks like た,  more than ku:ク looks like く  TT: L-pinky< to the 'Q' char"},
		{"チ", "ち", "chi",
			" chi:ち:チ ,   ち チ ",
			" ...  chi:ち:チ ,  Compare:  chi:チ  te:テ ",
			" TT: left-pinky on the 'A' char ",
			" ?:ち:チ   Kata bares 'some' resemblance to ち   TT: L-pinky on the 'A' char "},

		{"ニ", "に", "ni",
			" ni:に:ニ  I can see a knee-cap in the Hiragana ",
			" ni:に:ニ  ",
			" TT: right-middle< to the 'I' char ",
			" ?:に:ニ  TT: R-middle< to the 'I' char,  I can see a knee-cap in the Hiragana "},

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

		{"レ", "れ", "re",
			" re:レ:れ  laying on the little finger",
			" re:レ:れ  the れ looks like someone getting blown, laid ",
			" TT: right-pinky on it's ';' char",
			" レ:れ   TT: right-pinky on it's ';' char"},
		{"ロ", "ろ", "ro",
			" ro:ロ:ろ  no loop on ro:ろ  Compare ru:る",
			" ro",
			" TT: right-pinky slide down --> right of the '?' char",
			" Looks like a square O,  TT: right-pinky slide down --> right of the '?' char"},

		//

	*/
}
