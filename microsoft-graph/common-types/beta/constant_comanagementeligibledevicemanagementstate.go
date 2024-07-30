package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceManagementState string

const (
	ComanagementEligibleDeviceManagementState_DeletePending  ComanagementEligibleDeviceManagementState = "deletePending"
	ComanagementEligibleDeviceManagementState_Discovered     ComanagementEligibleDeviceManagementState = "discovered"
	ComanagementEligibleDeviceManagementState_Managed        ComanagementEligibleDeviceManagementState = "managed"
	ComanagementEligibleDeviceManagementState_RetireCanceled ComanagementEligibleDeviceManagementState = "retireCanceled"
	ComanagementEligibleDeviceManagementState_RetireFailed   ComanagementEligibleDeviceManagementState = "retireFailed"
	ComanagementEligibleDeviceManagementState_RetireIssued   ComanagementEligibleDeviceManagementState = "retireIssued"
	ComanagementEligibleDeviceManagementState_RetirePending  ComanagementEligibleDeviceManagementState = "retirePending"
	ComanagementEligibleDeviceManagementState_Unhealthy      ComanagementEligibleDeviceManagementState = "unhealthy"
	ComanagementEligibleDeviceManagementState_WipeCanceled   ComanagementEligibleDeviceManagementState = "wipeCanceled"
	ComanagementEligibleDeviceManagementState_WipeFailed     ComanagementEligibleDeviceManagementState = "wipeFailed"
	ComanagementEligibleDeviceManagementState_WipeIssued     ComanagementEligibleDeviceManagementState = "wipeIssued"
	ComanagementEligibleDeviceManagementState_WipePending    ComanagementEligibleDeviceManagementState = "wipePending"
)

func PossibleValuesForComanagementEligibleDeviceManagementState() []string {
	return []string{
		string(ComanagementEligibleDeviceManagementState_DeletePending),
		string(ComanagementEligibleDeviceManagementState_Discovered),
		string(ComanagementEligibleDeviceManagementState_Managed),
		string(ComanagementEligibleDeviceManagementState_RetireCanceled),
		string(ComanagementEligibleDeviceManagementState_RetireFailed),
		string(ComanagementEligibleDeviceManagementState_RetireIssued),
		string(ComanagementEligibleDeviceManagementState_RetirePending),
		string(ComanagementEligibleDeviceManagementState_Unhealthy),
		string(ComanagementEligibleDeviceManagementState_WipeCanceled),
		string(ComanagementEligibleDeviceManagementState_WipeFailed),
		string(ComanagementEligibleDeviceManagementState_WipeIssued),
		string(ComanagementEligibleDeviceManagementState_WipePending),
	}
}

func (s *ComanagementEligibleDeviceManagementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComanagementEligibleDeviceManagementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComanagementEligibleDeviceManagementState(input string) (*ComanagementEligibleDeviceManagementState, error) {
	vals := map[string]ComanagementEligibleDeviceManagementState{
		"deletepending":  ComanagementEligibleDeviceManagementState_DeletePending,
		"discovered":     ComanagementEligibleDeviceManagementState_Discovered,
		"managed":        ComanagementEligibleDeviceManagementState_Managed,
		"retirecanceled": ComanagementEligibleDeviceManagementState_RetireCanceled,
		"retirefailed":   ComanagementEligibleDeviceManagementState_RetireFailed,
		"retireissued":   ComanagementEligibleDeviceManagementState_RetireIssued,
		"retirepending":  ComanagementEligibleDeviceManagementState_RetirePending,
		"unhealthy":      ComanagementEligibleDeviceManagementState_Unhealthy,
		"wipecanceled":   ComanagementEligibleDeviceManagementState_WipeCanceled,
		"wipefailed":     ComanagementEligibleDeviceManagementState_WipeFailed,
		"wipeissued":     ComanagementEligibleDeviceManagementState_WipeIssued,
		"wipepending":    ComanagementEligibleDeviceManagementState_WipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceManagementState(input)
	return &out, nil
}
