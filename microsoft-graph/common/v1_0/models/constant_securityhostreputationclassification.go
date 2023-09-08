package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostReputationClassification string

const (
	SecurityHostReputationClassificationmalicious  SecurityHostReputationClassification = "Malicious"
	SecurityHostReputationClassificationneutral    SecurityHostReputationClassification = "Neutral"
	SecurityHostReputationClassificationsuspicious SecurityHostReputationClassification = "Suspicious"
	SecurityHostReputationClassificationunknown    SecurityHostReputationClassification = "Unknown"
)

func PossibleValuesForSecurityHostReputationClassification() []string {
	return []string{
		string(SecurityHostReputationClassificationmalicious),
		string(SecurityHostReputationClassificationneutral),
		string(SecurityHostReputationClassificationsuspicious),
		string(SecurityHostReputationClassificationunknown),
	}
}

func parseSecurityHostReputationClassification(input string) (*SecurityHostReputationClassification, error) {
	vals := map[string]SecurityHostReputationClassification{
		"malicious":  SecurityHostReputationClassificationmalicious,
		"neutral":    SecurityHostReputationClassificationneutral,
		"suspicious": SecurityHostReputationClassificationsuspicious,
		"unknown":    SecurityHostReputationClassificationunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostReputationClassification(input)
	return &out, nil
}
