package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestContentType string

const (
	MailAssessmentRequestContentTypefile MailAssessmentRequestContentType = "File"
	MailAssessmentRequestContentTypemail MailAssessmentRequestContentType = "Mail"
	MailAssessmentRequestContentTypeurl  MailAssessmentRequestContentType = "Url"
)

func PossibleValuesForMailAssessmentRequestContentType() []string {
	return []string{
		string(MailAssessmentRequestContentTypefile),
		string(MailAssessmentRequestContentTypemail),
		string(MailAssessmentRequestContentTypeurl),
	}
}

func parseMailAssessmentRequestContentType(input string) (*MailAssessmentRequestContentType, error) {
	vals := map[string]MailAssessmentRequestContentType{
		"file": MailAssessmentRequestContentTypefile,
		"mail": MailAssessmentRequestContentTypemail,
		"url":  MailAssessmentRequestContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestContentType(input)
	return &out, nil
}
