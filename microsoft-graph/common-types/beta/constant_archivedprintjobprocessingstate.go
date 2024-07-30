package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArchivedPrintJobProcessingState string

const (
	ArchivedPrintJobProcessingState_Aborted    ArchivedPrintJobProcessingState = "aborted"
	ArchivedPrintJobProcessingState_Canceled   ArchivedPrintJobProcessingState = "canceled"
	ArchivedPrintJobProcessingState_Completed  ArchivedPrintJobProcessingState = "completed"
	ArchivedPrintJobProcessingState_Paused     ArchivedPrintJobProcessingState = "paused"
	ArchivedPrintJobProcessingState_Pending    ArchivedPrintJobProcessingState = "pending"
	ArchivedPrintJobProcessingState_Processing ArchivedPrintJobProcessingState = "processing"
	ArchivedPrintJobProcessingState_Stopped    ArchivedPrintJobProcessingState = "stopped"
	ArchivedPrintJobProcessingState_Unknown    ArchivedPrintJobProcessingState = "unknown"
)

func PossibleValuesForArchivedPrintJobProcessingState() []string {
	return []string{
		string(ArchivedPrintJobProcessingState_Aborted),
		string(ArchivedPrintJobProcessingState_Canceled),
		string(ArchivedPrintJobProcessingState_Completed),
		string(ArchivedPrintJobProcessingState_Paused),
		string(ArchivedPrintJobProcessingState_Pending),
		string(ArchivedPrintJobProcessingState_Processing),
		string(ArchivedPrintJobProcessingState_Stopped),
		string(ArchivedPrintJobProcessingState_Unknown),
	}
}

func (s *ArchivedPrintJobProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseArchivedPrintJobProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseArchivedPrintJobProcessingState(input string) (*ArchivedPrintJobProcessingState, error) {
	vals := map[string]ArchivedPrintJobProcessingState{
		"aborted":    ArchivedPrintJobProcessingState_Aborted,
		"canceled":   ArchivedPrintJobProcessingState_Canceled,
		"completed":  ArchivedPrintJobProcessingState_Completed,
		"paused":     ArchivedPrintJobProcessingState_Paused,
		"pending":    ArchivedPrintJobProcessingState_Pending,
		"processing": ArchivedPrintJobProcessingState_Processing,
		"stopped":    ArchivedPrintJobProcessingState_Stopped,
		"unknown":    ArchivedPrintJobProcessingState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArchivedPrintJobProcessingState(input)
	return &out, nil
}
