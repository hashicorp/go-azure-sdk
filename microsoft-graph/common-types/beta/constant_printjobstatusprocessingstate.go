package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatusProcessingState string

const (
	PrintJobStatusProcessingState_Aborted    PrintJobStatusProcessingState = "aborted"
	PrintJobStatusProcessingState_Canceled   PrintJobStatusProcessingState = "canceled"
	PrintJobStatusProcessingState_Completed  PrintJobStatusProcessingState = "completed"
	PrintJobStatusProcessingState_Paused     PrintJobStatusProcessingState = "paused"
	PrintJobStatusProcessingState_Pending    PrintJobStatusProcessingState = "pending"
	PrintJobStatusProcessingState_Processing PrintJobStatusProcessingState = "processing"
	PrintJobStatusProcessingState_Stopped    PrintJobStatusProcessingState = "stopped"
	PrintJobStatusProcessingState_Unknown    PrintJobStatusProcessingState = "unknown"
)

func PossibleValuesForPrintJobStatusProcessingState() []string {
	return []string{
		string(PrintJobStatusProcessingState_Aborted),
		string(PrintJobStatusProcessingState_Canceled),
		string(PrintJobStatusProcessingState_Completed),
		string(PrintJobStatusProcessingState_Paused),
		string(PrintJobStatusProcessingState_Pending),
		string(PrintJobStatusProcessingState_Processing),
		string(PrintJobStatusProcessingState_Stopped),
		string(PrintJobStatusProcessingState_Unknown),
	}
}

func (s *PrintJobStatusProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobStatusProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobStatusProcessingState(input string) (*PrintJobStatusProcessingState, error) {
	vals := map[string]PrintJobStatusProcessingState{
		"aborted":    PrintJobStatusProcessingState_Aborted,
		"canceled":   PrintJobStatusProcessingState_Canceled,
		"completed":  PrintJobStatusProcessingState_Completed,
		"paused":     PrintJobStatusProcessingState_Paused,
		"pending":    PrintJobStatusProcessingState_Pending,
		"processing": PrintJobStatusProcessingState_Processing,
		"stopped":    PrintJobStatusProcessingState_Stopped,
		"unknown":    PrintJobStatusProcessingState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusProcessingState(input)
	return &out, nil
}
