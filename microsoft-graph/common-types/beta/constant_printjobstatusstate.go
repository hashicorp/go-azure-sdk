package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatusState string

const (
	PrintJobStatusState_Aborted    PrintJobStatusState = "aborted"
	PrintJobStatusState_Canceled   PrintJobStatusState = "canceled"
	PrintJobStatusState_Completed  PrintJobStatusState = "completed"
	PrintJobStatusState_Paused     PrintJobStatusState = "paused"
	PrintJobStatusState_Pending    PrintJobStatusState = "pending"
	PrintJobStatusState_Processing PrintJobStatusState = "processing"
	PrintJobStatusState_Stopped    PrintJobStatusState = "stopped"
	PrintJobStatusState_Unknown    PrintJobStatusState = "unknown"
)

func PossibleValuesForPrintJobStatusState() []string {
	return []string{
		string(PrintJobStatusState_Aborted),
		string(PrintJobStatusState_Canceled),
		string(PrintJobStatusState_Completed),
		string(PrintJobStatusState_Paused),
		string(PrintJobStatusState_Pending),
		string(PrintJobStatusState_Processing),
		string(PrintJobStatusState_Stopped),
		string(PrintJobStatusState_Unknown),
	}
}

func (s *PrintJobStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobStatusState(input string) (*PrintJobStatusState, error) {
	vals := map[string]PrintJobStatusState{
		"aborted":    PrintJobStatusState_Aborted,
		"canceled":   PrintJobStatusState_Canceled,
		"completed":  PrintJobStatusState_Completed,
		"paused":     PrintJobStatusState_Paused,
		"pending":    PrintJobStatusState_Pending,
		"processing": PrintJobStatusState_Processing,
		"stopped":    PrintJobStatusState_Stopped,
		"unknown":    PrintJobStatusState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobStatusState(input)
	return &out, nil
}
