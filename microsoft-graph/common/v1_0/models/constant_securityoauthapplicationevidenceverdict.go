package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOauthApplicationEvidenceVerdict string

const (
	SecurityOauthApplicationEvidenceVerdictmalicious      SecurityOauthApplicationEvidenceVerdict = "Malicious"
	SecurityOauthApplicationEvidenceVerdictnoThreatsFound SecurityOauthApplicationEvidenceVerdict = "NoThreatsFound"
	SecurityOauthApplicationEvidenceVerdictsuspicious     SecurityOauthApplicationEvidenceVerdict = "Suspicious"
	SecurityOauthApplicationEvidenceVerdictunknown        SecurityOauthApplicationEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityOauthApplicationEvidenceVerdict() []string {
	return []string{
		string(SecurityOauthApplicationEvidenceVerdictmalicious),
		string(SecurityOauthApplicationEvidenceVerdictnoThreatsFound),
		string(SecurityOauthApplicationEvidenceVerdictsuspicious),
		string(SecurityOauthApplicationEvidenceVerdictunknown),
	}
}

func parseSecurityOauthApplicationEvidenceVerdict(input string) (*SecurityOauthApplicationEvidenceVerdict, error) {
	vals := map[string]SecurityOauthApplicationEvidenceVerdict{
		"malicious":      SecurityOauthApplicationEvidenceVerdictmalicious,
		"nothreatsfound": SecurityOauthApplicationEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityOauthApplicationEvidenceVerdictsuspicious,
		"unknown":        SecurityOauthApplicationEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOauthApplicationEvidenceVerdict(input)
	return &out, nil
}
