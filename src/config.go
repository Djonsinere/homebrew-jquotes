package src

type Quote struct {
	Used           bool
	Quote_english  string
	Quote_japanese string
}

var Quote_bank = []Quote{
	{Used: false, Quote_english: "The best way to predict the future is to create it.", Quote_japanese: "未来を予測する最良の方法は、それを創造することです。"},
	{Used: false, Quote_english: "Do what you can, with what you have, where you are.", Quote_japanese: "できることを、あるもので、いる場所でやりなさい。"},
	{Used: false, Quote_english: "It always seems impossible until it’s done.", Quote_japanese: "成し遂げるまでは、いつも不可能に見える。"},
	{Used: false, Quote_english: "Happiness depends upon ourselves.", Quote_japanese: "幸せは自分自身にかかっている。"},
	{Used: false, Quote_english: "Turn your wounds into wisdom.", Quote_japanese: "傷を知恵に変えなさい。"},
	{Used: false, Quote_english: "The harder you work for something, the greater you’ll feel when you achieve it.", Quote_japanese: "努力すればするほど、達成したときの喜びは大きくなる。"},
	{Used: false, Quote_english: "Don’t count the days, make the days count.", Quote_japanese: "日々を数えるのではなく、日々を意味あるものにしなさい。"},
	{Used: false, Quote_english: "Your limitation—it’s only your imagination.", Quote_japanese: "限界とは、あなたの想像にすぎない。"},
	{Used: false, Quote_english: "Dream bigger. Do bigger.", Quote_japanese: "もっと大きな夢を持ち、もっと大きな行動を起こしなさい。"},
	{Used: false, Quote_english: "Sometimes later becomes never. Do it now.", Quote_japanese: "時には『あとで』が『決して』になってしまう。今やりなさい。"},
}
