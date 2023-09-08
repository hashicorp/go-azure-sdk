package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10TeamGeneralConfigurationMiracastChannel string

const (
	Windows10TeamGeneralConfigurationMiracastChanneleight                Windows10TeamGeneralConfigurationMiracastChannel = "Eight"
	Windows10TeamGeneralConfigurationMiracastChanneleleven               Windows10TeamGeneralConfigurationMiracastChannel = "Eleven"
	Windows10TeamGeneralConfigurationMiracastChannelfive                 Windows10TeamGeneralConfigurationMiracastChannel = "Five"
	Windows10TeamGeneralConfigurationMiracastChannelforty                Windows10TeamGeneralConfigurationMiracastChannel = "Forty"
	Windows10TeamGeneralConfigurationMiracastChannelfortyEight           Windows10TeamGeneralConfigurationMiracastChannel = "FortyEight"
	Windows10TeamGeneralConfigurationMiracastChannelfortyFour            Windows10TeamGeneralConfigurationMiracastChannel = "FortyFour"
	Windows10TeamGeneralConfigurationMiracastChannelfour                 Windows10TeamGeneralConfigurationMiracastChannel = "Four"
	Windows10TeamGeneralConfigurationMiracastChannelnine                 Windows10TeamGeneralConfigurationMiracastChannel = "Nine"
	Windows10TeamGeneralConfigurationMiracastChannelone                  Windows10TeamGeneralConfigurationMiracastChannel = "One"
	Windows10TeamGeneralConfigurationMiracastChanneloneHundredFiftySeven Windows10TeamGeneralConfigurationMiracastChannel = "OneHundredFiftySeven"
	Windows10TeamGeneralConfigurationMiracastChanneloneHundredFiftyThree Windows10TeamGeneralConfigurationMiracastChannel = "OneHundredFiftyThree"
	Windows10TeamGeneralConfigurationMiracastChanneloneHundredFortyNine  Windows10TeamGeneralConfigurationMiracastChannel = "OneHundredFortyNine"
	Windows10TeamGeneralConfigurationMiracastChanneloneHundredSixtyFive  Windows10TeamGeneralConfigurationMiracastChannel = "OneHundredSixtyFive"
	Windows10TeamGeneralConfigurationMiracastChanneloneHundredSixtyOne   Windows10TeamGeneralConfigurationMiracastChannel = "OneHundredSixtyOne"
	Windows10TeamGeneralConfigurationMiracastChannelseven                Windows10TeamGeneralConfigurationMiracastChannel = "Seven"
	Windows10TeamGeneralConfigurationMiracastChannelsix                  Windows10TeamGeneralConfigurationMiracastChannel = "Six"
	Windows10TeamGeneralConfigurationMiracastChannelten                  Windows10TeamGeneralConfigurationMiracastChannel = "Ten"
	Windows10TeamGeneralConfigurationMiracastChannelthirtySix            Windows10TeamGeneralConfigurationMiracastChannel = "ThirtySix"
	Windows10TeamGeneralConfigurationMiracastChannelthree                Windows10TeamGeneralConfigurationMiracastChannel = "Three"
	Windows10TeamGeneralConfigurationMiracastChanneltwo                  Windows10TeamGeneralConfigurationMiracastChannel = "Two"
	Windows10TeamGeneralConfigurationMiracastChanneluserDefined          Windows10TeamGeneralConfigurationMiracastChannel = "UserDefined"
)

func PossibleValuesForWindows10TeamGeneralConfigurationMiracastChannel() []string {
	return []string{
		string(Windows10TeamGeneralConfigurationMiracastChanneleight),
		string(Windows10TeamGeneralConfigurationMiracastChanneleleven),
		string(Windows10TeamGeneralConfigurationMiracastChannelfive),
		string(Windows10TeamGeneralConfigurationMiracastChannelforty),
		string(Windows10TeamGeneralConfigurationMiracastChannelfortyEight),
		string(Windows10TeamGeneralConfigurationMiracastChannelfortyFour),
		string(Windows10TeamGeneralConfigurationMiracastChannelfour),
		string(Windows10TeamGeneralConfigurationMiracastChannelnine),
		string(Windows10TeamGeneralConfigurationMiracastChannelone),
		string(Windows10TeamGeneralConfigurationMiracastChanneloneHundredFiftySeven),
		string(Windows10TeamGeneralConfigurationMiracastChanneloneHundredFiftyThree),
		string(Windows10TeamGeneralConfigurationMiracastChanneloneHundredFortyNine),
		string(Windows10TeamGeneralConfigurationMiracastChanneloneHundredSixtyFive),
		string(Windows10TeamGeneralConfigurationMiracastChanneloneHundredSixtyOne),
		string(Windows10TeamGeneralConfigurationMiracastChannelseven),
		string(Windows10TeamGeneralConfigurationMiracastChannelsix),
		string(Windows10TeamGeneralConfigurationMiracastChannelten),
		string(Windows10TeamGeneralConfigurationMiracastChannelthirtySix),
		string(Windows10TeamGeneralConfigurationMiracastChannelthree),
		string(Windows10TeamGeneralConfigurationMiracastChanneltwo),
		string(Windows10TeamGeneralConfigurationMiracastChanneluserDefined),
	}
}

func parseWindows10TeamGeneralConfigurationMiracastChannel(input string) (*Windows10TeamGeneralConfigurationMiracastChannel, error) {
	vals := map[string]Windows10TeamGeneralConfigurationMiracastChannel{
		"eight":                Windows10TeamGeneralConfigurationMiracastChanneleight,
		"eleven":               Windows10TeamGeneralConfigurationMiracastChanneleleven,
		"five":                 Windows10TeamGeneralConfigurationMiracastChannelfive,
		"forty":                Windows10TeamGeneralConfigurationMiracastChannelforty,
		"fortyeight":           Windows10TeamGeneralConfigurationMiracastChannelfortyEight,
		"fortyfour":            Windows10TeamGeneralConfigurationMiracastChannelfortyFour,
		"four":                 Windows10TeamGeneralConfigurationMiracastChannelfour,
		"nine":                 Windows10TeamGeneralConfigurationMiracastChannelnine,
		"one":                  Windows10TeamGeneralConfigurationMiracastChannelone,
		"onehundredfiftyseven": Windows10TeamGeneralConfigurationMiracastChanneloneHundredFiftySeven,
		"onehundredfiftythree": Windows10TeamGeneralConfigurationMiracastChanneloneHundredFiftyThree,
		"onehundredfortynine":  Windows10TeamGeneralConfigurationMiracastChanneloneHundredFortyNine,
		"onehundredsixtyfive":  Windows10TeamGeneralConfigurationMiracastChanneloneHundredSixtyFive,
		"onehundredsixtyone":   Windows10TeamGeneralConfigurationMiracastChanneloneHundredSixtyOne,
		"seven":                Windows10TeamGeneralConfigurationMiracastChannelseven,
		"six":                  Windows10TeamGeneralConfigurationMiracastChannelsix,
		"ten":                  Windows10TeamGeneralConfigurationMiracastChannelten,
		"thirtysix":            Windows10TeamGeneralConfigurationMiracastChannelthirtySix,
		"three":                Windows10TeamGeneralConfigurationMiracastChannelthree,
		"two":                  Windows10TeamGeneralConfigurationMiracastChanneltwo,
		"userdefined":          Windows10TeamGeneralConfigurationMiracastChanneluserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10TeamGeneralConfigurationMiracastChannel(input)
	return &out, nil
}
