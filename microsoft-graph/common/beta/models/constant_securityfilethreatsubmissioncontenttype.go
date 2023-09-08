package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionContentType string

const (
	SecurityFileThreatSubmissionContentTypeapp   SecurityFileThreatSubmissionContentType = "App"
	SecurityFileThreatSubmissionContentTypeemail SecurityFileThreatSubmissionContentType = "Email"
	SecurityFileThreatSubmissionContentTypefile  SecurityFileThreatSubmissionContentType = "File"
	SecurityFileThreatSubmissionContentTypeurl   SecurityFileThreatSubmissionContentType = "Url"
)

func PossibleValuesForSecurityFileThreatSubmissionContentType() []string {
	return []string{
		string(SecurityFileThreatSubmissionContentTypeapp),
		string(SecurityFileThreatSubmissionContentTypeemail),
		string(SecurityFileThreatSubmissionContentTypefile),
		string(SecurityFileThreatSubmissionContentTypeurl),
	}
}

func parseSecurityFileThreatSubmissionContentType(input string) (*SecurityFileThreatSubmissionContentType, error) {
	vals := map[string]SecurityFileThreatSubmissionContentType{
		"app":   SecurityFileThreatSubmissionContentTypeapp,
		"email": SecurityFileThreatSubmissionContentTypeemail,
		"file":  SecurityFileThreatSubmissionContentTypefile,
		"url":   SecurityFileThreatSubmissionContentTypeurl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionContentType(input)
	return &out, nil
}
