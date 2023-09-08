package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionSource string

const (
	SecurityEmailThreatSubmissionSourceadministrator SecurityEmailThreatSubmissionSource = "Administrator"
	SecurityEmailThreatSubmissionSourceuser          SecurityEmailThreatSubmissionSource = "User"
)

func PossibleValuesForSecurityEmailThreatSubmissionSource() []string {
	return []string{
		string(SecurityEmailThreatSubmissionSourceadministrator),
		string(SecurityEmailThreatSubmissionSourceuser),
	}
}

func parseSecurityEmailThreatSubmissionSource(input string) (*SecurityEmailThreatSubmissionSource, error) {
	vals := map[string]SecurityEmailThreatSubmissionSource{
		"administrator": SecurityEmailThreatSubmissionSourceadministrator,
		"user":          SecurityEmailThreatSubmissionSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionSource(input)
	return &out, nil
}
