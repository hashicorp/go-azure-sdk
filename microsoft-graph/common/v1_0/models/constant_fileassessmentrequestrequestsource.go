package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestRequestSource string

const (
	FileAssessmentRequestRequestSourceadministrator FileAssessmentRequestRequestSource = "Administrator"
	FileAssessmentRequestRequestSourceundefined     FileAssessmentRequestRequestSource = "Undefined"
	FileAssessmentRequestRequestSourceuser          FileAssessmentRequestRequestSource = "User"
)

func PossibleValuesForFileAssessmentRequestRequestSource() []string {
	return []string{
		string(FileAssessmentRequestRequestSourceadministrator),
		string(FileAssessmentRequestRequestSourceundefined),
		string(FileAssessmentRequestRequestSourceuser),
	}
}

func parseFileAssessmentRequestRequestSource(input string) (*FileAssessmentRequestRequestSource, error) {
	vals := map[string]FileAssessmentRequestRequestSource{
		"administrator": FileAssessmentRequestRequestSourceadministrator,
		"undefined":     FileAssessmentRequestRequestSourceundefined,
		"user":          FileAssessmentRequestRequestSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestRequestSource(input)
	return &out, nil
}
