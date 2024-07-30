package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordOperationCompletionReason string

const (
	RecordOperationCompletionReason_InitialSilenceTimeout    RecordOperationCompletionReason = "initialSilenceTimeout"
	RecordOperationCompletionReason_MaxRecordDurationReached RecordOperationCompletionReason = "maxRecordDurationReached"
	RecordOperationCompletionReason_MaxSilenceTimeout        RecordOperationCompletionReason = "maxSilenceTimeout"
	RecordOperationCompletionReason_MediaReceiveTimeout      RecordOperationCompletionReason = "mediaReceiveTimeout"
	RecordOperationCompletionReason_OperationCanceled        RecordOperationCompletionReason = "operationCanceled"
	RecordOperationCompletionReason_PlayBeepFailed           RecordOperationCompletionReason = "playBeepFailed"
	RecordOperationCompletionReason_PlayPromptFailed         RecordOperationCompletionReason = "playPromptFailed"
	RecordOperationCompletionReason_StopToneDetected         RecordOperationCompletionReason = "stopToneDetected"
	RecordOperationCompletionReason_UnspecifiedError         RecordOperationCompletionReason = "unspecifiedError"
)

func PossibleValuesForRecordOperationCompletionReason() []string {
	return []string{
		string(RecordOperationCompletionReason_InitialSilenceTimeout),
		string(RecordOperationCompletionReason_MaxRecordDurationReached),
		string(RecordOperationCompletionReason_MaxSilenceTimeout),
		string(RecordOperationCompletionReason_MediaReceiveTimeout),
		string(RecordOperationCompletionReason_OperationCanceled),
		string(RecordOperationCompletionReason_PlayBeepFailed),
		string(RecordOperationCompletionReason_PlayPromptFailed),
		string(RecordOperationCompletionReason_StopToneDetected),
		string(RecordOperationCompletionReason_UnspecifiedError),
	}
}

func (s *RecordOperationCompletionReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecordOperationCompletionReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecordOperationCompletionReason(input string) (*RecordOperationCompletionReason, error) {
	vals := map[string]RecordOperationCompletionReason{
		"initialsilencetimeout":    RecordOperationCompletionReason_InitialSilenceTimeout,
		"maxrecorddurationreached": RecordOperationCompletionReason_MaxRecordDurationReached,
		"maxsilencetimeout":        RecordOperationCompletionReason_MaxSilenceTimeout,
		"mediareceivetimeout":      RecordOperationCompletionReason_MediaReceiveTimeout,
		"operationcanceled":        RecordOperationCompletionReason_OperationCanceled,
		"playbeepfailed":           RecordOperationCompletionReason_PlayBeepFailed,
		"playpromptfailed":         RecordOperationCompletionReason_PlayPromptFailed,
		"stoptonedetected":         RecordOperationCompletionReason_StopToneDetected,
		"unspecifiederror":         RecordOperationCompletionReason_UnspecifiedError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordOperationCompletionReason(input)
	return &out, nil
}
