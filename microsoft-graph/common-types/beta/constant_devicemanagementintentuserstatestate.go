package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentUserStateState string

const (
	DeviceManagementIntentUserStateState_Compliant     DeviceManagementIntentUserStateState = "compliant"
	DeviceManagementIntentUserStateState_Conflict      DeviceManagementIntentUserStateState = "conflict"
	DeviceManagementIntentUserStateState_Error         DeviceManagementIntentUserStateState = "error"
	DeviceManagementIntentUserStateState_NonCompliant  DeviceManagementIntentUserStateState = "nonCompliant"
	DeviceManagementIntentUserStateState_NotApplicable DeviceManagementIntentUserStateState = "notApplicable"
	DeviceManagementIntentUserStateState_NotAssigned   DeviceManagementIntentUserStateState = "notAssigned"
	DeviceManagementIntentUserStateState_Remediated    DeviceManagementIntentUserStateState = "remediated"
	DeviceManagementIntentUserStateState_Unknown       DeviceManagementIntentUserStateState = "unknown"
)

func PossibleValuesForDeviceManagementIntentUserStateState() []string {
	return []string{
		string(DeviceManagementIntentUserStateState_Compliant),
		string(DeviceManagementIntentUserStateState_Conflict),
		string(DeviceManagementIntentUserStateState_Error),
		string(DeviceManagementIntentUserStateState_NonCompliant),
		string(DeviceManagementIntentUserStateState_NotApplicable),
		string(DeviceManagementIntentUserStateState_NotAssigned),
		string(DeviceManagementIntentUserStateState_Remediated),
		string(DeviceManagementIntentUserStateState_Unknown),
	}
}

func (s *DeviceManagementIntentUserStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementIntentUserStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementIntentUserStateState(input string) (*DeviceManagementIntentUserStateState, error) {
	vals := map[string]DeviceManagementIntentUserStateState{
		"compliant":     DeviceManagementIntentUserStateState_Compliant,
		"conflict":      DeviceManagementIntentUserStateState_Conflict,
		"error":         DeviceManagementIntentUserStateState_Error,
		"noncompliant":  DeviceManagementIntentUserStateState_NonCompliant,
		"notapplicable": DeviceManagementIntentUserStateState_NotApplicable,
		"notassigned":   DeviceManagementIntentUserStateState_NotAssigned,
		"remediated":    DeviceManagementIntentUserStateState_Remediated,
		"unknown":       DeviceManagementIntentUserStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementIntentUserStateState(input)
	return &out, nil
}
