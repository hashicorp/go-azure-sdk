package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileAssessmentRequestContentType string

const (
	FileAssessmentRequestContentTypefile FileAssessmentRequestContentType = "File"
	FileAssessmentRequestContentTypemail FileAssessmentRequestContentType = "Mail"
	FileAssessmentRequestContentTypeurl  FileAssessmentRequestContentType = "Url"
)

func PossibleValuesForFileAssessmentRequestContentType() []string {
	return []string{
		string(FileAssessmentRequestContentTypefile),
		string(FileAssessmentRequestContentTypemail),
		string(FileAssessmentRequestContentTypeurl),
	}
}

func parseFileAssessmentRequestContentType(input string) (*FileAssessmentRequestContentType, error) {
	vals := map[string]FileAssessmentRequestContentType{
		"file": FileAssessmentRequestContentTypefile,
		"mail": FileAssessmentRequestContentTypemail,
		"url":  FileAssessmentRequestContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileAssessmentRequestContentType(input)
	return &out, nil
}
