package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentClassification string

const (
	SecurityIncidentClassificationfalsePositive                 SecurityIncidentClassification = "FalsePositive"
	SecurityIncidentClassificationinformationalExpectedActivity SecurityIncidentClassification = "InformationalExpectedActivity"
	SecurityIncidentClassificationtruePositive                  SecurityIncidentClassification = "TruePositive"
	SecurityIncidentClassificationunknown                       SecurityIncidentClassification = "Unknown"
)

func PossibleValuesForSecurityIncidentClassification() []string {
	return []string{
		string(SecurityIncidentClassificationfalsePositive),
		string(SecurityIncidentClassificationinformationalExpectedActivity),
		string(SecurityIncidentClassificationtruePositive),
		string(SecurityIncidentClassificationunknown),
	}
}

func parseSecurityIncidentClassification(input string) (*SecurityIncidentClassification, error) {
	vals := map[string]SecurityIncidentClassification{
		"falsepositive":                 SecurityIncidentClassificationfalsePositive,
		"informationalexpectedactivity": SecurityIncidentClassificationinformationalExpectedActivity,
		"truepositive":                  SecurityIncidentClassificationtruePositive,
		"unknown":                       SecurityIncidentClassificationunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIncidentClassification(input)
	return &out, nil
}
