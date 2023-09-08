package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveTypeClassificationMethod string

const (
	SensitiveTypeClassificationMethodexactDataMatch  SensitiveTypeClassificationMethod = "ExactDataMatch"
	SensitiveTypeClassificationMethodfingerprint     SensitiveTypeClassificationMethod = "Fingerprint"
	SensitiveTypeClassificationMethodmachineLearning SensitiveTypeClassificationMethod = "MachineLearning"
	SensitiveTypeClassificationMethodpatternMatch    SensitiveTypeClassificationMethod = "PatternMatch"
)

func PossibleValuesForSensitiveTypeClassificationMethod() []string {
	return []string{
		string(SensitiveTypeClassificationMethodexactDataMatch),
		string(SensitiveTypeClassificationMethodfingerprint),
		string(SensitiveTypeClassificationMethodmachineLearning),
		string(SensitiveTypeClassificationMethodpatternMatch),
	}
}

func parseSensitiveTypeClassificationMethod(input string) (*SensitiveTypeClassificationMethod, error) {
	vals := map[string]SensitiveTypeClassificationMethod{
		"exactdatamatch":  SensitiveTypeClassificationMethodexactDataMatch,
		"fingerprint":     SensitiveTypeClassificationMethodfingerprint,
		"machinelearning": SensitiveTypeClassificationMethodmachineLearning,
		"patternmatch":    SensitiveTypeClassificationMethodpatternMatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitiveTypeClassificationMethod(input)
	return &out, nil
}
