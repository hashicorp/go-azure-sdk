package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentInfoState string

const (
	SecurityContentInfoStatemotion SecurityContentInfoState = "Motion"
	SecurityContentInfoStaterest   SecurityContentInfoState = "Rest"
	SecurityContentInfoStateuse    SecurityContentInfoState = "Use"
)

func PossibleValuesForSecurityContentInfoState() []string {
	return []string{
		string(SecurityContentInfoStatemotion),
		string(SecurityContentInfoStaterest),
		string(SecurityContentInfoStateuse),
	}
}

func parseSecurityContentInfoState(input string) (*SecurityContentInfoState, error) {
	vals := map[string]SecurityContentInfoState{
		"motion": SecurityContentInfoStatemotion,
		"rest":   SecurityContentInfoStaterest,
		"use":    SecurityContentInfoStateuse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContentInfoState(input)
	return &out, nil
}
