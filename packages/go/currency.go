// GENERATED CODE - DO NOT MODIFY BY HAND

package intl

import "strings"

// CurrencyCode represents an ISO 4217 currency code.
type CurrencyCode string

// Currency codes.
const (
	CurrencyAED CurrencyCode = "AED"
	CurrencyAFN CurrencyCode = "AFN"
	CurrencyALL CurrencyCode = "ALL"
	CurrencyAMD CurrencyCode = "AMD"
	CurrencyANG CurrencyCode = "ANG"
	CurrencyAOA CurrencyCode = "AOA"
	CurrencyARS CurrencyCode = "ARS"
	CurrencyAUD CurrencyCode = "AUD"
	CurrencyAWG CurrencyCode = "AWG"
	CurrencyAZN CurrencyCode = "AZN"
	CurrencyBAM CurrencyCode = "BAM"
	CurrencyBBD CurrencyCode = "BBD"
	CurrencyBDT CurrencyCode = "BDT"
	CurrencyBGN CurrencyCode = "BGN"
	CurrencyBHD CurrencyCode = "BHD"
	CurrencyBIF CurrencyCode = "BIF"
	CurrencyBMD CurrencyCode = "BMD"
	CurrencyBND CurrencyCode = "BND"
	CurrencyBOB CurrencyCode = "BOB"
	CurrencyBOV CurrencyCode = "BOV"
	CurrencyBRL CurrencyCode = "BRL"
	CurrencyBSD CurrencyCode = "BSD"
	CurrencyBTN CurrencyCode = "BTN"
	CurrencyBWP CurrencyCode = "BWP"
	CurrencyBYN CurrencyCode = "BYN"
	CurrencyBZD CurrencyCode = "BZD"
	CurrencyCAD CurrencyCode = "CAD"
	CurrencyCDF CurrencyCode = "CDF"
	CurrencyCHE CurrencyCode = "CHE"
	CurrencyCHF CurrencyCode = "CHF"
	CurrencyCHW CurrencyCode = "CHW"
	CurrencyCLF CurrencyCode = "CLF"
	CurrencyCLP CurrencyCode = "CLP"
	CurrencyCNY CurrencyCode = "CNY"
	CurrencyCOP CurrencyCode = "COP"
	CurrencyCOU CurrencyCode = "COU"
	CurrencyCRC CurrencyCode = "CRC"
	CurrencyCUP CurrencyCode = "CUP"
	CurrencyCVE CurrencyCode = "CVE"
	CurrencyCZK CurrencyCode = "CZK"
	CurrencyDJF CurrencyCode = "DJF"
	CurrencyDKK CurrencyCode = "DKK"
	CurrencyDOP CurrencyCode = "DOP"
	CurrencyDZD CurrencyCode = "DZD"
	CurrencyEGP CurrencyCode = "EGP"
	CurrencyERN CurrencyCode = "ERN"
	CurrencyETB CurrencyCode = "ETB"
	CurrencyEUR CurrencyCode = "EUR"
	CurrencyFJD CurrencyCode = "FJD"
	CurrencyFKP CurrencyCode = "FKP"
	CurrencyGBP CurrencyCode = "GBP"
	CurrencyGEL CurrencyCode = "GEL"
	CurrencyGHS CurrencyCode = "GHS"
	CurrencyGIP CurrencyCode = "GIP"
	CurrencyGMD CurrencyCode = "GMD"
	CurrencyGNF CurrencyCode = "GNF"
	CurrencyGTQ CurrencyCode = "GTQ"
	CurrencyGYD CurrencyCode = "GYD"
	CurrencyHKD CurrencyCode = "HKD"
	CurrencyHNL CurrencyCode = "HNL"
	CurrencyHTG CurrencyCode = "HTG"
	CurrencyHUF CurrencyCode = "HUF"
	CurrencyIDR CurrencyCode = "IDR"
	CurrencyILS CurrencyCode = "ILS"
	CurrencyINR CurrencyCode = "INR"
	CurrencyIQD CurrencyCode = "IQD"
	CurrencyIRR CurrencyCode = "IRR"
	CurrencyISK CurrencyCode = "ISK"
	CurrencyJMD CurrencyCode = "JMD"
	CurrencyJOD CurrencyCode = "JOD"
	CurrencyJPY CurrencyCode = "JPY"
	CurrencyKES CurrencyCode = "KES"
	CurrencyKGS CurrencyCode = "KGS"
	CurrencyKHR CurrencyCode = "KHR"
	CurrencyKMF CurrencyCode = "KMF"
	CurrencyKPW CurrencyCode = "KPW"
	CurrencyKRW CurrencyCode = "KRW"
	CurrencyKWD CurrencyCode = "KWD"
	CurrencyKYD CurrencyCode = "KYD"
	CurrencyKZT CurrencyCode = "KZT"
	CurrencyLAK CurrencyCode = "LAK"
	CurrencyLBP CurrencyCode = "LBP"
	CurrencyLKR CurrencyCode = "LKR"
	CurrencyLRD CurrencyCode = "LRD"
	CurrencyLSL CurrencyCode = "LSL"
	CurrencyLYD CurrencyCode = "LYD"
	CurrencyMAD CurrencyCode = "MAD"
	CurrencyMDL CurrencyCode = "MDL"
	CurrencyMGA CurrencyCode = "MGA"
	CurrencyMKD CurrencyCode = "MKD"
	CurrencyMMK CurrencyCode = "MMK"
	CurrencyMNT CurrencyCode = "MNT"
	CurrencyMOP CurrencyCode = "MOP"
	CurrencyMRU CurrencyCode = "MRU"
	CurrencyMUR CurrencyCode = "MUR"
	CurrencyMVR CurrencyCode = "MVR"
	CurrencyMWK CurrencyCode = "MWK"
	CurrencyMXN CurrencyCode = "MXN"
	CurrencyMXV CurrencyCode = "MXV"
	CurrencyMYR CurrencyCode = "MYR"
	CurrencyMZN CurrencyCode = "MZN"
	CurrencyNAD CurrencyCode = "NAD"
	CurrencyNGN CurrencyCode = "NGN"
	CurrencyNIO CurrencyCode = "NIO"
	CurrencyNOK CurrencyCode = "NOK"
	CurrencyNPR CurrencyCode = "NPR"
	CurrencyNZD CurrencyCode = "NZD"
	CurrencyOMR CurrencyCode = "OMR"
	CurrencyPAB CurrencyCode = "PAB"
	CurrencyPEN CurrencyCode = "PEN"
	CurrencyPGK CurrencyCode = "PGK"
	CurrencyPHP CurrencyCode = "PHP"
	CurrencyPKR CurrencyCode = "PKR"
	CurrencyPLN CurrencyCode = "PLN"
	CurrencyPYG CurrencyCode = "PYG"
	CurrencyQAR CurrencyCode = "QAR"
	CurrencyRON CurrencyCode = "RON"
	CurrencyRSD CurrencyCode = "RSD"
	CurrencyRUB CurrencyCode = "RUB"
	CurrencyRWF CurrencyCode = "RWF"
	CurrencySAR CurrencyCode = "SAR"
	CurrencySBD CurrencyCode = "SBD"
	CurrencySCR CurrencyCode = "SCR"
	CurrencySDG CurrencyCode = "SDG"
	CurrencySEK CurrencyCode = "SEK"
	CurrencySGD CurrencyCode = "SGD"
	CurrencySHP CurrencyCode = "SHP"
	CurrencySLE CurrencyCode = "SLE"
	CurrencySOS CurrencyCode = "SOS"
	CurrencySRD CurrencyCode = "SRD"
	CurrencySSP CurrencyCode = "SSP"
	CurrencySTN CurrencyCode = "STN"
	CurrencySVC CurrencyCode = "SVC"
	CurrencySYP CurrencyCode = "SYP"
	CurrencySZL CurrencyCode = "SZL"
	CurrencyTHB CurrencyCode = "THB"
	CurrencyTJS CurrencyCode = "TJS"
	CurrencyTMT CurrencyCode = "TMT"
	CurrencyTND CurrencyCode = "TND"
	CurrencyTOP CurrencyCode = "TOP"
	CurrencyTRY CurrencyCode = "TRY"
	CurrencyTTD CurrencyCode = "TTD"
	CurrencyTWD CurrencyCode = "TWD"
	CurrencyTZS CurrencyCode = "TZS"
	CurrencyUAH CurrencyCode = "UAH"
	CurrencyUGX CurrencyCode = "UGX"
	CurrencyUSD CurrencyCode = "USD"
	CurrencyUSN CurrencyCode = "USN"
	CurrencyUYI CurrencyCode = "UYI"
	CurrencyUYU CurrencyCode = "UYU"
	CurrencyUYW CurrencyCode = "UYW"
	CurrencyUZS CurrencyCode = "UZS"
	CurrencyVED CurrencyCode = "VED"
	CurrencyVES CurrencyCode = "VES"
	CurrencyVND CurrencyCode = "VND"
	CurrencyVUV CurrencyCode = "VUV"
	CurrencyWST CurrencyCode = "WST"
	CurrencyXAF CurrencyCode = "XAF"
	CurrencyXAG CurrencyCode = "XAG"
	CurrencyXAU CurrencyCode = "XAU"
	CurrencyXBA CurrencyCode = "XBA"
	CurrencyXBB CurrencyCode = "XBB"
	CurrencyXBC CurrencyCode = "XBC"
	CurrencyXBD CurrencyCode = "XBD"
	CurrencyXCD CurrencyCode = "XCD"
	CurrencyXDR CurrencyCode = "XDR"
	CurrencyXOF CurrencyCode = "XOF"
	CurrencyXPD CurrencyCode = "XPD"
	CurrencyXPF CurrencyCode = "XPF"
	CurrencyXPT CurrencyCode = "XPT"
	CurrencyXSU CurrencyCode = "XSU"
	CurrencyXTS CurrencyCode = "XTS"
	CurrencyXUA CurrencyCode = "XUA"
	CurrencyXXX CurrencyCode = "XXX"
	CurrencyYER CurrencyCode = "YER"
	CurrencyZAR CurrencyCode = "ZAR"
	CurrencyZMW CurrencyCode = "ZMW"
	CurrencyZMG CurrencyCode = "ZMG"
	CurrencyZWL CurrencyCode = "ZWL"
)

