package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftItemTheme string

const (
	ShiftItemTheme_Blue       ShiftItemTheme = "blue"
	ShiftItemTheme_DarkBlue   ShiftItemTheme = "darkBlue"
	ShiftItemTheme_DarkGreen  ShiftItemTheme = "darkGreen"
	ShiftItemTheme_DarkPink   ShiftItemTheme = "darkPink"
	ShiftItemTheme_DarkPurple ShiftItemTheme = "darkPurple"
	ShiftItemTheme_DarkYellow ShiftItemTheme = "darkYellow"
	ShiftItemTheme_Gray       ShiftItemTheme = "gray"
	ShiftItemTheme_Green      ShiftItemTheme = "green"
	ShiftItemTheme_Pink       ShiftItemTheme = "pink"
	ShiftItemTheme_Purple     ShiftItemTheme = "purple"
	ShiftItemTheme_White      ShiftItemTheme = "white"
	ShiftItemTheme_Yellow     ShiftItemTheme = "yellow"
)

func PossibleValuesForShiftItemTheme() []string {
	return []string{
		string(ShiftItemTheme_Blue),
		string(ShiftItemTheme_DarkBlue),
		string(ShiftItemTheme_DarkGreen),
		string(ShiftItemTheme_DarkPink),
		string(ShiftItemTheme_DarkPurple),
		string(ShiftItemTheme_DarkYellow),
		string(ShiftItemTheme_Gray),
		string(ShiftItemTheme_Green),
		string(ShiftItemTheme_Pink),
		string(ShiftItemTheme_Purple),
		string(ShiftItemTheme_White),
		string(ShiftItemTheme_Yellow),
	}
}

func (s *ShiftItemTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseShiftItemTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseShiftItemTheme(input string) (*ShiftItemTheme, error) {
	vals := map[string]ShiftItemTheme{
		"blue":       ShiftItemTheme_Blue,
		"darkblue":   ShiftItemTheme_DarkBlue,
		"darkgreen":  ShiftItemTheme_DarkGreen,
		"darkpink":   ShiftItemTheme_DarkPink,
		"darkpurple": ShiftItemTheme_DarkPurple,
		"darkyellow": ShiftItemTheme_DarkYellow,
		"gray":       ShiftItemTheme_Gray,
		"green":      ShiftItemTheme_Green,
		"pink":       ShiftItemTheme_Pink,
		"purple":     ShiftItemTheme_Purple,
		"white":      ShiftItemTheme_White,
		"yellow":     ShiftItemTheme_Yellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShiftItemTheme(input)
	return &out, nil
}
