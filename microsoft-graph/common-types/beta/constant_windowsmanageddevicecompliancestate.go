package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceComplianceState string

const (
	WindowsManagedDeviceComplianceState_Compliant     WindowsManagedDeviceComplianceState = "compliant"
	WindowsManagedDeviceComplianceState_ConfigManager WindowsManagedDeviceComplianceState = "configManager"
	WindowsManagedDeviceComplianceState_Conflict      WindowsManagedDeviceComplianceState = "conflict"
	WindowsManagedDeviceComplianceState_Error         WindowsManagedDeviceComplianceState = "error"
	WindowsManagedDeviceComplianceState_InGracePeriod WindowsManagedDeviceComplianceState = "inGracePeriod"
	WindowsManagedDeviceComplianceState_Noncompliant  WindowsManagedDeviceComplianceState = "noncompliant"
	WindowsManagedDeviceComplianceState_Unknown       WindowsManagedDeviceComplianceState = "unknown"
)

func PossibleValuesForWindowsManagedDeviceComplianceState() []string {
	return []string{
		string(WindowsManagedDeviceComplianceState_Compliant),
		string(WindowsManagedDeviceComplianceState_ConfigManager),
		string(WindowsManagedDeviceComplianceState_Conflict),
		string(WindowsManagedDeviceComplianceState_Error),
		string(WindowsManagedDeviceComplianceState_InGracePeriod),
		string(WindowsManagedDeviceComplianceState_Noncompliant),
		string(WindowsManagedDeviceComplianceState_Unknown),
	}
}

func (s *WindowsManagedDeviceComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceComplianceState(input string) (*WindowsManagedDeviceComplianceState, error) {
	vals := map[string]WindowsManagedDeviceComplianceState{
		"compliant":     WindowsManagedDeviceComplianceState_Compliant,
		"configmanager": WindowsManagedDeviceComplianceState_ConfigManager,
		"conflict":      WindowsManagedDeviceComplianceState_Conflict,
		"error":         WindowsManagedDeviceComplianceState_Error,
		"ingraceperiod": WindowsManagedDeviceComplianceState_InGracePeriod,
		"noncompliant":  WindowsManagedDeviceComplianceState_Noncompliant,
		"unknown":       WindowsManagedDeviceComplianceState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceComplianceState(input)
	return &out, nil
}