// Currency represents a currency with its metadata.
type Currency struct {
	ID               string
	Code             string
	NativeName       string
	NativeNamePlural string
	Symbol           string
}

// String returns the currency code.
func (c CurrencyCode) String() string {
	return string(c)
}

// Currency returns the currency data for this code.
func (c CurrencyCode) Currency() Currency {
	return currencies[c]
}

var currencies = map[CurrencyCode]Currency{
	CurrencyAED: {
		ID:               "aed",
		Code:             "AED",
		NativeName:       "United Arab Emirates Dirham",
		NativeNamePlural: "United Arab Emirates Dirhams",
		Symbol:           "د.إ",
	},
	CurrencyAFN: {
		ID:               "afn",
		Code:             "AFN",
		NativeName:       "Afghan Afghani",
		NativeNamePlural: "Afghan Afghanis",
		Symbol:           "؋",
	},
	CurrencyALL: {
		ID:               "all",
		Code:             "ALL",
		NativeName:       "Albanian Lek",
		NativeNamePlural: "Albanian Leke",
		Symbol:           "L",
	},
	CurrencyAMD: {
		ID:               "amd",
		Code:             "AMD",
		NativeName:       "Armenian Dram",
		NativeNamePlural: "Armenian Drams",
		Symbol:           "֏",
	},
	CurrencyANG: {
		ID:               "ang",
		Code:             "ANG",
		NativeName:       "Netherlands Antillean Guilder",
		NativeNamePlural: "Netherlands Antillean Guilders",
		Symbol:           "ƒ",
	},
	CurrencyAOA: {
		ID:               "aoa",
		Code:             "AOA",
		NativeName:       "Angolan Kwanza",
		NativeNamePlural: "Angolan Kwanza",
		Symbol:           "Kz",
	},
	CurrencyARS: {
		ID:               "ars",
		Code:             "ARS",
		NativeName:       "Argentine Peso",
		NativeNamePlural: "Argentine Pesos",
		Symbol:           "\\$",
	},
	CurrencyAUD: {
		ID:               "aud",
		Code:             "AUD",
		NativeName:       "Australian Dollar",
		NativeNamePlural: "Australian Dollars",
		Symbol:           "\\$",
	},
	CurrencyAWG: {
		ID:               "awg",
		Code:             "AWG",
		NativeName:       "Aruban Florin",
		NativeNamePlural: "Aruban Florin",
		Symbol:           "ƒ",
	},
	CurrencyAZN: {
		ID:               "azn",
		Code:             "AZN",
		NativeName:       "Azerbaijani Manat",
		NativeNamePlural: "Azerbaijani Manat",
		Symbol:           "₼",
	},
	CurrencyBAM: {
		ID:               "bam",
		Code:             "BAM",
		NativeName:       "Bosnia-Herzegovina Convertible Mark",
		NativeNamePlural: "Bosnia-Herzegovina Convertible Marks",
		Symbol:           "KM",
	},
	CurrencyBBD: {
		ID:               "bbd",
		Code:             "BBD",
		NativeName:       "Barbadian Dollar",
		NativeNamePlural: "Barbadian Dollars",
		Symbol:           "\\$",
	},
	CurrencyBDT: {
		ID:               "bdt",
		Code:             "BDT",
		NativeName:       "Bangladeshi Taka",
		NativeNamePlural: "Bangladeshi Takas",
		Symbol:           "৳",
	},
	CurrencyBGN: {
		ID:               "bgn",
		Code:             "BGN",
		NativeName:       "Bulgarian Lev",
		NativeNamePlural: "Bulgarian Leva",
		Symbol:           "лв",
	},
	CurrencyBHD: {
		ID:               "bhd",
		Code:             "BHD",
		NativeName:       "Bahraini Dinar",
		NativeNamePlural: "Bahraini Dinars",
		Symbol:           "ب.د",
	},
	CurrencyBIF: {
		ID:               "bif",
		Code:             "BIF",
		NativeName:       "Burundian Franc",
		NativeNamePlural: "Burundian Francs",
		Symbol:           "FBu",
	},
	CurrencyBMD: {
		ID:               "bmd",
		Code:             "BMD",
		NativeName:       "Bermudian Dollar",
		NativeNamePlural: "Bermudian Dollars",
		Symbol:           "\\$",
	},
	CurrencyBND: {
		ID:               "bnd",
		Code:             "BND",
		NativeName:       "Brunei Dollar",
		NativeNamePlural: "Brunei Dollars",
		Symbol:           "\\$",
	},
	CurrencyBOB: {
		ID:               "bob",
		Code:             "BOB",
		NativeName:       "Bolivian Boliviano",
		NativeNamePlural: "Bolivian Bolivianos",
		Symbol:           "Bs.",
	},
	CurrencyBOV: {
		ID:               "bov",
		Code:             "BOV",
		NativeName:       "Bolivian Mvdol",
		NativeNamePlural: "Bolivian Mvdols",
		Symbol:           "BOV",
	},
	CurrencyBRL: {
		ID:               "brl",
		Code:             "BRL",
		NativeName:       "Brazilian Real",
		NativeNamePlural: "Brazilian Reals",
		Symbol:           "R\\$",
	},
	CurrencyBSD: {
		ID:               "bsd",
		Code:             "BSD",
		NativeName:       "Bahamian Dollar",
		NativeNamePlural: "Bahamian Dollars",
		Symbol:           "\\$",
	},
	CurrencyBTN: {
		ID:               "btn",
		Code:             "BTN",
		NativeName:       "Bhutanese Ngultrum",
		NativeNamePlural: "Bhutanese Ngultrums",
		Symbol:           "Nu.",
	},
	CurrencyBWP: {
		ID:               "bwp",
		Code:             "BWP",
		NativeName:       "Botswana Pula",
		NativeNamePlural: "Botswana Pulas",
		Symbol:           "P",
	},
	CurrencyBYN: {
		ID:               "byn",
		Code:             "BYN",
		NativeName:       "Belarusian Ruble",
		NativeNamePlural: "Belarusian Rubles",
		Symbol:           "Br",
	},
	CurrencyBZD: {
		ID:               "bzd",
		Code:             "BZD",
		NativeName:       "Belize Dollar",
		NativeNamePlural: "Belize Dollars",
		Symbol:           "BZ\\$",
	},
	CurrencyCAD: {
		ID:               "cad",
		Code:             "CAD",
		NativeName:       "Canadian Dollar",
		NativeNamePlural: "Canadian Dollars",
		Symbol:           "\\$",
	},
	CurrencyCDF: {
		ID:               "cdf",
		Code:             "CDF",
		NativeName:       "Congolese Franc",
		NativeNamePlural: "Congolese Francs",
		Symbol:           "FC",
	},
	CurrencyCHE: {
		ID:               "che",
		Code:             "CHE",
		NativeName:       "WIR Euro",
		NativeNamePlural: "WIR Euros",
		Symbol:           "CHE",
	},
	CurrencyCHF: {
		ID:               "chf",
		Code:             "CHF",
		NativeName:       "Swiss Franc",
		NativeNamePlural: "Swiss Francs",
		Symbol:           "CHF",
	},
	CurrencyCHW: {
		ID:               "chw",
		Code:             "CHW",
		NativeName:       "WIR Franc",
		NativeNamePlural: "WIR Francs",
		Symbol:           "CHW",
	},
	CurrencyCLF: {
		ID:               "clf",
		Code:             "CLF",
		NativeName:       "Chilean Unit of Account (UF)",
		NativeNamePlural: "Chilean Units of Account (UF)",
		Symbol:           "CLF",
	},
	CurrencyCLP: {
		ID:               "clp",
		Code:             "CLP",
		NativeName:       "Chilean Peso",
		NativeNamePlural: "Chilean Pesos",
		Symbol:           "\\$",
	},
	CurrencyCNY: {
		ID:               "cny",
		Code:             "CNY",
		NativeName:       "Chinese Yuan",
		NativeNamePlural: "Chinese Yuan",
		Symbol:           "¥",
	},
	CurrencyCOP: {
		ID:               "cop",
		Code:             "COP",
		NativeName:       "Colombian Peso",
		NativeNamePlural: "Colombian Pesos",
		Symbol:           "\\$",
	},
	CurrencyCOU: {
		ID:               "cou",
		Code:             "COU",
		NativeName:       "Unidad de Valor Real",
		NativeNamePlural: "Unidad de Valor Reals",
		Symbol:           "COU",
	},
	CurrencyCRC: {
		ID:               "crc",
		Code:             "CRC",
		NativeName:       "Costa Rican Colón",
		NativeNamePlural: "Costa Rican Colóns",
		Symbol:           "₡",
	},
	CurrencyCUP: {
		ID:               "cup",
		Code:             "CUP",
		NativeName:       "Cuban Peso",
		NativeNamePlural: "Cuban Pesos",
		Symbol:           "\\$",
	},
	CurrencyCVE: {
		ID:               "cve",
		Code:             "CVE",
		NativeName:       "Cape Verdean Escudo",
		NativeNamePlural: "Cape Verdean Escudos",
		Symbol:           "Esc",
	},
	CurrencyCZK: {
		ID:               "czk",
		Code:             "CZK",
		NativeName:       "Czech Koruna",
		NativeNamePlural: "Czech Korunas",
		Symbol:           "Kč",
	},
	CurrencyDJF: {
		ID:               "djf",
		Code:             "DJF",
		NativeName:       "Djiboutian Franc",
		NativeNamePlural: "Djiboutian Francs",
		Symbol:           "Fdj",
	},
	CurrencyDKK: {
		ID:               "dkk",
		Code:             "DKK",
		NativeName:       "Danske Kroner",
		NativeNamePlural: "Danske Kroner",
		Symbol:           "kr",
	},
	CurrencyDOP: {
		ID:               "dop",
		Code:             "DOP",
		NativeName:       "Dominican Peso",
		NativeNamePlural: "Dominican Pesos",
		Symbol:           "RD\\$",
	},
	CurrencyDZD: {
		ID:               "dzd",
		Code:             "DZD",
		NativeName:       "Algerian Dinar",
		NativeNamePlural: "Algerian Dinars",
		Symbol:           "د.ج",
	},
	CurrencyEGP: {
		ID:               "egp",
		Code:             "EGP",
		NativeName:       "Egyptian Pound",
		NativeNamePlural: "Egyptian Pounds",
		Symbol:           "E£",
	},
	CurrencyERN: {
		ID:               "ern",
		Code:             "ERN",
		NativeName:       "Eritrean Nakfa",
		NativeNamePlural: "Eritrean Nakfas",
		Symbol:           "Nfk",
	},
	CurrencyETB: {
		ID:               "etb",
		Code:             "ETB",
		NativeName:       "Ethiopian Birr",
		NativeNamePlural: "Ethiopian Birrs",
		Symbol:           "Br",
	},
	CurrencyEUR: {
		ID:               "eur",
		Code:             "EUR",
		NativeName:       "Euro",
		NativeNamePlural: "Euros",
		Symbol:           "€",
	},
	CurrencyFJD: {
		ID:               "fjd",
		Code:             "FJD",
		NativeName:       "Fijian Dollar",
		NativeNamePlural: "Fijian Dollars",
		Symbol:           "\\$",
	},
	CurrencyFKP: {
		ID:               "fkp",
		Code:             "FKP",
		NativeName:       "Falkland Islands Pound",
		NativeNamePlural: "Falkland Islands Pounds",
		Symbol:           "£",
	},
	CurrencyGBP: {
		ID:               "gbp",
		Code:             "GBP",
		NativeName:       "British Pound",
		NativeNamePlural: "British Pounds",
		Symbol:           "£",
	},
	CurrencyGEL: {
		ID:               "gel",
		Code:             "GEL",
		NativeName:       "Georgian Lari",
		NativeNamePlural: "Georgian Laris",
		Symbol:           "₾",
	},
	CurrencyGHS: {
		ID:               "ghs",
		Code:             "GHS",
		NativeName:       "Ghanaian Cedi",
		NativeNamePlural: "Ghanaian Cedis",
		Symbol:           "₵",
	},
	CurrencyGIP: {
		ID:               "gip",
		Code:             "GIP",
		NativeName:       "Gibraltar Pound",
		NativeNamePlural: "Gibraltar Pounds",
		Symbol:           "£",
	},
	CurrencyGMD: {
		ID:               "gmd",
		Code:             "GMD",
		NativeName:       "Gambian Dalasi",
		NativeNamePlural: "Gambian Dalasis",
		Symbol:           "D",
	},
	CurrencyGNF: {
		ID:               "gnf",
		Code:             "GNF",
		NativeName:       "Guinean Franc",
		NativeNamePlural: "Guinean Francs",
		Symbol:           "FG",
	},
	CurrencyGTQ: {
		ID:               "gtq",
		Code:             "GTQ",
		NativeName:       "Guatemalan Quetzal",
		NativeNamePlural: "Guatemalan Quetzals",
		Symbol:           "Q",
	},
	CurrencyGYD: {
		ID:               "gyd",
		Code:             "GYD",
		NativeName:       "Guyanese Dollar",
		NativeNamePlural: "Guyanese Dollars",
		Symbol:           "\\$",
	},
	CurrencyHKD: {
		ID:               "hkd",
		Code:             "HKD",
		NativeName:       "Hong Kong Dollar",
		NativeNamePlural: "Hong Kong Dollars",
		Symbol:           "\\$",
	},
	CurrencyHNL: {
		ID:               "hnl",
		Code:             "HNL",
		NativeName:       "Honduran Lempira",
		NativeNamePlural: "Honduran Lempiras",
		Symbol:           "L",
	},
	CurrencyHTG: {
		ID:               "htg",
		Code:             "HTG",
		NativeName:       "Haitian Gourde",
		NativeNamePlural: "Haitian Gourdes",
		Symbol:           "G",
	},
	CurrencyHUF: {
		ID:               "huf",
		Code:             "HUF",
		NativeName:       "Hungarian Forint",
		NativeNamePlural: "Hungarian Forints",
		Symbol:           "Ft",
	},
	CurrencyIDR: {
		ID:               "idr",
		Code:             "IDR",
		NativeName:       "Indonesian Rupiah",
		NativeNamePlural: "Indonesian Rupiahs",
		Symbol:           "Rp",
	},
	CurrencyILS: {
		ID:               "ils",
		Code:             "ILS",
		NativeName:       "Israeli New Shekel",
		NativeNamePlural: "Israeli New Shekels",
		Symbol:           "₪",
	},
	CurrencyINR: {
		ID:               "inr",
		Code:             "INR",
		NativeName:       "Indian Rupee",
		NativeNamePlural: "Indian Rupees",
		Symbol:           "₹",
	},
	CurrencyIQD: {
		ID:               "iqd",
		Code:             "IQD",
		NativeName:       "Iraqi Dinar",
		NativeNamePlural: "Iraqi Dinars",
		Symbol:           "ع.د",
	},
	CurrencyIRR: {
		ID:               "irr",
		Code:             "IRR",
		NativeName:       "Iranian Rial",
		NativeNamePlural: "Iranian Rials",
		Symbol:           "﷼",
	},
	CurrencyISK: {
		ID:               "isk",
		Code:             "ISK",
		NativeName:       "Icelandic Króna",
		NativeNamePlural: "Icelandic Krónur",
		Symbol:           "kr",
	},
	CurrencyJMD: {
		ID:               "jmd",
		Code:             "JMD",
		NativeName:       "Jamaican Dollar",
		NativeNamePlural: "Jamaican Dollars",
		Symbol:           "J\\$",
	},
	CurrencyJOD: {
		ID:               "jod",
		Code:             "JOD",
		NativeName:       "Jordanian Dinar",
		NativeNamePlural: "Jordanian Dinars",
		Symbol:           "د.ا",
	},
	CurrencyJPY: {
		ID:               "jpy",
		Code:             "JPY",
		NativeName:       "Japanese Yen",
		NativeNamePlural: "Japanese Yen",
		Symbol:           "¥",
	},
	CurrencyKES: {
		ID:               "kes",
		Code:             "KES",
		NativeName:       "Kenyan Shilling",
		NativeNamePlural: "Kenyan Shillings",
		Symbol:           "KSh",
	},
	CurrencyKGS: {
		ID:               "kgs",
		Code:             "KGS",
		NativeName:       "Kyrgyzstani Som",
		NativeNamePlural: "Kyrgyzstani Soms",
		Symbol:           "с",
	},
	CurrencyKHR: {
		ID:               "khr",
		Code:             "KHR",
		NativeName:       "Cambodian Riel",
		NativeNamePlural: "Cambodian Riels",
		Symbol:           "៛",
	},
	CurrencyKMF: {
		ID:               "kmf",
		Code:             "KMF",
		NativeName:       "Comorian Franc",
		NativeNamePlural: "Comorian Francs",
		Symbol:           "CF",
	},
	CurrencyKPW: {
		ID:               "kpw",
		Code:             "KPW",
		NativeName:       "North Korean Won",
		NativeNamePlural: "North Korean Won",
		Symbol:           "₩",
	},
	CurrencyKRW: {
		ID:               "krw",
		Code:             "KRW",
		NativeName:       "South Korean Won",
		NativeNamePlural: "South Korean Won",
		Symbol:           "₩",
	},
	CurrencyKWD: {
		ID:               "kwd",
		Code:             "KWD",
		NativeName:       "Kuwaiti Dinar",
		NativeNamePlural: "Kuwaiti Dinars",
		Symbol:           "د.ك",
	},
	CurrencyKYD: {
		ID:               "kyd",
		Code:             "KYD",
		NativeName:       "Cayman Islands Dollar",
		NativeNamePlural: "Cayman Islands Dollars",
		Symbol:           "\\$",
	},
	CurrencyKZT: {
		ID:               "kzt",
		Code:             "KZT",
		NativeName:       "Kazakhstani Tenge",
		NativeNamePlural: "Kazakhstani Tenges",
		Symbol:           "₸",
	},
	CurrencyLAK: {
		ID:               "lak",
		Code:             "LAK",
		NativeName:       "Lao Kip",
		NativeNamePlural: "Lao Kips",
		Symbol:           "₭",
	},
	CurrencyLBP: {
		ID:               "lbp",
		Code:             "LBP",
		NativeName:       "Lebanese Pound",
		NativeNamePlural: "Lebanese Pounds",
		Symbol:           "ل.ل",
	},
	CurrencyLKR: {
		ID:               "lkr",
		Code:             "LKR",
		NativeName:       "Sri Lankan Rupee",
		NativeNamePlural: "Sri Lankan Rupees",
		Symbol:           "Rs",
	},
	CurrencyLRD: {
		ID:               "lrd",
		Code:             "LRD",
		NativeName:       "Liberian Dollar",
		NativeNamePlural: "Liberian Dollars",
		Symbol:           "\\$",
	},
	CurrencyLSL: {
		ID:               "lsl",
		Code:             "LSL",
		NativeName:       "Lesotho Loti",
		NativeNamePlural: "Lesotho Lotis",
		Symbol:           "L",
	},
	CurrencyLYD: {
		ID:               "lyd",
		Code:             "LYD",
		NativeName:       "Libyan Dinar",
		NativeNamePlural: "Libyan Dinars",
		Symbol:           "ل.د",
	},
	CurrencyMAD: {
		ID:               "mad",
		Code:             "MAD",
		NativeName:       "Moroccan Dirham",
		NativeNamePlural: "Moroccan Dirhams",
		Symbol:           "د.م.",
	},
	CurrencyMDL: {
		ID:               "mdl",
		Code:             "MDL",
		NativeName:       "Moldovan Leu",
		NativeNamePlural: "Moldovan Lei",
		Symbol:           "L",
	},
	CurrencyMGA: {
		ID:               "mga",
		Code:             "MGA",
		NativeName:       "Malagasy Ariary",
		NativeNamePlural: "Malagasy Ariaries",
		Symbol:           "Ar",
	},
	CurrencyMKD: {
		ID:               "mkd",
		Code:             "MKD",
		NativeName:       "Macedonian Denar",
		NativeNamePlural: "Macedonian Denari",
		Symbol:           "ден",
	},
	CurrencyMMK: {
		ID:               "mmk",
		Code:             "MMK",
		NativeName:       "Myanmar Kyat",
		NativeNamePlural: "Myanmar Kyats",
		Symbol:           "K",
	},
	CurrencyMNT: {
		ID:               "mnt",
		Code:             "MNT",
		NativeName:       "Mongolian Tugrik",
		NativeNamePlural: "Mongolian Tugriks",
		Symbol:           "₮",
	},
	CurrencyMOP: {
		ID:               "mop",
		Code:             "MOP",
		NativeName:       "Macanese Pataca",
		NativeNamePlural: "Macanese Patacas",
		Symbol:           "MOP\\$",
	},
	CurrencyMRU: {
		ID:               "mru",
		Code:             "MRU",
		NativeName:       "Mauritanian Ouguiya",
		NativeNamePlural: "Mauritanian Ouguiyas",
		Symbol:           "UM",
	},
	CurrencyMUR: {
		ID:               "mur",
		Code:             "MUR",
		NativeName:       "Mauritian Rupee",
		NativeNamePlural: "Mauritian Rupees",
		Symbol:           "₨",
	},
	CurrencyMVR: {
		ID:               "mvr",
		Code:             "MVR",
		NativeName:       "Maldivian Rufiyaa",
		NativeNamePlural: "Maldivian Rufiyaas",
		Symbol:           "Rf",
	},
	CurrencyMWK: {
		ID:               "mwk",
		Code:             "MWK",
		NativeName:       "Malawian Kwacha",
		NativeNamePlural: "Malawian Kwachas",
		Symbol:           "MK",
	},
	CurrencyMXN: {
		ID:               "mxn",
		Code:             "MXN",
		NativeName:       "Mexican Peso",
		NativeNamePlural: "Mexican Pesos",
		Symbol:           "\\$",
	},
	CurrencyMXV: {
		ID:               "mxv",
		Code:             "MXV",
		NativeName:       "Mexican Unidad de Inversion (UDI)",
		NativeNamePlural: "Mexican Unidad de Inversions (UDI)",
		Symbol:           "MXV",
	},
	CurrencyMYR: {
		ID:               "myr",
		Code:             "MYR",
		NativeName:       "Malaysian Ringgit",
		NativeNamePlural: "Malaysian Ringgits",
		Symbol:           "RM",
	},
	CurrencyMZN: {
		ID:               "mzn",
		Code:             "MZN",
		NativeName:       "Mozambican Metical",
		NativeNamePlural: "Mozambican Meticals",
		Symbol:           "MT",
	},
	CurrencyNAD: {
		ID:               "nad",
		Code:             "NAD",
		NativeName:       "Namibian Dollar",
		NativeNamePlural: "Namibian Dollars",
		Symbol:           "\\$",
	},
	CurrencyNGN: {
		ID:               "ngn",
		Code:             "NGN",
		NativeName:       "Nigerian Naira",
		NativeNamePlural: "Nigerian Nairas",
		Symbol:           "₦",
	},
	CurrencyNIO: {
		ID:               "nio",
		Code:             "NIO",
		NativeName:       "Nicaraguan Córdoba",
		NativeNamePlural: "Nicaraguan Córdobas",
		Symbol:           "C\\$",
	},
	CurrencyNOK: {
		ID:               "nok",
		Code:             "NOK",
		NativeName:       "Norwegian Krone",
		NativeNamePlural: "Norwegian Kroner",
		Symbol:           "kr",
	},
	CurrencyNPR: {
		ID:               "npr",
		Code:             "NPR",
		NativeName:       "Nepalese Rupee",
		NativeNamePlural: "Nepalese Rupees",
		Symbol:           "रू",
	},
	CurrencyNZD: {
		ID:               "nzd",
		Code:             "NZD",
		NativeName:       "New Zealand Dollar",
		NativeNamePlural: "New Zealand Dollars",
		Symbol:           "\\$",
	},
	CurrencyOMR: {
		ID:               "omr",
		Code:             "OMR",
		NativeName:       "Omani Rial",
		NativeNamePlural: "Omani Rials",
		Symbol:           "ر.ع.",
	},
	CurrencyPAB: {
		ID:               "pab",
		Code:             "PAB",
		NativeName:       "Panamanian Balboa",
		NativeNamePlural: "Panamanian Balboas",
		Symbol:           "B/.",
	},
	CurrencyPEN: {
		ID:               "pen",
		Code:             "PEN",
		NativeName:       "Peruvian Nuevo Sol",
		NativeNamePlural: "Peruvian Nuevo Sols",
		Symbol:           "S/.",
	},
	CurrencyPGK: {
		ID:               "pgk",
		Code:             "PGK",
		NativeName:       "Papua New Guinean Kina",
		NativeNamePlural: "Papua New Guinean Kina",
		Symbol:           "K",
	},
	CurrencyPHP: {
		ID:               "php",
		Code:             "PHP",
		NativeName:       "Philippine Peso",
		NativeNamePlural: "Philippine Pesos",
		Symbol:           "₱",
	},
	CurrencyPKR: {
		ID:               "pkr",
		Code:             "PKR",
		NativeName:       "Pakistani Rupee",
		NativeNamePlural: "Pakistani Rupees",
		Symbol:           "₨",
	},
	CurrencyPLN: {
		ID:               "pln",
		Code:             "PLN",
		NativeName:       "Polish Złoty",
		NativeNamePlural: "Polish Złotys",
		Symbol:           "zł",
	},
	CurrencyPYG: {
		ID:               "pyg",
		Code:             "PYG",
		NativeName:       "Paraguayan Guaraní",
		NativeNamePlural: "Paraguayan Guaranís",
		Symbol:           "₲",
	},
	CurrencyQAR: {
		ID:               "qar",
		Code:             "QAR",
		NativeName:       "Qatari Rial",
		NativeNamePlural: "Qatari Rials",
		Symbol:           "ر.ق",
	},
	CurrencyRON: {
		ID:               "ron",
		Code:             "RON",
		NativeName:       "Romanian Leu",
		NativeNamePlural: "Romanian Lei",
		Symbol:           "lei",
	},
	CurrencyRSD: {
		ID:               "rsd",
		Code:             "RSD",
		NativeName:       "Serbian Dinar",
		NativeNamePlural: "Serbian Dinars",
		Symbol:           "дин",
	},
	CurrencyRUB: {
		ID:               "rub",
		Code:             "RUB",
		NativeName:       "Russian Ruble",
		NativeNamePlural: "Russian Rubles",
		Symbol:           "₽",
	},
	CurrencyRWF: {
		ID:               "rwf",
		Code:             "RWF",
		NativeName:       "Rwandan Franc",
		NativeNamePlural: "Rwandan Francs",
		Symbol:           "FRw",
	},
	CurrencySAR: {
		ID:               "sar",
		Code:             "SAR",
		NativeName:       "Saudi Riyal",
		NativeNamePlural: "Saudi Riyals",
		Symbol:           "ر.س",
	},
	CurrencySBD: {
		ID:               "sbd",
		Code:             "SBD",
		NativeName:       "Solomon Islands Dollar",
		NativeNamePlural: "Solomon Islands Dollars",
		Symbol:           "\\$",
	},
	CurrencySCR: {
		ID:               "scr",
		Code:             "SCR",
		NativeName:       "Seychellois Rupee",
		NativeNamePlural: "Seychellois Rupees",
		Symbol:           "₨",
	},
	CurrencySDG: {
		ID:               "sdg",
		Code:             "SDG",
		NativeName:       "Sudanese Pound",
		NativeNamePlural: "Sudanese Pounds",
		Symbol:           "ج.س.",
	},
	CurrencySEK: {
		ID:               "sek",
		Code:             "SEK",
		NativeName:       "Swedish Krona",
		NativeNamePlural: "Swedish Kronor",
		Symbol:           "kr",
	},
	CurrencySGD: {
		ID:               "sgd",
		Code:             "SGD",
		NativeName:       "Singapore Dollar",
		NativeNamePlural: "Singapore Dollars",
		Symbol:           "\\$",
	},
	CurrencySHP: {
		ID:               "shp",
		Code:             "SHP",
		NativeName:       "Saint Helena Pound",
		NativeNamePlural: "Saint Helena Pounds",
		Symbol:           "£",
	},
	CurrencySLE: {
		ID:               "sle",
		Code:             "SLE",
		NativeName:       "Sierra Leonean Leone",
		NativeNamePlural: "Sierra Leonean Leones",
		Symbol:           "Le",
	},
	CurrencySOS: {
		ID:               "sos",
		Code:             "SOS",
		NativeName:       "Somali Shilling",
		NativeNamePlural: "Somali Shillings",
		Symbol:           "Sh.So.",
	},
	CurrencySRD: {
		ID:               "srd",
		Code:             "SRD",
		NativeName:       "Surinamese Dollar",
		NativeNamePlural: "Surinamese Dollars",
		Symbol:           "\\$",
	},
	CurrencySSP: {
		ID:               "ssp",
		Code:             "SSP",
		NativeName:       "South Sudanese Pound",
		NativeNamePlural: "South Sudanese Pounds",
		Symbol:           "£",
	},
	CurrencySTN: {
		ID:               "stn",
		Code:             "STN",
		NativeName:       "São Tomé and Príncipe Dobra",
		NativeNamePlural: "São Tomé and Príncipe Dobras",
		Symbol:           "Db",
	},
	CurrencySVC: {
		ID:               "svc",
		Code:             "SVC",
		NativeName:       "Salvadoran Colón",
		NativeNamePlural: "Salvadoran Colóns",
		Symbol:           "₡",
	},
	CurrencySYP: {
		ID:               "syp",
		Code:             "SYP",
		NativeName:       "Syrian Pound",
		NativeNamePlural: "Syrian Pounds",
		Symbol:           "ل.س",
	},
	CurrencySZL: {
		ID:               "szl",
		Code:             "SZL",
		NativeName:       "Swazi Lilangeni",
		NativeNamePlural: "Swazi Emalangeni",
		Symbol:           "L",
	},
	CurrencyTHB: {
		ID:               "thb",
		Code:             "THB",
		NativeName:       "Thai Baht",
		NativeNamePlural: "Thai Baht",
		Symbol:           "฿",
	},
	CurrencyTJS: {
		ID:               "tjs",
		Code:             "TJS",
		NativeName:       "Tajikistani Somoni",
		NativeNamePlural: "Tajikistani Somonis",
		Symbol:           "ЅМ",
	},
	CurrencyTMT: {
		ID:               "tmt",
		Code:             "TMT",
		NativeName:       "Turkmenistani Manat",
		NativeNamePlural: "Turkmenistani Manats",
		Symbol:           "T",
	},
	CurrencyTND: {
		ID:               "tnd",
		Code:             "TND",
		NativeName:       "Tunisian Dinar",
		NativeNamePlural: "Tunisian Dinars",
		Symbol:           "د.ت",
	},
	CurrencyTOP: {
		ID:               "top",
		Code:             "TOP",
		NativeName:       "Tongan Pa'anga",
		NativeNamePlural: "Tongan Pa'anga",
		Symbol:           "T\\$",
	},
	CurrencyTRY: {
		ID:               "turkishLira",
		Code:             "TRY",
		NativeName:       "Turkish Lira",
		NativeNamePlural: "Turkish Lira",
		Symbol:           "₺",
	},
	CurrencyTTD: {
		ID:               "ttd",
		Code:             "TTD",
		NativeName:       "Trinidad and Tobago Dollar",
		NativeNamePlural: "Trinidad and Tobago Dollars",
		Symbol:           "TT\\$",
	},
	CurrencyTWD: {
		ID:               "twd",
		Code:             "TWD",
		NativeName:       "New Taiwan Dollar",
		NativeNamePlural: "New Taiwan Dollars",
		Symbol:           "NT\\$",
	},
	CurrencyTZS: {
		ID:               "tzs",
		Code:             "TZS",
		NativeName:       "Tanzanian Shilling",
		NativeNamePlural: "Tanzanian Shillings",
		Symbol:           "TSh",
	},
	CurrencyUAH: {
		ID:               "uah",
		Code:             "UAH",
		NativeName:       "Ukrainian Hryvnia",
		NativeNamePlural: "Ukrainian Hryvnias",
		Symbol:           "₴",
	},
	CurrencyUGX: {
		ID:               "ugx",
		Code:             "UGX",
		NativeName:       "Ugandan Shilling",
		NativeNamePlural: "Ugandan Shillings",
		Symbol:           "USh",
	},
	CurrencyUSD: {
		ID:               "usd",
		Code:             "USD",
		NativeName:       "United States Dollar",
		NativeNamePlural: "United States Dollars",
		Symbol:           "\\$",
	},
	CurrencyUSN: {
		ID:               "usn",
		Code:             "USN",
		NativeName:       "United States Dollar (Next day)",
		NativeNamePlural: "United States Dollars (Next day)",
		Symbol:           "USN",
	},
	CurrencyUYI: {
		ID:               "uyi",
		Code:             "UYI",
		NativeName:       "Uruguayan Peso en Unidades Indexadas",
		NativeNamePlural: "Uruguayan Pesos en Unidades Indexadas",
		Symbol:           "UYI",
	},
	CurrencyUYU: {
		ID:               "uyu",
		Code:             "UYU",
		NativeName:       "Uruguayan Peso",
		NativeNamePlural: "Uruguayan Pesos",
		Symbol:           "\\$U",
	},
	CurrencyUYW: {
		ID:               "uyw",
		Code:             "UYW",
		NativeName:       "Unidad Previsional",
		NativeNamePlural: "Unidad Previsionals",
		Symbol:           "UYW",
	},
	CurrencyUZS: {
		ID:               "uzs",
		Code:             "UZS",
		NativeName:       "Uzbekistan Som",
		NativeNamePlural: "Uzbekistan Soms",
		Symbol:           "лв",
	},
	CurrencyVED: {
		ID:               "ved",
		Code:             "VED",
		NativeName:       "Venezuelan Bolívar Soberano",
		NativeNamePlural: "Venezuelan Bolívar Soberanos",
		Symbol:           "VES",
	},
	CurrencyVES: {
		ID:               "ves",
		Code:             "VES",
		NativeName:       "Venezuelan Bolívar",
		NativeNamePlural: "Venezuelan Bolívars",
		Symbol:           "Bs.",
	},
	CurrencyVND: {
		ID:               "vnd",
		Code:             "VND",
		NativeName:       "Vietnamese Dong",
		NativeNamePlural: "Vietnamese Dong",
		Symbol:           "₫",
	},
	CurrencyVUV: {
		ID:               "vuv",
		Code:             "VUV",
		NativeName:       "Vanuatu Vatu",
		NativeNamePlural: "Vanuatu Vatus",
		Symbol:           "VT",
	},
	CurrencyWST: {
		ID:               "wst",
		Code:             "WST",
		NativeName:       "Samoan Tala",
		NativeNamePlural: "Samoan Tala",
		Symbol:           "WS\\$",
	},
	CurrencyXAF: {
		ID:               "xaf",
		Code:             "XAF",
		NativeName:       "Central African CFA Franc",
		NativeNamePlural: "Central African CFA Francs",
		Symbol:           "FCFA",
	},
	CurrencyXAG: {
		ID:               "xag",
		Code:             "XAG",
		NativeName:       "Silver (one troy ounce)",
		NativeNamePlural: "Silver (one troy ounce)",
		Symbol:           "XAG",
	},
	CurrencyXAU: {
		ID:               "xau",
		Code:             "XAU",
		NativeName:       "Gold (one troy ounce)",
		NativeNamePlural: "Gold (one troy ounce)",
		Symbol:           "XAU",
	},
	CurrencyXBA: {
		ID:               "xba",
		Code:             "XBA",
		NativeName:       "European Composite Unit",
		NativeNamePlural: "European Composite Units",
		Symbol:           "XBA",
	},
	CurrencyXBB: {
		ID:               "xbb",
		Code:             "XBB",
		NativeName:       "European Monetary Unit",
		NativeNamePlural: "European Monetary Units",
		Symbol:           "XBB",
	},
	CurrencyXBC: {
		ID:               "xbc",
		Code:             "XBC",
		NativeName:       "European Unit of Account 9",
		NativeNamePlural: "European Unit of Account 9",
		Symbol:           "XBC",
	},
	CurrencyXBD: {
		ID:               "xbd",
		Code:             "XBD",
		NativeName:       "European Unit of Account 17",
		NativeNamePlural: "European Unit of Account 17",
		Symbol:           "XBD",
	},
	CurrencyXCD: {
		ID:               "xcd",
		Code:             "XCD",
		NativeName:       "East Caribbean Dollar",
		NativeNamePlural: "East Caribbean Dollars",
		Symbol:           "\\$",
	},
	CurrencyXDR: {
		ID:               "xdr",
		Code:             "XDR",
		NativeName:       "Special Drawing Rights",
		NativeNamePlural: "Special Drawing Rights",
		Symbol:           "XDR",
	},
	CurrencyXOF: {
		ID:               "xof",
		Code:             "XOF",
		NativeName:       "West African CFA Franc",
		NativeNamePlural: "West African CFA Francs",
		Symbol:           "CFA",
	},
	CurrencyXPD: {
		ID:               "xpd",
		Code:             "XPD",
		NativeName:       "Palladium (one troy ounce)",
		NativeNamePlural: "Palladium (one troy ounce)",
		Symbol:           "XPD",
	},
	CurrencyXPF: {
		ID:               "xpf",
		Code:             "XPF",
		NativeName:       "CFP Franc",
		NativeNamePlural: "CFP Francs",
		Symbol:           "CFPF",
	},
	CurrencyXPT: {
		ID:               "xpt",
		Code:             "XPT",
		NativeName:       "Platinum (one troy ounce)",
		NativeNamePlural: "Platinum (one troy ounce)",
		Symbol:           "XPT",
	},
	CurrencyXSU: {
		ID:               "xsu",
		Code:             "XSU",
		NativeName:       "SUCRE",
		NativeNamePlural: "SUCRE",
		Symbol:           "XSU",
	},
	CurrencyXTS: {
		ID:               "xts",
		Code:             "XTS",
		NativeName:       "Codes specifically reserved for testing purposes",
		NativeNamePlural: "Codes specifically reserved for testing purposes",
		Symbol:           "XTS",
	},
	CurrencyXUA: {
		ID:               "xua",
		Code:             "XUA",
		NativeName:       "ADB Unit of Account",
		NativeNamePlural: "ADB Unit of Accounts",
		Symbol:           "XUA",
	},
	CurrencyXXX: {
		ID:               "xxx",
		Code:             "XXX",
		NativeName:       "No currency",
		NativeNamePlural: "No currency",
		Symbol:           "XXX",
	},
	CurrencyYER: {
		ID:               "yer",
		Code:             "YER",
		NativeName:       "Yemeni Rial",
		NativeNamePlural: "Yemeni Rials",
		Symbol:           "﷼",
	},
	CurrencyZAR: {
		ID:               "zar",
		Code:             "ZAR",
		NativeName:       "South African Rand",
		NativeNamePlural: "South African Rand",
		Symbol:           "R",
	},
	CurrencyZMW: {
		ID:               "zmw",
		Code:             "ZMW",
		NativeName:       "Zambian Kwacha",
		NativeNamePlural: "Zambian Kwachas",
		Symbol:           "ZK",
	},
	CurrencyZMG: {
		ID:               "zmg",
		Code:             "ZMG",
		NativeName:       "Zambian Kwacha (pre-2013)",
		NativeNamePlural: "Zambian Kwachas (pre-2013)",
		Symbol:           "ZMG",
	},
	CurrencyZWL: {
		ID:               "zwl",
		Code:             "ZWL",
		NativeName:       "Zimbabwean Dollar",
		NativeNamePlural: "Zimbabwean Dollars",
		Symbol:           "ZWL",
	},
}

// CurrencyByCode returns a currency by its code.
func CurrencyByCode(code string) (Currency, bool) {
	c, ok := currencies[CurrencyCode(strings.ToUpper(code))]

	return c, ok
}

// AllCurrencies returns all currencies.
func AllCurrencies() []Currency {
	result := make([]Currency, 0, len(currencies))
	for _, c := range currencies {
		result = append(result, c)
	}

	return result
}
