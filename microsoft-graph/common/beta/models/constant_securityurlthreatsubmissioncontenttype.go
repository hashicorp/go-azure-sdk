package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionContentType string

const (
	SecurityUrlThreatSubmissionContentTypeapp   SecurityUrlThreatSubmissionContentType = "App"
	SecurityUrlThreatSubmissionContentTypeemail SecurityUrlThreatSubmissionContentType = "Email"
	SecurityUrlThreatSubmissionContentTypefile  SecurityUrlThreatSubmissionContentType = "File"
	SecurityUrlThreatSubmissionContentTypeurl   SecurityUrlThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityUrlThreatSubmissionContentType() []string {
	return []string{
		string(SecurityUrlThreatSubmissionContentTypeapp),
		string(SecurityUrlThreatSubmissionContentTypeemail),
		string(SecurityUrlThreatSubmissionContentTypefile),
		string(SecurityUrlThreatSubmissionContentTypeurl),
	}
}

func parseSecurityUrlThreatSubmissionContentType(input string) (*SecurityUrlThreatSubmissionContentType, error) {
	vals := map[string]SecurityUrlThreatSubmissionContentType{
		"app":   SecurityUrlThreatSubmissionContentTypeapp,
		"email": SecurityUrlThreatSubmissionContentTypeemail,
		"file":  SecurityUrlThreatSubmissionContentTypefile,
		"url":   SecurityUrlThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionContentType(input)
	return &out, nil
}
