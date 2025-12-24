// GENERATED CODE - DO NOT MODIFY BY HAND

package intl

import "strings"

// ContinentCode represents a continent code.
type ContinentCode string

// Continent codes.
const (
	ContinentAF ContinentCode = "AF"
	ContinentAQ ContinentCode = "AQ"
	ContinentAS ContinentCode = "AS"
	ContinentEU ContinentCode = "EU"
	ContinentNA ContinentCode = "NA"
	ContinentOS ContinentCode = "OS"
	ContinentSA ContinentCode = "SA"
)

// Continent represents a continent.
type Continent struct {
	ID   string
	Code string
	Name string
}

// String returns the continent code.
func (c ContinentCode) String() string {
	return string(c)
}

// Continent returns the continent data for this code.
func (c ContinentCode) Continent() Continent {
	return continents[c]
}

var continents = map[ContinentCode]Continent{
	ContinentAF: {
		ID:   "africa",
		Code: "AF",
		Name: "Africa",
	},
	ContinentAQ: {
		ID:   "antarctica",
		Code: "AQ",
		Name: "Antarctica",
	},
	ContinentAS: {
		ID:   "asia",
		Code: "AS",
		Name: "Asia",
	},
	ContinentEU: {
		ID:   "europe",
		Code: "EU",
		Name: "Europe",
	},
	ContinentNA: {
		ID:   "northAmerica",
		Code: "NA",
		Name: "North America",
	},
	ContinentOS: {
		ID:   "oceania",
		Code: "OS",
		Name: "Oceania",
	},
	ContinentSA: {
		ID:   "southAmerica",
		Code: "SA",
		Name: "South America",
	},
}

// ContinentByCode returns a continent by its code.
func ContinentByCode(code string) (Continent, bool) {
	c, ok := continents[ContinentCode(strings.ToUpper(code))]

	return c, ok
}

// AllContinents returns all continents.
func AllContinents() []Continent {
	result := make([]Continent, 0, len(continents))
	for _, c := range continents {
		result = append(result, c)
	}

	return result
}
