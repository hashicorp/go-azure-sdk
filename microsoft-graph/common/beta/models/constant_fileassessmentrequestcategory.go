package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestCategory string

const (
	FileAssessmentRequestCategorymalware   FileAssessmentRequestCategory = "Malware"
	FileAssessmentRequestCategoryphishing  FileAssessmentRequestCategory = "Phishing"
	FileAssessmentRequestCategoryspam      FileAssessmentRequestCategory = "Spam"
	FileAssessmentRequestCategoryundefined FileAssessmentRequestCategory = "Undefined"
)

func PossibleValuesForFileAssessmentRequestCategory() []string {
	return []string{
		string(FileAssessmentRequestCategorymalware),
		string(FileAssessmentRequestCategoryphishing),
		string(FileAssessmentRequestCategoryspam),
		string(FileAssessmentRequestCategoryundefined),
	}
}

func parseFileAssessmentRequestCategory(input string) (*FileAssessmentRequestCategory, error) {
	vals := map[string]FileAssessmentRequestCategory{
		"malware":   FileAssessmentRequestCategorymalware,
		"phishing":  FileAssessmentRequestCategoryphishing,
		"spam":      FileAssessmentRequestCategoryspam,
		"undefined": FileAssessmentRequestCategoryundefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestCategory(input)
	return &out, nil
}
