package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionContentType string

const (
	SecurityEmailThreatSubmissionContentTypeapp   SecurityEmailThreatSubmissionContentType = "App"
	SecurityEmailThreatSubmissionContentTypeemail SecurityEmailThreatSubmissionContentType = "Email"
	SecurityEmailThreatSubmissionContentTypefile  SecurityEmailThreatSubmissionContentType = "File"
	SecurityEmailThreatSubmissionContentTypeurl   SecurityEmailThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityEmailThreatSubmissionContentType() []string {
	return []string{
		string(SecurityEmailThreatSubmissionContentTypeapp),
		string(SecurityEmailThreatSubmissionContentTypeemail),
		string(SecurityEmailThreatSubmissionContentTypefile),
		string(SecurityEmailThreatSubmissionContentTypeurl),
	}
}

func parseSecurityEmailThreatSubmissionContentType(input string) (*SecurityEmailThreatSubmissionContentType, error) {
	vals := map[string]SecurityEmailThreatSubmissionContentType{
		"app":   SecurityEmailThreatSubmissionContentTypeapp,
		"email": SecurityEmailThreatSubmissionContentTypeemail,
		"file":  SecurityEmailThreatSubmissionContentTypefile,
		"url":   SecurityEmailThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionContentType(input)
	return &out, nil
}
