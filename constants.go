package main

/*
There are four files (or decks) of cards in this file: constants.go
	fileOfCardsE
	fileOfCardsS
	fileOfCardsMostDifficult
	fileOfCards
*/
// All decks will draw cards per this aCard var
var aCard = charSetStruct{}

var fileOfCardsE = []charSetStruct{

	{"イィ", "やぃ", "yi",
		" やぃ",
		" イィ (yi)",
		" ",
		" "},
	{"イェ", "やェ", "ye",
		" やェ",
		" イェ (ye)",
		" ",
		" "},

	{"ウァ", "うわ", "wa",
		" うぃ",
		" ウァ (wa)",
		" ",
		" "},
	{"ウィ", "うぃ", "wi",
		" ",
		" ウィ (wi)",
		" ",
		" "},
	{"ウゥ", "うぅ", "wu",
		" うぅ",
		" ウゥ (wu)",
		" ",
		" "},
	{"ウェ", "うぇ", "we",
		" うぇ",
		" ウェ (we)",
		" ",
		" "},
	{"ウォ", "うぉ", "wo",
		" うぉ",
		" ウォ (wo)",
		" ",
		" "},

	{"ヴァ", "ゔぁ", "va",
		" ゔぁ",
		" ヴァ (va)",
		" ",
		" "},
	{"ヴィ", "ゔぃ", "",
		" ゔぃ",
		" ヴィ (vi)",
		" ",
		" "},
	{"ヴ", "ゔ", "vu",
		" ゔ",
		" ヴ (vu)",
		" ",
		" "},
	{"ヴェ", "ゔぇ", "ve",
		" ゔぇ",
		" ヴェ (ve)",
		" ",
		" "},
	{"ヴォ", "ゔぉ", "vo",
		" ゔぉ",
		" ヴォ (vo)",
		" ",
		" "},
	{"ヴィェ", "ゔぃぇ", "vye",
		" ゔぃぇ",
		" ヴィェ (vye)",
		" ",
		" "},
}

/*
			{"", "", "",
				" ",
				" ",
				" ",
				" "},

	キェ (kye)
	ギェ (gye)

	クァ (kwa)
	クィ(kwi)
	クェ (kwe)
	クォ (kwo)

	to get closer to the pronunciation of foreign words, a list of not so common extended katakana is used:

	グァ (gwa)
	グィ (gwi)
	グェ (gwe)
	グォ (gwo)

	シェ (she)
	ジェ (je)

	スィ (si)
	ズィ (zi)

	チェ (che)

	ツァ (tsa)
	ツィ (tsi)
	ツェ (tse)
	ツォ (tso)

	ティ (ti)
	テゥ (tu)
	ディ (di)
	デゥ (du)

	ニェ (nye)
	ヒェ (hye)
	ビェ (bye)
	ピェ (pye)

	ファ (fa)
	フィ (fi)
	フェ (fe)
	フォ (fo)

	フィェ (fye)

	ホゥ (hu)

	ミェ (mye)

	リェ (rye)

	ラ゜(la)
	リ゜(li)
	ル゜(lu)
	レ゜(le)
	ロ゜(lo)
*/

