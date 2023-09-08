package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionSource string

const (
	SecurityFileThreatSubmissionSourceadministrator SecurityFileThreatSubmissionSource = "Administrator"
	SecurityFileThreatSubmissionSourceuser          SecurityFileThreatSubmissionSource = "User"
)

func PossibleValuesForSecurityFileThreatSubmissionSource() []string {
	return []string{
		string(SecurityFileThreatSubmissionSourceadministrator),
		string(SecurityFileThreatSubmissionSourceuser),
	}
}

func parseSecurityFileThreatSubmissionSource(input string) (*SecurityFileThreatSubmissionSource, error) {
	vals := map[string]SecurityFileThreatSubmissionSource{
		"administrator": SecurityFileThreatSubmissionSourceadministrator,
		"user":          SecurityFileThreatSubmissionSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionSource(input)
	return &out, nil
}
