package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10TeamGeneralConfigurationMiracastChannel string

const (
	Windows10TeamGeneralConfigurationMiracastChannel_Eight                Windows10TeamGeneralConfigurationMiracastChannel = "eight"
	Windows10TeamGeneralConfigurationMiracastChannel_Eleven               Windows10TeamGeneralConfigurationMiracastChannel = "eleven"
	Windows10TeamGeneralConfigurationMiracastChannel_Five                 Windows10TeamGeneralConfigurationMiracastChannel = "five"
	Windows10TeamGeneralConfigurationMiracastChannel_Forty                Windows10TeamGeneralConfigurationMiracastChannel = "forty"
	Windows10TeamGeneralConfigurationMiracastChannel_FortyEight           Windows10TeamGeneralConfigurationMiracastChannel = "fortyEight"
	Windows10TeamGeneralConfigurationMiracastChannel_FortyFour            Windows10TeamGeneralConfigurationMiracastChannel = "fortyFour"
	Windows10TeamGeneralConfigurationMiracastChannel_Four                 Windows10TeamGeneralConfigurationMiracastChannel = "four"
	Windows10TeamGeneralConfigurationMiracastChannel_Nine                 Windows10TeamGeneralConfigurationMiracastChannel = "nine"
	Windows10TeamGeneralConfigurationMiracastChannel_One                  Windows10TeamGeneralConfigurationMiracastChannel = "one"
	Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFiftySeven Windows10TeamGeneralConfigurationMiracastChannel = "oneHundredFiftySeven"
	Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFiftyThree Windows10TeamGeneralConfigurationMiracastChannel = "oneHundredFiftyThree"
	Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFortyNine  Windows10TeamGeneralConfigurationMiracastChannel = "oneHundredFortyNine"
	Windows10TeamGeneralConfigurationMiracastChannel_OneHundredSixtyFive  Windows10TeamGeneralConfigurationMiracastChannel = "oneHundredSixtyFive"
	Windows10TeamGeneralConfigurationMiracastChannel_OneHundredSixtyOne   Windows10TeamGeneralConfigurationMiracastChannel = "oneHundredSixtyOne"
	Windows10TeamGeneralConfigurationMiracastChannel_Seven                Windows10TeamGeneralConfigurationMiracastChannel = "seven"
	Windows10TeamGeneralConfigurationMiracastChannel_Six                  Windows10TeamGeneralConfigurationMiracastChannel = "six"
	Windows10TeamGeneralConfigurationMiracastChannel_Ten                  Windows10TeamGeneralConfigurationMiracastChannel = "ten"
	Windows10TeamGeneralConfigurationMiracastChannel_ThirtySix            Windows10TeamGeneralConfigurationMiracastChannel = "thirtySix"
	Windows10TeamGeneralConfigurationMiracastChannel_Three                Windows10TeamGeneralConfigurationMiracastChannel = "three"
	Windows10TeamGeneralConfigurationMiracastChannel_Two                  Windows10TeamGeneralConfigurationMiracastChannel = "two"
	Windows10TeamGeneralConfigurationMiracastChannel_UserDefined          Windows10TeamGeneralConfigurationMiracastChannel = "userDefined"
)

func PossibleValuesForWindows10TeamGeneralConfigurationMiracastChannel() []string {
	return []string{
		string(Windows10TeamGeneralConfigurationMiracastChannel_Eight),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Eleven),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Five),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Forty),
		string(Windows10TeamGeneralConfigurationMiracastChannel_FortyEight),
		string(Windows10TeamGeneralConfigurationMiracastChannel_FortyFour),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Four),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Nine),
		string(Windows10TeamGeneralConfigurationMiracastChannel_One),
		string(Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFiftySeven),
		string(Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFiftyThree),
		string(Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFortyNine),
		string(Windows10TeamGeneralConfigurationMiracastChannel_OneHundredSixtyFive),
		string(Windows10TeamGeneralConfigurationMiracastChannel_OneHundredSixtyOne),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Seven),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Six),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Ten),
		string(Windows10TeamGeneralConfigurationMiracastChannel_ThirtySix),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Three),
		string(Windows10TeamGeneralConfigurationMiracastChannel_Two),
		string(Windows10TeamGeneralConfigurationMiracastChannel_UserDefined),
	}
}

func (s *Windows10TeamGeneralConfigurationMiracastChannel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10TeamGeneralConfigurationMiracastChannel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10TeamGeneralConfigurationMiracastChannel(input string) (*Windows10TeamGeneralConfigurationMiracastChannel, error) {
	vals := map[string]Windows10TeamGeneralConfigurationMiracastChannel{
		"eight":                Windows10TeamGeneralConfigurationMiracastChannel_Eight,
		"eleven":               Windows10TeamGeneralConfigurationMiracastChannel_Eleven,
		"five":                 Windows10TeamGeneralConfigurationMiracastChannel_Five,
		"forty":                Windows10TeamGeneralConfigurationMiracastChannel_Forty,
		"fortyeight":           Windows10TeamGeneralConfigurationMiracastChannel_FortyEight,
		"fortyfour":            Windows10TeamGeneralConfigurationMiracastChannel_FortyFour,
		"four":                 Windows10TeamGeneralConfigurationMiracastChannel_Four,
		"nine":                 Windows10TeamGeneralConfigurationMiracastChannel_Nine,
		"one":                  Windows10TeamGeneralConfigurationMiracastChannel_One,
		"onehundredfiftyseven": Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFiftySeven,
		"onehundredfiftythree": Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFiftyThree,
		"onehundredfortynine":  Windows10TeamGeneralConfigurationMiracastChannel_OneHundredFortyNine,
		"onehundredsixtyfive":  Windows10TeamGeneralConfigurationMiracastChannel_OneHundredSixtyFive,
		"onehundredsixtyone":   Windows10TeamGeneralConfigurationMiracastChannel_OneHundredSixtyOne,
		"seven":                Windows10TeamGeneralConfigurationMiracastChannel_Seven,
		"six":                  Windows10TeamGeneralConfigurationMiracastChannel_Six,
		"ten":                  Windows10TeamGeneralConfigurationMiracastChannel_Ten,
		"thirtysix":            Windows10TeamGeneralConfigurationMiracastChannel_ThirtySix,
		"three":                Windows10TeamGeneralConfigurationMiracastChannel_Three,
		"two":                  Windows10TeamGeneralConfigurationMiracastChannel_Two,
		"userdefined":          Windows10TeamGeneralConfigurationMiracastChannel_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10TeamGeneralConfigurationMiracastChannel(input)
	return &out, nil
}
