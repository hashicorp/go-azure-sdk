package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskStatusState string

const (
	PrintTaskStatusState_Aborted    PrintTaskStatusState = "aborted"
	PrintTaskStatusState_Completed  PrintTaskStatusState = "completed"
	PrintTaskStatusState_Pending    PrintTaskStatusState = "pending"
	PrintTaskStatusState_Processing PrintTaskStatusState = "processing"
)

func PossibleValuesForPrintTaskStatusState() []string {
	return []string{
		string(PrintTaskStatusState_Aborted),
		string(PrintTaskStatusState_Completed),
		string(PrintTaskStatusState_Pending),
		string(PrintTaskStatusState_Processing),
	}
}

func (s *PrintTaskStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintTaskStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintTaskStatusState(input string) (*PrintTaskStatusState, error) {
	vals := map[string]PrintTaskStatusState{
		"aborted":    PrintTaskStatusState_Aborted,
		"completed":  PrintTaskStatusState_Completed,
		"pending":    PrintTaskStatusState_Pending,
		"processing": PrintTaskStatusState_Processing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintTaskStatusState(input)
	return &out, nil
}
