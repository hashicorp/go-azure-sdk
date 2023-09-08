package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestContentType string

const (
	EmailFileAssessmentRequestContentTypefile EmailFileAssessmentRequestContentType = "File"
	EmailFileAssessmentRequestContentTypemail EmailFileAssessmentRequestContentType = "Mail"
	EmailFileAssessmentRequestContentTypeurl  EmailFileAssessmentRequestContentType = "Url"
)

func PossibleValuesForEmailFileAssessmentRequestContentType() []string {
	return []string{
		string(EmailFileAssessmentRequestContentTypefile),
		string(EmailFileAssessmentRequestContentTypemail),
		string(EmailFileAssessmentRequestContentTypeurl),
	}
}

func parseEmailFileAssessmentRequestContentType(input string) (*EmailFileAssessmentRequestContentType, error) {
	vals := map[string]EmailFileAssessmentRequestContentType{
		"file": EmailFileAssessmentRequestContentTypefile,
		"mail": EmailFileAssessmentRequestContentTypemail,
		"url":  EmailFileAssessmentRequestContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestContentType(input)
	return &out, nil
}
