package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachineLearningDetectedSensitiveContentSensitiveTypeSource string

const (
	MachineLearningDetectedSensitiveContentSensitiveTypeSourceoutOfBox MachineLearningDetectedSensitiveContentSensitiveTypeSource = "OutOfBox"
	MachineLearningDetectedSensitiveContentSensitiveTypeSourcetenant   MachineLearningDetectedSensitiveContentSensitiveTypeSource = "Tenant"
)

func PossibleValuesForMachineLearningDetectedSensitiveContentSensitiveTypeSource() []string {
	return []string{
		string(MachineLearningDetectedSensitiveContentSensitiveTypeSourceoutOfBox),
		string(MachineLearningDetectedSensitiveContentSensitiveTypeSourcetenant),
	}
}

func parseMachineLearningDetectedSensitiveContentSensitiveTypeSource(input string) (*MachineLearningDetectedSensitiveContentSensitiveTypeSource, error) {
	vals := map[string]MachineLearningDetectedSensitiveContentSensitiveTypeSource{
		"outofbox": MachineLearningDetectedSensitiveContentSensitiveTypeSourceoutOfBox,
		"tenant":   MachineLearningDetectedSensitiveContentSensitiveTypeSourcetenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MachineLearningDetectedSensitiveContentSensitiveTypeSource(input)
	return &out, nil
}