var fileOfCardsS = []charSetStruct{

	{"ザ", "ざ", "za",
		" sa-->za:ざ  dakuten  濁点",
		" sa-->za   za:ザ:ざ  so it still protrudes to the left",
		" TT: down on the 'X' char plus a dakuten 濁点 for sa-->za",
		" ?:ざ:ザ  TT: down on the 'X' char, plus a dakuten 濁点"},

	{"デ", "で", "de",
		" de:で:デ ",
		" de from ta-->da , plus a dakuten 濁点",
		" TT: left-ring to the 'W' char, plus a dakuten 濁点",
		" ?:で:デ  , plus a dakuten, TT: L-ring to the 'W' char, plus a dakuten 濁点"},

	{"セ", "せ", "se",
		" se:せ:セ  ... セ is just a bit more angular, as is the way with most Katakana ",
		" se:せ looks like a face saying something ... 'se'ing something ",
		" TT: right-pinky to the 'P' char",
		" ?:せ:セ  TT: R-pinky to the 'P' char,  セ is more angular, as with most Katakana"},

	{"キ", "き", "ki",
		" ki:き:キ ... is an easy one!   キ has the same top as ki:き, and they look like 'keys' ",
		" ki:き ... they both have the same top:  き キ  looks like a chi-key",
		" TT: L-index> to the 'G' char,  Think: ki-of-G  き,  Compare:  sa:さ-ki:き  さき ... saki ",
		"き:キ  TT: L-index > to the 'G' char,  キ same top as き, looks like a 'chi-key'"},

	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},

	{"ダ", "だ", "da",
		" ta-->da  だ:ダ ",
		" da is ta with a dakuten 濁点",
		" TT: left-pinky to the 'Q' char",
		" だ:ダ  ta-->?  TT: left-pinky to the 'Q' char, plus a dakuten 濁点"},

	{"ブ", "ぶ", "bu",
		" bu:ブ:ぶ",
		" fu-->bu , plus a dakuten 濁点",
		" TT: left-ring to the '2' char is fu-->bu",
		" ブ:ぶ   TT: left-ring to the '2' char, plus a dakuten 濁点 :: fu-->?"},

	{"サ", "さ", "sa",
		" sa:さ:サ , if you 'se':せ そサ ,  Compare:  Hira se:せ  to  Kata sa:サ  せサ ",
		" ... , but it looks an awful-lot like Hira se:せ, I'd say   ... just 'Hir'-say ?? ",
		" TT: ring>v to the 'X' char  --  at least Kata 'sa' still goes to the left ... ",
		" さ:サ ,  Compare:  Hira:se:せ  to  Kata:?:サ  せサ se:? TT: ring>v to the 'X' char"},

	{"ソ", "そ", "so",
		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},

	{"マ", "ま", "ma",
		" ma:マ:ま ",
		" mama is using a brest pump ",
		" TT: right-index on the 'J' char :: ma",
		" ?:マ:ま   lady is using a brest pump    TT: right-index on the 'J' char"},

	{"ゼ", "ぜ", "ze",
		" sa-->za   se-->ze , plus a dakuten 濁点",
		" ze:ゼ:ぜ",
		" TT: pinky to the 'P' char, plus a dakuten 濁点   ze:ゼ:ぜ",
		" ?:ゼ:ぜ   TT: pinky to the 'P' char, plus a dakuten 濁点   sa-->za  se:? "},

	{"チ", "ち", "chi",
		" chi:ち:チ ,   ち チ ",
		" ...  chi:ち:チ ,  Compare:  chi:チ  te:テ ",
		" TT: left-pinky on the 'A' char ",
		" ?:ち:チ   Kata bares 'some' resemblance to ち   TT: L-pinky on the 'A' char "},

	{"ヌ", "ぬ", "nu",
		" nu:ぬ:ヌ   Compare:  me:メ  &  na:ナ ",
		" nu flew ヌ -- till it hit a ceiling ",
		" TT: R-ring<--- to the '1' char, at least two lines cross, nu flew ヌ -- hit a ceiling ",
		" ?:ぬ:ヌ   ?-flew ヌ -- till it hit a ceiling  TT: R-ring<--- to the '1' char"},

	{"ニ", "に", "ni",
		" ni:に:ニ  I can see a knee-cap in the Hiragana ",
		" ni:に:ニ  ",
		" TT: right-middle< to the 'I' char ",
		" ?:に:ニ  TT: R-middle< to the 'I' char,  I can see a knee-cap in the Hiragana "},

	{"ピ", "ぴ", "pi",
		" pi:ピ:ぴ  ",
		" hi-->pi  becomes with a handakuten゜半濁点",
		" TT: L-index to the 'V' char, plus a handakuten゜半濁点",
		" ?:ピ:ぴ   TT: L-index to the 'V' char, plus a handakuten゜半濁点"},

	{"ホ", "ほ", "ho",
		" ho:ホ:ほ  looks like a hoe in a dress",
		" ho",
		" TT: right-ring-way-up to the '-' char",
		" ?:ホ:ほ   TT: right-ring-way-up to the '-' char"},

	{"ケ", "け", "ke",
		" ke:け:ケ  bits of it are there, just as many curves though ,  Compare to:  ku:ク  ta:タ ",
		" ke:け  ... looks enough like a Keg ...  Foam-ment THAT image in your brain ",
		" TT: right-pinky-> one-over to the ':*' chars   け:ケ ",
		"け:ケ  bits still there  Compare:  ku:ク  ta:タ  クケタ  TT: R-pinky-> to the ':*' chars"},

	{"ロ", "ろ", "ro",
		" ro:ロ:ろ  no loop on ro:ろ  Compare ru:る",
		" ro",
		" TT: right-pinky slide down --> right of the '?' char",
		" Looks like a square O,  TT: right-pinky slide down --> right of the '?' char"},

	{"バ", "ば", "ba",
		" ba:バ:ば  ha-->ba-->pa   dakuten 濁点, handakuten゜半濁点",
		" ba",
		" TT: L-index on the 'F' char, plus a dakuten 濁点 ",
		" ?:バ:ば  ha-->?  TT: L-index on the 'F' char, plus a dakuten 濁点"},

	{"タ", "た", "ta",
		" ta:た:タ  ... it's a ku:くク with a drool クタ... and that's くool I guess ",
		" ta:た:タ ,  Compare:  ku:ク  &  ke:ケ ",
		" TT: left-pinky< to the 'Q' char ",
		" top of タ looks like た,  more than ku:ク looks like く  TT: L-pinky< to the 'Q' char"},

	{"ユ", "ゆ", "yu",
		" yu:ユ:ゆ  Katakana looks like number 1, you are #1 ",
		" yu",
		" TT: on the '8' char, shiftless when naked  ",
		" ?:ユ:ゆ  TT: on the '8' char, shiftless when naked "},

	{"パ", "ぱ", "pa",
		" pa:パ:ぱ  ha-->ba-->pa   dakuten 濁点, handakuten゜半濁点",
		" pa",
		" TT: left-index on the 'F' char, plus a handakuten 半濁点",
		" ha-->?  TT: L-index on the 'F' char, plus a handakuten 半濁点"},

	{"ト", "と", "to",
		" to:と:ト   Kata 'toe' is flipped-out. Kicked in the balls. on its head: と -> ト ",
		" Katakana to:と is flipped-out. Kicked in the balls, On its head ",
		" TT: left-ring on the 'S' char, Kata toe is flipped-out. Kicked on its head: と -> ト ",
		" Kata char was kicked on its head: と --> ト  TT: L-ring on the 'S' char"},

	{"ポ", "ぽ", "po",
		" po:ポ:ぽ  ho-->bo-->po  dakuten, handakuten゜半濁点",
		" po",
		" TT: right-ring-way-up to the '-' char, plus a handakuten゜半濁点",
		" ?:ポ:ぽ   TT: right-ring-way-up to the '-' char, plus a handakuten゜半濁点"},

	{"ア", "あ", "a",
		" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
		" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
		" TT: left-ring-> to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
		" Compare:  me:め  あ  nu:ぬ  ne:ね   あ-Fuck mae-te!  TT: L-middle<- to the '3' char"},

	{"オ", "お", "o",
		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" お:オ, on-the-go', cloak trailing behind   TT: index--> to the '6' char"},

	{"ク", "く", "ku",
		" ku:く:ク ,  Kata: starting with one angle, they settled for this??  く:ク ",
		" ku:く:ク ,  Compare:  ta:タ  ke:ケ  ku:ク  クケタ,  it's kuku, kookoo I tell you! ",
		" TT: R-index<- to the 'H' char  く:ク ",
		"く:ク ,  Compare:  クケタ ta:タ  ke:ケ  ?:ク   TT: R-index<- to the 'H' char"},

	{"メ", "め", "me",
		" me:メ:め   Compare: nu:ヌ",
		" me",
		" TT: right-pinky-slide down to the '?' char",
		" メ:め messy-lesser ver of nuddle:ぬ　 TT: R-pinky-slide down to the '?' char"},

	{"ヨ", "よ", "yo",
		" yo:ヨ:よ  Looks like a toy",
		" yo",
		" TT: right-middle to the '9' char, un-shifted",
		" ヨ:よ  Looks like a toy  TT: right-middle to the '9' char, un-shifted"},

	{"ル", "る", "ru",
		" ru:ル:る  ru-two, and the Hira has a luup for ru ",
		" ru",
		" TT: right-pinky-slide under to the left, on the '.' char",
		" ル:る  ?:two    TT: right-pinky-slide under to the left, on the '.' char"},

	{"テ", "て", "te",
		" te:て:テ  Looks like a T",
		" te:て:テ ,  Compare:  chi:チ & te:テ:て ",
		" TT: left-ring< to the 'W' char, te:テ:て  went the wrong way, and gained a flat hat ",
		" ?:て:テ ,  Compare:  chi:チ & テ:て  TT: left-ring< to the 'W' char"},

	{"カ", "か", "ka",
		" ka:か:カ ... is an easy one!  It simply looks like a 'K' ",
		" ... Kata is the same カ , albeit more-angular and with one-less line to draw than か ",
		" TT: left-index--> to the 'T' char   か or カ",
		" か:カ ... looks like a 'K'    TT: left-index--> to the 'T' char"},

	{"ヤ", "や", "ya",
		" ya:ヤ:や  look very similar, which is too rare",
		" ya",
		" TT: right-index to the '7' char, un-shifted",
		" ヤ:や   TT: right-index to the '7' char, un-shifted"},

	{"ミ", "み", "mi",
		" mi:ミ:み  mi is three, or phallic 'me' is mi ",
		" mi is three; lines",
		" TT: right-index to the 'N' char",
		" ミ:み   TT: right-index to the 'N' char"},

	{"エ", "え", "e",
		" e:え:'エ' does have a vague, angular resemblance: eye see it as an, eye; may-bey, may-b-eh e ",
		" e:え ... see 'エ' as a ... an ... eye eh; may-bey, m'ay'-b-eh e   e:え:エ ",
		" TT: left-index> to the '5' char   e:え:エ ",
		" エ:え   TT: L-index> to the '5' char,  エ, eye see it as a ... an, eye ??  eh?:え:エ"},

	{"ゴ", "ご", "go",
		" go:ご:ゴ   looks like two Koy fish ko-y, and as ka-->ga, ko-->go ",
		" go:ご:ゴ    ご:ゴ makes sense, 'cause angles'",
		" TT: index <--- to the 'B' char, plus a dakuten 濁点  ご",
		" ご:ゴ  TT: index <--- to the 'B' char, plus a dakuten 濁点 ka,ki,ku,ke,ko-->ga,gi,gu,ge,?"},

	{"ギ", "ぎ", "gi",
		" gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し ",
		" ぎ:ギ is always the sound gi, and is Seldom ji ( that usually being じ:ji:ジ )",
		" TT: L-index on 'G' char, plus a dakuten 濁点, ぎ:ギ, compare to za:ざ from sa:さ",
		" ぎ:ギ  TT: 'G' char,  sound is from き ,and NEVER from し or ち  ka-->ga,?,gu,ge,go"},

	{"ウ", "う", "u",
		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ooh:lays-atop tsu:つ  u:ウ is more-angular with a tick for its top line ",
		" TT: L-middle> to the '4' char  う : u it's the middle-Finger, or wa:ワ with a mohawk ' ",
		" う:ウ ,  TT: left-middle> to the '4' char, う:ウ not ra:ら | wa:ワ with a mohawk"},

	{"ビ", "び", "bi",
		" bi:ビ:び   hi-bi-pi ",
		" bi is from hi which looks like a smiling hi-hi びび",
		" TT: L-index to the 'V' char, plus a dakuten 濁点",
		" ?:ビ:び from hi, looks like smiling hi-hi びび   TT: L-index to 'V' char, plus a dakuten"},

	{"ジ", "じ", "ji",
		" ji:ジ:じ    Seldom is it fr chi:ち ,it's the sound ji, NEVER gi (that being ぎ:gi:ギ)",
		" ji, the sound, VERY RARELY is from chi:ち:ヂ ,but NEVER the sound gi (that being ぎ:gi:ギ)",
		" TT: 'D' char, remember, the sound:ji is nearly-always from shi:し with a dakuten 濁点 じ:ジ",
		" ジ:じ  TT: 'D' char, with a dakuten 濁点     it's the sound:gee:jeep"},

	{"イ", "い", "i",
		" i:イ:い  the Kata イ looks more like a hiragana te:て, but still two mostly-vertical lines: イ",
		" i:い:イ  for the Kata ... maybe shift the two lines of the hiragana い to forms an 'I'  i:イ",
		" TT: left-middle< to the 'E' char  い:イ  Longer on the left is i, Compare: ri:リ:り",
		" い:イ at least it is still two mostly-vertical lines   TT: L-middle< to the 'E' char"},

	{"ス", "す", "su",
		" su:す:ス ",
		" ス sue:す looks like she is running (and they were looking for angles)",
		" TT: left-index < to the 'R' char,  (sue is running ス)",
		" ス:す  she is running  TT: left-index < to the 'R' char"},

	{"ゲ", "げ", "ge",
		" ge:げ:ゲ bits still there (as many curves though)  Compare  ku:gu:グ,  &  ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ   don't forget the dakuten 濁点",
		" TT: ':*' chars, 濁点    げ:ゲ  ダグゲ    ka,ki,ku,ke,ko-->ga,gi,gu,?,go  濁点",
		" げ:ゲ   TT: L-pinky> to ':*' chars   げ Compare to ku:gu:グ, and ta:da:ダ "},

	{"ネ", "ね", "ne",
		" ne:ね:ネ  Kata looks like a hoe that got nailed ネ ",
		" something got nailed-down by that spike at the left ",
		" TT: R-ring<v to the ',<' chars  ...  Yea, to me it is a hoe:ほ that got nailed ネ ",
		" ?:ね:ネ  Kata looks like a hoe that got nailed ネ  TT: R-ring<v to the ',<' chars"},

	{"モ", "も", "mo",
		" mo:モ:も ",
		" mo",
		" TT: right-middle<v to the 'M' char",
		" モ:も   TT: right-middle<v to the 'M' char,  m is for ?"},

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

	{"ヘ", "へ", "he",
		" he:ヘ:へ  Hira and Kata are pretty-much the same, both hay:he stacks",
		" he",
		" TT: left-pinky way-up to the '^' char",
		" ヘ:へ   TT: left-pinky way-up to the '^' char, clearly stacks of feed"},

	{"ズ", "ず", "zu",
		" zu:ず:ズ:づ:ヅ: , is either ず:ズ or づ:ヅ , as they are both the same sound",
		" zu: 'Sue' ズ is running with her hair flowing behind her (from su:す) ",
		"  TT: L-index< to the 'R' char, same sound: a big wave つ  or a  snake ず at the zoo",
		" ?:ず:ズ づ:ヅ  same sound   TT: left-index< to the 'R' char, dakuten 濁点 "},

	{"ド", "ど", "do",
		" do:ド:ど   ta-->da  to --> do    plus a dakuten 濁点",
		" do:ど:ド ",
		" TT: left-ring on the 'S' char, plus a dakuten 濁点 ",
		" ?:ど:ド  TT: left-ring on the 'S' char, plus a dakuten 濁点   ta-->da  to-->? "},

	{"ツ", "つ", "tsu",
		" tsu:つ:ツ  see water crashing on the she shore",
		" tsu",
		" TT: left-pinky on 'Z' char",
		" ?:つ:ツ see water crashing on the she shore  TT: left-pinky on 'Z' char"},

	{"ペ", "ぺ", "pe",
		" pe:ペ:ぺ  same same, , clearly stacks of feed, hay:he stacks, handakuten゜半濁点",
		" pe:  he-->be-->pe , dakuten 濁点, handakuten゜半濁点",
		" TT: left-pinky way-up to the '^' char, plus handakuten゜半濁点",
		" ペ:ぺ   TT: left-pinky way-up to the '^' char, stack of feed, handakuten゜半濁点"},

	{"ノ", "の", "no",
		" no:の:ノ   it looks like a 'no' symbol ",
		" no:の:ノ  and the Kata retains the slash ",
		" TT: right-middle on 'K' char",
		" ?:の:ノ  TT: right-middle on 'K' char   ?-thank-you "},

	{"ヲ", "を", "wo",
		" wo:ヲ:を   Compare: fu:フ　wa:ワ　u:ウ　a:ア　te:テ:て　 ",
		" wo:ヲ, at least it looks something like wa:ワ, though shifted",
		" TT: right-ring up to the '0' char (shifted)",
		" ヲ:を   Compare: wa:ワ   TT: right-ring up to the '0' char (shifted)"},

	{"ヅ", "づ", "zu",
		" zu:づ:ヅ ず:ズ , is either ず:ズ or づ:ヅ , as they are both the same sound ",
		" zu: for the Kataす think 'Sue' ズ is running with her hair flowing behind her",
		" TT: L-pinky>v to the 'Z' char, plus dakuten 濁点 ",
		" ?:づ:ヅ ず:ズ  that big wave つ sounds the same, TT: L-pinky>v to the 'Z' char"},

	{"ム", "む", "mu",
		" mu:ム:む  somehow, it does look like a cow",
		" mu",
		" TT: right-pinky way-over to the '}]' chars",
		" ム:む   TT: right-pinky way-over to the '}]' chars"},

	{"シ", "し", "shi",
		" shi:し:シ  'she' sleeps & snores゛,  Compare: tsu:ツ & so:ソ　to the more-laid-back shi:シ ",
		" shi:し:シ  し looks like a sheep hook,  Compare: ツシ  tsu-shi ",
		" TT: on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" し:シ  TT: on the 'D' char   Compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し"},

	{"グ", "ぐ", "gu",
		" gu:ぐ:グ Starting with one angle, they settled for this?  Compare:  ta:da:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ  Compare:  da:ダ, ge:ゲ    ダグゲ　",
		" TT: R-index to 'H' char, plus dakuten 濁点     ka,ki,ku,ke,ko-->ga,gi,gu,ge,go",
		" ぐ:グ  TT: R-index to 'H' char, plus dakuten 濁点   Compare:  da:ダ, ge:ゲ  ダグゲ  "},

	{"リ", "り", "ri",
		" ri:リ:り  actually, very very similar, actually!",
		" ri:リ:り Longer on the Right is Ri",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},

	{"ヒ", "ひ", "hi",
		" hi:ひ:ヒ  a smiling mouth doing a hee-hee hi-hi",
		" hi, pronounced hee ",
		" TT: left-index > to the 'V' char",
		" ?:ひ:ヒ   '?'-man   TT: left-index > to the 'V' char"},

	{"ゾ", "ぞ", "zo",
		" zo:ぞ:ゾ  it's soo big!   Compare: n:ん:ン which lays-down more",
		" zo:ぞ:ゾ  long hair and a big one,  Compare:  し:シ  &  no:ノ,  or  n:ん:ン ",
		" TT: left-index <-- to the C char, plus a dakuten 濁点 , a big one!",
		" ?:ぞ:ゾ  TT: left-index <-- to the C char, not no:ノ | n:ん:ン which lays-down more"},

	{"ベ", "べ", "be",
		" be:ベ:べ  same same, , clearly stacks of feed, hay:he stacks, ",
		" be:  he-->be-->pe , dakuten 濁点, handakuten゜半濁点",
		" TT: left-pinky way-up to the '^' char, plus dakuten 濁点",
		" TT: left-pinky way-up to the '^' char, plus dakuten 濁点   stack of feed, "},

	{"ワ", "わ", "wa",
		" wa:ワ:わ  Water-fall, pissing in the wind (making わしこ) ",
		" wa",
		" TT: right-ring to the '0' char",
		" ワ:わ  aqua-fall   TT: right-ring to the '0' char"},

	{"レ", "れ", "re",
		" re:レ:れ  laying on the little finger",
		" re:レ:れ  the れ looks like someone getting blown, laid ",
		" TT: right-pinky on it's ';' char",
		" レ:れ   TT: right-pinky on it's ';' char"},

	{"フ", "ふ", "fu",
		" fu:ふ:フ  ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" fu: think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" ふ:フ  TT: left-ring to the '2' char,  it is the big mountain in Japan "},

	{"ボ", "ぼ", "bo",
		" bo:ボ:ぼ  looks like a hoe in a dress, ho-->bo-->po  dakuten, handakuten゜半濁点",
		" bo",
		" TT: right-ring-way-up to the '-' char, plus dakuten 濁点",
		" ボ:ぼ  TT: right-ring-way-up to the '-' char, plus dakuten 濁点"},

	{"ガ", "が", "ga",
		" ga:ガ:が  same ガ albeit more angular with one-less line to draw than が",
		" ga:が is an easy one. ka--> ga,gi,gu,ge,go with a dakuten 濁点",
		" TT: index --> to the 'T' char, plus a dakuten 濁点 ",
		" ガ:が   TT: index --> to the 'T' char,  Kata has one-less line to draw"},

	{"ハ", "は", "ha",
		" ha:ハ:は:ha  looks a bit like the letter H, and a ハ ha-haystack ハ ",
		" ha: ハ ha-haystack ハ ",
		" TT: left-index on it's 'F' char",
		" ハ:は   looks a bit like the letter    TT: left-index on it's 'F' char "},

	{"プ", "ぷ", "pu",
		" pu:プ:ぷ   ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" pu:fu:  think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" プ:ぷ   TT: left-ring to the '2' char,  it is the big mountain in Japan"},

	{"コ", "こ", "ko",
		" ko:こ:コ ... it makes sense, 'cause angles:  ko:コ ,  Compare:  ni:に:ニ ",
		" ... see the two Koy fish?   'ko'oy fish? ",
		" TT: index <--- to the 'B' char,  ko way-down below  こ ",
		" こ:コ   TT: index <--- to the 'B' char    Compare:  ni:に:ニ"},
}

