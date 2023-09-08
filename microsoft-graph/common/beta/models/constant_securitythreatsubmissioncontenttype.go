package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionContentType string

const (
	SecurityThreatSubmissionContentTypeapp   SecurityThreatSubmissionContentType = "App"
	SecurityThreatSubmissionContentTypeemail SecurityThreatSubmissionContentType = "Email"
	SecurityThreatSubmissionContentTypefile  SecurityThreatSubmissionContentType = "File"
	SecurityThreatSubmissionContentTypeurl   SecurityThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityThreatSubmissionContentType() []string {
	return []string{
		string(SecurityThreatSubmissionContentTypeapp),
		string(SecurityThreatSubmissionContentTypeemail),
		string(SecurityThreatSubmissionContentTypefile),
		string(SecurityThreatSubmissionContentTypeurl),
	}
}

func parseSecurityThreatSubmissionContentType(input string) (*SecurityThreatSubmissionContentType, error) {
	vals := map[string]SecurityThreatSubmissionContentType{
		"app":   SecurityThreatSubmissionContentTypeapp,
		"email": SecurityThreatSubmissionContentTypeemail,
		"file":  SecurityThreatSubmissionContentTypefile,
		"url":   SecurityThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionContentType(input)
	return &out, nil
}
