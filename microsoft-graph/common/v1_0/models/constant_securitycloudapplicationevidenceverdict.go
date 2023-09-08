package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudApplicationEvidenceVerdict string

const (
	SecurityCloudApplicationEvidenceVerdictmalicious      SecurityCloudApplicationEvidenceVerdict = "Malicious"
	SecurityCloudApplicationEvidenceVerdictnoThreatsFound SecurityCloudApplicationEvidenceVerdict = "NoThreatsFound"
	SecurityCloudApplicationEvidenceVerdictsuspicious     SecurityCloudApplicationEvidenceVerdict = "Suspicious"
	SecurityCloudApplicationEvidenceVerdictunknown        SecurityCloudApplicationEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityCloudApplicationEvidenceVerdict() []string {
	return []string{
		string(SecurityCloudApplicationEvidenceVerdictmalicious),
		string(SecurityCloudApplicationEvidenceVerdictnoThreatsFound),
		string(SecurityCloudApplicationEvidenceVerdictsuspicious),
		string(SecurityCloudApplicationEvidenceVerdictunknown),
	}
}

func parseSecurityCloudApplicationEvidenceVerdict(input string) (*SecurityCloudApplicationEvidenceVerdict, error) {
	vals := map[string]SecurityCloudApplicationEvidenceVerdict{
		"malicious":      SecurityCloudApplicationEvidenceVerdictmalicious,
		"nothreatsfound": SecurityCloudApplicationEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityCloudApplicationEvidenceVerdictsuspicious,
		"unknown":        SecurityCloudApplicationEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCloudApplicationEvidenceVerdict(input)
	return &out, nil
}
