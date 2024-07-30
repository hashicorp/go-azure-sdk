package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationStateState string

const (
	DeviceConfigurationStateState_Compliant     DeviceConfigurationStateState = "compliant"
	DeviceConfigurationStateState_Conflict      DeviceConfigurationStateState = "conflict"
	DeviceConfigurationStateState_Error         DeviceConfigurationStateState = "error"
	DeviceConfigurationStateState_NonCompliant  DeviceConfigurationStateState = "nonCompliant"
	DeviceConfigurationStateState_NotApplicable DeviceConfigurationStateState = "notApplicable"
	DeviceConfigurationStateState_NotAssigned   DeviceConfigurationStateState = "notAssigned"
	DeviceConfigurationStateState_Remediated    DeviceConfigurationStateState = "remediated"
	DeviceConfigurationStateState_Unknown       DeviceConfigurationStateState = "unknown"
)

func PossibleValuesForDeviceConfigurationStateState() []string {
	return []string{
		string(DeviceConfigurationStateState_Compliant),
		string(DeviceConfigurationStateState_Conflict),
		string(DeviceConfigurationStateState_Error),
		string(DeviceConfigurationStateState_NonCompliant),
		string(DeviceConfigurationStateState_NotApplicable),
		string(DeviceConfigurationStateState_NotAssigned),
		string(DeviceConfigurationStateState_Remediated),
		string(DeviceConfigurationStateState_Unknown),
	}
}

func (s *DeviceConfigurationStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationStateState(input string) (*DeviceConfigurationStateState, error) {
	vals := map[string]DeviceConfigurationStateState{
		"compliant":     DeviceConfigurationStateState_Compliant,
		"conflict":      DeviceConfigurationStateState_Conflict,
		"error":         DeviceConfigurationStateState_Error,
		"noncompliant":  DeviceConfigurationStateState_NonCompliant,
		"notapplicable": DeviceConfigurationStateState_NotApplicable,
		"notassigned":   DeviceConfigurationStateState_NotAssigned,
		"remediated":    DeviceConfigurationStateState_Remediated,
		"unknown":       DeviceConfigurationStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationStateState(input)
	return &out, nil
}
