// GENERATED CODE - DO NOT MODIFY BY HAND

package intl

import "strings"

// LanguageCode represents an ISO 639 language code.
type LanguageCode string

// Language codes.
const (
	LanguageDA LanguageCode = "da"
	LanguageEN LanguageCode = "en"
	LanguageAZ LanguageCode = "az"
	LanguageDE LanguageCode = "de"
	LanguageAB LanguageCode = "ab"
	LanguageAA LanguageCode = "aa"
	LanguageAF LanguageCode = "af"
	LanguageAK LanguageCode = "ak"
	LanguageSQ LanguageCode = "sq"
	LanguageAM LanguageCode = "am"
	LanguageAR LanguageCode = "ar"
	LanguageAN LanguageCode = "an"
	LanguageHY LanguageCode = "hy"
	LanguageAS LanguageCode = "as"
	LanguageAV LanguageCode = "av"
	LanguageAE LanguageCode = "ae"
	LanguageAY LanguageCode = "ay"
	LanguageBM LanguageCode = "bm"
	LanguageBA LanguageCode = "ba"
	LanguageEU LanguageCode = "eu"
	LanguageBE LanguageCode = "be"
	LanguageBN LanguageCode = "bn"
	LanguageBH LanguageCode = "bh"
	LanguageBI LanguageCode = "bi"
	LanguageBS LanguageCode = "bs"
	LanguageBR LanguageCode = "br"
	LanguageBG LanguageCode = "bg"
	LanguageMY LanguageCode = "my"
	LanguageCA LanguageCode = "ca"
	LanguageCH LanguageCode = "ch"
	LanguageCE LanguageCode = "ce"
	LanguageCU LanguageCode = "cu"
	LanguageNY LanguageCode = "ny"
	LanguageZH LanguageCode = "zh"
	LanguageCV LanguageCode = "cv"
	LanguageKW LanguageCode = "kw"
	LanguageCO LanguageCode = "co"
	LanguageCR LanguageCode = "cr"
	LanguageHR LanguageCode = "hr"
	LanguageCS LanguageCode = "cs"
	LanguageDV LanguageCode = "dv"
	LanguageNL LanguageCode = "nl"
	LanguageDZ LanguageCode = "dz"
	LanguageEO LanguageCode = "eo"
	LanguageET LanguageCode = "et"
	LanguageEE LanguageCode = "ee"
	LanguageFO LanguageCode = "fo"
	LanguageFJ LanguageCode = "fj"
	LanguageFI LanguageCode = "fi"
	LanguageFR LanguageCode = "fr"
	LanguageFY LanguageCode = "fy"
	LanguageFF LanguageCode = "ff"
	LanguageGD LanguageCode = "gd"
	LanguageGL LanguageCode = "gl"
	LanguageLG LanguageCode = "lg"
	LanguageKA LanguageCode = "ka"
	LanguageEL LanguageCode = "el"
	LanguageKL LanguageCode = "kl"
	LanguageGN LanguageCode = "gn"
	LanguageGU LanguageCode = "gu"
	LanguageHT LanguageCode = "ht"
	LanguageHA LanguageCode = "ha"
	LanguageHE LanguageCode = "he"
	LanguageHZ LanguageCode = "hz"
	LanguageHI LanguageCode = "hi"
	LanguageHO LanguageCode = "ho"
	LanguageHU LanguageCode = "hu"
	LanguageIS LanguageCode = "is"
	LanguageIO LanguageCode = "io"
	LanguageIG LanguageCode = "ig"
	LanguageID LanguageCode = "id"
	LanguageIA LanguageCode = "ia"
	LanguageIE LanguageCode = "ie"
	LanguageII LanguageCode = "ii"
	LanguageIU LanguageCode = "iu"
	LanguageIK LanguageCode = "ik"
	LanguageGA LanguageCode = "ga"
	LanguageIT LanguageCode = "it"
	LanguageJA LanguageCode = "ja"
	LanguageJV LanguageCode = "jv"
	LanguageKN LanguageCode = "kn"
	LanguageKR LanguageCode = "kr"
	LanguageKS LanguageCode = "ks"
	LanguageKK LanguageCode = "kk"
	LanguageKM LanguageCode = "km"
	LanguageKI LanguageCode = "ki"
	LanguageRW LanguageCode = "rw"
	LanguageKY LanguageCode = "ky"
	LanguageKV LanguageCode = "kv"
	LanguageKG LanguageCode = "kg"
	LanguageKO LanguageCode = "ko"
	LanguageKJ LanguageCode = "kj"
	LanguageKU LanguageCode = "ku"
	LanguageLO LanguageCode = "lo"
	LanguageLA LanguageCode = "la"
	LanguageLV LanguageCode = "lv"
	LanguageLI LanguageCode = "li"
	LanguageLN LanguageCode = "ln"
	LanguageLT LanguageCode = "lt"
	LanguageLU LanguageCode = "lu"
	LanguageLB LanguageCode = "lb"
	LanguageMK LanguageCode = "mk"
	LanguageMG LanguageCode = "mg"
	LanguageMS LanguageCode = "ms"
	LanguageML LanguageCode = "ml"
	LanguageMT LanguageCode = "mt"
	LanguageGV LanguageCode = "gv"
	LanguageMI LanguageCode = "mi"
	LanguageMR LanguageCode = "mr"
	LanguageMH LanguageCode = "mh"
	LanguageMN LanguageCode = "mn"
	LanguageNA LanguageCode = "na"
	LanguageNV LanguageCode = "nv"
	LanguageND LanguageCode = "nd"
	LanguageNR LanguageCode = "nr"
	LanguageNG LanguageCode = "ng"
	LanguageNE LanguageCode = "ne"
	LanguageNO LanguageCode = "no"
	LanguageNB LanguageCode = "nb"
	LanguageNN LanguageCode = "nn"
	LanguageOC LanguageCode = "oc"
	LanguageOJ LanguageCode = "oj"
	LanguageOR LanguageCode = "or"
	LanguageOM LanguageCode = "om"
	LanguageOS LanguageCode = "os"
	LanguagePI LanguageCode = "pi"
	LanguagePS LanguageCode = "ps"
	LanguageFA LanguageCode = "fa"
	LanguagePL LanguageCode = "pl"
	LanguagePT LanguageCode = "pt"
	LanguagePA LanguageCode = "pa"
	LanguageQU LanguageCode = "qu"
	LanguageRO LanguageCode = "ro"
	LanguageRM LanguageCode = "rm"
	LanguageRN LanguageCode = "rn"
	LanguageRU LanguageCode = "ru"
	LanguageSE LanguageCode = "se"
	LanguageSM LanguageCode = "sm"
	LanguageSG LanguageCode = "sg"
	LanguageSA LanguageCode = "sa"
	LanguageSC LanguageCode = "sc"
	LanguageSR LanguageCode = "sr"
	LanguageSD LanguageCode = "sd"
	LanguageSI LanguageCode = "si"
	LanguageSK LanguageCode = "sk"
	LanguageSL LanguageCode = "sl"
	LanguageSO LanguageCode = "so"
	LanguageST LanguageCode = "st"
	LanguageSN LanguageCode = "sn"
	LanguageES LanguageCode = "es"
	LanguageSU LanguageCode = "su"
	LanguageSW LanguageCode = "sw"
	LanguageSS LanguageCode = "ss"
	LanguageSV LanguageCode = "sv"
	LanguageTL LanguageCode = "tl"
	LanguageTY LanguageCode = "ty"
	LanguageTG LanguageCode = "tg"
	LanguageTA LanguageCode = "ta"
	LanguageTT LanguageCode = "tt"
	LanguageTE LanguageCode = "te"
	LanguageTH LanguageCode = "th"
	LanguageBO LanguageCode = "bo"
	LanguageTI LanguageCode = "ti"
	LanguageTO LanguageCode = "to"
	LanguageTS LanguageCode = "ts"
	LanguageTN LanguageCode = "tn"
	LanguageTR LanguageCode = "tr"
	LanguageTK LanguageCode = "tk"
	LanguageTW LanguageCode = "tw"
	LanguageUG LanguageCode = "ug"
	LanguageUK LanguageCode = "uk"
	LanguageUR LanguageCode = "ur"
	LanguageUZ LanguageCode = "uz"
	LanguageVE LanguageCode = "ve"
	LanguageVI LanguageCode = "vi"
	LanguageVO LanguageCode = "vo"
	LanguageWA LanguageCode = "wa"
	LanguageCY LanguageCode = "cy"
	LanguageWO LanguageCode = "wo"
	LanguageXH LanguageCode = "xh"
	LanguageYI LanguageCode = "yi"
	LanguageYO LanguageCode = "yo"
	LanguageZA LanguageCode = "za"
	LanguageZU LanguageCode = "zu"
)

