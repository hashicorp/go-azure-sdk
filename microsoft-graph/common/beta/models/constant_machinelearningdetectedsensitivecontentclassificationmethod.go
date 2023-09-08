package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentClassificationMethod string

const (
	MachineLearningDetectedSensitiveContentClassificationMethodexactDataMatch  MachineLearningDetectedSensitiveContentClassificationMethod = "ExactDataMatch"
	MachineLearningDetectedSensitiveContentClassificationMethodfingerprint     MachineLearningDetectedSensitiveContentClassificationMethod = "Fingerprint"
	MachineLearningDetectedSensitiveContentClassificationMethodmachineLearning MachineLearningDetectedSensitiveContentClassificationMethod = "MachineLearning"
	MachineLearningDetectedSensitiveContentClassificationMethodpatternMatch    MachineLearningDetectedSensitiveContentClassificationMethod = "PatternMatch"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentClassificationMethod() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentClassificationMethodexactDataMatch),
		string(MachineLearningDetectedSensitiveContentClassificationMethodfingerprint),
		string(MachineLearningDetectedSensitiveContentClassificationMethodmachineLearning),
		string(MachineLearningDetectedSensitiveContentClassificationMethodpatternMatch),
	}
}

func parseMachineLearningDetectedSensitiveContentClassificationMethod(input string) (*MachineLearningDetectedSensitiveContentClassificationMethod, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentClassificationMethod{
		"exactdatamatch":  MachineLearningDetectedSensitiveContentClassificationMethodexactDataMatch,
		"fingerprint":     MachineLearningDetectedSensitiveContentClassificationMethodfingerprint,
		"machinelearning": MachineLearningDetectedSensitiveContentClassificationMethodmachineLearning,
		"patternmatch":    MachineLearningDetectedSensitiveContentClassificationMethodpatternMatch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentClassificationMethod(input)
	return &out, nil
}
