package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentScope string

const (
	DetectedSensitiveContentScopefullDocument    DetectedSensitiveContentScope = "FullDocument"
	DetectedSensitiveContentScopepartialDocument DetectedSensitiveContentScope = "PartialDocument"
)

func PossibleValuesForDetectedSensitiveContentScope() []string {
	return []string{
		string(DetectedSensitiveContentScopefullDocument),
		string(DetectedSensitiveContentScopepartialDocument),
	}
}

func parseDetectedSensitiveContentScope(input string) (*DetectedSensitiveContentScope, error) {
	vals := map[string]DetectedSensitiveContentScope{
		"fulldocument":    DetectedSensitiveContentScopefullDocument,
		"partialdocument": DetectedSensitiveContentScopepartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedSensitiveContentScope(input)
	return &out, nil
}
