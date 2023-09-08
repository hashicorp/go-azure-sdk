package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionSource string

const (
	SecurityEmailUrlThreatSubmissionSourceadministrator SecurityEmailUrlThreatSubmissionSource = "Administrator"
	SecurityEmailUrlThreatSubmissionSourceuser          SecurityEmailUrlThreatSubmissionSource = "User"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionSource() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionSourceadministrator),
		string(SecurityEmailUrlThreatSubmissionSourceuser),
	}
}

func parseSecurityEmailUrlThreatSubmissionSource(input string) (*SecurityEmailUrlThreatSubmissionSource, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionSource{
		"administrator": SecurityEmailUrlThreatSubmissionSourceadministrator,
		"user":          SecurityEmailUrlThreatSubmissionSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionSource(input)
	return &out, nil
}
