package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidenceVerdict string

const (
	SecurityGoogleCloudResourceEvidenceVerdictmalicious      SecurityGoogleCloudResourceEvidenceVerdict = "Malicious"
	SecurityGoogleCloudResourceEvidenceVerdictnoThreatsFound SecurityGoogleCloudResourceEvidenceVerdict = "NoThreatsFound"
	SecurityGoogleCloudResourceEvidenceVerdictsuspicious     SecurityGoogleCloudResourceEvidenceVerdict = "Suspicious"
	SecurityGoogleCloudResourceEvidenceVerdictunknown        SecurityGoogleCloudResourceEvidenceVerdict = "Unknown"
)

func PossibleValuesForSecurityGoogleCloudResourceEvidenceVerdict() []string {
	return []string{
		string(SecurityGoogleCloudResourceEvidenceVerdictmalicious),
		string(SecurityGoogleCloudResourceEvidenceVerdictnoThreatsFound),
		string(SecurityGoogleCloudResourceEvidenceVerdictsuspicious),
		string(SecurityGoogleCloudResourceEvidenceVerdictunknown),
	}
}

func parseSecurityGoogleCloudResourceEvidenceVerdict(input string) (*SecurityGoogleCloudResourceEvidenceVerdict, error) {
	vals := map[string]SecurityGoogleCloudResourceEvidenceVerdict{
		"malicious":      SecurityGoogleCloudResourceEvidenceVerdictmalicious,
		"nothreatsfound": SecurityGoogleCloudResourceEvidenceVerdictnoThreatsFound,
		"suspicious":     SecurityGoogleCloudResourceEvidenceVerdictsuspicious,
		"unknown":        SecurityGoogleCloudResourceEvidenceVerdictunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudResourceEvidenceVerdict(input)
	return &out, nil
}
