package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionClientSource string

const (
	SecurityThreatSubmissionClientSourcemicrosoft SecurityThreatSubmissionClientSource = "Microsoft"
	SecurityThreatSubmissionClientSourceother     SecurityThreatSubmissionClientSource = "Other"
)

func PossibleValuesForSecurityThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityThreatSubmissionClientSourcemicrosoft),
		string(SecurityThreatSubmissionClientSourceother),
	}
}

func parseSecurityThreatSubmissionClientSource(input string) (*SecurityThreatSubmissionClientSource, error) {
	vals := map[string]SecurityThreatSubmissionClientSource{
		"microsoft": SecurityThreatSubmissionClientSourcemicrosoft,
		"other":     SecurityThreatSubmissionClientSourceother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionClientSource(input)
	return &out, nil
}
