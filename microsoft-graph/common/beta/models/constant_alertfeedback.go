package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertFeedback string

const (
	AlertFeedbackbenignPositive AlertFeedback = "BenignPositive"
	AlertFeedbackfalsePositive  AlertFeedback = "FalsePositive"
	AlertFeedbacktruePositive   AlertFeedback = "TruePositive"
	AlertFeedbackunknown        AlertFeedback = "Unknown"
)

func PossibleValuesForAlertFeedback() []string {
	return []string{
		string(AlertFeedbackbenignPositive),
		string(AlertFeedbackfalsePositive),
		string(AlertFeedbacktruePositive),
		string(AlertFeedbackunknown),
	}
}

func parseAlertFeedback(input string) (*AlertFeedback, error) {
	vals := map[string]AlertFeedback{
		"benignpositive": AlertFeedbackbenignPositive,
		"falsepositive":  AlertFeedbackfalsePositive,
		"truepositive":   AlertFeedbacktruePositive,
		"unknown":        AlertFeedbackunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertFeedback(input)
	return &out, nil
}