/*
.
.
.
.
.
.
.
.
.
.
.
. 7 Most difficult Kata:  type Romaji or Hira
. 10 Most difficult: Romaji: type Hira:  /nu, /ne, /na, /me, /ri, /ra, /a, /u, /o
*/
// var aCardD = charSetStruct{}

var fileOfCardsMostDifficult = []charSetStruct{
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
	/*

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

	*/
	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},

	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},
	/*

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

	*/
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
	/*

		{"ツ", "つ", "tsu",
			" tsu:つ:ツ  see water crashing on the she shore",
			" tsu",
			" TT: left-pinky on 'Z' char",
			" ?:つ:ツ see water crashing on the she shore  TT: left-pinky on 'Z' char"},

		{"ツ", "つ", "tsu",
			" tsu:つ:ツ  see water crashing on the she shore",
			" tsu",
			" TT: left-pinky on 'Z' char",
			" ?:つ:ツ see water crashing on the she shore  TT: left-pinky on 'Z' char"},

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

	*/
	{"オ", "お", "o",
		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" お:オ, on-the-go', cloak trailing behind   TT: index--> to the '6' char"},

	/*

		{"ン", "ん", "n"
			" n:ン:ん  looks like a cursive n, and sounds like it too",
			" n:ン:ん  ン looks like 'no' has a nose",
			" TT: index to the 'Y' char",
			" ン:ん   TT: index to the 'Y' char, sounds the way it looks"},

		{"ン", "ん", "n"
			" n:ン:ん  looks like a cursive n, and sounds like it too",
			" n:ン:ん  ン looks like 'no' has a nose",
			" TT: index to the 'Y' char",
			" ン:ん   TT: index to the 'Y' char, sounds the way it looks"},

		{"ア", "あ", "a"
			" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
			" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
			" TT: left-ring-> to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
			" Compare:  me:め  あ  nu:ぬ  ne:ね   あ-Fuck mae-te!  TT: L-middle<- to the '3' char"},

	*/
	{"ウ", "う", "u",
		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ooh:lays-atop tsu:つ  u:ウ is more-angular with a tick for its top line ",
		" TT: L-middle> to the '4' char  う : u it's the middle-Finger, or wa:ワ with a mohawk ' ",
		" う:ウ ,  TT: left-middle> to the '4' char, う:ウ not ra:ら | wa:ワ with a mohawk"},

	{"ウ", "う", "u",
		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ooh:lays-atop tsu:つ  u:ウ is more-angular with a tick for its top line ",
		" TT: L-middle> to the '4' char  う : u it's the middle-Finger, or wa:ワ with a mohawk ' ",
		" う:ウ ,  TT: left-middle> to the '4' char, う:ウ not ra:ら | wa:ワ with a mohawk"},
	//
}

