package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventPropagationResultStatus string

const (
	SecurityEventPropagationResultStatusfailed       SecurityEventPropagationResultStatus = "Failed"
	SecurityEventPropagationResultStatusinProcessing SecurityEventPropagationResultStatus = "InProcessing"
	SecurityEventPropagationResultStatusnone         SecurityEventPropagationResultStatus = "None"
	SecurityEventPropagationResultStatussuccess      SecurityEventPropagationResultStatus = "Success"
)

func PossibleValuesForSecurityEventPropagationResultStatus() []string {
	return []string{
		string(SecurityEventPropagationResultStatusfailed),
		string(SecurityEventPropagationResultStatusinProcessing),
		string(SecurityEventPropagationResultStatusnone),
		string(SecurityEventPropagationResultStatussuccess),
	}
}

func parseSecurityEventPropagationResultStatus(input string) (*SecurityEventPropagationResultStatus, error) {
	vals := map[string]SecurityEventPropagationResultStatus{
		"failed":       SecurityEventPropagationResultStatusfailed,
		"inprocessing": SecurityEventPropagationResultStatusinProcessing,
		"none":         SecurityEventPropagationResultStatusnone,
		"success":      SecurityEventPropagationResultStatussuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventPropagationResultStatus(input)
	return &out, nil
}
