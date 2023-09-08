package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxEvidenceVerdict string

const (
	SecurityMailboxEvidenceVerdictmalicious      SecurityMailboxEvidenceVerdict = "Malicious"
	SecurityMailboxEvidenceVerdictnoThreatsFound SecurityMailboxEvidenceVerdict = "NoThreatsFound"
	SecurityMailboxEvidenceVerdictsuspicious     SecurityMailboxEvidenceVerdict = "Suspicious"
	SecurityMailboxEvidenceVerdictunknown        SecurityMailboxEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityMailboxEvidenceVerdict() []string {
	return []string{
		string(SecurityMailboxEvidenceVerdictmalicious),
		string(SecurityMailboxEvidenceVerdictnoThreatsFound),
		string(SecurityMailboxEvidenceVerdictsuspicious),
		string(SecurityMailboxEvidenceVerdictunknown),
	}
}

func parseSecurityMailboxEvidenceVerdict(input string) (*SecurityMailboxEvidenceVerdict, error) {
	vals := map[string]SecurityMailboxEvidenceVerdict{
		"malicious":      SecurityMailboxEvidenceVerdictmalicious,
		"nothreatsfound": SecurityMailboxEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityMailboxEvidenceVerdictsuspicious,
		"unknown":        SecurityMailboxEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailboxEvidenceVerdict(input)
	return &out, nil
}