/*
.
.
.
.
.
.
.
.
.
.
*/

var aCardA = charSetStruct{}

var fileOfCards = []charSetStruct{
	// Did not work :
	// const fileOfCards = []charSetStruct{

	// vowels: (a-i-u-e-o) ==========================================================================================
	// 2 of あ ア a -- 2 times for a:あ:ア
	{"ア", "あ", "a",
		" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
		" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
		" TT: left-ring-> to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
		" Compare:  me:め  あ  nu:ぬ  ne:ね   あ-Fuck mae-te!  TT: L-middle<- to the '3' char"},
	{"ア", "あ", "a",
		" a:あ:ア  ahh-'no-got-stabbed',  Compare: me:め  a:あ  nu:ぬ  ne:ね    あア-Fuck-mae-te:て!",
		" The Kata a:ア looks nothing-like the hiragana あ, but like hira te:て   あア-Fuck-mae-te:て! ",
		" TT: left-ring-> to the '3' char for あ:ア  ?is maybe a grotesque A:ア ?? ",
		" Compare:  me:め  あ  nu:ぬ  ne:ね   あ-Fuck mae-te!  TT: L-middle<- to the '3' char"},
	//
	// 2 of い イ i -- 2 times for i:い:イ
	{"イ", "い", "i",
		" i:イ:い  the Kata イ looks more like a hiragana te:て, but still two mostly-vertical lines: イ",
		" i:い:イ  for the Kata ... maybe shift the two lines of the hiragana い to forms an 'I'  i:イ",
		" TT: left-middle< to the 'E' char  い:イ  Longer on the left is i, Compare: ri:リ:り",
		" い:イ at least it is still two mostly-vertical lines   TT: L-middle< to the 'E' char"},
	{"イ", "い", "i",
		" i:イ:い  the Kata イ looks more like a hiragana te:て, but still two mostly-vertical lines: イ",
		" i:い:イ  for the Kata ... maybe shift the two lines of the hiragana い to forms an 'I'  i:イ",
		" TT: left-middle< to the 'E' char  い:イ  Longer on the left is i, Compare: ri:リ:り",
		" い:イ at least it is still two mostly-vertical lines   TT: L-middle< to the 'E' char"},
	//
	// 2 of u:う:ウ  2 times for u:う:ウ
	{"ウ", "う", "u",
		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ooh:lays-atop tsu:つ  u:ウ is more-angular with a tick for its top line ",
		" TT: L-middle> to the '4' char  う : u it's the middle-Finger, or wa:ワ with a mohawk ' ",
		" う:ウ ,  TT: left-middle> to the '4' char, う:ウ not ra:ら | wa:ワ with a mohawk"},
	{"ウ", "う", "u",
		" u:う:ウ , they just added angles for the Kata -- looks like the wa:ワ albeit with a mohawk ' ",
		" u:ooh:lays-atop tsu:つ  u:ウ is more-angular with a tick for its top line ",
		" TT: L-middle> to the '4' char  う : u it's the middle-Finger, or wa:ワ with a mohawk ' ",
		" う:ウ ,  TT: left-middle> to the '4' char, う:ウ not ra:ら | wa:ワ with a mohawk"},
	//
	// 2 of e:え:'エ'   2 times for e:え
	{"エ", "え", "e",
		" e:え:'エ' does have a vague, angular resemblance: eye see it as an, eye; may-bey, may-b-eh e ",
		" e:え ... see 'エ' as a ... an ... eye eh; may-bey, m'ay'-b-eh e   e:え:エ ",
		" TT: left-index> to the '5' char   e:え:エ ",
		" エ:え   TT: L-index> to the '5' char,  エ, eye see it as a ... an, eye ??  eh?:え:エ"},
	{"エ", "え", "e",
		" e:え:'エ' does have a vague, angular resemblance: eye see it as an, eye; may-bey, may-b-eh e ",
		" e:え ... see 'エ' as a ... an ... eye eh; may-bey, m'ay'-b-eh e   e:え:エ ",
		" TT: left-index> to the '5' char   e:え:エ ",
		" エ:え   TT: L-index> to the '5' char,  エ, eye see it as a ... an, eye ??  eh?:え:エ"},
	//
	// 2 of o:お:オ
	{"オ", "お", "o",
		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" お:オ, on-the-go', cloak trailing behind   TT: index--> to the '6' char"},
	{"オ", "お", "o",
		" o:お:オ , and オ does have a vague resemblance, albeit with less curves: お オ",
		" o:お , for the Kata オ ... think: someone is 'on-the-go' ... see the fl'o'w ",
		" TT: index--> to the '6' char     o:お:オ ",
		" お:オ, on-the-go', cloak trailing behind   TT: index--> to the '6' char"},
	//
	// ka group: (ka-ga) ============================================================================================
	// 2 ka
	{"カ", "か", "ka",
		" ka:か:カ ... is an easy one!  It simply looks like a 'K' ",
		" ... Kata is the same カ , albeit more-angular and with one-less line to draw than か ",
		" TT: left-index--> to the 'T' char   か or カ",
		" か:カ ... looks like a 'K'    TT: left-index--> to the 'T' char"},
	{"カ", "か", "ka",
		" ka:か:カ ... is an easy one!  It simply looks like a 'K' ",
		" ... Kata is the same カ , albeit more-angular and with one-less line to draw than か ",
		" TT: left-index--> to the 'T' char   か or カ",
		" か:カ ... looks like a 'K'    TT: left-index--> to the 'T' char"},
	//
	// 2 of ki, ku, ke, ko
	{"キ", "き", "ki",
		" ki:き:キ ... is an easy one!   キ has the same top as ki:き, and they look like 'keys' ",
		" ki:き ... they both have the same top:  き キ  looks like a chi-key",
		" TT: L-index> to the 'G' char,  Think: ki-of-G  き,  Compare:  sa:さ-ki:き  さき ... saki ",
		"き:キ  TT: L-index > to the 'G' char,  キ same top as き, looks like a 'chi-key'"},
	{"キ", "き", "ki",
		" ki:き:キ ... is an easy one!   キ has the same top as ki:き, and they look like 'keys' ",
		" ki:き ... they both have the same top:  き キ  looks like a chi-key",
		" TT: L-index> to the 'G' char,  Think: ki-of-G  き,  Compare:  sa:さ-ki:き  さき ... saki ",
		"き:キ  TT: L-index > to the 'G' char,  キ same top as き, looks like a 'chi-key'"},
	// 2 ku
	{"ク", "く", "ku",
		" ku:く:ク ,  Kata: starting with one angle, they settled for this??  く:ク ",
		" ku:く:ク ,  Compare:  ta:タ  ke:ケ  ku:ク  クケタ,  it's kuku, kookoo I tell you! ",
		" TT: R-index<- to the 'H' char  く:ク ",
		"く:ク ,  Compare:  クケタ ta:タ  ke:ケ  ?:ク   TT: R-index<- to the 'H' char"},
	{"ク", "く", "ku",
		" ku:く:ク ,  Kata: starting with one angle, they settled for this??  く:ク ",
		" ku:く:ク ,  Compare:  ta:タ  ke:ケ  ku:ク  クケタ,  it's kuku, kookoo I tell you! ",
		" TT: R-index<- to the 'H' char  く:ク ",
		"く:ク ,  Compare:  クケタ ta:タ  ke:ケ  ?:ク   TT: R-index<- to the 'H' char"},
	// 2 ke
	{"ケ", "け", "ke",
		" ke:け:ケ  bits of it are there, just as many curves though ,  Compare to:  ku:ク  ta:タ ",
		" ke:け  ... looks enough like a Keg ...  Foam-ment THAT image in your brain ",
		" TT: right-pinky-> one-over to the ':*' chars   け:ケ ",
		"け:ケ  bits still there  Compare:  ku:ク  ta:タ  クケタ  TT: R-pinky-> to the ':*' chars"},
	{"ケ", "け", "ke",
		" ke:け:ケ  bits of it are there, just as many curves though ,  Compare to:  ku:ク  ta:タ ",
		" ke:け  ... looks enough like a Keg ...  Foam-ment THAT image in your brain ",
		" TT: right-pinky-> one-over to the ':*' chars   け:ケ ",
		"け:ケ  bits still there  Compare:  ku:ク  ta:タ  クケタ  TT: R-pinky-> to the ':*' chars"},
	// 2 ko   2 times for ko:こ:コ
	{"コ", "こ", "ko",
		" ko:こ:コ ... it makes sense, 'cause angles:  ko:コ ,  Compare:  ni:に:ニ ",
		" ... see the two Koy fish?   'ko'oy fish? ",
		" TT: index <--- to the 'B' char,  ko way-down below  こ ",
		" こ:コ   TT: index <--- to the 'B' char    Compare:  ni:に:ニ"},
	{"コ", "こ", "ko",
		" ko:こ:コ ... it makes sense, 'cause angles:  ko:コ ,  Compare:  ni:に:ニ ",
		" ... see the two Koy fish?   'ko'oy fish? ",
		" TT: index <--- to the 'B' char,  ko way-down below  こ ",
		" こ:コ   TT: index <--- to the 'B' char    Compare:  ni:に:ニ"},
	//
	// ka becomes ga group: -----------------------------------------------------------------------------------------
	// 2 ga
	{"ガ", "が", "ga",
		" ga:ガ:が  same ガ albeit more angular with one-less line to draw than が",
		" ga:が is an easy one. ka--> ga,gi,gu,ge,go with a dakuten 濁点",
		" TT: index --> to the 'T' char, plus a dakuten 濁点 ",
		" ガ:が   TT: index --> to the 'T' char,  Kata has one-less line to draw"},
	{"ガ", "が", "ga",
		" ga:ガ:が  same ガ albeit more angular with one-less line to draw than が",
		" ga:が is an easy one. ka--> ga,gi,gu,ge,go with a dakuten 濁点",
		" TT: index --> to the 'T' char, plus a dakuten 濁点 ",
		" ガ:が   TT: index --> to the 'T' char,  Kata has one-less line to draw"},
	// 2 gi
	{"ギ", "ぎ", "gi",
		" gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し ",
		" ぎ:ギ is always the sound gi, and is Seldom ji ( that usually being じ:ji:ジ )",
		" TT: L-index on 'G' char, plus a dakuten 濁点, ぎ:ギ, compare to za:ざ from sa:さ",
		" ぎ:ギ  TT: 'G' char,  sound is from き ,and NEVER from し or ち  ka-->ga,?,gu,ge,go"},
	{"ギ", "ぎ", "gi",
		" gi:ぎ:ギ is an easy Kata, ギ has the same top as ki:き, and is NEVER from shi:し ",
		" ぎ:ギ is always the sound gi, and is Seldom ji ( that usually being じ:ji:ジ )",
		" TT: L-index on 'G' char, plus a dakuten 濁点, ぎ:ギ, compare to za:ざ from sa:さ",
		" ぎ:ギ  TT: 'G' char,  sound is from き ,and NEVER from し or ち  ka-->ga,?,gu,ge,go"},
	//
	// 4 times for gu:ぐ:グ
	{"グ", "ぐ", "gu",
		" gu:ぐ:グ Starting with one angle, they settled for this?  Compare:  ta:da:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ  Compare:  da:ダ, ge:ゲ    ダグゲ　",
		" TT: R-index to 'H' char, plus dakuten 濁点     ka,ki,ku,ke,ko-->ga,gi,gu,ge,go",
		" ぐ:グ  TT: R-index to 'H' char, plus dakuten 濁点   Compare:  da:ダ, ge:ゲ  ダグゲ  "},
	{"グ", "ぐ", "gu",
		" gu:ぐ:グ Starting with one angle, they settled for this?  Compare:  ta:da:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ  Compare:  da:ダ, ge:ゲ    ダグゲ　",
		" TT: R-index to 'H' char, plus dakuten 濁点     ka,ki,ku,ke,ko-->ga,gi,gu,ge,go",
		" ぐ:グ  TT: R-index to 'H' char, plus dakuten 濁点   Compare:  da:ダ, ge:ゲ  ダグゲ  "},
	{"グ", "ぐ", "gu",
		" gu:ぐ:グ Starting with one angle, they settled for this?  Compare:  ta:da:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ  Compare:  da:ダ, ge:ゲ    ダグゲ　",
		" TT: R-index to 'H' char, plus dakuten 濁点     ka,ki,ku,ke,ko-->ga,gi,gu,ge,go",
		" ぐ:グ  TT: R-index to 'H' char, plus dakuten 濁点   Compare:  da:ダ, ge:ゲ  ダグゲ  "},
	{"グ", "ぐ", "gu",
		" gu:ぐ:グ Starting with one angle, they settled for this?  Compare:  ta:da:ダ, and ke:ge:ゲ",
		" gu:ぐ:グ  Compare:  da:ダ, ge:ゲ    ダグゲ　",
		" TT: R-index to 'H' char, plus dakuten 濁点     ka,ki,ku,ke,ko-->ga,gi,gu,ge,go",
		" ぐ:グ  TT: R-index to 'H' char, plus dakuten 濁点   Compare:  da:ダ, ge:ゲ  ダグゲ  "},
	//
	// 4 times for ge:げ:ゲ   times for ge:げ:ゲ
	{"ゲ", "げ", "ge",
		" ge:げ:ゲ bits still there (as many curves though)  Compare  ku:gu:グ,  &  ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ   don't forget the dakuten 濁点",
		" TT: ':*' chars, 濁点    げ:ゲ  ダグゲ    ka,ki,ku,ke,ko-->ga,gi,gu,?,go  濁点",
		" げ:ゲ   TT: L-pinky> to ':*' chars   げ Compare to ku:gu:グ, and ta:da:ダ "},
	{"ゲ", "げ", "ge",
		" ge:げ:ゲ bits still there (as many curves though)  Compare  ku:gu:グ,  &  ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ   don't forget the dakuten 濁点",
		" TT: ':*' chars, 濁点    げ:ゲ  ダグゲ    ka,ki,ku,ke,ko-->ga,gi,gu,?,go  濁点",
		" げ:ゲ   TT: L-pinky> to ':*' chars   げ Compare to ku:gu:グ, and ta:da:ダ "},
	{"ゲ", "げ", "ge",
		" ge:げ:ゲ bits still there (as many curves though)  Compare  ku:gu:グ,  &  ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ   don't forget the dakuten 濁点",
		" TT: ':*' chars, 濁点    げ:ゲ  ダグゲ    ka,ki,ku,ke,ko-->ga,gi,gu,?,go  濁点",
		" げ:ゲ   TT: L-pinky> to ':*' chars   げ Compare to ku:gu:グ, and ta:da:ダ "},
	{"ゲ", "げ", "ge",
		" ge:げ:ゲ bits still there (as many curves though)  Compare  ku:gu:グ,  &  ta:za:ダ",
		" ge:げ Compare to ku:gu:グ, and ta:za:ダ   don't forget the dakuten 濁点",
		" TT: ':*' chars, 濁点    げ:ゲ  ダグゲ    ka,ki,ku,ke,ko-->ga,gi,gu,?,go  濁点",
		" げ:ゲ   TT: L-pinky> to ':*' chars   げ Compare to ku:gu:グ, and ta:da:ダ "},
	//
	// 3 go
	{"ゴ", "ご", "go",
		" go:ご:ゴ   looks like two Koy fish ko-y, and as ka-->ga, ko-->go ",
		" go:ご:ゴ    ご:ゴ makes sense, 'cause angles'",
		" TT: index <--- to the 'B' char, plus a dakuten 濁点  ご",
		" ご:ゴ  TT: index <--- to the 'B' char, plus a dakuten 濁点 ka,ki,ku,ke,ko-->ga,gi,gu,ge,?"},
	{"ゴ", "ご", "go",
		" go:ご:ゴ   looks like two Koy fish ko-y, and as ka-->ga, ko-->go ",
		" go:ご:ゴ    ご:ゴ makes sense, 'cause angles'",
		" TT: index <--- to the 'B' char, plus a dakuten 濁点  ご",
		" ご:ゴ  TT: index <--- to the 'B' char, plus a dakuten 濁点 ka,ki,ku,ke,ko-->ga,gi,gu,ge,?"},
	{"ゴ", "ご", "go",
		" go:ご:ゴ   looks like two Koy fish ko-y, and as ka-->ga, ko-->go ",
		" go:ご:ゴ    ご:ゴ makes sense, 'cause angles'",
		" TT: index <--- to the 'B' char, plus a dakuten 濁点  ご",
		" ご:ゴ  TT: index <--- to the 'B' char, plus a dakuten 濁点 ka,ki,ku,ke,ko-->ga,gi,gu,ge,?"},
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
	// sa group: (sa-za,ji,zu,za,zo) ================================================================================
	// 2 of each [2 sets]: sa, shi, su, se ***
	{"サ", "さ", "sa",
		" sa:さ:サ , if you 'se':せ そサ ,  Compare:  Hira se:せ  to  Kata sa:サ  せサ ",
		" ... , but it looks an awful-lot like Hira se:せ, I'd say   ... just 'Hir'-say ?? ",
		" TT: ring>v to the 'X' char  --  at least Kata 'sa' still goes to the left ... ",
		" さ:サ ,  Compare:  Hira:se:せ  to  Kata:?:サ  せサ se:? TT: ring>v to the 'X' char"},
	{"シ", "し", "shi",
		" shi:し:シ  'she' sleeps & snores゛,  Compare: tsu:ツ & so:ソ　to the more-laid-back shi:シ ",
		" shi:し:シ  し looks like a sheep hook,  Compare: ツシ  tsu-shi ",
		" TT: on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" し:シ  TT: on the 'D' char   Compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し"},
	{"ス", "す", "su",
		" su:す:ス ",
		" ス sue:す looks like she is running (and they were looking for angles)",
		" TT: left-index < to the 'R' char,  (sue is running ス)",
		" ス:す  she is running  TT: left-index < to the 'R' char"},
	{"セ", "せ", "se",
		" se:せ:セ  ... セ is just a bit more angular, as is the way with most Katakana ",
		" se:せ looks like a face saying something ... 'se'ing something ",
		" TT: right-pinky to the 'P' char",
		" ?:せ:セ  TT: R-pinky to the 'P' char,  セ is more angular, as with most Katakana"},
	//
	// Second set ::sa, shi, su, se ***
	{"サ", "さ", "sa",
		" sa:さ:サ , if you 'se':せ そサ ,  Compare:  Hira se:せ  to  Kata sa:サ  せサ ",
		" ... , but it looks an awful-lot like Hira se:せ, I'd say   ... just 'Hir'-say ?? ",
		" TT: ring>v to the 'X' char  --  at least Kata 'sa' still goes to the left ... ",
		" さ:サ ,  Compare:  Hira:se:せ  to  Kata:?:サ  せサ se:? TT: ring>v to the 'X' char"},
	{"シ", "し", "shi",
		" shi:し:シ  'she' sleeps & snores゛,  Compare: tsu:ツ & so:ソ　to the more-laid-back shi:シ ",
		" shi:し:シ  し looks like a sheep hook,  Compare: ツシ  tsu-shi ",
		" TT: on the 'D' char, シ 'she' is sleeping/snoring,  (no angles here, less curve though)",
		" し:シ  TT: on the 'D' char   Compare: tsu:ツ  & so:ソ　to the more-laid-back シ:し"},
	{"ス", "す", "su",
		" su:す:ス ",
		" ス sue:す looks like she is running (and they were looking for angles)",
		" TT: left-index < to the 'R' char,  (sue is running ス)",
		" ス:す  she is running  TT: left-index < to the 'R' char"},
	{"セ", "せ", "se",
		" se:せ:セ  ... セ is just a bit more angular, as is the way with most Katakana ",
		" se:せ looks like a face saying something ... 'se'ing something ",
		" TT: right-pinky to the 'P' char",
		" ?:せ:セ  TT: R-pinky to the 'P' char,  セ is more angular, as with most Katakana"},
	//
	//  4 times so:そ:ソ
	// Seriously?  so:ソ vs nh:ン  ... I guess nh got nh-ocked-down by an upper-cut to the chin
	{"ソ", "そ", "so",
		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},
	{"ソ", "そ", "so",
		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},
	{"ソ", "そ", "so",
		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},
	{"ソ", "そ", "so",
		" so:そ:ソ  starting with all angles -- I guess they `had to `backtrack? ",
		" so:そ:ソ ,  Compare:  n:ん:ン which lays-down more, and has a more-laid-down back-tick ",
		" TT: index<-- to the 'C' char,  so:ソ Looks like: し:シ  & no:ノ but MOSTLY like:  n:ん:ン ",
		" ?:そ:ソ  ,  Compare:  n:ん:ン     TT: L-index<-- to the 'C' char"},
	//
	// sa group becomes za, ji, zu, ze, zo ----------------------------------------------------------------------------
	// -- and again we have two sets
	{"ザ", "ざ", "za",
		" sa-->za:ざ  dakuten  濁点",
		" sa-->za   za:ザ:ざ  so it still protrudes to the left",
		" TT: down on the 'X' char plus a dakuten 濁点 for sa-->za",
		" ?:ざ:ザ  TT: down on the 'X' char, plus a dakuten 濁点"},
	{"ジ", "じ", "ji",
		" ji:ジ:じ    Seldom is it fr chi:ち ,it's the sound ji, NEVER gi (that being ぎ:gi:ギ)",
		" ji, the sound, VERY RARELY is from chi:ち:ヂ ,but NEVER the sound gi (that being ぎ:gi:ギ)",
		" TT: 'D' char, remember, the sound:ji is nearly-always from shi:し with a dakuten 濁点 じ:ジ",
		" ジ:じ  TT: 'D' char, with a dakuten 濁点     it's the sound:gee:jeep"},
	// zu has two Hira values ず, づ
	{"ズ", "ず", "zu",
		" zu:ず:ズ:づ:ヅ: , is either ず:ズ or づ:ヅ , as they are both the same sound",
		" zu: 'Sue' ズ is running with her hair flowing behind her (from su:す) ",
		"  TT: L-index< to the 'R' char, same sound: a big wave つ  or a  snake ず at the zoo",
		" ?:ず:ズ づ:ヅ  same sound   TT: left-index< to the 'R' char, dakuten 濁点 "},
	{"ゼ", "ぜ", "ze",
		" sa-->za   se-->ze , plus a dakuten 濁点",
		" ze:ゼ:ぜ",
		" TT: pinky to the 'P' char, plus a dakuten 濁点   ze:ゼ:ぜ",
		" ?:ゼ:ぜ   TT: pinky to the 'P' char, plus a dakuten 濁点   sa-->za  se:? "},
	// 3 zo
	{"ゾ", "ぞ", "zo",
		" zo:ぞ:ゾ  it's soo big!   Compare: n:ん:ン which lays-down more",
		" zo:ぞ:ゾ  long hair and a big one,  Compare:  し:シ  &  no:ノ,  or  n:ん:ン ",
		" TT: left-index <-- to the C char, plus a dakuten 濁点 , a big one!",
		" ?:ぞ:ゾ  TT: left-index <-- to the C char, not no:ノ | n:ん:ン which lays-down more"},
	{"ゾ", "ぞ", "zo",
		" zo:ぞ:ゾ  it's soo big!   Compare: n:ん:ン which lays-down more",
		" zo:ぞ:ゾ  long hair and a big one,  Compare:  し:シ  &  no:ノ,  or  n:ん:ン ",
		" TT: left-index <-- to the C char, plus a dakuten 濁点 , a big one!",
		" ?:ぞ:ゾ  TT: left-index <-- to the C char, not no:ノ | n:ん:ン which lays-down more"},
	{"ゾ", "ぞ", "zo",
		" zo:ぞ:ゾ  it's soo big!   Compare: n:ん:ン which lays-down more",
		" zo:ぞ:ゾ  long hair and a big one,  Compare:  し:シ  &  no:ノ,  or  n:ん:ン ",
		" TT: left-index <-- to the C char, plus a dakuten 濁点 , a big one!",
		" ?:ぞ:ゾ  TT: left-index <-- to the C char, not no:ノ | n:ん:ン which lays-down more"},
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
	// ta group:  ============================================================================================
	//
	// 2 of each ta,chi,tsu,te,to
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
	{"ト", "と", "to",
		" to:と:ト   Kata 'toe' is flipped-out. Kicked in the balls. on its head: と -> ト ",
		" Katakana to:と is flipped-out. Kicked in the balls, On its head ",
		" TT: left-ring on the 'S' char, Kata toe is flipped-out. Kicked on its head: と -> ト ",
		" Kata char was kicked on its head: と --> ト  TT: L-ring on the 'S' char"},
	//
	// Second set:
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
	{"ト", "と", "to",
		" to:と:ト   Kata 'toe' is flipped-out. Kicked in the balls. on its head: と -> ト ",
		" Katakana to:と is flipped-out. Kicked in the balls, On its head ",
		" TT: left-ring on the 'S' char, Kata toe is flipped-out. Kicked on its head: と -> ト ",
		" Kata char was kicked on its head: と --> ト  TT: L-ring on the 'S' char"},
	//
	// ta group becomes da, ji, zu, de, do ----------------------------------------------------------------------------
	//
	// 2 da
	{"ダ", "だ", "da",
		" ta-->da  だ:ダ ",
		" da is ta with a dakuten 濁点",
		" TT: left-pinky to the 'Q' char",
		" だ:ダ  ta-->?  TT: left-pinky to the 'Q' char, plus a dakuten 濁点"},
	{"ダ", "だ", "da",
		" ta-->da  だ:ダ ",
		" da is ta with a dakuten 濁点",
		" TT: left-pinky to the 'Q' char",
		" だ:ダ  ta-->?  TT: left-pinky to the 'Q' char, plus a dakuten 濁点"},
	// 2 zu
	// : zu has 2 Hira values: づ, ず
	{"ヅ", "づ", "zu",
		" zu:づ:ヅ ず:ズ , is either ず:ズ or づ:ヅ , as they are both the same sound ",
		" zu: for the Kataす think 'Sue' ズ is running with her hair flowing behind her",
		" TT: L-pinky>v to the 'Z' char, plus dakuten 濁点 ",
		" ?:づ:ヅ ず:ズ  that big wave つ sounds the same, TT: L-pinky>v to the 'Z' char"},
	// : zu has 2 Hira values: づ, ず
	{"ヅ", "づ", "zu",
		" zu:づ:ヅ ず:ズ , is either ず:ズ or づ:ヅ , as they are both the same sound ",
		" zu: for the Kataす think 'Sue' ズ is running with her hair flowing behind her",
		" TT: L-pinky>v to the 'Z' char, plus dakuten 濁点 ",
		" ?:づ:ヅ ず:ズ  that big wave つ sounds the same, TT: L-pinky>v to the 'Z' char"},
	// 2 de
	{"デ", "で", "de",
		" de:で:デ ",
		" de from ta-->da , plus a dakuten 濁点",
		" TT: left-ring to the 'W' char, plus a dakuten 濁点",
		" ?:で:デ  , plus a dakuten, TT: L-ring to the 'W' char, plus a dakuten 濁点"},
	{"デ", "で", "de",
		" de:で:デ ",
		" de from ta-->da , plus a dakuten 濁点",
		" TT: left-ring to the 'W' char, plus a dakuten 濁点",
		" ?:で:デ  , plus a dakuten, TT: L-ring to the 'W' char, plus a dakuten 濁点"},
	// 2 do
	{"ド", "ど", "do",
		" do:ド:ど   ta-->da  to --> do    plus a dakuten 濁点",
		" do:ど:ド ",
		" TT: left-ring on the 'S' char, plus a dakuten 濁点 ",
		" ?:ど:ド  TT: left-ring on the 'S' char, plus a dakuten 濁点   ta-->da  to-->? "},
	{"ド", "ど", "do",
		" do:ド:ど   ta-->da  to --> do    plus a dakuten 濁点",
		" do:ど:ド ",
		" TT: left-ring on the 'S' char, plus a dakuten 濁点 ",
		" ?:ど:ド  TT: left-ring on the 'S' char, plus a dakuten 濁点   ta-->da  to-->? "},
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
	// na group: (always a naked group) =============================================================================
	// 3 na
	{"ナ", "な", "na",
		" na:な:ナ  Kata ナ is just like the left side of the Hira な -- な:ナ ",
		" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  ( め ぬ )",
		" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
		" な:ナ  TT: index< to the 'U' char, Hiragana t-ies a knot, see the two chars thar "},
	{"ナ", "な", "na",
		" na:な:ナ  Kata ナ is just like the left side of the Hira な -- な:ナ ",
		" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  ( め ぬ )",
		" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
		" な:ナ  TT: index< to the 'U' char, Hiragana t-ies a knot, see the two chars thar "},
	{"ナ", "な", "na",
		" na:な:ナ  Kata ナ is just like the left side of the Hira な -- な:ナ ",
		" na:な:ナ , Compare:  na:ナ  me:メ  nu:ヌ  ( め ぬ )",
		" TT: index< to the 'U' char,  ナ is very simple ... na-t complex at-all :: な ",
		" な:ナ  TT: index< to the 'U' char, Hiragana t-ies a knot, see the two chars thar "},
	// 1 ni
	{"ニ", "に", "ni",
		" ni:に:ニ  I can see a knee-cap in the Hiragana ",
		" ni:に:ニ  ",
		" TT: right-middle< to the 'I' char ",
		" ?:に:ニ  TT: R-middle< to the 'I' char,  I can see a knee-cap in the Hiragana "},
	// 2 each of: nu,ne,no
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
	{"ノ", "の", "no",
		" no:の:ノ   it looks like a 'no' symbol ",
		" no:の:ノ  and the Kata retains the slash ",
		" TT: right-middle on 'K' char",
		" ?:の:ノ  TT: right-middle on 'K' char   ?-thank-you "},
	// second set:
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
	{"ノ", "の", "no",
		" no:の:ノ   it looks like a 'no' symbol ",
		" no:の:ノ  and the Kata retains the slash ",
		" TT: right-middle on 'K' char",
		" ?:の:ノ  TT: right-middle on 'K' char   ?-thank-you "},
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
	// ha group: ha:hi:fu:he:ho =============================================================================
	// 2 ha
	{"ハ", "は", "ha",
		" ha:ハ:は:ha  looks a bit like the letter H, and a ハ ha-haystack ハ ",
		" ha: ハ ha-haystack ハ ",
		" TT: left-index on it's 'F' char",
		" ハ:は   looks a bit like the letter    TT: left-index on it's 'F' char "},
	{"ハ", "は", "ha",
		" ha:ハ:は:ha  looks a bit like the letter H, and a ハ ha-haystack ハ ",
		" ha: ハ ha-haystack ハ ",
		" TT: left-index on it's 'F' char",
		" ハ:は   looks a bit like the letter    TT: left-index on it's 'F' char "},
	// 1 hi
	{"ヒ", "ひ", "hi",
		" hi:ひ:ヒ  a smiling mouth doing a hee-hee hi-hi",
		" hi, pronounced hee ",
		" TT: left-index > to the 'V' char",
		" ?:ひ:ヒ   '?'-man   TT: left-index > to the 'V' char"},
	// 3 fu
	{"フ", "ふ", "fu",
		" fu:ふ:フ  ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" fu: think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" ふ:フ  TT: left-ring to the '2' char,  it is the big mountain in Japan "},
	{"フ", "ふ", "fu",
		" fu:ふ:フ  ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" fu: think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" ふ:フ  TT: left-ring to the '2' char,  it is the big mountain in Japan "},
	{"フ", "ふ", "fu",
		" fu:ふ:フ  ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" fu: think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" ふ:フ  TT: left-ring to the '2' char,  it is the big mountain in Japan "},
	// 1 he
	{"ヘ", "へ", "he",
		" he:ヘ:へ  Hira and Kata are pretty-much the same, both hay:he stacks",
		" he",
		" TT: left-pinky way-up to the '^' char",
		" ヘ:へ   TT: left-pinky way-up to the '^' char, clearly stacks of feed"},
	// 2 ho
	{"ホ", "ほ", "ho",
		" ho:ホ:ほ  looks like a hoe in a dress",
		" ho",
		" TT: right-ring-way-up to the '-' char",
		" ?:ホ:ほ   TT: right-ring-way-up to the '-' char"},
	{"ホ", "ほ", "ho",
		" ho:ホ:ほ  looks like a hoe in a dress",
		" ho",
		" TT: right-ring-way-up to the '-' char",
		" ?:ホ:ほ   TT: right-ring-way-up to the '-' char"},
	// ha group becomes ba group:
	// ba, bi, bu, be, bo -------------------------------------------------------------------------------------------
	// 2 ba
	{"バ", "ば", "ba",
		" ba:バ:ば  ha-->ba-->pa   dakuten 濁点, handakuten゜半濁点",
		" ba",
		" TT: L-index on the 'F' char, plus a dakuten 濁点 ",
		" ?:バ:ば  ha-->?  TT: L-index on the 'F' char, plus a dakuten 濁点"},
	{"バ", "ば", "ba",
		" ba:バ:ば  ha-->ba-->pa   dakuten 濁点, handakuten゜半濁点",
		" ba",
		" TT: L-index on the 'F' char, plus a dakuten 濁点 ",
		" ?:バ:ば  ha-->?  TT: L-index on the 'F' char, plus a dakuten 濁点"},
	// 1 bi
	{"ビ", "び", "bi",
		" bi:ビ:び   hi-bi-pi ",
		" bi is from hi which looks like a smiling hi-hi びび",
		" TT: L-index to the 'V' char, plus a dakuten 濁点",
		" ?:ビ:び from hi, looks like smiling hi-hi びび   TT: L-index to 'V' char, plus a dakuten"},
	// 3 bu
	{"ブ", "ぶ", "bu",
		" bu:ブ:ぶ",
		" fu-->bu , plus a dakuten 濁点",
		" TT: left-ring to the '2' char is fu-->bu",
		" ブ:ぶ   TT: left-ring to the '2' char, plus a dakuten 濁点 :: fu-->?"},
	{"ブ", "ぶ", "bu",
		" bu:ブ:ぶ",
		" fu-->bu , plus a dakuten゛濁点",
		" TT: left-ring to the '2' char is fu-->bu",
		" ブ:ぶ   TT: left-ring to the '2' char, plus a dakuten 濁点 :: fu-->?"},
	{"ブ", "ぶ", "bu",
		" bu:ブ:ぶ",
		" fu-->bu , plus a dakuten゛濁点",
		" TT: left-ring to the '2' char is fu-->bu",
		" ブ:ぶ   TT: left-ring to the '2' char, plus a dakuten 濁点 :: fu-->?"},
	// 1 be
	{"ベ", "べ", "be",
		" be:ベ:べ  same same, , clearly stacks of feed, hay:he stacks, ",
		" be:  he-->be-->pe , dakuten゛濁点, handakuten゜半濁点",
		" TT: left-pinky way-up to the '^' char, plus dakuten 濁点",
		" TT: left-pinky way-up to the '^' char, plus dakuten 濁点   stack of feed, "},
	// 2 bo
	{"ボ", "ぼ", "bo",
		" bo:ボ:ぼ  looks like a hoe in a dress, ho-->bo-->po  dakuten゛濁点, handakuten゜半濁点",
		" bo",
		" TT: right-ring-way-up to the '-' char, plus a dakuten゛濁点",
		" ボ:ぼ  TT: right-ring-way-up to the '-' char, plus dakuten 濁点"},
	{"ボ", "ぼ", "bo",
		" bo:ボ:ぼ  looks like a hoe in a dress, ho-->bo-->po  dakuten゛濁点, handakuten゜半濁点",
		" bo",
		" TT: right-ring-way-up to the '-' char, plus dakuten 濁点",
		" ボ:ぼ  TT: right-ring-way-up to the '-' char, plus dakuten 濁点"},
	//
	// pa, pi, pu, pe, po -------------------------------------------------------------------------------------------
	// 2 of each :: pa, pi, pu, pe, po
	{"パ", "ぱ", "pa",
		" pa:パ:ぱ  ha-->ba-->pa   dakuten゛濁点, handakuten゜半濁点",
		" pa",
		" TT: left-index on the 'F' char, plus a handakuten 半濁点",
		" ha-->?  TT: L-index on the 'F' char, plus a handakuten 半濁点"},
	{"ピ", "ぴ", "pi",
		" pi:ピ:ぴ  ",
		" hi-->pi  becomes with a handakuten゜半濁点",
		" TT: L-index to the 'V' char, plus a handakuten゜半濁点",
		" ?:ピ:ぴ   TT: L-index to the 'V' char, plus a handakuten゜半濁点"},
	{"プ", "ぷ", "pu",
		" pu:プ:ぷ   ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" pu:fu:  think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" プ:ぷ   TT: left-ring to the '2' char,  it is the big mountain in Japan"},
	{"ペ", "ぺ", "pe",
		" pe:ペ:ぺ  same same, , clearly stacks of feed, hay:he stacks, handakuten゜半濁点",
		" pe:  he-->be-->pe , dakuten゛濁点, handakuten゜半濁点",
		" TT: left-pinky way-up to the '^' char, plus handakuten゜半濁点",
		" ペ:ぺ   TT: left-pinky way-up to the '^' char, stack of feed, handakuten゜半濁点"},
	{"ポ", "ぽ", "po",
		" po:ポ:ぽ  ho-->bo-->po  dakuten゛濁点, handakuten゜半濁点",
		" po",
		" TT: right-ring-way-up to the '-' char, plus a handakuten゜半濁点",
		" ?:ポ:ぽ   TT: right-ring-way-up to the '-' char, plus a handakuten゜半濁点"},
	//
	// Second set:
	{"パ", "ぱ", "pa",
		" pa:パ:ぱ  ha-->ba-->pa   dakuten゛濁点, handakuten゜半濁点",
		" pa",
		" TT: left-index on the 'F' char, plus a handakuten 半濁点",
		" ha-->?  TT: L-index on the 'F' char, plus a handakuten 半濁点"},
	{"ピ", "ぴ", "pi",
		" pi:ピ:ぴ  ",
		" hi-->pi  becomes with a handakuten゜半濁点",
		" TT: L-index to the 'V' char, plus a handakuten゜半濁点",
		" ?:ピ:ぴ   TT: L-index to the 'V' char, plus a handakuten゜半濁点"},
	{"プ", "ぷ", "pu",
		" pu:プ:ぷ   ふじ　Mount Fuji 富士山  富:Fu: wealth, abundant  士:Ji: gentleman, samurai, warrior",
		" pu:fu:  think Mt Fuji 'fu'ji,  fu or, alternatively, hu ",
		" TT: left-ring to the '2' char",
		" プ:ぷ   TT: left-ring to the '2' char,  it is the big mountain in Japan"},
	{"ペ", "ぺ", "pe",
		" pe:ペ:ぺ  same same, , clearly stacks of feed, hay:he stacks, handakuten゜半濁点",
		" pe:  he-->be-->pe , dakuten゛濁点, handakuten゜半濁点",
		" TT: left-pinky way-up to the '^' char, plus handakuten゜半濁点",
		" ペ:ぺ   TT: left-pinky way-up to the '^' char, stack of feed, handakuten゜半濁点"},
	{"ポ", "ぽ", "po",
		" po:ポ:ぽ  ho-->bo-->po  dakuten゛濁点, handakuten゜半濁点",
		" po",
		" TT: right-ring-way-up to the '-' char, plus a handakuten゜半濁点",
		" ?:ポ:ぽ   TT: right-ring-way-up to the '-' char, plus a handakuten゜半濁点"},
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
	// ma group: ====================================================================================================
	// 2 of each ::
	{"マ", "ま", "ma",
		" ma:マ:ま ",
		" mama is using a brest pump ",
		" TT: right-index on the 'J' char :: ma",
		" ?:マ:ま   lady is using a brest pump    TT: right-index on the 'J' char"},
	{"ミ", "み", "mi",
		" mi:ミ:み  mi is three, or phallic 'me' is mi ",
		" mi is three; lines",
		" TT: right-index to the 'N' char",
		" ミ:み   TT: right-index to the 'N' char"},
	{"ム", "む", "mu",
		" mu:ム:む  somehow, it does look like a cow",
		" mu",
		" TT: right-pinky way-over to the '}]' chars",
		" ム:む   TT: right-pinky way-over to the '}]' chars"},
	{"メ", "め", "me",
		" me:メ:め   Compare: nu:ヌ",
		" me",
		" TT: right-pinky-slide down to the '?' char",
		" メ:め messy-lesser ver of nuddle:ぬ　 TT: R-pinky-slide down to the '?' char"},
	{"モ", "も", "mo",
		" mo:モ:も ",
		" mo",
		" TT: right-middle<v to the 'M' char",
		" モ:も   TT: right-middle<v to the 'M' char,  m is for ?"},
	//
	// Second set:
	{"マ", "ま", "ma",
		" ma:マ:ま ",
		" mama is using a brest pump ",
		" TT: right-index on the 'J' char :: ma",
		" ?:マ:ま   lady is using a brest pump    TT: right-index on the 'J' char"},
	{"ミ", "み", "mi",
		" mi:ミ:み  mi is three, or phallic 'me' is mi ",
		" mi is three; lines",
		" TT: right-index to the 'N' char",
		" ミ:み   TT: right-index to the 'N' char"},
	{"ム", "む", "mu",
		" mu:ム:む  somehow, it does look like a cow",
		" mu",
		" TT: right-pinky way-over to the '}]' chars",
		" ム:む   TT: right-pinky way-over to the '}]' chars"},
	{"メ", "め", "me",
		" me:メ:め   Compare: nu:ヌ",
		" me",
		" TT: right-pinky-slide down to the '?' char",
		" メ:め messy-lesser ver of nuddle:ぬ　 TT: R-pinky-slide down to the '?' char"},
	// 2 mo's
	{"モ", "も", "mo",
		" mo:モ:も ",
		" mo",
		" TT: right-middle<v to the 'M' char",
		" モ:も   TT: right-middle<v to the 'M' char,  m is for ?"},
	{"モ", "も", "mo",
		" mo:モ:も ",
		" mo",
		" TT: right-middle<v to the 'M' char",
		" モ:も   TT: right-middle<v to the 'M' char,  m is for ?"},
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
	// ra group: (pronounced more like 'la') ========================================================================
	// 4 ra
	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},
	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},
	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},
	{"ラ", "ら", "ra",
		" ra:ら:ラ    Compare:  hu:フ u:ウ  ラ:ら:ra:la",
		" ra is usually pronounced as la",
		" TT: right-ring up to the 'O' char",
		" Compare:  fu:フ, u:ウ:う,  & our:ラ:ら:?  TT: right-ring up to the 'O' char"},
	// 3 ri
	{"リ", "り", "ri",
		" ri:リ:り  actually, very very similar, actually!",
		" ri:リ:り Longer on the Right is Ri",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},
	{"リ", "り", "ri",
		" ri:リ:り  actually, very very similar, actually!",
		" ri:リ:り Longer on the Right is Ri",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},
	{"リ", "り", "ri",
		" ri:リ:り  actually, very very similar, actually!",
		" ri:リ:り Longer on the Right is Ri",
		" TT: left-ring-finger on it's 'L' char",
		" リ:り very similar   TT: left-ring-finger on it's 'L' char"},
	// 2 ru
	{"ル", "る", "ru",
		" ru:ル:る  ru-two, and the Hira has a luup for ru ",
		" ru",
		" TT: right-pinky-slide under to the left, on the '.' char",
		" ル:る  ?:two    TT: right-pinky-slide under to the left, on the '.' char"},
	{"ル", "る", "ru",
		" ru:ル:る  ru-two, and the Hira has a luup for ru ",
		" ru",
		" TT: right-pinky-slide under to the left, on the '.' char",
		" ル:る  ?:two    TT: right-pinky-slide under to the left, on the '.' char"},
	// 3 re
	{"レ", "れ", "re",
		" re:レ:れ  laying on the little finger",
		" re:レ:れ  the れ looks like someone getting blown, laid ",
		" TT: right-pinky on it's ';' char",
		" レ:れ   TT: right-pinky on it's ';' char"},
	{"レ", "れ", "re",
		" re:レ:れ  laying on the little finger",
		" re:レ:れ  the れ looks like someone getting blown, laid ",
		" TT: right-pinky on it's ';' char",
		" レ:れ   TT: right-pinky on it's ';' char"},
	{"レ", "れ", "re",
		" re:レ:れ  laying on the little finger",
		" re:レ:れ  the れ looks like someone getting blown, laid ",
		" TT: right-pinky on it's ';' char",
		" レ:れ   TT: right-pinky on it's ';' char"},
	// 2 ro
	{"ロ", "ろ", "ro",
		" ro:ロ:ろ  no loop on ro:ろ  Compare ru:る",
		" ro",
		" TT: right-pinky slide down --> right of the '?' char",
		" Looks like a square O,  TT: right-pinky slide down --> right of the '?' char"},
	{"ロ", "ろ", "ro",
		" ro:ロ:ろ  no loop on ro:ろ  Compare ru:る",
		" ro",
		" TT: right-pinky slide down --> right of the '?' char",
		" Looks like a square O,  TT: right-pinky slide down --> right of the '?' char"},
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
	// wa and wo ====================================================================================================
	// 2 of each ::
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
	//
	// Second set:
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
	//
	// n ============================================================================================================
	// 2 n ::
	{"ン", "ん", "n",
		" n:ン:ん  looks like a cursive n, and sounds like it too",
		" n:ン:ん  ン looks like 'no' has a nose, or was knocked-down",
		" TT: index to the 'Y' char",
		" ン:ん   TT: index to the 'Y'  Hira sounds the way it looks"},
	{"ン", "ん", "n",
		" n:ン:ん  looks like a cursive n, and sounds like it too",
		" n:ン:ん  ン looks like 'no' has a nose, or was knocked-down",
		" TT: index to the 'Y' char",
		" ン:ん   TT: index to the 'Y'  Hira sounds the way it looks"},
}

// const aCard = charSetStruct{} // did not work

// The structure of a single 'card' (aCard.) from fileOfCards
type charSetStruct struct {
	Kata   string
	Hira   string
	Romaji string

	HiraHint   string
	KataHint   string
	TT_Hint    string
	SansR_Hint string
}

/*
 Above, we instantiate a series of struct objects as slices of instances of the charSetStruct type
 e.g., fileOfCards being a "slice" of structures of type charSetStruct
 i.e., We are creating a slice named fileOfCards that holds instances of the charSetStruct type

Each element in a slice is initialized using the composite literal syntax whereby
... we are providing values for each field of the charSetStruct struct: i.e.,
... each set of values enclosed in curly braces { ... } represents an instance of the struct
*/

// Constants:
const colorReset = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorCyan = "\033[36m"
const colorPurple = "\033[35m"
