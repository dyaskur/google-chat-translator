package utils

type Languages struct {
	Language string   `json:"language,omitempty"`
	Code     string   `json:"code,omitempty"`
	IDs      []int16  `json:"ids,omitempty"`
	Commands []string `json:"commands,omitempty"`
}

var LanguagesData = []Languages{
	{Language: "Afrikaans", Code: "af", IDs: []int16{11}, Commands: []string{"/af"}},
	{Language: "Albanian", Code: "sq", IDs: []int16{12}, Commands: []string{"/albanian"}},
	{Language: "Amharic", Code: "am", IDs: []int16{13}, Commands: []string{"/amharic"}},
	{Language: "Arabic", Code: "ar", IDs: []int16{14}, Commands: []string{"/arabic"}},
	{Language: "Armenian", Code: "hy", IDs: []int16{15}, Commands: []string{"/armenian"}},
	{Language: "Assamese", Code: "as", IDs: []int16{17}, Commands: []string{"/assamese"}},
	{Language: "Aymara", Code: "ay", IDs: []int16{18}, Commands: []string{"/aymara"}},
	{Language: "Azerbaijani", Code: "az", IDs: []int16{19, 201}, Commands: []string{"/azerbaijani"}},
	{Language: "Bambara", Code: "bm", IDs: []int16{21, 202}, Commands: []string{"/bambara"}},
	{Language: "Basque", Code: "eu", IDs: []int16{22, 203}, Commands: []string{"/basque"}},
	{Language: "Belarusian", Code: "be", IDs: []int16{23, 204}, Commands: []string{"/belarusian"}},
	{Language: "Bengali", Code: "bn", IDs: []int16{24, 205}, Commands: []string{"/bengali"}},
	{Language: "Bhojpuri", Code: "bho", IDs: []int16{25, 206}, Commands: []string{"/bhojpuri"}},
	{Language: "Bosnian", Code: "bs", IDs: []int16{26, 207}, Commands: []string{"/bosnian"}},
	{Language: "Bulgarian", Code: "bg", IDs: []int16{27, 208}, Commands: []string{"/bulgarian"}},
	{Language: "Catalan", Code: "ca", IDs: []int16{28, 209}, Commands: []string{"/catalan"}},
	{Language: "Cebuano", Code: "ceb", IDs: []int16{29}, Commands: []string{"/cebuano"}},
	{Language: "Chinese_Simplified", Code: "zh", IDs: []int16{30, 210}, Commands: []string{"/chinese_simplified"}},
	{Language: "Chinese_Traditional", Code: "zh-TW", IDs: []int16{31, 211}, Commands: []string{"/chinese_traditional"}},
	{Language: "Corsican", Code: "co", IDs: []int16{32}, Commands: []string{"/corsican"}},
	{Language: "Croatian", Code: "hr", IDs: []int16{33, 212}, Commands: []string{"/croatian"}},
	{Language: "Czech", Code: "cs", IDs: []int16{34, 213}, Commands: []string{"/czech"}},
	{Language: "Danish", Code: "da", IDs: []int16{35, 214}, Commands: []string{"/danish"}},
	{Language: "Divehi", Code: "dv", IDs: []int16{36, 215}, Commands: []string{"/divehi"}},
	{Language: "Dutch", Code: "nl", IDs: []int16{37, 216}, Commands: []string{"/dutch"}},
	{Language: "English", Code: "en", IDs: []int16{38}, Commands: []string{"/english"}},
	{Language: "Esperanto", Code: "eo", IDs: []int16{39}, Commands: []string{"/esperanto"}},
	{Language: "Estonian", Code: "et", IDs: []int16{40, 217}, Commands: []string{"/estonian"}},
	{Language: "Ewe", Code: "ee", IDs: []int16{41}, Commands: []string{"/ewe"}},
	{Language: "Filipino (Tagalog)", Code: "fil", IDs: []int16{45}, Commands: []string{"/filipino_tagalog"}},
	{Language: "Finnish", Code: "fi", IDs: []int16{42, 218}, Commands: []string{"/finnish"}},
	{Language: "French", Code: "fr", IDs: []int16{43}, Commands: []string{"/french"}},
	{Language: "Frisian", Code: "fy", IDs: []int16{44}, Commands: []string{"/frisian"}},
	{Language: "Galician", Code: "gy", IDs: []int16{48}, Commands: []string{"/galician"}},
	{Language: "Georgian", Code: "ka", IDs: []int16{46}, Commands: []string{"/georgian"}},
	{Language: "German", Code: "de", IDs: []int16{47, 219}, Commands: []string{"/german", "/deutch"}},
	{Language: "Greek", Code: "el", IDs: []int16{49}, Commands: []string{"/greek"}},
	{Language: "Gujarati", Code: "gu", IDs: []int16{50}, Commands: []string{"/gujarati"}},
	{Language: "Guarani", Code: "gn", IDs: []int16{51, 220}, Commands: []string{"/guarani", "/avançado"}},
	{Language: "Haitian Creole", Code: "ht", IDs: []int16{52, 221}, Commands: []string{"/haitian_creole", "/kreyol"}},
	{Language: "Hausa", Code: "ha", IDs: []int16{53}, Commands: []string{"/hausa"}},
	{Language: "Hawaiian", Code: "haw", IDs: []int16{54}, Commands: []string{"/hawaiian"}},
	{Language: "Hebrew", Code: "he", IDs: []int16{55}, Commands: []string{"/hebrew"}},
	{Language: "Hindi", Code: "hi", IDs: []int16{56}, Commands: []string{"/hindi"}},
	{Language: "Hmong", Code: "hmn", IDs: []int16{57}, Commands: []string{"/hmong"}},
	{Language: "Hungarian", Code: "hu", IDs: []int16{58, 222}, Commands: []string{"/hungarian", "/magyar"}},
	{Language: "Icelandic", Code: "is", IDs: []int16{59}, Commands: []string{"/icelandic"}},
	{Language: "Igbo", Code: "ig", IDs: []int16{60}, Commands: []string{"/igbo"}},
	{Language: "Ilocano", Code: "ilo", IDs: []int16{61}, Commands: []string{"/ilocano"}},
	{Language: "Indonesian", Code: "id", IDs: []int16{62}, Commands: []string{"/indonesian"}},
	{Language: "Irish", Code: "ga", IDs: []int16{63, 224}, Commands: []string{"/irish", "/gaeilge"}},
	{Language: "Italian", Code: "it", IDs: []int16{64}, Commands: []string{"/italian"}},
	{Language: "Japanese", Code: "ja", IDs: []int16{65}, Commands: []string{"/japanese"}},
	{Language: "Javanese", Code: "jv", IDs: []int16{66}, Commands: []string{"/javanese"}},
	{Language: "Kannada", Code: "kn", IDs: []int16{67}, Commands: []string{"/kannada"}},
	{Language: "Kazakh", Code: "kk", IDs: []int16{68}, Commands: []string{"/kazakh"}},
	{Language: "Khmer", Code: "km", IDs: []int16{69}, Commands: []string{"/khmer"}},
	{Language: "Kinyarwanda", Code: "rw", IDs: []int16{70}, Commands: []string{"/kinyarwanda", "/rwanda"}},
	{Language: "Konkani", Code: "gom", IDs: []int16{71, 225}, Commands: []string{"/konkani"}},
	{Language: "Korean", Code: "ko", IDs: []int16{72}, Commands: []string{"/korean"}},
	{Language: "Krio", Code: "kri", IDs: []int16{73}, Commands: []string{"/krio"}},
	{Language: "Kurdish", Code: "ku", IDs: []int16{74}, Commands: []string{"/kurdish"}},
	{Language: "Kurdish (Sorani)", Code: "ckb", IDs: []int16{75}, Commands: []string{"/kurdish_sorani"}},
	{Language: "Kyrgyz", Code: "ky", IDs: []int16{76}, Commands: []string{"/kyrgyz"}},
	{Language: "Lao", Code: "lo", IDs: []int16{77}, Commands: []string{"/lao"}},
	{Language: "Latin", Code: "la", IDs: []int16{78}, Commands: []string{"/latin"}},
	{Language: "Latvian", Code: "lv", IDs: []int16{79}, Commands: []string{"/latvian"}},
	{Language: "Lingala", Code: "ln", IDs: []int16{80}, Commands: []string{"/lingala"}},
	{Language: "Lithuanian", Code: "lt", IDs: []int16{81}, Commands: []string{"/lithuanian"}},
	{Language: "Luganda", Code: "lg", IDs: []int16{82}, Commands: []string{"/luganda"}},
	{Language: "Luxembourgish", Code: "lb", IDs: []int16{83}, Commands: []string{"/luxembourgish"}},
	{Language: "Macedonian", Code: "mk", IDs: []int16{84}, Commands: []string{"/macedonian"}},
	{Language: "Maithili", Code: "mai", IDs: []int16{85}, Commands: []string{"/maithili"}},
	{Language: "Malagasy", Code: "mg", IDs: []int16{86}, Commands: []string{"/malagasy"}},
	{Language: "Malay", Code: "ms", IDs: []int16{87}, Commands: []string{"/malay"}},
	{Language: "Malayalam", Code: "ml", IDs: []int16{88}, Commands: []string{"/malayalam"}},
	{Language: "Maltese", Code: "mt", IDs: []int16{89}, Commands: []string{"/maltese"}},
	{Language: "Maori", Code: "mi", IDs: []int16{90}, Commands: []string{"/maori"}},
	{Language: "Marathi", Code: "mr", IDs: []int16{91}, Commands: []string{"/marathi"}},
	{Language: "Meiteilon (Manipuri)", Code: "mni-Mtei", IDs: []int16{92}, Commands: []string{"/meiteilon_manipuri"}},
	{Language: "Mizo", Code: "lus", IDs: []int16{93}, Commands: []string{"/mizo"}},
	{Language: "Mongolian", Code: "mn", IDs: []int16{94}, Commands: []string{"/mongolian"}},
	{Language: "Myanmar (Burmese)", Code: "my", IDs: []int16{95}, Commands: []string{"/burmese"}},
	{Language: "Nepali", Code: "ne", IDs: []int16{96}, Commands: []string{"/nepali"}},
	{Language: "Norwegian", Code: "no", IDs: []int16{97}, Commands: []string{"/norwegian"}},
	{Language: "Nyanja (Chichewa)", Code: "ny", IDs: []int16{98}, Commands: []string{"/nyanja_chichewa"}},
	{Language: "Odia (Oriya)", Code: "or", IDs: []int16{99}, Commands: []string{"/odia_oriya"}},
	{Language: "Oromo", Code: "om", IDs: []int16{100}, Commands: []string{"/oromo"}},
	{Language: "Pashto", Code: "ps", IDs: []int16{101}, Commands: []string{"/pashto"}},
	{Language: "Persian", Code: "fa", IDs: []int16{102}, Commands: []string{"/persian"}},
	{Language: "Polish", Code: "pl", IDs: []int16{103}, Commands: []string{"/polish"}},
	{Language: "Portuguese (Portugal, Brazil)", Code: "pt", IDs: []int16{104}, Commands: []string{"/portuguese"}},
	{Language: "Punjabi", Code: "pa", IDs: []int16{105}, Commands: []string{"/punjabi"}},
	{Language: "Quechua", Code: "qu", IDs: []int16{106}, Commands: []string{"/quechua"}},
	{Language: "Romanian", Code: "ro", IDs: []int16{107}, Commands: []string{"/romanian"}},
	{Language: "Russian", Code: "ru", IDs: []int16{108}, Commands: []string{"/russian"}},
	{Language: "Samoan", Code: "sm", IDs: []int16{109}, Commands: []string{"/samoan"}},
	{Language: "Sanskrit", Code: "sa", IDs: []int16{110}, Commands: []string{"/sanskrit"}},
	{Language: "Scots Gaelic", Code: "gd", IDs: []int16{111}, Commands: []string{"/scots_gaelic"}},
	{Language: "Sepedi", Code: "nso", IDs: []int16{112}, Commands: []string{"/sepedi"}},
	{Language: "Serbian", Code: "sr", IDs: []int16{113}, Commands: []string{"/serbian"}},
	{Language: "Sesotho", Code: "st", IDs: []int16{114}, Commands: []string{"/sesotho"}},
	{Language: "Shona", Code: "sn", IDs: []int16{115}, Commands: []string{"/shona"}},
	{Language: "Sindhi", Code: "sd", IDs: []int16{116}, Commands: []string{"/sindhi"}},
	{Language: "Sinhala (Sinhalese)", Code: "si", IDs: []int16{117}, Commands: []string{"/sinhala_sinhalese"}},
	{Language: "Slovak", Code: "sk", IDs: []int16{118}, Commands: []string{"/slovak"}},
	{Language: "Slovenian", Code: "sl", IDs: []int16{119}, Commands: []string{"/slovenian"}},
	{Language: "Somali", Code: "so", IDs: []int16{120}, Commands: []string{"/somali"}},
	{Language: "Spanish", Code: "es", IDs: []int16{121}, Commands: []string{"/spanish"}},
	{Language: "Sundanese", Code: "su", IDs: []int16{122}, Commands: []string{"/sundanese"}},
	{Language: "Swahili", Code: "sw", IDs: []int16{123}, Commands: []string{"/swahili"}},
	{Language: "Swedish", Code: "sv", IDs: []int16{124}, Commands: []string{"/swedish"}},
	{Language: "Tagalog (Filipino)", Code: "tl", IDs: []int16{125}, Commands: []string{"/tagalog"}},
	{Language: "Tajik", Code: "tg", IDs: []int16{126}, Commands: []string{"/tajik"}},
	{Language: "Tamil", Code: "ta", IDs: []int16{127}, Commands: []string{"/tamil"}},
	{Language: "Tatar", Code: "tt", IDs: []int16{128}, Commands: []string{"/tatar"}},
	{Language: "Telugu", Code: "te", IDs: []int16{129}, Commands: []string{"/telugu"}},
	{Language: "Thai", Code: "th", IDs: []int16{130}, Commands: []string{"/thai"}},
	{Language: "Tigrinya", Code: "ti", IDs: []int16{131}, Commands: []string{"/tigrinya"}},
	{Language: "Tsonga", Code: "ts", IDs: []int16{132}, Commands: []string{"/tsonga"}},
	{Language: "Turkish", Code: "tr", IDs: []int16{133}, Commands: []string{"/turkish"}},
	{Language: "Turkmen", Code: "tk", IDs: []int16{134}, Commands: []string{"/turkmen"}},
	{Language: "Twi (Akan)", Code: "ak", IDs: []int16{135}, Commands: []string{"/twi_akan"}},
	{Language: "Ukrainian", Code: "uk", IDs: []int16{136}, Commands: []string{"/ukrainian"}},
	{Language: "Urdu", Code: "ur", IDs: []int16{137}, Commands: []string{"/urdu"}},
	{Language: "Uyghur", Code: "ug", IDs: []int16{138}, Commands: []string{"/uyghur"}},
	{Language: "Uzbek", Code: "uz", IDs: []int16{139}, Commands: []string{"/uzbek"}},
	{Language: "Vietnamese", Code: "vi", IDs: []int16{140}, Commands: []string{"/vietnamese"}},
	{Language: "Welsh", Code: "cy", IDs: []int16{141}, Commands: []string{"/welsh"}},
	{Language: "Xhosa", Code: "xh", IDs: []int16{142}, Commands: []string{"/xhosa"}},
	{Language: "Yiddish", Code: "yi", IDs: []int16{143}, Commands: []string{"/yiddish"}},
	{Language: "Yoruba", Code: "yo", IDs: []int16{144}, Commands: []string{"/yoruba"}},
	{Language: "Zulu", Code: "zu", IDs: []int16{145}, Commands: []string{"/zulu"}},
}

// group by code
func GroupByCode() map[string]Languages {
	grouped := make(map[string]Languages)
	for _, l := range LanguagesData {
		grouped[l.Code] = l
	}
	return grouped
}

// group by ID
func GroupById() map[int16]Languages {
	grouped := make(map[int16]Languages)
	for _, l := range LanguagesData {
		for _, id := range l.IDs {
			grouped[id] = l
		}
	}
	return grouped
}

// GetByCode group by ID
func GetByCode(code string) Languages {
	for _, l := range LanguagesData {
		if l.Code == code {
			return l
		}
	}
	return LanguagesData[0]
}

// group by ID
func GetById(id int16) Languages {
	for _, l := range LanguagesData {
		for _, i := range l.IDs {
			if i == id {
				return l
			}
		}
	}
	return LanguagesData[0]
}

func GroupByCommand() map[string]Languages {
	grouped := make(map[string]Languages)
	for _, l := range LanguagesData {
		for _, command := range l.Commands {
			grouped[command] = l
		}
	}
	return grouped
}
