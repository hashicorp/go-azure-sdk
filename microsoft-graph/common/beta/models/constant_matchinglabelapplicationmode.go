package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingLabelApplicationMode string

const (
	MatchingLabelApplicationModeautomatic   MatchingLabelApplicationMode = "Automatic"
	MatchingLabelApplicationModemanual      MatchingLabelApplicationMode = "Manual"
	MatchingLabelApplicationModerecommended MatchingLabelApplicationMode = "Recommended"
)

func PossibleValuesForMatchingLabelApplicationMode() []string {
	return []string{
		string(MatchingLabelApplicationModeautomatic),
		string(MatchingLabelApplicationModemanual),
		string(MatchingLabelApplicationModerecommended),
	}
}

func parseMatchingLabelApplicationMode(input string) (*MatchingLabelApplicationMode, error) {
	vals := map[string]MatchingLabelApplicationMode{
		"automatic":   MatchingLabelApplicationModeautomatic,
		"manual":      MatchingLabelApplicationModemanual,
		"recommended": MatchingLabelApplicationModerecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MatchingLabelApplicationMode(input)
	return &out, nil
}
