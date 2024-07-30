package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagementState string

const (
	ManagedDeviceManagementState_DeletePending  ManagedDeviceManagementState = "deletePending"
	ManagedDeviceManagementState_Discovered     ManagedDeviceManagementState = "discovered"
	ManagedDeviceManagementState_Managed        ManagedDeviceManagementState = "managed"
	ManagedDeviceManagementState_RetireCanceled ManagedDeviceManagementState = "retireCanceled"
	ManagedDeviceManagementState_RetireFailed   ManagedDeviceManagementState = "retireFailed"
	ManagedDeviceManagementState_RetireIssued   ManagedDeviceManagementState = "retireIssued"
	ManagedDeviceManagementState_RetirePending  ManagedDeviceManagementState = "retirePending"
	ManagedDeviceManagementState_Unhealthy      ManagedDeviceManagementState = "unhealthy"
	ManagedDeviceManagementState_WipeCanceled   ManagedDeviceManagementState = "wipeCanceled"
	ManagedDeviceManagementState_WipeFailed     ManagedDeviceManagementState = "wipeFailed"
	ManagedDeviceManagementState_WipeIssued     ManagedDeviceManagementState = "wipeIssued"
	ManagedDeviceManagementState_WipePending    ManagedDeviceManagementState = "wipePending"
)

func PossibleValuesForManagedDeviceManagementState() []string {
	return []string{
		string(ManagedDeviceManagementState_DeletePending),
		string(ManagedDeviceManagementState_Discovered),
		string(ManagedDeviceManagementState_Managed),
		string(ManagedDeviceManagementState_RetireCanceled),
		string(ManagedDeviceManagementState_RetireFailed),
		string(ManagedDeviceManagementState_RetireIssued),
		string(ManagedDeviceManagementState_RetirePending),
		string(ManagedDeviceManagementState_Unhealthy),
		string(ManagedDeviceManagementState_WipeCanceled),
		string(ManagedDeviceManagementState_WipeFailed),
		string(ManagedDeviceManagementState_WipeIssued),
		string(ManagedDeviceManagementState_WipePending),
	}
}

func (s *ManagedDeviceManagementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceManagementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceManagementState(input string) (*ManagedDeviceManagementState, error) {
	vals := map[string]ManagedDeviceManagementState{
		"deletepending":  ManagedDeviceManagementState_DeletePending,
		"discovered":     ManagedDeviceManagementState_Discovered,
		"managed":        ManagedDeviceManagementState_Managed,
		"retirecanceled": ManagedDeviceManagementState_RetireCanceled,
		"retirefailed":   ManagedDeviceManagementState_RetireFailed,
		"retireissued":   ManagedDeviceManagementState_RetireIssued,
		"retirepending":  ManagedDeviceManagementState_RetirePending,
		"unhealthy":      ManagedDeviceManagementState_Unhealthy,
		"wipecanceled":   ManagedDeviceManagementState_WipeCanceled,
		"wipefailed":     ManagedDeviceManagementState_WipeFailed,
		"wipeissued":     ManagedDeviceManagementState_WipeIssued,
		"wipepending":    ManagedDeviceManagementState_WipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementState(input)
	return &out, nil
}
