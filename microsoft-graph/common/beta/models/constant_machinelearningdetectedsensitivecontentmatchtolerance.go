package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentMatchTolerance string

const (
	MachineLearningDetectedSensitiveContentMatchToleranceexact MachineLearningDetectedSensitiveContentMatchTolerance = "Exact"
	MachineLearningDetectedSensitiveContentMatchTolerancenear  MachineLearningDetectedSensitiveContentMatchTolerance = "Near"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentMatchTolerance() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentMatchToleranceexact),
		string(MachineLearningDetectedSensitiveContentMatchTolerancenear),
	}
}

func parseMachineLearningDetectedSensitiveContentMatchTolerance(input string) (*MachineLearningDetectedSensitiveContentMatchTolerance, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentMatchTolerance{
		"exact": MachineLearningDetectedSensitiveContentMatchToleranceexact,
		"near":  MachineLearningDetectedSensitiveContentMatchTolerancenear,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentMatchTolerance(input)
	return &out, nil
}
