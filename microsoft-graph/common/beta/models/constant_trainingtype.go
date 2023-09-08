package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingType string

const (
	TrainingTypephishing TrainingType = "Phishing"
	TrainingTypeunknown  TrainingType = "Unknown"
)

func PossibleValuesForTrainingType() []string {
	return []string{
		string(TrainingTypephishing),
		string(TrainingTypeunknown),
	}
}

func parseTrainingType(input string) (*TrainingType, error) {
	vals := map[string]TrainingType{
		"phishing": TrainingTypephishing,
		"unknown":  TrainingTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingType(input)
	return &out, nil
}
