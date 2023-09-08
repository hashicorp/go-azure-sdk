package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionClientSource string

const (
	SecurityEmailUrlThreatSubmissionClientSourcemicrosoft SecurityEmailUrlThreatSubmissionClientSource = "Microsoft"
	SecurityEmailUrlThreatSubmissionClientSourceother     SecurityEmailUrlThreatSubmissionClientSource = "Other"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionClientSourcemicrosoft),
		string(SecurityEmailUrlThreatSubmissionClientSourceother),
	}
}

func parseSecurityEmailUrlThreatSubmissionClientSource(input string) (*SecurityEmailUrlThreatSubmissionClientSource, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionClientSource{
		"microsoft": SecurityEmailUrlThreatSubmissionClientSourcemicrosoft,
		"other":     SecurityEmailUrlThreatSubmissionClientSourceother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionClientSource(input)
	return &out, nil
}
