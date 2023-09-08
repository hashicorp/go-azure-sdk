package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestExpectedAssessment string

const (
	FileAssessmentRequestExpectedAssessmentblock   FileAssessmentRequestExpectedAssessment = "Block"
	FileAssessmentRequestExpectedAssessmentunblock FileAssessmentRequestExpectedAssessment = "Unblock"
)

func PossibleValuesForFileAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(FileAssessmentRequestExpectedAssessmentblock),
		string(FileAssessmentRequestExpectedAssessmentunblock),
	}
}

func parseFileAssessmentRequestExpectedAssessment(input string) (*FileAssessmentRequestExpectedAssessment, error) {
	vals := map[string]FileAssessmentRequestExpectedAssessment{
		"block":   FileAssessmentRequestExpectedAssessmentblock,
		"unblock": FileAssessmentRequestExpectedAssessmentunblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
