package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatusProcessingState string

const (
	PrinterStatusProcessingState_Idle       PrinterStatusProcessingState = "idle"
	PrinterStatusProcessingState_Processing PrinterStatusProcessingState = "processing"
	PrinterStatusProcessingState_Stopped    PrinterStatusProcessingState = "stopped"
	PrinterStatusProcessingState_Unknown    PrinterStatusProcessingState = "unknown"
)

func PossibleValuesForPrinterStatusProcessingState() []string {
	return []string{
		string(PrinterStatusProcessingState_Idle),
		string(PrinterStatusProcessingState_Processing),
		string(PrinterStatusProcessingState_Stopped),
		string(PrinterStatusProcessingState_Unknown),
	}
}

func (s *PrinterStatusProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterStatusProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterStatusProcessingState(input string) (*PrinterStatusProcessingState, error) {
	vals := map[string]PrinterStatusProcessingState{
		"idle":       PrinterStatusProcessingState_Idle,
		"processing": PrinterStatusProcessingState_Processing,
		"stopped":    PrinterStatusProcessingState_Stopped,
		"unknown":    PrinterStatusProcessingState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterStatusProcessingState(input)
	return &out, nil
}
