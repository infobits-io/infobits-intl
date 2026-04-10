// GENERATED CODE - DO NOT MODIFY BY HAND

package intl

import "strings"

// LanguageCode represents an ISO 639 language code
type LanguageCode string

// Language codes
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

// LanguageDialect represents a dialect of a language
type LanguageDialect struct {
	ID         string
	Code       string
	NativeName string
	FlagCode   string
}

// Language represents a language with its metadata
type Language struct {
	ID              string
	Code            string
	NativeName      string
	Dialects        []LanguageDialect
	DefaultFlagCode string
}

// String returns the language code
func (l LanguageCode) String() string {
	return string(l)
}

// Language returns the language data for this code
func (l LanguageCode) Language() Language {
	return languages[l]
}

var languages = map[LanguageCode]Language{
	LanguageDA: {
		ID:         "danish",
		Code:       "da",
		NativeName: "dansk",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "dk",
	},
	LanguageEN: {
		ID:         "english",
		Code:       "en",
		NativeName: "english",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gb",
	},
	LanguageAZ: {
		ID:         "azerbaijani",
		Code:       "az",
		NativeName: "azərbaycan dili",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "az",
	},
	LanguageDE: {
		ID:         "german",
		Code:       "de",
		NativeName: "Deutsch",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "de",
	},
	LanguageAB: {
		ID:         "abkhazian",
		Code:       "ab",
		NativeName: "аҧсуа",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ge",
	},
	LanguageAA: {
		ID:         "afar",
		Code:       "aa",
		NativeName: "Afaraf",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "et",
	},
	LanguageAF: {
		ID:         "afrikaans",
		Code:       "af",
		NativeName: "Afrikaans",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "za",
	},
	LanguageAK: {
		ID:         "akan",
		Code:       "ak",
		NativeName: "Akan",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gh",
	},
	LanguageSQ: {
		ID:         "albanian",
		Code:       "sq",
		NativeName: "Shqip",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "al",
	},
	LanguageAM: {
		ID:         "amharic",
		Code:       "am",
		NativeName: "አማርኛ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "et",
	},
	LanguageAR: {
		ID:         "arabic",
		Code:       "ar",
		NativeName: "العربية",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "sa",
	},
	LanguageAN: {
		ID:         "aragonese",
		Code:       "an",
		NativeName: "aragonés",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "es",
	},
	LanguageHY: {
		ID:         "armenian",
		Code:       "hy",
		NativeName: "Հայերեն",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "am",
	},
	LanguageAS: {
		ID:         "assamese",
		Code:       "as",
		NativeName: "অসমীয়া",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageAV: {
		ID:         "avaric",
		Code:       "av",
		NativeName: "авар мацӀ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ru",
	},
	LanguageAE: {
		ID:         "avestan",
		Code:       "ae",
		NativeName: "avesta",
		Dialects: []LanguageDialect{
		},
	},
	LanguageAY: {
		ID:         "aymara",
		Code:       "ay",
		NativeName: "aymar aru",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "bo",
	},
	LanguageBM: {
		ID:         "bambara",
		Code:       "bm",
		NativeName: "bamanankan",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ml",
	},
	LanguageBA: {
		ID:         "bashkir",
		Code:       "ba",
		NativeName: "башҡорт теле",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ru",
	},
	LanguageEU: {
		ID:         "basque",
		Code:       "eu",
		NativeName: "euskara",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "es",
	},
	LanguageBE: {
		ID:         "belarusian",
		Code:       "be",
		NativeName: "беларуская мова",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "by",
	},
	LanguageBN: {
		ID:         "bengali",
		Code:       "bn",
		NativeName: "বাংলা",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "bd",
	},
	LanguageBH: {
		ID:         "bihari",
		Code:       "bh",
		NativeName: "भोजपुरी",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageBI: {
		ID:         "bislama",
		Code:       "bi",
		NativeName: "Bislama",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "vu",
	},
	LanguageBS: {
		ID:         "bosnian",
		Code:       "bs",
		NativeName: "bosanski jezik",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ba",
	},
	LanguageBR: {
		ID:         "breton",
		Code:       "br",
		NativeName: "brezhoneg",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "fr",
	},
	LanguageBG: {
		ID:         "bulgarian",
		Code:       "bg",
		NativeName: "български език",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "bg",
	},
	LanguageMY: {
		ID:         "burmese",
		Code:       "my",
		NativeName: "ဗမာစာ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mm",
	},
	LanguageCA: {
		ID:         "catalan",
		Code:       "ca",
		NativeName: "català",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "es",
	},
	LanguageCH: {
		ID:         "chamorro",
		Code:       "ch",
		NativeName: "Chamoru",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gu",
	},
	LanguageCE: {
		ID:         "chechen",
		Code:       "ce",
		NativeName: "нохчийн мотт",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ru",
	},
	LanguageCU: {
		ID:         "churchSlavic",
		Code:       "cu",
		NativeName: "ѩзыкъ словѣньскъ",
		Dialects: []LanguageDialect{
		},
	},
	LanguageNY: {
		ID:         "chichewa",
		Code:       "ny",
		NativeName: "chiCheŵa",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mw",
	},
	LanguageZH: {
		ID:         "chinese",
		Code:       "zh",
		NativeName: "中文",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cn",
	},
	LanguageCV: {
		ID:         "chuvash",
		Code:       "cv",
		NativeName: "чӑваш чӗлхи",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ru",
	},
	LanguageKW: {
		ID:         "cornish",
		Code:       "kw",
		NativeName: "Kernewek",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gb",
	},
	LanguageCO: {
		ID:         "corsican",
		Code:       "co",
		NativeName: "corsu",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "fr",
	},
	LanguageCR: {
		ID:         "cree",
		Code:       "cr",
		NativeName: "ᓀᐦᐃᔭᐍᐏᐣ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ca",
	},
	LanguageHR: {
		ID:         "croatian",
		Code:       "hr",
		NativeName: "hrvatski jezik",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "hr",
	},
	LanguageCS: {
		ID:         "czech",
		Code:       "cs",
		NativeName: "čeština",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cz",
	},
	LanguageDV: {
		ID:         "divehi",
		Code:       "dv",
		NativeName: "ދިވެހި",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mv",
	},
	LanguageNL: {
		ID:         "dutch",
		Code:       "nl",
		NativeName: "Nederlands",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "nl",
	},
	LanguageDZ: {
		ID:         "dzongkha",
		Code:       "dz",
		NativeName: "རྫོང་ཁ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "bt",
	},
	LanguageEO: {
		ID:         "esperanto",
		Code:       "eo",
		NativeName: "Esperanto",
		Dialects: []LanguageDialect{
		},
	},
	LanguageET: {
		ID:         "estonian",
		Code:       "et",
		NativeName: "eesti",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ee",
	},
	LanguageEE: {
		ID:         "ewe",
		Code:       "ee",
		NativeName: "Eʋegbe",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gh",
	},
	LanguageFO: {
		ID:         "faroese",
		Code:       "fo",
		NativeName: "føroyskt",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "fo",
	},
	LanguageFJ: {
		ID:         "fijian",
		Code:       "fj",
		NativeName: "vosa Vakaviti",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "fj",
	},
	LanguageFI: {
		ID:         "finnish",
		Code:       "fi",
		NativeName: "suomi",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "fi",
	},
	LanguageFR: {
		ID:         "french",
		Code:       "fr",
		NativeName: "français",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "fr",
	},
	LanguageFY: {
		ID:         "westernFrisian",
		Code:       "fy",
		NativeName: "Frysk",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "nl",
	},
	LanguageFF: {
		ID:         "fulah",
		Code:       "ff",
		NativeName: "Fulfulde",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gn",
	},
	LanguageGD: {
		ID:         "gaelic",
		Code:       "gd",
		NativeName: "Gàidhlig",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gb",
	},
	LanguageGL: {
		ID:         "galician",
		Code:       "gl",
		NativeName: "galego",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "es",
	},
	LanguageLG: {
		ID:         "ganda",
		Code:       "lg",
		NativeName: "Luganda",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ug",
	},
	LanguageKA: {
		ID:         "georgian",
		Code:       "ka",
		NativeName: "ქართული",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ge",
	},
	LanguageEL: {
		ID:         "greek",
		Code:       "el",
		NativeName: "Ελληνικά",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gr",
	},
	LanguageKL: {
		ID:         "greenlandic",
		Code:       "kl",
		NativeName: "kalaallisut",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gl",
	},
	LanguageGN: {
		ID:         "guarani",
		Code:       "gn",
		NativeName: "Avañe'ẽ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "py",
	},
	LanguageGU: {
		ID:         "gujarati",
		Code:       "gu",
		NativeName: "ગુજરાતી",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageHT: {
		ID:         "haitian",
		Code:       "ht",
		NativeName: "Kreyòl ayisyen",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ht",
	},
	LanguageHA: {
		ID:         "hausa",
		Code:       "ha",
		NativeName: "هَوُسَ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ng",
	},
	LanguageHE: {
		ID:         "hebrew",
		Code:       "he",
		NativeName: "עברית",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "il",
	},
	LanguageHZ: {
		ID:         "herero",
		Code:       "hz",
		NativeName: "Otjiherero",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "na",
	},
	LanguageHI: {
		ID:         "hindi",
		Code:       "hi",
		NativeName: "हिन्दी",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageHO: {
		ID:         "hiriMotu",
		Code:       "ho",
		NativeName: "Hiri Motu",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "pg",
	},
	LanguageHU: {
		ID:         "hungarian",
		Code:       "hu",
		NativeName: "magyar",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "hu",
	},
	LanguageIS: {
		ID:         "icelandic",
		Code:       "is",
		NativeName: "Íslenska",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "is",
	},
	LanguageIO: {
		ID:         "ido",
		Code:       "io",
		NativeName: "Ido",
		Dialects: []LanguageDialect{
		},
	},
	LanguageIG: {
		ID:         "igbo",
		Code:       "ig",
		NativeName: "Asụsụ Igbo",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ng",
	},
	LanguageID: {
		ID:         "indonesian",
		Code:       "id",
		NativeName: "Bahasa Indonesia",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "id",
	},
	LanguageIA: {
		ID:         "interlingua",
		Code:       "ia",
		NativeName: "Interlingua",
		Dialects: []LanguageDialect{
		},
	},
	LanguageIE: {
		ID:         "interlingue",
		Code:       "ie",
		NativeName: "Interlingue",
		Dialects: []LanguageDialect{
		},
	},
	LanguageII: {
		ID:         "sichuanYi",
		Code:       "ii",
		NativeName: "ꆈꌠ꒿ Nuosuhxop",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cn",
	},
	LanguageIU: {
		ID:         "inuktitut",
		Code:       "iu",
		NativeName: "ᐃᓄᒃᑎᑐᑦ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ca",
	},
	LanguageIK: {
		ID:         "inupiaq",
		Code:       "ik",
		NativeName: "Iñupiaq",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "us",
	},
	LanguageGA: {
		ID:         "irish",
		Code:       "ga",
		NativeName: "Gaeilge",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ie",
	},
	LanguageIT: {
		ID:         "italian",
		Code:       "it",
		NativeName: "italiano",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "it",
	},
	LanguageJA: {
		ID:         "japanese",
		Code:       "ja",
		NativeName: "日本語",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "jp",
	},
	LanguageJV: {
		ID:         "javanese",
		Code:       "jv",
		NativeName: "basa Jawa",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "id",
	},
	LanguageKN: {
		ID:         "kannada",
		Code:       "kn",
		NativeName: "ಕನ್ನಡ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageKR: {
		ID:         "kanuri",
		Code:       "kr",
		NativeName: "Kanuri",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ng",
	},
	LanguageKS: {
		ID:         "kashmiri",
		Code:       "ks",
		NativeName: "कश्मीरी",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageKK: {
		ID:         "kazakh",
		Code:       "kk",
		NativeName: "қазақ тілі",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "kz",
	},
	LanguageKM: {
		ID:         "khmer",
		Code:       "km",
		NativeName: "ភាសាខ្មែរ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "kh",
	},
	LanguageKI: {
		ID:         "kikuyu",
		Code:       "ki",
		NativeName: "Gĩkũyũ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ke",
	},
	LanguageRW: {
		ID:         "kinyarwanda",
		Code:       "rw",
		NativeName: "Ikinyarwanda",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "rw",
	},
	LanguageKY: {
		ID:         "kirghiz",
		Code:       "ky",
		NativeName: "Кыргызча",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "kg",
	},
	LanguageKV: {
		ID:         "komi",
		Code:       "kv",
		NativeName: "коми кыв",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ru",
	},
	LanguageKG: {
		ID:         "kongo",
		Code:       "kg",
		NativeName: "Kikongo",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cd",
	},
	LanguageKO: {
		ID:         "korean",
		Code:       "ko",
		NativeName: "한국어",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "kr",
	},
	LanguageKJ: {
		ID:         "kuanyama",
		Code:       "kj",
		NativeName: "Kuanyama",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "na",
	},
	LanguageKU: {
		ID:         "kurdish",
		Code:       "ku",
		NativeName: "Kurdî",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "iq",
	},
	LanguageLO: {
		ID:         "lao",
		Code:       "lo",
		NativeName: "ພາສາລາວ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "la",
	},
	LanguageLA: {
		ID:         "latin",
		Code:       "la",
		NativeName: "latine",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "va",
	},
	LanguageLV: {
		ID:         "latvian",
		Code:       "lv",
		NativeName: "latviešu valoda",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "lv",
	},
	LanguageLI: {
		ID:         "limburgish",
		Code:       "li",
		NativeName: "Limburgs",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "nl",
	},
	LanguageLN: {
		ID:         "lingala",
		Code:       "ln",
		NativeName: "Lingála",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cd",
	},
	LanguageLT: {
		ID:         "lithuanian",
		Code:       "lt",
		NativeName: "lietuvių kalba",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "lt",
	},
	LanguageLU: {
		ID:         "lubaKatanga",
		Code:       "lu",
		NativeName: "Tshiluba",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cd",
	},
	LanguageLB: {
		ID:         "luxembourgish",
		Code:       "lb",
		NativeName: "Lëtzebuergesch",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "lu",
	},
	LanguageMK: {
		ID:         "macedonian",
		Code:       "mk",
		NativeName: "македонски јазик",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mk",
	},
	LanguageMG: {
		ID:         "malagasy",
		Code:       "mg",
		NativeName: "fiteny malagasy",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mg",
	},
	LanguageMS: {
		ID:         "malay",
		Code:       "ms",
		NativeName: "Bahasa Melayu",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "my",
	},
	LanguageML: {
		ID:         "malayalam",
		Code:       "ml",
		NativeName: "മലയാളം",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageMT: {
		ID:         "maltese",
		Code:       "mt",
		NativeName: "Malti",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mt",
	},
	LanguageGV: {
		ID:         "manx",
		Code:       "gv",
		NativeName: "Gaelg",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "im",
	},
	LanguageMI: {
		ID:         "maori",
		Code:       "mi",
		NativeName: "te reo Māori",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "nz",
	},
	LanguageMR: {
		ID:         "marathi",
		Code:       "mr",
		NativeName: "मराठी",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageMH: {
		ID:         "marshallese",
		Code:       "mh",
		NativeName: "Kajin M̧ajeļ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mh",
	},
	LanguageMN: {
		ID:         "mongolian",
		Code:       "mn",
		NativeName: "Монгол хэл",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "mn",
	},
	LanguageNA: {
		ID:         "nauru",
		Code:       "na",
		NativeName: "Ekakairũ Naoero",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "nr",
	},
	LanguageNV: {
		ID:         "navajo",
		Code:       "nv",
		NativeName: "Diné bizaad",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "us",
	},
	LanguageND: {
		ID:         "northNdebele",
		Code:       "nd",
		NativeName: "isiNdebele",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "zw",
	},
	LanguageNR: {
		ID:         "southNdebele",
		Code:       "nr",
		NativeName: "isiNdebele",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "za",
	},
	LanguageNG: {
		ID:         "ndonga",
		Code:       "ng",
		NativeName: "Owambo",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "na",
	},
	LanguageNE: {
		ID:         "nepali",
		Code:       "ne",
		NativeName: "नेपाली",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "np",
	},
	LanguageNO: {
		ID:         "norwegian",
		Code:       "no",
		NativeName: "Norsk",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "no",
	},
	LanguageNB: {
		ID:         "norwegianBokmal",
		Code:       "nb",
		NativeName: "Norsk bokmål",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "no",
	},
	LanguageNN: {
		ID:         "norwegianNynorsk",
		Code:       "nn",
		NativeName: "Norsk nynorsk",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "no",
	},
	LanguageOC: {
		ID:         "occitan",
		Code:       "oc",
		NativeName: "occitan",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "fr",
	},
	LanguageOJ: {
		ID:         "ojibwa",
		Code:       "oj",
		NativeName: "ᐊᓂᔑᓈᐯᒧᐎᓐ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ca",
	},
	LanguageOR: {
		ID:         "oriya",
		Code:       "or",
		NativeName: "ଓଡ଼ିଆ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageOM: {
		ID:         "oromo",
		Code:       "om",
		NativeName: "Afaan Oromoo",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "et",
	},
	LanguageOS: {
		ID:         "ossetian",
		Code:       "os",
		NativeName: "ирон æвзаг",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ge",
	},
	LanguagePI: {
		ID:         "pali",
		Code:       "pi",
		NativeName: "पाऴि",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguagePS: {
		ID:         "pashto",
		Code:       "ps",
		NativeName: "پښتو",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "af",
	},
	LanguageFA: {
		ID:         "persian",
		Code:       "fa",
		NativeName: "فارسی",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ir",
	},
	LanguagePL: {
		ID:         "polish",
		Code:       "pl",
		NativeName: "język polski",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "pl",
	},
	LanguagePT: {
		ID:         "portuguese",
		Code:       "pt",
		NativeName: "português",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "pt",
	},
	LanguagePA: {
		ID:         "punjabi",
		Code:       "pa",
		NativeName: "ਪੰਜਾਬੀ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageQU: {
		ID:         "quechua",
		Code:       "qu",
		NativeName: "Runa Simi",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "pe",
	},
	LanguageRO: {
		ID:         "romanian",
		Code:       "ro",
		NativeName: "limba română",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ro",
	},
	LanguageRM: {
		ID:         "romansh",
		Code:       "rm",
		NativeName: "rumantsch grischun",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ch",
	},
	LanguageRN: {
		ID:         "rundi",
		Code:       "rn",
		NativeName: "Ikirundi",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "bi",
	},
	LanguageRU: {
		ID:         "russian",
		Code:       "ru",
		NativeName: "русский язык",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ru",
	},
	LanguageSE: {
		ID:         "northernSami",
		Code:       "se",
		NativeName: "Davvisámegiella",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "no",
	},
	LanguageSM: {
		ID:         "samoan",
		Code:       "sm",
		NativeName: "gagana fa'a Samoa",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ws",
	},
	LanguageSG: {
		ID:         "sango",
		Code:       "sg",
		NativeName: "yângâ tî sängö",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cf",
	},
	LanguageSA: {
		ID:         "sanskrit",
		Code:       "sa",
		NativeName: "संस्कृतम्",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageSC: {
		ID:         "sardinian",
		Code:       "sc",
		NativeName: "sardu",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "it",
	},
	LanguageSR: {
		ID:         "serbian",
		Code:       "sr",
		NativeName: "српски језик",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "rs",
	},
	LanguageSD: {
		ID:         "sindhi",
		Code:       "sd",
		NativeName: "सिन्धी",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "pk",
	},
	LanguageSI: {
		ID:         "sinhala",
		Code:       "si",
		NativeName: "සිංහල",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "lk",
	},
	LanguageSK: {
		ID:         "slovak",
		Code:       "sk",
		NativeName: "slovenčina",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "sk",
	},
	LanguageSL: {
		ID:         "slovenian",
		Code:       "sl",
		NativeName: "slovenščina",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "si",
	},
	LanguageSO: {
		ID:         "somali",
		Code:       "so",
		NativeName: "Soomaaliga",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "so",
	},
	LanguageST: {
		ID:         "southernSotho",
		Code:       "st",
		NativeName: "Sesotho",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "za",
	},
	LanguageSN: {
		ID:         "shona",
		Code:       "sn",
		NativeName: "chiShona",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "zw",
	},
	LanguageES: {
		ID:         "spanish",
		Code:       "es",
		NativeName: "español",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "es",
	},
	LanguageSU: {
		ID:         "sundanese",
		Code:       "su",
		NativeName: "Basa Sunda",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "id",
	},
	LanguageSW: {
		ID:         "swahili",
		Code:       "sw",
		NativeName: "Kiswahili",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "tz",
	},
	LanguageSS: {
		ID:         "swati",
		Code:       "ss",
		NativeName: "SiSwati",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "sz",
	},
	LanguageSV: {
		ID:         "swedish",
		Code:       "sv",
		NativeName: "svenska",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "se",
	},
	LanguageTL: {
		ID:         "tagalog",
		Code:       "tl",
		NativeName: "Wikang Tagalog",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ph",
	},
	LanguageTY: {
		ID:         "tahitian",
		Code:       "ty",
		NativeName: "Reo Tahiti",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "pf",
	},
	LanguageTG: {
		ID:         "tajik",
		Code:       "tg",
		NativeName: "тоҷикӣ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "tj",
	},
	LanguageTA: {
		ID:         "tamil",
		Code:       "ta",
		NativeName: "தமிழ்",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageTT: {
		ID:         "tatar",
		Code:       "tt",
		NativeName: "татар теле",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ru",
	},
	LanguageTE: {
		ID:         "telugu",
		Code:       "te",
		NativeName: "తెలుగు",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "in",
	},
	LanguageTH: {
		ID:         "thai",
		Code:       "th",
		NativeName: "ไทย",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "th",
	},
	LanguageBO: {
		ID:         "tibetan",
		Code:       "bo",
		NativeName: "བོད་ཡིག",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cn",
	},
	LanguageTI: {
		ID:         "tigrinya",
		Code:       "ti",
		NativeName: "ትግርኛ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "er",
	},
	LanguageTO: {
		ID:         "tonga",
		Code:       "to",
		NativeName: "faka Tonga",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "to",
	},
	LanguageTS: {
		ID:         "tsonga",
		Code:       "ts",
		NativeName: "Xitsonga",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "za",
	},
	LanguageTN: {
		ID:         "tswana",
		Code:       "tn",
		NativeName: "Setswana",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "bw",
	},
	LanguageTR: {
		ID:         "turkish",
		Code:       "tr",
		NativeName: "Türkçe",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "tr",
	},
	LanguageTK: {
		ID:         "turkmen",
		Code:       "tk",
		NativeName: "Türkmen",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "tm",
	},
	LanguageTW: {
		ID:         "twi",
		Code:       "tw",
		NativeName: "Twi",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gh",
	},
	LanguageUG: {
		ID:         "uighur",
		Code:       "ug",
		NativeName: "ئۇيغۇرچە‎",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cn",
	},
	LanguageUK: {
		ID:         "ukrainian",
		Code:       "uk",
		NativeName: "українська мова",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ua",
	},
	LanguageUR: {
		ID:         "urdu",
		Code:       "ur",
		NativeName: "اردو",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "pk",
	},
	LanguageUZ: {
		ID:         "uzbek",
		Code:       "uz",
		NativeName: "Oʻzbek",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "uz",
	},
	LanguageVE: {
		ID:         "venda",
		Code:       "ve",
		NativeName: "Tshivenḓa",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "za",
	},
	LanguageVI: {
		ID:         "vietnamese",
		Code:       "vi",
		NativeName: "Tiếng Việt",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "vn",
	},
	LanguageVO: {
		ID:         "volapuk",
		Code:       "vo",
		NativeName: "Volapük",
		Dialects: []LanguageDialect{
		},
	},
	LanguageWA: {
		ID:         "walloon",
		Code:       "wa",
		NativeName: "walon",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "be",
	},
	LanguageCY: {
		ID:         "welsh",
		Code:       "cy",
		NativeName: "Cymraeg",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "gb",
	},
	LanguageWO: {
		ID:         "wolof",
		Code:       "wo",
		NativeName: "Wollof",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "sn",
	},
	LanguageXH: {
		ID:         "xhosa",
		Code:       "xh",
		NativeName: "isiXhosa",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "za",
	},
	LanguageYI: {
		ID:         "yiddish",
		Code:       "yi",
		NativeName: "ייִדיש",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "il",
	},
	LanguageYO: {
		ID:         "yoruba",
		Code:       "yo",
		NativeName: "Yorùbá",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "ng",
	},
	LanguageZA: {
		ID:         "zhuang",
		Code:       "za",
		NativeName: "Saɯ cueŋƅ",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "cn",
	},
	LanguageZU: {
		ID:         "zulu",
		Code:       "zu",
		NativeName: "isiZulu",
		Dialects: []LanguageDialect{
		},
		DefaultFlagCode: "za",
	},
}

// LanguageByCode returns a language by its code
func LanguageByCode(code string) (Language, bool) {
	l, ok := languages[LanguageCode(strings.ToLower(code))]
	return l, ok
}

// AllLanguages returns all languages
func AllLanguages() []Language {
	result := make([]Language, 0, len(languages))
	for _, l := range languages {
		result = append(result, l)
	}
	return result
}

// EmojiFlagFromLanguage returns the emoji flag for a language code.
// Returns the emoji flag string and true if found, empty string and false otherwise.
func EmojiFlagFromLanguage(languageCode string) (string, bool) {
	lang, ok := LanguageByCode(languageCode)
	if !ok || lang.DefaultFlagCode == "" {
		return "", false
	}
	code := strings.ToUpper(lang.DefaultFlagCode)
	if len(code) != 2 {
		return "", false
	}
	base := rune(0x1F1E6 - 65)
	return string([]rune{rune(code[0]) + base, rune(code[1]) + base}), true
}
