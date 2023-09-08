package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionSource string

const (
	SecurityThreatSubmissionSourceadministrator SecurityThreatSubmissionSource = "Administrator"
	SecurityThreatSubmissionSourceuser          SecurityThreatSubmissionSource = "User"
)

func PossibleValuesForSecurityThreatSubmissionSource() []string {
	return []string{
		string(SecurityThreatSubmissionSourceadministrator),
		string(SecurityThreatSubmissionSourceuser),
	}
}

func parseSecurityThreatSubmissionSource(input string) (*SecurityThreatSubmissionSource, error) {
	vals := map[string]SecurityThreatSubmissionSource{
		"administrator": SecurityThreatSubmissionSourceadministrator,
		"user":          SecurityThreatSubmissionSourceuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionSource(input)
	return &out, nil
}
