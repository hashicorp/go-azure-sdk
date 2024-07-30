package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderScanActionResultActionState string

const (
	WindowsDefenderScanActionResultActionState_Active       WindowsDefenderScanActionResultActionState = "active"
	WindowsDefenderScanActionResultActionState_Canceled     WindowsDefenderScanActionResultActionState = "canceled"
	WindowsDefenderScanActionResultActionState_Done         WindowsDefenderScanActionResultActionState = "done"
	WindowsDefenderScanActionResultActionState_Failed       WindowsDefenderScanActionResultActionState = "failed"
	WindowsDefenderScanActionResultActionState_None         WindowsDefenderScanActionResultActionState = "none"
	WindowsDefenderScanActionResultActionState_NotSupported WindowsDefenderScanActionResultActionState = "notSupported"
	WindowsDefenderScanActionResultActionState_Pending      WindowsDefenderScanActionResultActionState = "pending"
)

func PossibleValuesForWindowsDefenderScanActionResultActionState() []string {
	return []string{
		string(WindowsDefenderScanActionResultActionState_Active),
		string(WindowsDefenderScanActionResultActionState_Canceled),
		string(WindowsDefenderScanActionResultActionState_Done),
		string(WindowsDefenderScanActionResultActionState_Failed),
		string(WindowsDefenderScanActionResultActionState_None),
		string(WindowsDefenderScanActionResultActionState_NotSupported),
		string(WindowsDefenderScanActionResultActionState_Pending),
	}
}

func (s *WindowsDefenderScanActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDefenderScanActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDefenderScanActionResultActionState(input string) (*WindowsDefenderScanActionResultActionState, error) {
	vals := map[string]WindowsDefenderScanActionResultActionState{
		"active":       WindowsDefenderScanActionResultActionState_Active,
		"canceled":     WindowsDefenderScanActionResultActionState_Canceled,
		"done":         WindowsDefenderScanActionResultActionState_Done,
		"failed":       WindowsDefenderScanActionResultActionState_Failed,
		"none":         WindowsDefenderScanActionResultActionState_None,
		"notsupported": WindowsDefenderScanActionResultActionState_NotSupported,
		"pending":      WindowsDefenderScanActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDefenderScanActionResultActionState(input)
	return &out, nil
}
