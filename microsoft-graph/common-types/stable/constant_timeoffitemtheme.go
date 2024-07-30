package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffItemTheme string

const (
	TimeOffItemTheme_Blue       TimeOffItemTheme = "blue"
	TimeOffItemTheme_DarkBlue   TimeOffItemTheme = "darkBlue"
	TimeOffItemTheme_DarkGreen  TimeOffItemTheme = "darkGreen"
	TimeOffItemTheme_DarkPink   TimeOffItemTheme = "darkPink"
	TimeOffItemTheme_DarkPurple TimeOffItemTheme = "darkPurple"
	TimeOffItemTheme_DarkYellow TimeOffItemTheme = "darkYellow"
	TimeOffItemTheme_Gray       TimeOffItemTheme = "gray"
	TimeOffItemTheme_Green      TimeOffItemTheme = "green"
	TimeOffItemTheme_Pink       TimeOffItemTheme = "pink"
	TimeOffItemTheme_Purple     TimeOffItemTheme = "purple"
	TimeOffItemTheme_White      TimeOffItemTheme = "white"
	TimeOffItemTheme_Yellow     TimeOffItemTheme = "yellow"
)

func PossibleValuesForTimeOffItemTheme() []string {
	return []string{
		string(TimeOffItemTheme_Blue),
		string(TimeOffItemTheme_DarkBlue),
		string(TimeOffItemTheme_DarkGreen),
		string(TimeOffItemTheme_DarkPink),
		string(TimeOffItemTheme_DarkPurple),
		string(TimeOffItemTheme_DarkYellow),
		string(TimeOffItemTheme_Gray),
		string(TimeOffItemTheme_Green),
		string(TimeOffItemTheme_Pink),
		string(TimeOffItemTheme_Purple),
		string(TimeOffItemTheme_White),
		string(TimeOffItemTheme_Yellow),
	}
}

func (s *TimeOffItemTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffItemTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffItemTheme(input string) (*TimeOffItemTheme, error) {
	vals := map[string]TimeOffItemTheme{
		"blue":       TimeOffItemTheme_Blue,
		"darkblue":   TimeOffItemTheme_DarkBlue,
		"darkgreen":  TimeOffItemTheme_DarkGreen,
		"darkpink":   TimeOffItemTheme_DarkPink,
		"darkpurple": TimeOffItemTheme_DarkPurple,
		"darkyellow": TimeOffItemTheme_DarkYellow,
		"gray":       TimeOffItemTheme_Gray,
		"green":      TimeOffItemTheme_Green,
		"pink":       TimeOffItemTheme_Pink,
		"purple":     TimeOffItemTheme_Purple,
		"white":      TimeOffItemTheme_White,
		"yellow":     TimeOffItemTheme_Yellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffItemTheme(input)
	return &out, nil
}
