package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveTypeScope string

const (
	SensitiveTypeScopefullDocument    SensitiveTypeScope = "FullDocument"
	SensitiveTypeScopepartialDocument SensitiveTypeScope = "PartialDocument"
)

func PossibleValuesForSensitiveTypeScope() []string {
	return []string{
		string(SensitiveTypeScopefullDocument),
		string(SensitiveTypeScopepartialDocument),
	}
}

func parseSensitiveTypeScope(input string) (*SensitiveTypeScope, error) {
	vals := map[string]SensitiveTypeScope{
		"fulldocument":    SensitiveTypeScopefullDocument,
		"partialdocument": SensitiveTypeScopepartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitiveTypeScope(input)
	return &out, nil
}
