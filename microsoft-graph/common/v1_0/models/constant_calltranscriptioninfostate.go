package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallTranscriptionInfoState string

const (
	CallTranscriptionInfoStateactive     CallTranscriptionInfoState = "Active"
	CallTranscriptionInfoStateinactive   CallTranscriptionInfoState = "Inactive"
	CallTranscriptionInfoStatenotStarted CallTranscriptionInfoState = "NotStarted"
)

func PossibleValuesForCallTranscriptionInfoState() []string {
	return []string{
		string(CallTranscriptionInfoStateactive),
		string(CallTranscriptionInfoStateinactive),
		string(CallTranscriptionInfoStatenotStarted),
	}
}

func parseCallTranscriptionInfoState(input string) (*CallTranscriptionInfoState, error) {
	vals := map[string]CallTranscriptionInfoState{
		"active":     CallTranscriptionInfoStateactive,
		"inactive":   CallTranscriptionInfoStateinactive,
		"notstarted": CallTranscriptionInfoStatenotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallTranscriptionInfoState(input)
	return &out, nil
}
