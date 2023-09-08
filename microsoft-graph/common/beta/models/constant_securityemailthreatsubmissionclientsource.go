package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionClientSource string

const (
	SecurityEmailThreatSubmissionClientSourcemicrosoft SecurityEmailThreatSubmissionClientSource = "Microsoft"
	SecurityEmailThreatSubmissionClientSourceother     SecurityEmailThreatSubmissionClientSource = "Other"
)

func PossibleValuesForSecurityEmailThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityEmailThreatSubmissionClientSourcemicrosoft),
		string(SecurityEmailThreatSubmissionClientSourceother),
	}
}

func parseSecurityEmailThreatSubmissionClientSource(input string) (*SecurityEmailThreatSubmissionClientSource, error) {
	vals := map[string]SecurityEmailThreatSubmissionClientSource{
		"microsoft": SecurityEmailThreatSubmissionClientSourcemicrosoft,
		"other":     SecurityEmailThreatSubmissionClientSourceother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionClientSource(input)
	return &out, nil
}
