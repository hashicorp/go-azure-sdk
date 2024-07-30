package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentDeviceStateState string

const (
	DeviceManagementIntentDeviceStateState_Compliant     DeviceManagementIntentDeviceStateState = "compliant"
	DeviceManagementIntentDeviceStateState_Conflict      DeviceManagementIntentDeviceStateState = "conflict"
	DeviceManagementIntentDeviceStateState_Error         DeviceManagementIntentDeviceStateState = "error"
	DeviceManagementIntentDeviceStateState_NonCompliant  DeviceManagementIntentDeviceStateState = "nonCompliant"
	DeviceManagementIntentDeviceStateState_NotApplicable DeviceManagementIntentDeviceStateState = "notApplicable"
	DeviceManagementIntentDeviceStateState_NotAssigned   DeviceManagementIntentDeviceStateState = "notAssigned"
	DeviceManagementIntentDeviceStateState_Remediated    DeviceManagementIntentDeviceStateState = "remediated"
	DeviceManagementIntentDeviceStateState_Unknown       DeviceManagementIntentDeviceStateState = "unknown"
)

func PossibleValuesForDeviceManagementIntentDeviceStateState() []string {
	return []string{
		string(DeviceManagementIntentDeviceStateState_Compliant),
		string(DeviceManagementIntentDeviceStateState_Conflict),
		string(DeviceManagementIntentDeviceStateState_Error),
		string(DeviceManagementIntentDeviceStateState_NonCompliant),
		string(DeviceManagementIntentDeviceStateState_NotApplicable),
		string(DeviceManagementIntentDeviceStateState_NotAssigned),
		string(DeviceManagementIntentDeviceStateState_Remediated),
		string(DeviceManagementIntentDeviceStateState_Unknown),
	}
}

func (s *DeviceManagementIntentDeviceStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementIntentDeviceStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementIntentDeviceStateState(input string) (*DeviceManagementIntentDeviceStateState, error) {
	vals := map[string]DeviceManagementIntentDeviceStateState{
		"compliant":     DeviceManagementIntentDeviceStateState_Compliant,
		"conflict":      DeviceManagementIntentDeviceStateState_Conflict,
		"error":         DeviceManagementIntentDeviceStateState_Error,
		"noncompliant":  DeviceManagementIntentDeviceStateState_NonCompliant,
		"notapplicable": DeviceManagementIntentDeviceStateState_NotApplicable,
		"notassigned":   DeviceManagementIntentDeviceStateState_NotAssigned,
		"remediated":    DeviceManagementIntentDeviceStateState_Remediated,
		"unknown":       DeviceManagementIntentDeviceStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementIntentDeviceStateState(input)
	return &out, nil
}
