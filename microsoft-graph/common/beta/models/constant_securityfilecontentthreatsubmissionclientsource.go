package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionClientSource string

const (
	SecurityFileContentThreatSubmissionClientSourcemicrosoft SecurityFileContentThreatSubmissionClientSource = "Microsoft"
	SecurityFileContentThreatSubmissionClientSourceother     SecurityFileContentThreatSubmissionClientSource = "Other"
)

func PossibleValuesForSecurityFileContentThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionClientSourcemicrosoft),
		string(SecurityFileContentThreatSubmissionClientSourceother),
	}
}

func parseSecurityFileContentThreatSubmissionClientSource(input string) (*SecurityFileContentThreatSubmissionClientSource, error) {
	vals := map[string]SecurityFileContentThreatSubmissionClientSource{
		"microsoft": SecurityFileContentThreatSubmissionClientSourcemicrosoft,
		"other":     SecurityFileContentThreatSubmissionClientSourceother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionClientSource(input)
	return &out, nil
}
