package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionContentType string

const (
	SecurityEmailUrlThreatSubmissionContentTypeapp   SecurityEmailUrlThreatSubmissionContentType = "App"
	SecurityEmailUrlThreatSubmissionContentTypeemail SecurityEmailUrlThreatSubmissionContentType = "Email"
	SecurityEmailUrlThreatSubmissionContentTypefile  SecurityEmailUrlThreatSubmissionContentType = "File"
	SecurityEmailUrlThreatSubmissionContentTypeurl   SecurityEmailUrlThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionContentType() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionContentTypeapp),
		string(SecurityEmailUrlThreatSubmissionContentTypeemail),
		string(SecurityEmailUrlThreatSubmissionContentTypefile),
		string(SecurityEmailUrlThreatSubmissionContentTypeurl),
	}
}

func parseSecurityEmailUrlThreatSubmissionContentType(input string) (*SecurityEmailUrlThreatSubmissionContentType, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionContentType{
		"app":   SecurityEmailUrlThreatSubmissionContentTypeapp,
		"email": SecurityEmailUrlThreatSubmissionContentTypeemail,
		"file":  SecurityEmailUrlThreatSubmissionContentTypefile,
		"url":   SecurityEmailUrlThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionContentType(input)
	return &out, nil
}
