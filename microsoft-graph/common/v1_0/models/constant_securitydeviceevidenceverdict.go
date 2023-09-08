package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceVerdict string

const (
	SecurityDeviceEvidenceVerdictmalicious      SecurityDeviceEvidenceVerdict = "Malicious"
	SecurityDeviceEvidenceVerdictnoThreatsFound SecurityDeviceEvidenceVerdict = "NoThreatsFound"
	SecurityDeviceEvidenceVerdictsuspicious     SecurityDeviceEvidenceVerdict = "Suspicious"
	SecurityDeviceEvidenceVerdictunknown        SecurityDeviceEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityDeviceEvidenceVerdict() []string {
	return []string{
		string(SecurityDeviceEvidenceVerdictmalicious),
		string(SecurityDeviceEvidenceVerdictnoThreatsFound),
		string(SecurityDeviceEvidenceVerdictsuspicious),
		string(SecurityDeviceEvidenceVerdictunknown),
	}
}

func parseSecurityDeviceEvidenceVerdict(input string) (*SecurityDeviceEvidenceVerdict, error) {
	vals := map[string]SecurityDeviceEvidenceVerdict{
		"malicious":      SecurityDeviceEvidenceVerdictmalicious,
		"nothreatsfound": SecurityDeviceEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityDeviceEvidenceVerdictsuspicious,
		"unknown":        SecurityDeviceEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceVerdict(input)
	return &out, nil
}
