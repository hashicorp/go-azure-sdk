package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordOperationCompletionReason string

const (
	RecordOperationCompletionReasoninitialSilenceTimeout    RecordOperationCompletionReason = "InitialSilenceTimeout"
	RecordOperationCompletionReasonmaxRecordDurationReached RecordOperationCompletionReason = "MaxRecordDurationReached"
	RecordOperationCompletionReasonmaxSilenceTimeout        RecordOperationCompletionReason = "MaxSilenceTimeout"
	RecordOperationCompletionReasonmediaReceiveTimeout      RecordOperationCompletionReason = "MediaReceiveTimeout"
	RecordOperationCompletionReasonoperationCanceled        RecordOperationCompletionReason = "OperationCanceled"
	RecordOperationCompletionReasonplayBeepFailed           RecordOperationCompletionReason = "PlayBeepFailed"
	RecordOperationCompletionReasonplayPromptFailed         RecordOperationCompletionReason = "PlayPromptFailed"
	RecordOperationCompletionReasonstopToneDetected         RecordOperationCompletionReason = "StopToneDetected"
	RecordOperationCompletionReasonunspecifiedError         RecordOperationCompletionReason = "UnspecifiedError"
)

func PossibleValuesForRecordOperationCompletionReason() []string {
	return []string{
		string(RecordOperationCompletionReasoninitialSilenceTimeout),
		string(RecordOperationCompletionReasonmaxRecordDurationReached),
		string(RecordOperationCompletionReasonmaxSilenceTimeout),
		string(RecordOperationCompletionReasonmediaReceiveTimeout),
		string(RecordOperationCompletionReasonoperationCanceled),
		string(RecordOperationCompletionReasonplayBeepFailed),
		string(RecordOperationCompletionReasonplayPromptFailed),
		string(RecordOperationCompletionReasonstopToneDetected),
		string(RecordOperationCompletionReasonunspecifiedError),
	}
}

func parseRecordOperationCompletionReason(input string) (*RecordOperationCompletionReason, error) {
	vals := map[string]RecordOperationCompletionReason{
		"initialsilencetimeout":    RecordOperationCompletionReasoninitialSilenceTimeout,
		"maxrecorddurationreached": RecordOperationCompletionReasonmaxRecordDurationReached,
		"maxsilencetimeout":        RecordOperationCompletionReasonmaxSilenceTimeout,
		"mediareceivetimeout":      RecordOperationCompletionReasonmediaReceiveTimeout,
		"operationcanceled":        RecordOperationCompletionReasonoperationCanceled,
		"playbeepfailed":           RecordOperationCompletionReasonplayBeepFailed,
		"playpromptfailed":         RecordOperationCompletionReasonplayPromptFailed,
		"stoptonedetected":         RecordOperationCompletionReasonstopToneDetected,
		"unspecifiederror":         RecordOperationCompletionReasonunspecifiedError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordOperationCompletionReason(input)
	return &out, nil
}
