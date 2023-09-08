package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyLabelActionActionSource string

const (
	ApplyLabelActionActionSourceautomatic   ApplyLabelActionActionSource = "Automatic"
	ApplyLabelActionActionSourcedefault     ApplyLabelActionActionSource = "Default"
	ApplyLabelActionActionSourcemanual      ApplyLabelActionActionSource = "Manual"
	ApplyLabelActionActionSourcerecommended ApplyLabelActionActionSource = "Recommended"
)

func PossibleValuesForApplyLabelActionActionSource() []string {
	return []string{
		string(ApplyLabelActionActionSourceautomatic),
		string(ApplyLabelActionActionSourcedefault),
		string(ApplyLabelActionActionSourcemanual),
		string(ApplyLabelActionActionSourcerecommended),
	}
}

func parseApplyLabelActionActionSource(input string) (*ApplyLabelActionActionSource, error) {
	vals := map[string]ApplyLabelActionActionSource{
		"automatic":   ApplyLabelActionActionSourceautomatic,
		"default":     ApplyLabelActionActionSourcedefault,
		"manual":      ApplyLabelActionActionSourcemanual,
		"recommended": ApplyLabelActionActionSourcerecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplyLabelActionActionSource(input)
	return &out, nil
}
