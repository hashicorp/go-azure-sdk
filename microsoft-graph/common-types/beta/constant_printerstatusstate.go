package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusState string

const (
	PrinterStatusState_Idle       PrinterStatusState = "idle"
	PrinterStatusState_Processing PrinterStatusState = "processing"
	PrinterStatusState_Stopped    PrinterStatusState = "stopped"
	PrinterStatusState_Unknown    PrinterStatusState = "unknown"
)

func PossibleValuesForPrinterStatusState() []string {
	return []string{
		string(PrinterStatusState_Idle),
		string(PrinterStatusState_Processing),
		string(PrinterStatusState_Stopped),
		string(PrinterStatusState_Unknown),
	}
}

func (s *PrinterStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusState(input string) (*PrinterStatusState, error) {
	vals := map[string]PrinterStatusState{
		"idle":       PrinterStatusState_Idle,
		"processing": PrinterStatusState_Processing,
		"stopped":    PrinterStatusState_Stopped,
		"unknown":    PrinterStatusState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusState(input)
	return &out, nil
}
