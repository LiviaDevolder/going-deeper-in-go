package address

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var caser = cases.Title(language.BrazilianPortuguese)

// AddressType verify if the address has a valid type and returns it
func AddressType(address string) string {
	validTypes := []string{"street", "avenue", "road", "route", "highway"}

	firstAddressWord := strings.ToLower(strings.Split(address, " ")[0])

	addressHasValidType := false

	for _, dataType := range validTypes {
		if dataType == firstAddressWord {
			addressHasValidType = true
		}
	}

	if addressHasValidType {
		return caser.String(firstAddressWord)
	}

	return "Invalid Type"
}
