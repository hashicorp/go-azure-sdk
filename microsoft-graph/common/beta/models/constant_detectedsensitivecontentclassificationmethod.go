package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentClassificationMethod string

const (
	DetectedSensitiveContentClassificationMethodexactDataMatch  DetectedSensitiveContentClassificationMethod = "ExactDataMatch"
	DetectedSensitiveContentClassificationMethodfingerprint     DetectedSensitiveContentClassificationMethod = "Fingerprint"
	DetectedSensitiveContentClassificationMethodmachineLearning DetectedSensitiveContentClassificationMethod = "MachineLearning"
	DetectedSensitiveContentClassificationMethodpatternMatch    DetectedSensitiveContentClassificationMethod = "PatternMatch"
)

func PossibleValuesForDetectedSensitiveContentClassificationMethod() []string {
	return []string{
		string(DetectedSensitiveContentClassificationMethodexactDataMatch),
		string(DetectedSensitiveContentClassificationMethodfingerprint),
		string(DetectedSensitiveContentClassificationMethodmachineLearning),
		string(DetectedSensitiveContentClassificationMethodpatternMatch),
	}
}

func parseDetectedSensitiveContentClassificationMethod(input string) (*DetectedSensitiveContentClassificationMethod, error) {
	vals := map[string]DetectedSensitiveContentClassificationMethod{
		"exactdatamatch":  DetectedSensitiveContentClassificationMethodexactDataMatch,
		"fingerprint":     DetectedSensitiveContentClassificationMethodfingerprint,
		"machinelearning": DetectedSensitiveContentClassificationMethodmachineLearning,
		"patternmatch":    DetectedSensitiveContentClassificationMethodpatternMatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedSensitiveContentClassificationMethod(input)
	return &out, nil
}
