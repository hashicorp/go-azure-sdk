package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestStatus string

const (
	FileAssessmentRequestStatuscompleted FileAssessmentRequestStatus = "Completed"
	FileAssessmentRequestStatuspending   FileAssessmentRequestStatus = "Pending"
)

func PossibleValuesForFileAssessmentRequestStatus() []string {
	return []string{
		string(FileAssessmentRequestStatuscompleted),
		string(FileAssessmentRequestStatuspending),
	}
}

func parseFileAssessmentRequestStatus(input string) (*FileAssessmentRequestStatus, error) {
	vals := map[string]FileAssessmentRequestStatus{
		"completed": FileAssessmentRequestStatuscompleted,
		"pending":   FileAssessmentRequestStatuspending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestStatus(input)
	return &out, nil
}
