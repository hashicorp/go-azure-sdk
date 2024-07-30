package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceComplianceState string

const (
	ManagedDeviceComplianceState_Compliant     ManagedDeviceComplianceState = "compliant"
	ManagedDeviceComplianceState_ConfigManager ManagedDeviceComplianceState = "configManager"
	ManagedDeviceComplianceState_Conflict      ManagedDeviceComplianceState = "conflict"
	ManagedDeviceComplianceState_Error         ManagedDeviceComplianceState = "error"
	ManagedDeviceComplianceState_InGracePeriod ManagedDeviceComplianceState = "inGracePeriod"
	ManagedDeviceComplianceState_Noncompliant  ManagedDeviceComplianceState = "noncompliant"
	ManagedDeviceComplianceState_Unknown       ManagedDeviceComplianceState = "unknown"
)

func PossibleValuesForManagedDeviceComplianceState() []string {
	return []string{
		string(ManagedDeviceComplianceState_Compliant),
		string(ManagedDeviceComplianceState_ConfigManager),
		string(ManagedDeviceComplianceState_Conflict),
		string(ManagedDeviceComplianceState_Error),
		string(ManagedDeviceComplianceState_InGracePeriod),
		string(ManagedDeviceComplianceState_Noncompliant),
		string(ManagedDeviceComplianceState_Unknown),
	}
}

func (s *ManagedDeviceComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceComplianceState(input string) (*ManagedDeviceComplianceState, error) {
	vals := map[string]ManagedDeviceComplianceState{
		"compliant":     ManagedDeviceComplianceState_Compliant,
		"configmanager": ManagedDeviceComplianceState_ConfigManager,
		"conflict":      ManagedDeviceComplianceState_Conflict,
		"error":         ManagedDeviceComplianceState_Error,
		"ingraceperiod": ManagedDeviceComplianceState_InGracePeriod,
		"noncompliant":  ManagedDeviceComplianceState_Noncompliant,
		"unknown":       ManagedDeviceComplianceState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceComplianceState(input)
	return &out, nil
}
