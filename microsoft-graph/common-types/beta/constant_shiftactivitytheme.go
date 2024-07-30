package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftActivityTheme string

const (
	ShiftActivityTheme_Blue       ShiftActivityTheme = "blue"
	ShiftActivityTheme_DarkBlue   ShiftActivityTheme = "darkBlue"
	ShiftActivityTheme_DarkGreen  ShiftActivityTheme = "darkGreen"
	ShiftActivityTheme_DarkPink   ShiftActivityTheme = "darkPink"
	ShiftActivityTheme_DarkPurple ShiftActivityTheme = "darkPurple"
	ShiftActivityTheme_DarkYellow ShiftActivityTheme = "darkYellow"
	ShiftActivityTheme_Gray       ShiftActivityTheme = "gray"
	ShiftActivityTheme_Green      ShiftActivityTheme = "green"
	ShiftActivityTheme_Pink       ShiftActivityTheme = "pink"
	ShiftActivityTheme_Purple     ShiftActivityTheme = "purple"
	ShiftActivityTheme_White      ShiftActivityTheme = "white"
	ShiftActivityTheme_Yellow     ShiftActivityTheme = "yellow"
)

func PossibleValuesForShiftActivityTheme() []string {
	return []string{
		string(ShiftActivityTheme_Blue),
		string(ShiftActivityTheme_DarkBlue),
		string(ShiftActivityTheme_DarkGreen),
		string(ShiftActivityTheme_DarkPink),
		string(ShiftActivityTheme_DarkPurple),
		string(ShiftActivityTheme_DarkYellow),
		string(ShiftActivityTheme_Gray),
		string(ShiftActivityTheme_Green),
		string(ShiftActivityTheme_Pink),
		string(ShiftActivityTheme_Purple),
		string(ShiftActivityTheme_White),
		string(ShiftActivityTheme_Yellow),
	}
}

func (s *ShiftActivityTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShiftActivityTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShiftActivityTheme(input string) (*ShiftActivityTheme, error) {
	vals := map[string]ShiftActivityTheme{
		"blue":       ShiftActivityTheme_Blue,
		"darkblue":   ShiftActivityTheme_DarkBlue,
		"darkgreen":  ShiftActivityTheme_DarkGreen,
		"darkpink":   ShiftActivityTheme_DarkPink,
		"darkpurple": ShiftActivityTheme_DarkPurple,
		"darkyellow": ShiftActivityTheme_DarkYellow,
		"gray":       ShiftActivityTheme_Gray,
		"green":      ShiftActivityTheme_Green,
		"pink":       ShiftActivityTheme_Pink,
		"purple":     ShiftActivityTheme_Purple,
		"white":      ShiftActivityTheme_White,
		"yellow":     ShiftActivityTheme_Yellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShiftActivityTheme(input)
	return &out, nil
}
