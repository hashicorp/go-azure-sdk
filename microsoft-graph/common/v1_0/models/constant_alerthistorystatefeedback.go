package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryStateFeedback string

const (
	AlertHistoryStateFeedbackbenignPositive AlertHistoryStateFeedback = "BenignPositive"
	AlertHistoryStateFeedbackfalsePositive  AlertHistoryStateFeedback = "FalsePositive"
	AlertHistoryStateFeedbacktruePositive   AlertHistoryStateFeedback = "TruePositive"
	AlertHistoryStateFeedbackunknown        AlertHistoryStateFeedback = "Unknown"
)

func PossibleValuesForAlertHistoryStateFeedback() []string {
	return []string{
		string(AlertHistoryStateFeedbackbenignPositive),
		string(AlertHistoryStateFeedbackfalsePositive),
		string(AlertHistoryStateFeedbacktruePositive),
		string(AlertHistoryStateFeedbackunknown),
	}
}

func parseAlertHistoryStateFeedback(input string) (*AlertHistoryStateFeedback, error) {
	vals := map[string]AlertHistoryStateFeedback{
		"benignpositive": AlertHistoryStateFeedbackbenignPositive,
		"falsepositive":  AlertHistoryStateFeedbackfalsePositive,
		"truepositive":   AlertHistoryStateFeedbacktruePositive,
		"unknown":        AlertHistoryStateFeedbackunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertHistoryStateFeedback(input)
	return &out, nil
}
