package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrentLabelApplicationMode string

const (
	CurrentLabelApplicationModeautomatic   CurrentLabelApplicationMode = "Automatic"
	CurrentLabelApplicationModemanual      CurrentLabelApplicationMode = "Manual"
	CurrentLabelApplicationModerecommended CurrentLabelApplicationMode = "Recommended"
)

func PossibleValuesForCurrentLabelApplicationMode() []string {
	return []string{
		string(CurrentLabelApplicationModeautomatic),
		string(CurrentLabelApplicationModemanual),
		string(CurrentLabelApplicationModerecommended),
	}
}

func parseCurrentLabelApplicationMode(input string) (*CurrentLabelApplicationMode, error) {
	vals := map[string]CurrentLabelApplicationMode{
		"automatic":   CurrentLabelApplicationModeautomatic,
		"manual":      CurrentLabelApplicationModemanual,
		"recommended": CurrentLabelApplicationModerecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CurrentLabelApplicationMode(input)
	return &out, nil
}
