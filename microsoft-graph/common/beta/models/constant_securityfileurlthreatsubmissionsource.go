package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmissionSource string

const (
	SecurityFileUrlThreatSubmissionSourceadministrator SecurityFileUrlThreatSubmissionSource = "Administrator"
	SecurityFileUrlThreatSubmissionSourceuser          SecurityFileUrlThreatSubmissionSource = "User"
)

func PossibleValuesForSecurityFileUrlThreatSubmissionSource() []string {
	return []string{
		string(SecurityFileUrlThreatSubmissionSourceadministrator),
		string(SecurityFileUrlThreatSubmissionSourceuser),
	}
}

func parseSecurityFileUrlThreatSubmissionSource(input string) (*SecurityFileUrlThreatSubmissionSource, error) {
	vals := map[string]SecurityFileUrlThreatSubmissionSource{
		"administrator": SecurityFileUrlThreatSubmissionSourceadministrator,
		"user":          SecurityFileUrlThreatSubmissionSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileUrlThreatSubmissionSource(input)
	return &out, nil
}
