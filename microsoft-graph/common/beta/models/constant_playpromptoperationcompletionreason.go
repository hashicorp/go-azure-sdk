package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlayPromptOperationCompletionReason string

const (
	PlayPromptOperationCompletionReasoncompletedSuccessfully  PlayPromptOperationCompletionReason = "CompletedSuccessfully"
	PlayPromptOperationCompletionReasonmediaOperationCanceled PlayPromptOperationCompletionReason = "MediaOperationCanceled"
	PlayPromptOperationCompletionReasonunknown                PlayPromptOperationCompletionReason = "Unknown"
)

func PossibleValuesForPlayPromptOperationCompletionReason() []string {
	return []string{
		string(PlayPromptOperationCompletionReasoncompletedSuccessfully),
		string(PlayPromptOperationCompletionReasonmediaOperationCanceled),
		string(PlayPromptOperationCompletionReasonunknown),
	}
}

func parsePlayPromptOperationCompletionReason(input string) (*PlayPromptOperationCompletionReason, error) {
	vals := map[string]PlayPromptOperationCompletionReason{
		"completedsuccessfully":  PlayPromptOperationCompletionReasoncompletedSuccessfully,
		"mediaoperationcanceled": PlayPromptOperationCompletionReasonmediaOperationCanceled,
		"unknown":                PlayPromptOperationCompletionReasonunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlayPromptOperationCompletionReason(input)
	return &out, nil
}
