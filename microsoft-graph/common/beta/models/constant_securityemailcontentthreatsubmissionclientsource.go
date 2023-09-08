package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionClientSource string

const (
	SecurityEmailContentThreatSubmissionClientSourcemicrosoft SecurityEmailContentThreatSubmissionClientSource = "Microsoft"
	SecurityEmailContentThreatSubmissionClientSourceother     SecurityEmailContentThreatSubmissionClientSource = "Other"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionClientSourcemicrosoft),
		string(SecurityEmailContentThreatSubmissionClientSourceother),
	}
}

func parseSecurityEmailContentThreatSubmissionClientSource(input string) (*SecurityEmailContentThreatSubmissionClientSource, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionClientSource{
		"microsoft": SecurityEmailContentThreatSubmissionClientSourcemicrosoft,
		"other":     SecurityEmailContentThreatSubmissionClientSourceother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionClientSource(input)
	return &out, nil
}
