package main

// Only used in locateCardAndDisplayHelpFieldsContainedInIt()
// ... and in func silentlyLocateCard(setKeyRequest string) {
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
		" se-->ze (sa-->za), plus a dakuten 濁点",
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