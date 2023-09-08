package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiIndicatorTlpLevel string

const (
	TiIndicatorTlpLevelamber   TiIndicatorTlpLevel = "Amber"
	TiIndicatorTlpLevelgreen   TiIndicatorTlpLevel = "Green"
	TiIndicatorTlpLevelred     TiIndicatorTlpLevel = "Red"
	TiIndicatorTlpLevelunknown TiIndicatorTlpLevel = "Unknown"
	TiIndicatorTlpLevelwhite   TiIndicatorTlpLevel = "White"
)

func PossibleValuesForTiIndicatorTlpLevel() []string {
	return []string{
		string(TiIndicatorTlpLevelamber),
		string(TiIndicatorTlpLevelgreen),
		string(TiIndicatorTlpLevelred),
		string(TiIndicatorTlpLevelunknown),
		string(TiIndicatorTlpLevelwhite),
	}
}

func parseTiIndicatorTlpLevel(input string) (*TiIndicatorTlpLevel, error) {
	vals := map[string]TiIndicatorTlpLevel{
		"amber":   TiIndicatorTlpLevelamber,
		"green":   TiIndicatorTlpLevelgreen,
		"red":     TiIndicatorTlpLevelred,
		"unknown": TiIndicatorTlpLevelunknown,
		"white":   TiIndicatorTlpLevelwhite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiIndicatorTlpLevel(input)
	return &out, nil
}
