package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityApplyLabelActionActionSource string

const (
	SecurityApplyLabelActionActionSourceautomatic   SecurityApplyLabelActionActionSource = "Automatic"
	SecurityApplyLabelActionActionSourcedefault     SecurityApplyLabelActionActionSource = "Default"
	SecurityApplyLabelActionActionSourcemanual      SecurityApplyLabelActionActionSource = "Manual"
	SecurityApplyLabelActionActionSourcerecommended SecurityApplyLabelActionActionSource = "Recommended"
)

func PossibleValuesForSecurityApplyLabelActionActionSource() []string {
	return []string{
		string(SecurityApplyLabelActionActionSourceautomatic),
		string(SecurityApplyLabelActionActionSourcedefault),
		string(SecurityApplyLabelActionActionSourcemanual),
		string(SecurityApplyLabelActionActionSourcerecommended),
	}
}

func parseSecurityApplyLabelActionActionSource(input string) (*SecurityApplyLabelActionActionSource, error) {
	vals := map[string]SecurityApplyLabelActionActionSource{
		"automatic":   SecurityApplyLabelActionActionSourceautomatic,
		"default":     SecurityApplyLabelActionActionSourcedefault,
		"manual":      SecurityApplyLabelActionActionSourcemanual,
		"recommended": SecurityApplyLabelActionActionSourcerecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityApplyLabelActionActionSource(input)
	return &out, nil
}
