package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintOperationStatusState string

const (
	PrintOperationStatusState_Failed     PrintOperationStatusState = "failed"
	PrintOperationStatusState_NotStarted PrintOperationStatusState = "notStarted"
	PrintOperationStatusState_Running    PrintOperationStatusState = "running"
	PrintOperationStatusState_Succeeded  PrintOperationStatusState = "succeeded"
)

func PossibleValuesForPrintOperationStatusState() []string {
	return []string{
		string(PrintOperationStatusState_Failed),
		string(PrintOperationStatusState_NotStarted),
		string(PrintOperationStatusState_Running),
		string(PrintOperationStatusState_Succeeded),
	}
}

func (s *PrintOperationStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintOperationStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintOperationStatusState(input string) (*PrintOperationStatusState, error) {
	vals := map[string]PrintOperationStatusState{
		"failed":     PrintOperationStatusState_Failed,
		"notstarted": PrintOperationStatusState_NotStarted,
		"running":    PrintOperationStatusState_Running,
		"succeeded":  PrintOperationStatusState_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintOperationStatusState(input)
	return &out, nil
}
