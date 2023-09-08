package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionContentType string

const (
	SecurityFileContentThreatSubmissionContentTypeapp   SecurityFileContentThreatSubmissionContentType = "App"
	SecurityFileContentThreatSubmissionContentTypeemail SecurityFileContentThreatSubmissionContentType = "Email"
	SecurityFileContentThreatSubmissionContentTypefile  SecurityFileContentThreatSubmissionContentType = "File"
	SecurityFileContentThreatSubmissionContentTypeurl   SecurityFileContentThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityFileContentThreatSubmissionContentType() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionContentTypeapp),
		string(SecurityFileContentThreatSubmissionContentTypeemail),
		string(SecurityFileContentThreatSubmissionContentTypefile),
		string(SecurityFileContentThreatSubmissionContentTypeurl),
	}
}

func parseSecurityFileContentThreatSubmissionContentType(input string) (*SecurityFileContentThreatSubmissionContentType, error) {
	vals := map[string]SecurityFileContentThreatSubmissionContentType{
		"app":   SecurityFileContentThreatSubmissionContentTypeapp,
		"email": SecurityFileContentThreatSubmissionContentTypeemail,
		"file":  SecurityFileContentThreatSubmissionContentTypefile,
		"url":   SecurityFileContentThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionContentType(input)
	return &out, nil
}
