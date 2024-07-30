package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceManagementState string

const (
	WindowsManagedDeviceManagementState_DeletePending  WindowsManagedDeviceManagementState = "deletePending"
	WindowsManagedDeviceManagementState_Discovered     WindowsManagedDeviceManagementState = "discovered"
	WindowsManagedDeviceManagementState_Managed        WindowsManagedDeviceManagementState = "managed"
	WindowsManagedDeviceManagementState_RetireCanceled WindowsManagedDeviceManagementState = "retireCanceled"
	WindowsManagedDeviceManagementState_RetireFailed   WindowsManagedDeviceManagementState = "retireFailed"
	WindowsManagedDeviceManagementState_RetireIssued   WindowsManagedDeviceManagementState = "retireIssued"
	WindowsManagedDeviceManagementState_RetirePending  WindowsManagedDeviceManagementState = "retirePending"
	WindowsManagedDeviceManagementState_Unhealthy      WindowsManagedDeviceManagementState = "unhealthy"
	WindowsManagedDeviceManagementState_WipeCanceled   WindowsManagedDeviceManagementState = "wipeCanceled"
	WindowsManagedDeviceManagementState_WipeFailed     WindowsManagedDeviceManagementState = "wipeFailed"
	WindowsManagedDeviceManagementState_WipeIssued     WindowsManagedDeviceManagementState = "wipeIssued"
	WindowsManagedDeviceManagementState_WipePending    WindowsManagedDeviceManagementState = "wipePending"
)

func PossibleValuesForWindowsManagedDeviceManagementState() []string {
	return []string{
		string(WindowsManagedDeviceManagementState_DeletePending),
		string(WindowsManagedDeviceManagementState_Discovered),
		string(WindowsManagedDeviceManagementState_Managed),
		string(WindowsManagedDeviceManagementState_RetireCanceled),
		string(WindowsManagedDeviceManagementState_RetireFailed),
		string(WindowsManagedDeviceManagementState_RetireIssued),
		string(WindowsManagedDeviceManagementState_RetirePending),
		string(WindowsManagedDeviceManagementState_Unhealthy),
		string(WindowsManagedDeviceManagementState_WipeCanceled),
		string(WindowsManagedDeviceManagementState_WipeFailed),
		string(WindowsManagedDeviceManagementState_WipeIssued),
		string(WindowsManagedDeviceManagementState_WipePending),
	}
}

func (s *WindowsManagedDeviceManagementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceManagementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceManagementState(input string) (*WindowsManagedDeviceManagementState, error) {
	vals := map[string]WindowsManagedDeviceManagementState{
		"deletepending":  WindowsManagedDeviceManagementState_DeletePending,
		"discovered":     WindowsManagedDeviceManagementState_Discovered,
		"managed":        WindowsManagedDeviceManagementState_Managed,
		"retirecanceled": WindowsManagedDeviceManagementState_RetireCanceled,
		"retirefailed":   WindowsManagedDeviceManagementState_RetireFailed,
		"retireissued":   WindowsManagedDeviceManagementState_RetireIssued,
		"retirepending":  WindowsManagedDeviceManagementState_RetirePending,
		"unhealthy":      WindowsManagedDeviceManagementState_Unhealthy,
		"wipecanceled":   WindowsManagedDeviceManagementState_WipeCanceled,
		"wipefailed":     WindowsManagedDeviceManagementState_WipeFailed,
		"wipeissued":     WindowsManagedDeviceManagementState_WipeIssued,
		"wipepending":    WindowsManagedDeviceManagementState_WipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceManagementState(input)
	return &out, nil
}
