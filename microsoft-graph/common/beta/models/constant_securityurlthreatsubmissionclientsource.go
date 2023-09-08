package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionClientSource string

const (
	SecurityUrlThreatSubmissionClientSourcemicrosoft SecurityUrlThreatSubmissionClientSource = "Microsoft"
	SecurityUrlThreatSubmissionClientSourceother     SecurityUrlThreatSubmissionClientSource = "Other"
)

func PossibleValuesForSecurityUrlThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityUrlThreatSubmissionClientSourcemicrosoft),
		string(SecurityUrlThreatSubmissionClientSourceother),
	}
}

func parseSecurityUrlThreatSubmissionClientSource(input string) (*SecurityUrlThreatSubmissionClientSource, error) {
	vals := map[string]SecurityUrlThreatSubmissionClientSource{
		"microsoft": SecurityUrlThreatSubmissionClientSourcemicrosoft,
		"other":     SecurityUrlThreatSubmissionClientSourceother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionClientSource(input)
	return &out, nil
}
