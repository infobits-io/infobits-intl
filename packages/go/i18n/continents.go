// GENERATED CODE - DO NOT MODIFY BY HAND

package i18n

// ContinentsTranslations contains translations keyed by locale then code.
var ContinentsTranslations = map[string]map[string]string{
	"da": {
		"AF": "Afrika",
		"AQ": "Antarktis",
		"AS": "Asien",
		"EU": "Europa",
		"NA": "Nordamerika",
		"OS": "Oceanien",
		"SA": "Sydamerika",
	},
	"de": {
		"AF": "Afrika",
		"AQ": "Antarktika",
		"AS": "Asien",
		"EU": "Europa",
		"NA": "Nordamerika",
		"OS": "Ozeanien",
		"SA": "Südamerika",
	},
	"en": {
		"AF": "Africa",
		"AQ": "Antarctica",
		"AS": "Asia",
		"EU": "Europe",
		"NA": "North America",
		"OS": "Oceania",
		"SA": "South America",
	},
	"es": {
		"AF": "África",
		"AQ": "Antártida",
		"AS": "Asia",
		"EU": "Europa",
		"NA": "América del Norte",
		"OS": "Oceanía",
		"SA": "América del Sur",
	},
	"fr": {
		"AF": "Afrique",
		"AQ": "Antarctique",
		"AS": "Asie",
		"EU": "Europe",
		"NA": "Amérique du Nord",
		"OS": "Océanie",
		"SA": "Amérique du Sud",
	},
	"it": {
		"AF": "Africa",
		"AQ": "Antartide",
		"AS": "Asia",
		"EU": "Europa",
		"NA": "Nord America",
		"OS": "Oceania",
		"SA": "Sud America",
	},
	"zh": {
		"AF": "非洲",
		"AQ": "南极洲",
		"AS": "亚洲",
		"EU": "欧洲",
		"NA": "北美洲",
		"OS": "大洋洲",
		"SA": "南美洲",
	},
}

// GetContinentsName returns the translated name for a code in a locale.
func GetContinentsName(code, locale string) (string, bool) {
	if localeMap, ok := ContinentsTranslations[locale]; ok {
		if name, ok := localeMap[code]; ok {
			return name, true
		}
	}

	return "", false
}
