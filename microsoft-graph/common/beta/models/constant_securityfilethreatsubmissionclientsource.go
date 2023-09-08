package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionClientSource string

const (
	SecurityFileThreatSubmissionClientSourcemicrosoft SecurityFileThreatSubmissionClientSource = "Microsoft"
	SecurityFileThreatSubmissionClientSourceother     SecurityFileThreatSubmissionClientSource = "Other"
)

func PossibleValuesForSecurityFileThreatSubmissionClientSource() []string {
	return []string{
		string(SecurityFileThreatSubmissionClientSourcemicrosoft),
		string(SecurityFileThreatSubmissionClientSourceother),
	}
}

func parseSecurityFileThreatSubmissionClientSource(input string) (*SecurityFileThreatSubmissionClientSource, error) {
	vals := map[string]SecurityFileThreatSubmissionClientSource{
		"microsoft": SecurityFileThreatSubmissionClientSourcemicrosoft,
		"other":     SecurityFileThreatSubmissionClientSourceother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionClientSource(input)
	return &out, nil
}
