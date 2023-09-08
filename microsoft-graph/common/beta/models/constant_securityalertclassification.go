package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertClassification string

const (
	SecurityAlertClassificationfalsePositive                 SecurityAlertClassification = "FalsePositive"
	SecurityAlertClassificationinformationalExpectedActivity SecurityAlertClassification = "InformationalExpectedActivity"
	SecurityAlertClassificationtruePositive                  SecurityAlertClassification = "TruePositive"
	SecurityAlertClassificationunknown                       SecurityAlertClassification = "Unknown"
)

func PossibleValuesForSecurityAlertClassification() []string {
	return []string{
		string(SecurityAlertClassificationfalsePositive),
		string(SecurityAlertClassificationinformationalExpectedActivity),
		string(SecurityAlertClassificationtruePositive),
		string(SecurityAlertClassificationunknown),
	}
}

func parseSecurityAlertClassification(input string) (*SecurityAlertClassification, error) {
	vals := map[string]SecurityAlertClassification{
		"falsepositive":                 SecurityAlertClassificationfalsePositive,
		"informationalexpectedactivity": SecurityAlertClassificationinformationalExpectedActivity,
		"truepositive":                  SecurityAlertClassificationtruePositive,
		"unknown":                       SecurityAlertClassificationunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertClassification(input)
	return &out, nil
}
