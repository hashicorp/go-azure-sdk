package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestContentType string

const (
	UrlAssessmentRequestContentTypefile UrlAssessmentRequestContentType = "File"
	UrlAssessmentRequestContentTypemail UrlAssessmentRequestContentType = "Mail"
	UrlAssessmentRequestContentTypeurl  UrlAssessmentRequestContentType = "Url"
)

func PossibleValuesForUrlAssessmentRequestContentType() []string {
	return []string{
		string(UrlAssessmentRequestContentTypefile),
		string(UrlAssessmentRequestContentTypemail),
		string(UrlAssessmentRequestContentTypeurl),
	}
}

func parseUrlAssessmentRequestContentType(input string) (*UrlAssessmentRequestContentType, error) {
	vals := map[string]UrlAssessmentRequestContentType{
		"file": UrlAssessmentRequestContentTypefile,
		"mail": UrlAssessmentRequestContentTypemail,
		"url":  UrlAssessmentRequestContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestContentType(input)
	return &out, nil
}
