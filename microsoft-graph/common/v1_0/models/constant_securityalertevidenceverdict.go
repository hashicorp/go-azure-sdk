package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertEvidenceVerdict string

const (
	SecurityAlertEvidenceVerdictmalicious      SecurityAlertEvidenceVerdict = "Malicious"
	SecurityAlertEvidenceVerdictnoThreatsFound SecurityAlertEvidenceVerdict = "NoThreatsFound"
	SecurityAlertEvidenceVerdictsuspicious     SecurityAlertEvidenceVerdict = "Suspicious"
	SecurityAlertEvidenceVerdictunknown        SecurityAlertEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityAlertEvidenceVerdict() []string {
	return []string{
		string(SecurityAlertEvidenceVerdictmalicious),
		string(SecurityAlertEvidenceVerdictnoThreatsFound),
		string(SecurityAlertEvidenceVerdictsuspicious),
		string(SecurityAlertEvidenceVerdictunknown),
	}
}

func parseSecurityAlertEvidenceVerdict(input string) (*SecurityAlertEvidenceVerdict, error) {
	vals := map[string]SecurityAlertEvidenceVerdict{
		"malicious":      SecurityAlertEvidenceVerdictmalicious,
		"nothreatsfound": SecurityAlertEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityAlertEvidenceVerdictsuspicious,
		"unknown":        SecurityAlertEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertEvidenceVerdict(input)
	return &out, nil
}
