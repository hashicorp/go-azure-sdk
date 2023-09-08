package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionContentType string

const (
	SecurityEmailContentThreatSubmissionContentTypeapp   SecurityEmailContentThreatSubmissionContentType = "App"
	SecurityEmailContentThreatSubmissionContentTypeemail SecurityEmailContentThreatSubmissionContentType = "Email"
	SecurityEmailContentThreatSubmissionContentTypefile  SecurityEmailContentThreatSubmissionContentType = "File"
	SecurityEmailContentThreatSubmissionContentTypeurl   SecurityEmailContentThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionContentType() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionContentTypeapp),
		string(SecurityEmailContentThreatSubmissionContentTypeemail),
		string(SecurityEmailContentThreatSubmissionContentTypefile),
		string(SecurityEmailContentThreatSubmissionContentTypeurl),
	}
}

func parseSecurityEmailContentThreatSubmissionContentType(input string) (*SecurityEmailContentThreatSubmissionContentType, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionContentType{
		"app":   SecurityEmailContentThreatSubmissionContentTypeapp,
		"email": SecurityEmailContentThreatSubmissionContentTypeemail,
		"file":  SecurityEmailContentThreatSubmissionContentTypefile,
		"url":   SecurityEmailContentThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionContentType(input)
	return &out, nil
}