// LanguageDialect represents a dialect of a language.
type LanguageDialect struct {
	ID         string
	Code       string
	NativeName string
	FlagCode   string
}

// Language represents a language with its metadata.
type Language struct {
	ID              string
	Code            string
	NativeName      string
	Dialects        []LanguageDialect
	DefaultFlagCode string
}

// String returns the language code.
func (l LanguageCode) String() string {
	return string(l)
}

// Language returns the language data for this code.
func (l LanguageCode) Language() Language {
	return languages[l]
}

var languages = map[LanguageCode]Language{
	LanguageDA: {
		ID:              "danish",
		Code:            "da",
		NativeName:      "dansk",
		Dialects:        []LanguageDialect{},
		DefaultFlagCode: "dk",
	},
	LanguageEN: {
		ID:              "english",
		Code:            "en",
		NativeName:      "english",
		Dialects:        []LanguageDialect{},
		DefaultFlagCode: "gb",
	},
	LanguageAZ: {
		ID:              "azerbaijani",
		Code:            "az",
		NativeName:      "azərbaycan dili",
		Dialects:        []LanguageDialect{},
		DefaultFlagCode: "az",
	},
	LanguageDE: {
		ID:              "german",
		Code:            "de",
		NativeName:      "Deutsch",
		Dialects:        []LanguageDialect{},
		DefaultFlagCode: "de",
	},
	LanguageAB: {
		ID:         "abkhazian",
		Code:       "ab",
		NativeName: "аҧсуа",
		Dialects:   []LanguageDialect{},
	},
	LanguageAA: {
		ID:         "afar",
		Code:       "aa",
		NativeName: "Afaraf",
		Dialects:   []LanguageDialect{},
	},
	LanguageAF: {
		ID:         "afrikaans",
		Code:       "af",
		NativeName: "Afrikaans",
		Dialects:   []LanguageDialect{},
	},
	LanguageAK: {
		ID:         "akan",
		Code:       "ak",
		NativeName: "Akan",
		Dialects:   []LanguageDialect{},
	},
	LanguageSQ: {
		ID:         "albanian",
		Code:       "sq",
		NativeName: "Shqip",
		Dialects:   []LanguageDialect{},
	},
	LanguageAM: {
		ID:         "amharic",
		Code:       "am",
		NativeName: "አማርኛ",
		Dialects:   []LanguageDialect{},
	},
	LanguageAR: {
		ID:         "arabic",
		Code:       "ar",
		NativeName: "العربية",
		Dialects:   []LanguageDialect{},
	},
	LanguageAN: {
		ID:         "aragonese",
		Code:       "an",
		NativeName: "aragonés",
		Dialects:   []LanguageDialect{},
	},
	LanguageHY: {
		ID:         "armenian",
		Code:       "hy",
		NativeName: "Հայերեն",
		Dialects:   []LanguageDialect{},
	},
	LanguageAS: {
		ID:         "assamese",
		Code:       "as",
		NativeName: "অসমীয়া",
		Dialects:   []LanguageDialect{},
	},
	LanguageAV: {
		ID:         "avaric",
		Code:       "av",
		NativeName: "авар мацӀ",
		Dialects:   []LanguageDialect{},
	},
	LanguageAE: {
		ID:         "avestan",
		Code:       "ae",
		NativeName: "avesta",
		Dialects:   []LanguageDialect{},
	},
	LanguageAY: {
		ID:         "aymara",
		Code:       "ay",
		NativeName: "aymar aru",
		Dialects:   []LanguageDialect{},
	},
	LanguageBM: {
		ID:         "bambara",
		Code:       "bm",
		NativeName: "bamanankan",
		Dialects:   []LanguageDialect{},
	},
	LanguageBA: {
		ID:         "bashkir",
		Code:       "ba",
		NativeName: "башҡорт теле",
		Dialects:   []LanguageDialect{},
	},
	LanguageEU: {
		ID:         "basque",
		Code:       "eu",
		NativeName: "euskara",
		Dialects:   []LanguageDialect{},
	},
	LanguageBE: {
		ID:         "belarusian",
		Code:       "be",
		NativeName: "беларуская мова",
		Dialects:   []LanguageDialect{},
	},
	LanguageBN: {
		ID:         "bengali",
		Code:       "bn",
		NativeName: "বাংলা",
		Dialects:   []LanguageDialect{},
	},
	LanguageBH: {
		ID:         "bihari",
		Code:       "bh",
		NativeName: "भोजपुरी",
		Dialects:   []LanguageDialect{},
	},
	LanguageBI: {
		ID:         "bislama",
		Code:       "bi",
		NativeName: "Bislama",
		Dialects:   []LanguageDialect{},
	},
	LanguageBS: {
		ID:         "bosnian",
		Code:       "bs",
		NativeName: "bosanski jezik",
		Dialects:   []LanguageDialect{},
	},
	LanguageBR: {
		ID:         "breton",
		Code:       "br",
		NativeName: "brezhoneg",
		Dialects:   []LanguageDialect{},
	},
	LanguageBG: {
		ID:         "bulgarian",
		Code:       "bg",
		NativeName: "български език",
		Dialects:   []LanguageDialect{},
	},
	LanguageMY: {
		ID:         "burmese",
		Code:       "my",
		NativeName: "ဗမာစာ",
		Dialects:   []LanguageDialect{},
	},
	LanguageCA: {
		ID:         "catalan",
		Code:       "ca",
		NativeName: "català",
		Dialects:   []LanguageDialect{},
	},
	LanguageCH: {
		ID:         "chamorro",
		Code:       "ch",
		NativeName: "Chamoru",
		Dialects:   []LanguageDialect{},
	},
	LanguageCE: {
		ID:         "chechen",
		Code:       "ce",
		NativeName: "нохчийн мотт",
		Dialects:   []LanguageDialect{},
	},
	LanguageCU: {
		ID:         "churchSlavic",
		Code:       "cu",
		NativeName: "ѩзыкъ словѣньскъ",
		Dialects:   []LanguageDialect{},
	},
	LanguageNY: {
		ID:         "chichewa",
		Code:       "ny",
		NativeName: "chiCheŵa",
		Dialects:   []LanguageDialect{},
	},
	LanguageZH: {
		ID:         "chinese",
		Code:       "zh",
		NativeName: "中文",
		Dialects:   []LanguageDialect{},
	},
	LanguageCV: {
		ID:         "chuvash",
		Code:       "cv",
		NativeName: "чӑваш чӗлхи",
		Dialects:   []LanguageDialect{},
	},
	LanguageKW: {
		ID:         "cornish",
		Code:       "kw",
		NativeName: "Kernewek",
		Dialects:   []LanguageDialect{},
	},
	LanguageCO: {
		ID:         "corsican",
		Code:       "co",
		NativeName: "corsu",
		Dialects:   []LanguageDialect{},
	},
	LanguageCR: {
		ID:         "cree",
		Code:       "cr",
		NativeName: "ᓀᐦᐃᔭᐍᐏᐣ",
		Dialects:   []LanguageDialect{},
	},
	LanguageHR: {
		ID:         "croatian",
		Code:       "hr",
		NativeName: "hrvatski jezik",
		Dialects:   []LanguageDialect{},
	},
	LanguageCS: {
		ID:         "czech",
		Code:       "cs",
		NativeName: "čeština",
		Dialects:   []LanguageDialect{},
	},
	LanguageDV: {
		ID:         "divehi",
		Code:       "dv",
		NativeName: "ދިވެހި",
		Dialects:   []LanguageDialect{},
	},
	LanguageNL: {
		ID:         "dutch",
		Code:       "nl",
		NativeName: "Nederlands",
		Dialects:   []LanguageDialect{},
	},
	LanguageDZ: {
		ID:         "dzongkha",
		Code:       "dz",
		NativeName: "རྫོང་ཁ",
		Dialects:   []LanguageDialect{},
	},
	LanguageEO: {
		ID:         "esperanto",
		Code:       "eo",
		NativeName: "Esperanto",
		Dialects:   []LanguageDialect{},
	},
	LanguageET: {
		ID:         "estonian",
		Code:       "et",
		NativeName: "eesti",
		Dialects:   []LanguageDialect{},
	},
	LanguageEE: {
		ID:         "ewe",
		Code:       "ee",
		NativeName: "Eʋegbe",
		Dialects:   []LanguageDialect{},
	},
	LanguageFO: {
		ID:         "faroese",
		Code:       "fo",
		NativeName: "føroyskt",
		Dialects:   []LanguageDialect{},
	},
	LanguageFJ: {
		ID:         "fijian",
		Code:       "fj",
		NativeName: "vosa Vakaviti",
		Dialects:   []LanguageDialect{},
	},
	LanguageFI: {
		ID:         "finnish",
		Code:       "fi",
		NativeName: "suomi",
		Dialects:   []LanguageDialect{},
	},
	LanguageFR: {
		ID:         "french",
		Code:       "fr",
		NativeName: "français",
		Dialects:   []LanguageDialect{},
	},
	LanguageFY: {
		ID:         "westernFrisian",
		Code:       "fy",
		NativeName: "Frysk",
		Dialects:   []LanguageDialect{},
	},
	LanguageFF: {
		ID:         "fulah",
		Code:       "ff",
		NativeName: "Fulfulde",
		Dialects:   []LanguageDialect{},
	},
	LanguageGD: {
		ID:         "gaelic",
		Code:       "gd",
		NativeName: "Gàidhlig",
		Dialects:   []LanguageDialect{},
	},
	LanguageGL: {
		ID:         "galician",
		Code:       "gl",
		NativeName: "galego",
		Dialects:   []LanguageDialect{},
	},
	LanguageLG: {
		ID:         "ganda",
		Code:       "lg",
		NativeName: "Luganda",
		Dialects:   []LanguageDialect{},
	},
	LanguageKA: {
		ID:         "georgian",
		Code:       "ka",
		NativeName: "ქართული",
		Dialects:   []LanguageDialect{},
	},
	LanguageEL: {
		ID:         "greek",
		Code:       "el",
		NativeName: "Ελληνικά",
		Dialects:   []LanguageDialect{},
	},
	LanguageKL: {
		ID:         "greenlandic",
		Code:       "kl",
		NativeName: "kalaallisut",
		Dialects:   []LanguageDialect{},
	},
	LanguageGN: {
		ID:         "guarani",
		Code:       "gn",
		NativeName: "Avañe'ẽ",
		Dialects:   []LanguageDialect{},
	},
	LanguageGU: {
		ID:         "gujarati",
		Code:       "gu",
		NativeName: "ગુજરાતી",
		Dialects:   []LanguageDialect{},
	},
	LanguageHT: {
		ID:         "haitian",
		Code:       "ht",
		NativeName: "Kreyòl ayisyen",
		Dialects:   []LanguageDialect{},
	},
	LanguageHA: {
		ID:         "hausa",
		Code:       "ha",
		NativeName: "هَوُسَ",
		Dialects:   []LanguageDialect{},
	},
	LanguageHE: {
		ID:         "hebrew",
		Code:       "he",
		NativeName: "עברית",
		Dialects:   []LanguageDialect{},
	},
	LanguageHZ: {
		ID:         "herero",
		Code:       "hz",
		NativeName: "Otjiherero",
		Dialects:   []LanguageDialect{},
	},
	LanguageHI: {
		ID:         "hindi",
		Code:       "hi",
		NativeName: "हिन्दी",
		Dialects:   []LanguageDialect{},
	},
	LanguageHO: {
		ID:         "hiriMotu",
		Code:       "ho",
		NativeName: "Hiri Motu",
		Dialects:   []LanguageDialect{},
	},
	LanguageHU: {
		ID:         "hungarian",
		Code:       "hu",
		NativeName: "magyar",
		Dialects:   []LanguageDialect{},
	},
	LanguageIS: {
		ID:         "icelandic",
		Code:       "is",
		NativeName: "Íslenska",
		Dialects:   []LanguageDialect{},
	},
	LanguageIO: {
		ID:         "ido",
		Code:       "io",
		NativeName: "Ido",
		Dialects:   []LanguageDialect{},
	},
	LanguageIG: {
		ID:         "igbo",
		Code:       "ig",
		NativeName: "Asụsụ Igbo",
		Dialects:   []LanguageDialect{},
	},
	LanguageID: {
		ID:         "indonesian",
		Code:       "id",
		NativeName: "Bahasa Indonesia",
		Dialects:   []LanguageDialect{},
	},
	LanguageIA: {
		ID:         "interlingua",
		Code:       "ia",
		NativeName: "Interlingua",
		Dialects:   []LanguageDialect{},
	},
	LanguageIE: {
		ID:         "interlingue",
		Code:       "ie",
		NativeName: "Interlingue",
		Dialects:   []LanguageDialect{},
	},
	LanguageII: {
		ID:         "sichuanYi",
		Code:       "ii",
		NativeName: "ꆈꌠ꒿ Nuosuhxop",
		Dialects:   []LanguageDialect{},
	},
	LanguageIU: {
		ID:         "inuktitut",
		Code:       "iu",
		NativeName: "ᐃᓄᒃᑎᑐᑦ",
		Dialects:   []LanguageDialect{},
	},
	LanguageIK: {
		ID:         "inupiaq",
		Code:       "ik",
		NativeName: "Iñupiaq",
		Dialects:   []LanguageDialect{},
	},
	LanguageGA: {
		ID:         "irish",
		Code:       "ga",
		NativeName: "Gaeilge",
		Dialects:   []LanguageDialect{},
	},
	LanguageIT: {
		ID:         "italian",
		Code:       "it",
		NativeName: "italiano",
		Dialects:   []LanguageDialect{},
	},
	LanguageJA: {
		ID:         "japanese",
		Code:       "ja",
		NativeName: "日本語",
		Dialects:   []LanguageDialect{},
	},
	LanguageJV: {
		ID:         "javanese",
		Code:       "jv",
		NativeName: "basa Jawa",
		Dialects:   []LanguageDialect{},
	},
	LanguageKN: {
		ID:         "kannada",
		Code:       "kn",
		NativeName: "ಕನ್ನಡ",
		Dialects:   []LanguageDialect{},
	},
	LanguageKR: {
		ID:         "kanuri",
		Code:       "kr",
		NativeName: "Kanuri",
		Dialects:   []LanguageDialect{},
	},
	LanguageKS: {
		ID:         "kashmiri",
		Code:       "ks",
		NativeName: "कश्मीरी",
		Dialects:   []LanguageDialect{},
	},
	LanguageKK: {
		ID:         "kazakh",
		Code:       "kk",
		NativeName: "қазақ тілі",
		Dialects:   []LanguageDialect{},
	},
	LanguageKM: {
		ID:         "khmer",
		Code:       "km",
		NativeName: "ភាសាខ្មែរ",
		Dialects:   []LanguageDialect{},
	},
	LanguageKI: {
		ID:         "kikuyu",
		Code:       "ki",
		NativeName: "Gĩkũyũ",
		Dialects:   []LanguageDialect{},
	},
	LanguageRW: {
		ID:         "kinyarwanda",
		Code:       "rw",
		NativeName: "Ikinyarwanda",
		Dialects:   []LanguageDialect{},
	},
	LanguageKY: {
		ID:         "kirghiz",
		Code:       "ky",
		NativeName: "Кыргызча",
		Dialects:   []LanguageDialect{},
	},
	LanguageKV: {
		ID:         "komi",
		Code:       "kv",
		NativeName: "коми кыв",
		Dialects:   []LanguageDialect{},
	},
	LanguageKG: {
		ID:         "kongo",
		Code:       "kg",
		NativeName: "Kikongo",
		Dialects:   []LanguageDialect{},
	},
	LanguageKO: {
		ID:         "korean",
		Code:       "ko",
		NativeName: "한국어",
		Dialects:   []LanguageDialect{},
	},
	LanguageKJ: {
		ID:         "kuanyama",
		Code:       "kj",
		NativeName: "Kuanyama",
		Dialects:   []LanguageDialect{},
	},
	LanguageKU: {
		ID:         "kurdish",
		Code:       "ku",
		NativeName: "Kurdî",
		Dialects:   []LanguageDialect{},
	},
	LanguageLO: {
		ID:         "lao",
		Code:       "lo",
		NativeName: "ພາສາລາວ",
		Dialects:   []LanguageDialect{},
	},
	LanguageLA: {
		ID:         "latin",
		Code:       "la",
		NativeName: "latine",
		Dialects:   []LanguageDialect{},
	},
	LanguageLV: {
		ID:         "latvian",
		Code:       "lv",
		NativeName: "latviešu valoda",
		Dialects:   []LanguageDialect{},
	},
	LanguageLI: {
		ID:         "limburgish",
		Code:       "li",
		NativeName: "Limburgs",
		Dialects:   []LanguageDialect{},
	},
	LanguageLN: {
		ID:         "lingala",
		Code:       "ln",
		NativeName: "Lingála",
		Dialects:   []LanguageDialect{},
	},
	LanguageLT: {
		ID:         "lithuanian",
		Code:       "lt",
		NativeName: "lietuvių kalba",
		Dialects:   []LanguageDialect{},
	},
	LanguageLU: {
		ID:         "lubaKatanga",
		Code:       "lu",
		NativeName: "Tshiluba",
		Dialects:   []LanguageDialect{},
	},
	LanguageLB: {
		ID:         "luxembourgish",
		Code:       "lb",
		NativeName: "Lëtzebuergesch",
		Dialects:   []LanguageDialect{},
	},
	LanguageMK: {
		ID:         "macedonian",
		Code:       "mk",
		NativeName: "македонски јазик",
		Dialects:   []LanguageDialect{},
	},
	LanguageMG: {
		ID:         "malagasy",
		Code:       "mg",
		NativeName: "fiteny malagasy",
		Dialects:   []LanguageDialect{},
	},
	LanguageMS: {
		ID:         "malay",
		Code:       "ms",
		NativeName: "Bahasa Melayu",
		Dialects:   []LanguageDialect{},
	},
	LanguageML: {
		ID:         "malayalam",
		Code:       "ml",
		NativeName: "മലയാളം",
		Dialects:   []LanguageDialect{},
	},
	LanguageMT: {
		ID:         "maltese",
		Code:       "mt",
		NativeName: "Malti",
		Dialects:   []LanguageDialect{},
	},
	LanguageGV: {
		ID:         "manx",
		Code:       "gv",
		NativeName: "Gaelg",
		Dialects:   []LanguageDialect{},
	},
	LanguageMI: {
		ID:         "maori",
		Code:       "mi",
		NativeName: "te reo Māori",
		Dialects:   []LanguageDialect{},
	},
	LanguageMR: {
		ID:         "marathi",
		Code:       "mr",
		NativeName: "मराठी",
		Dialects:   []LanguageDialect{},
	},
	LanguageMH: {
		ID:         "marshallese",
		Code:       "mh",
		NativeName: "Kajin M̧ajeļ",
		Dialects:   []LanguageDialect{},
	},
	LanguageMN: {
		ID:         "mongolian",
		Code:       "mn",
		NativeName: "Монгол хэл",
		Dialects:   []LanguageDialect{},
	},
	LanguageNA: {
		ID:         "nauru",
		Code:       "na",
		NativeName: "Ekakairũ Naoero",
		Dialects:   []LanguageDialect{},
	},
	LanguageNV: {
		ID:         "navajo",
		Code:       "nv",
		NativeName: "Diné bizaad",
		Dialects:   []LanguageDialect{},
	},
	LanguageND: {
		ID:         "northNdebele",
		Code:       "nd",
		NativeName: "isiNdebele",
		Dialects:   []LanguageDialect{},
	},
	LanguageNR: {
		ID:         "southNdebele",
		Code:       "nr",
		NativeName: "isiNdebele",
		Dialects:   []LanguageDialect{},
	},
	LanguageNG: {
		ID:         "ndonga",
		Code:       "ng",
		NativeName: "Owambo",
		Dialects:   []LanguageDialect{},
	},
	LanguageNE: {
		ID:         "nepali",
		Code:       "ne",
		NativeName: "नेपाली",
		Dialects:   []LanguageDialect{},
	},
	LanguageNO: {
		ID:         "norwegian",
		Code:       "no",
		NativeName: "Norsk",
		Dialects:   []LanguageDialect{},
	},
	LanguageNB: {
		ID:         "norwegianBokmal",
		Code:       "nb",
		NativeName: "Norsk bokmål",
		Dialects:   []LanguageDialect{},
	},
	LanguageNN: {
		ID:         "norwegianNynorsk",
		Code:       "nn",
		NativeName: "Norsk nynorsk",
		Dialects:   []LanguageDialect{},
	},
	LanguageOC: {
		ID:         "occitan",
		Code:       "oc",
		NativeName: "occitan",
		Dialects:   []LanguageDialect{},
	},
	LanguageOJ: {
		ID:         "ojibwa",
		Code:       "oj",
		NativeName: "ᐊᓂᔑᓈᐯᒧᐎᓐ",
		Dialects:   []LanguageDialect{},
	},
	LanguageOR: {
		ID:         "oriya",
		Code:       "or",
		NativeName: "ଓଡ଼ିଆ",
		Dialects:   []LanguageDialect{},
	},
	LanguageOM: {
		ID:         "oromo",
		Code:       "om",
		NativeName: "Afaan Oromoo",
		Dialects:   []LanguageDialect{},
	},
	LanguageOS: {
		ID:         "ossetian",
		Code:       "os",
		NativeName: "ирон æвзаг",
		Dialects:   []LanguageDialect{},
	},
	LanguagePI: {
		ID:         "pali",
		Code:       "pi",
		NativeName: "पाऴि",
		Dialects:   []LanguageDialect{},
	},
	LanguagePS: {
		ID:         "pashto",
		Code:       "ps",
		NativeName: "پښتو",
		Dialects:   []LanguageDialect{},
	},
	LanguageFA: {
		ID:         "persian",
		Code:       "fa",
		NativeName: "فارسی",
		Dialects:   []LanguageDialect{},
	},
	LanguagePL: {
		ID:         "polish",
		Code:       "pl",
		NativeName: "język polski",
		Dialects:   []LanguageDialect{},
	},
	LanguagePT: {
		ID:         "portuguese",
		Code:       "pt",
		NativeName: "português",
		Dialects:   []LanguageDialect{},
	},
	LanguagePA: {
		ID:         "punjabi",
		Code:       "pa",
		NativeName: "ਪੰਜਾਬੀ",
		Dialects:   []LanguageDialect{},
	},
	LanguageQU: {
		ID:         "quechua",
		Code:       "qu",
		NativeName: "Runa Simi",
		Dialects:   []LanguageDialect{},
	},
	LanguageRO: {
		ID:         "romanian",
		Code:       "ro",
		NativeName: "limba română",
		Dialects:   []LanguageDialect{},
	},
	LanguageRM: {
		ID:         "romansh",
		Code:       "rm",
		NativeName: "rumantsch grischun",
		Dialects:   []LanguageDialect{},
	},
	LanguageRN: {
		ID:         "rundi",
		Code:       "rn",
		NativeName: "Ikirundi",
		Dialects:   []LanguageDialect{},
	},
	LanguageRU: {
		ID:         "russian",
		Code:       "ru",
		NativeName: "русский язык",
		Dialects:   []LanguageDialect{},
	},
	LanguageSE: {
		ID:         "northernSami",
		Code:       "se",
		NativeName: "Davvisámegiella",
		Dialects:   []LanguageDialect{},
	},
	LanguageSM: {
		ID:         "samoan",
		Code:       "sm",
		NativeName: "gagana fa'a Samoa",
		Dialects:   []LanguageDialect{},
	},
	LanguageSG: {
		ID:         "sango",
		Code:       "sg",
		NativeName: "yângâ tî sängö",
		Dialects:   []LanguageDialect{},
	},
	LanguageSA: {
		ID:         "sanskrit",
		Code:       "sa",
		NativeName: "संस्कृतम्",
		Dialects:   []LanguageDialect{},
	},
	LanguageSC: {
		ID:         "sardinian",
		Code:       "sc",
		NativeName: "sardu",
		Dialects:   []LanguageDialect{},
	},
	LanguageSR: {
		ID:         "serbian",
		Code:       "sr",
		NativeName: "српски језик",
		Dialects:   []LanguageDialect{},
	},
	LanguageSD: {
		ID:         "sindhi",
		Code:       "sd",
		NativeName: "सिन्धी",
		Dialects:   []LanguageDialect{},
	},
	LanguageSI: {
		ID:         "sinhala",
		Code:       "si",
		NativeName: "සිංහල",
		Dialects:   []LanguageDialect{},
	},
	LanguageSK: {
		ID:         "slovak",
		Code:       "sk",
		NativeName: "slovenčina",
		Dialects:   []LanguageDialect{},
	},
	LanguageSL: {
		ID:         "slovenian",
		Code:       "sl",
		NativeName: "slovenščina",
		Dialects:   []LanguageDialect{},
	},
	LanguageSO: {
		ID:         "somali",
		Code:       "so",
		NativeName: "Soomaaliga",
		Dialects:   []LanguageDialect{},
	},
	LanguageST: {
		ID:         "southernSotho",
		Code:       "st",
		NativeName: "Sesotho",
		Dialects:   []LanguageDialect{},
	},
	LanguageSN: {
		ID:         "shona",
		Code:       "sn",
		NativeName: "chiShona",
		Dialects:   []LanguageDialect{},
	},
	LanguageES: {
		ID:         "spanish",
		Code:       "es",
		NativeName: "español",
		Dialects:   []LanguageDialect{},
	},
	LanguageSU: {
		ID:         "sundanese",
		Code:       "su",
		NativeName: "Basa Sunda",
		Dialects:   []LanguageDialect{},
	},
	LanguageSW: {
		ID:         "swahili",
		Code:       "sw",
		NativeName: "Kiswahili",
		Dialects:   []LanguageDialect{},
	},
	LanguageSS: {
		ID:         "swati",
		Code:       "ss",
		NativeName: "SiSwati",
		Dialects:   []LanguageDialect{},
	},
	LanguageSV: {
		ID:         "swedish",
		Code:       "sv",
		NativeName: "svenska",
		Dialects:   []LanguageDialect{},
	},
	LanguageTL: {
		ID:         "tagalog",
		Code:       "tl",
		NativeName: "Wikang Tagalog",
		Dialects:   []LanguageDialect{},
	},
	LanguageTY: {
		ID:         "tahitian",
		Code:       "ty",
		NativeName: "Reo Tahiti",
		Dialects:   []LanguageDialect{},
	},
	LanguageTG: {
		ID:         "tajik",
		Code:       "tg",
		NativeName: "тоҷикӣ",
		Dialects:   []LanguageDialect{},
	},
	LanguageTA: {
		ID:         "tamil",
		Code:       "ta",
		NativeName: "தமிழ்",
		Dialects:   []LanguageDialect{},
	},
	LanguageTT: {
		ID:         "tatar",
		Code:       "tt",
		NativeName: "татар теле",
		Dialects:   []LanguageDialect{},
	},
	LanguageTE: {
		ID:         "telugu",
		Code:       "te",
		NativeName: "తెలుగు",
		Dialects:   []LanguageDialect{},
	},
	LanguageTH: {
		ID:         "thai",
		Code:       "th",
		NativeName: "ไทย",
		Dialects:   []LanguageDialect{},
	},
	LanguageBO: {
		ID:         "tibetan",
		Code:       "bo",
		NativeName: "བོད་ཡིག",
		Dialects:   []LanguageDialect{},
	},
	LanguageTI: {
		ID:         "tigrinya",
		Code:       "ti",
		NativeName: "ትግርኛ",
		Dialects:   []LanguageDialect{},
	},
	LanguageTO: {
		ID:         "tonga",
		Code:       "to",
		NativeName: "faka Tonga",
		Dialects:   []LanguageDialect{},
	},
	LanguageTS: {
		ID:         "tsonga",
		Code:       "ts",
		NativeName: "Xitsonga",
		Dialects:   []LanguageDialect{},
	},
	LanguageTN: {
		ID:         "tswana",
		Code:       "tn",
		NativeName: "Setswana",
		Dialects:   []LanguageDialect{},
	},
	LanguageTR: {
		ID:         "turkish",
		Code:       "tr",
		NativeName: "Türkçe",
		Dialects:   []LanguageDialect{},
	},
	LanguageTK: {
		ID:         "turkmen",
		Code:       "tk",
		NativeName: "Türkmen",
		Dialects:   []LanguageDialect{},
	},
	LanguageTW: {
		ID:         "twi",
		Code:       "tw",
		NativeName: "Twi",
		Dialects:   []LanguageDialect{},
	},
	LanguageUG: {
		ID:         "uighur",
		Code:       "ug",
		NativeName: "ئۇيغۇرچە‎",
		Dialects:   []LanguageDialect{},
	},
	LanguageUK: {
		ID:         "ukrainian",
		Code:       "uk",
		NativeName: "українська мова",
		Dialects:   []LanguageDialect{},
	},
	LanguageUR: {
		ID:         "urdu",
		Code:       "ur",
		NativeName: "اردو",
		Dialects:   []LanguageDialect{},
	},
	LanguageUZ: {
		ID:         "uzbek",
		Code:       "uz",
		NativeName: "Oʻzbek",
		Dialects:   []LanguageDialect{},
	},
	LanguageVE: {
		ID:         "venda",
		Code:       "ve",
		NativeName: "Tshivenḓa",
		Dialects:   []LanguageDialect{},
	},
	LanguageVI: {
		ID:         "vietnamese",
		Code:       "vi",
		NativeName: "Tiếng Việt",
		Dialects:   []LanguageDialect{},
	},
	LanguageVO: {
		ID:         "volapuk",
		Code:       "vo",
		NativeName: "Volapük",
		Dialects:   []LanguageDialect{},
	},
	LanguageWA: {
		ID:         "walloon",
		Code:       "wa",
		NativeName: "walon",
		Dialects:   []LanguageDialect{},
	},
	LanguageCY: {
		ID:         "welsh",
		Code:       "cy",
		NativeName: "Cymraeg",
		Dialects:   []LanguageDialect{},
	},
	LanguageWO: {
		ID:         "wolof",
		Code:       "wo",
		NativeName: "Wollof",
		Dialects:   []LanguageDialect{},
	},
	LanguageXH: {
		ID:         "xhosa",
		Code:       "xh",
		NativeName: "isiXhosa",
		Dialects:   []LanguageDialect{},
	},
	LanguageYI: {
		ID:         "yiddish",
		Code:       "yi",
		NativeName: "ייִדיש",
		Dialects:   []LanguageDialect{},
	},
	LanguageYO: {
		ID:         "yoruba",
		Code:       "yo",
		NativeName: "Yorùbá",
		Dialects:   []LanguageDialect{},
	},
	LanguageZA: {
		ID:         "zhuang",
		Code:       "za",
		NativeName: "Saɯ cueŋƅ",
		Dialects:   []LanguageDialect{},
	},
	LanguageZU: {
		ID:         "zulu",
		Code:       "zu",
		NativeName: "isiZulu",
		Dialects:   []LanguageDialect{},
	},
}

// LanguageByCode returns a language by its code.
func LanguageByCode(code string) (Language, bool) {
	l, ok := languages[LanguageCode(strings.ToLower(code))]

	return l, ok
}

// AllLanguages returns all languages.
func AllLanguages() []Language {
	result := make([]Language, 0, len(languages))
	for _, l := range languages {
		result = append(result, l)
	}

	return result
}
