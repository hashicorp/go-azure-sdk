package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityRemediationState string

const (
	WindowsAutopilotDeviceIdentityRemediationState_AutomaticRemediationRequired WindowsAutopilotDeviceIdentityRemediationState = "automaticRemediationRequired"
	WindowsAutopilotDeviceIdentityRemediationState_ManualRemediationRequired    WindowsAutopilotDeviceIdentityRemediationState = "manualRemediationRequired"
	WindowsAutopilotDeviceIdentityRemediationState_NoRemediationRequired        WindowsAutopilotDeviceIdentityRemediationState = "noRemediationRequired"
	WindowsAutopilotDeviceIdentityRemediationState_Unknown                      WindowsAutopilotDeviceIdentityRemediationState = "unknown"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityRemediationState() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityRemediationState_AutomaticRemediationRequired),
		string(WindowsAutopilotDeviceIdentityRemediationState_ManualRemediationRequired),
		string(WindowsAutopilotDeviceIdentityRemediationState_NoRemediationRequired),
		string(WindowsAutopilotDeviceIdentityRemediationState_Unknown),
	}
}

func (s *WindowsAutopilotDeviceIdentityRemediationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeviceIdentityRemediationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeviceIdentityRemediationState(input string) (*WindowsAutopilotDeviceIdentityRemediationState, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityRemediationState{
		"automaticremediationrequired": WindowsAutopilotDeviceIdentityRemediationState_AutomaticRemediationRequired,
		"manualremediationrequired":    WindowsAutopilotDeviceIdentityRemediationState_ManualRemediationRequired,
		"noremediationrequired":        WindowsAutopilotDeviceIdentityRemediationState_NoRemediationRequired,
		"unknown":                      WindowsAutopilotDeviceIdentityRemediationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityRemediationState(input)
	return &out, nil
}
