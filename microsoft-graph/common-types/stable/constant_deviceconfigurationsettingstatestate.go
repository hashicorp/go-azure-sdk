package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationSettingStateState string

const (
	DeviceConfigurationSettingStateState_Compliant     DeviceConfigurationSettingStateState = "compliant"
	DeviceConfigurationSettingStateState_Conflict      DeviceConfigurationSettingStateState = "conflict"
	DeviceConfigurationSettingStateState_Error         DeviceConfigurationSettingStateState = "error"
	DeviceConfigurationSettingStateState_NonCompliant  DeviceConfigurationSettingStateState = "nonCompliant"
	DeviceConfigurationSettingStateState_NotApplicable DeviceConfigurationSettingStateState = "notApplicable"
	DeviceConfigurationSettingStateState_NotAssigned   DeviceConfigurationSettingStateState = "notAssigned"
	DeviceConfigurationSettingStateState_Remediated    DeviceConfigurationSettingStateState = "remediated"
	DeviceConfigurationSettingStateState_Unknown       DeviceConfigurationSettingStateState = "unknown"
)

func PossibleValuesForDeviceConfigurationSettingStateState() []string {
	return []string{
		string(DeviceConfigurationSettingStateState_Compliant),
		string(DeviceConfigurationSettingStateState_Conflict),
		string(DeviceConfigurationSettingStateState_Error),
		string(DeviceConfigurationSettingStateState_NonCompliant),
		string(DeviceConfigurationSettingStateState_NotApplicable),
		string(DeviceConfigurationSettingStateState_NotAssigned),
		string(DeviceConfigurationSettingStateState_Remediated),
		string(DeviceConfigurationSettingStateState_Unknown),
	}
}

func (s *DeviceConfigurationSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationSettingStateState(input string) (*DeviceConfigurationSettingStateState, error) {
	vals := map[string]DeviceConfigurationSettingStateState{
		"compliant":     DeviceConfigurationSettingStateState_Compliant,
		"conflict":      DeviceConfigurationSettingStateState_Conflict,
		"error":         DeviceConfigurationSettingStateState_Error,
		"noncompliant":  DeviceConfigurationSettingStateState_NonCompliant,
		"notapplicable": DeviceConfigurationSettingStateState_NotApplicable,
		"notassigned":   DeviceConfigurationSettingStateState_NotAssigned,
		"remediated":    DeviceConfigurationSettingStateState_Remediated,
		"unknown":       DeviceConfigurationSettingStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationSettingStateState(input)
	return &out, nil
}
