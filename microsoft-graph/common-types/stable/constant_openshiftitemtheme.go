package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftItemTheme string

const (
	OpenShiftItemTheme_Blue       OpenShiftItemTheme = "blue"
	OpenShiftItemTheme_DarkBlue   OpenShiftItemTheme = "darkBlue"
	OpenShiftItemTheme_DarkGreen  OpenShiftItemTheme = "darkGreen"
	OpenShiftItemTheme_DarkPink   OpenShiftItemTheme = "darkPink"
	OpenShiftItemTheme_DarkPurple OpenShiftItemTheme = "darkPurple"
	OpenShiftItemTheme_DarkYellow OpenShiftItemTheme = "darkYellow"
	OpenShiftItemTheme_Gray       OpenShiftItemTheme = "gray"
	OpenShiftItemTheme_Green      OpenShiftItemTheme = "green"
	OpenShiftItemTheme_Pink       OpenShiftItemTheme = "pink"
	OpenShiftItemTheme_Purple     OpenShiftItemTheme = "purple"
	OpenShiftItemTheme_White      OpenShiftItemTheme = "white"
	OpenShiftItemTheme_Yellow     OpenShiftItemTheme = "yellow"
)

func PossibleValuesForOpenShiftItemTheme() []string {
	return []string{
		string(OpenShiftItemTheme_Blue),
		string(OpenShiftItemTheme_DarkBlue),
		string(OpenShiftItemTheme_DarkGreen),
		string(OpenShiftItemTheme_DarkPink),
		string(OpenShiftItemTheme_DarkPurple),
		string(OpenShiftItemTheme_DarkYellow),
		string(OpenShiftItemTheme_Gray),
		string(OpenShiftItemTheme_Green),
		string(OpenShiftItemTheme_Pink),
		string(OpenShiftItemTheme_Purple),
		string(OpenShiftItemTheme_White),
		string(OpenShiftItemTheme_Yellow),
	}
}

func (s *OpenShiftItemTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOpenShiftItemTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOpenShiftItemTheme(input string) (*OpenShiftItemTheme, error) {
	vals := map[string]OpenShiftItemTheme{
		"blue":       OpenShiftItemTheme_Blue,
		"darkblue":   OpenShiftItemTheme_DarkBlue,
		"darkgreen":  OpenShiftItemTheme_DarkGreen,
		"darkpink":   OpenShiftItemTheme_DarkPink,
		"darkpurple": OpenShiftItemTheme_DarkPurple,
		"darkyellow": OpenShiftItemTheme_DarkYellow,
		"gray":       OpenShiftItemTheme_Gray,
		"green":      OpenShiftItemTheme_Green,
		"pink":       OpenShiftItemTheme_Pink,
		"purple":     OpenShiftItemTheme_Purple,
		"white":      OpenShiftItemTheme_White,
		"yellow":     OpenShiftItemTheme_Yellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenShiftItemTheme(input)
	return &out, nil
}
