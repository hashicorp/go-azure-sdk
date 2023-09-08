package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedMessageEvidenceVerdict string

const (
	SecurityAnalyzedMessageEvidenceVerdictmalicious      SecurityAnalyzedMessageEvidenceVerdict = "Malicious"
	SecurityAnalyzedMessageEvidenceVerdictnoThreatsFound SecurityAnalyzedMessageEvidenceVerdict = "NoThreatsFound"
	SecurityAnalyzedMessageEvidenceVerdictsuspicious     SecurityAnalyzedMessageEvidenceVerdict = "Suspicious"
	SecurityAnalyzedMessageEvidenceVerdictunknown        SecurityAnalyzedMessageEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityAnalyzedMessageEvidenceVerdict() []string {
	return []string{
		string(SecurityAnalyzedMessageEvidenceVerdictmalicious),
		string(SecurityAnalyzedMessageEvidenceVerdictnoThreatsFound),
		string(SecurityAnalyzedMessageEvidenceVerdictsuspicious),
		string(SecurityAnalyzedMessageEvidenceVerdictunknown),
	}
}

func parseSecurityAnalyzedMessageEvidenceVerdict(input string) (*SecurityAnalyzedMessageEvidenceVerdict, error) {
	vals := map[string]SecurityAnalyzedMessageEvidenceVerdict{
		"malicious":      SecurityAnalyzedMessageEvidenceVerdictmalicious,
		"nothreatsfound": SecurityAnalyzedMessageEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityAnalyzedMessageEvidenceVerdictsuspicious,
		"unknown":        SecurityAnalyzedMessageEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAnalyzedMessageEvidenceVerdict(input)
	return &out, nil
}
