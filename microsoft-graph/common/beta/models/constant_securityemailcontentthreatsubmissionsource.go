package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionSource string

const (
	SecurityEmailContentThreatSubmissionSourceadministrator SecurityEmailContentThreatSubmissionSource = "Administrator"
	SecurityEmailContentThreatSubmissionSourceuser          SecurityEmailContentThreatSubmissionSource = "User"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionSource() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionSourceadministrator),
		string(SecurityEmailContentThreatSubmissionSourceuser),
	}
}

func parseSecurityEmailContentThreatSubmissionSource(input string) (*SecurityEmailContentThreatSubmissionSource, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionSource{
		"administrator": SecurityEmailContentThreatSubmissionSourceadministrator,
		"user":          SecurityEmailContentThreatSubmissionSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionSource(input)
	return &out, nil
}
