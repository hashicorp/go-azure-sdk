package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmissionContentType string

const (
	SecurityFileUrlThreatSubmissionContentTypeapp   SecurityFileUrlThreatSubmissionContentType = "App"
	SecurityFileUrlThreatSubmissionContentTypeemail SecurityFileUrlThreatSubmissionContentType = "Email"
	SecurityFileUrlThreatSubmissionContentTypefile  SecurityFileUrlThreatSubmissionContentType = "File"
	SecurityFileUrlThreatSubmissionContentTypeurl   SecurityFileUrlThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityFileUrlThreatSubmissionContentType() []string {
	return []string{
		string(SecurityFileUrlThreatSubmissionContentTypeapp),
		string(SecurityFileUrlThreatSubmissionContentTypeemail),
		string(SecurityFileUrlThreatSubmissionContentTypefile),
		string(SecurityFileUrlThreatSubmissionContentTypeurl),
	}
}

func parseSecurityFileUrlThreatSubmissionContentType(input string) (*SecurityFileUrlThreatSubmissionContentType, error) {
	vals := map[string]SecurityFileUrlThreatSubmissionContentType{
		"app":   SecurityFileUrlThreatSubmissionContentTypeapp,
		"email": SecurityFileUrlThreatSubmissionContentTypeemail,
		"file":  SecurityFileUrlThreatSubmissionContentTypefile,
		"url":   SecurityFileUrlThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileUrlThreatSubmissionContentType(input)
	return &out, nil
}
