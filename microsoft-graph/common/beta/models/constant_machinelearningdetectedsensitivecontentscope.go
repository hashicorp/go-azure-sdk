package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentScope string

const (
	MachineLearningDetectedSensitiveContentScopefullDocument    MachineLearningDetectedSensitiveContentScope = "FullDocument"
	MachineLearningDetectedSensitiveContentScopepartialDocument MachineLearningDetectedSensitiveContentScope = "PartialDocument"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentScope() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentScopefullDocument),
		string(MachineLearningDetectedSensitiveContentScopepartialDocument),
	}
}

func parseMachineLearningDetectedSensitiveContentScope(input string) (*MachineLearningDetectedSensitiveContentScope, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentScope{
		"fulldocument":    MachineLearningDetectedSensitiveContentScopefullDocument,
		"partialdocument": MachineLearningDetectedSensitiveContentScopepartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentScope(input)
	return &out, nil
}
