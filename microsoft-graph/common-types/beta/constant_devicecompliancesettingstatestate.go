package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceSettingStateState string

const (
	DeviceComplianceSettingStateState_Compliant     DeviceComplianceSettingStateState = "compliant"
	DeviceComplianceSettingStateState_Conflict      DeviceComplianceSettingStateState = "conflict"
	DeviceComplianceSettingStateState_Error         DeviceComplianceSettingStateState = "error"
	DeviceComplianceSettingStateState_NonCompliant  DeviceComplianceSettingStateState = "nonCompliant"
	DeviceComplianceSettingStateState_NotApplicable DeviceComplianceSettingStateState = "notApplicable"
	DeviceComplianceSettingStateState_NotAssigned   DeviceComplianceSettingStateState = "notAssigned"
	DeviceComplianceSettingStateState_Remediated    DeviceComplianceSettingStateState = "remediated"
	DeviceComplianceSettingStateState_Unknown       DeviceComplianceSettingStateState = "unknown"
)

func PossibleValuesForDeviceComplianceSettingStateState() []string {
	return []string{
		string(DeviceComplianceSettingStateState_Compliant),
		string(DeviceComplianceSettingStateState_Conflict),
		string(DeviceComplianceSettingStateState_Error),
		string(DeviceComplianceSettingStateState_NonCompliant),
		string(DeviceComplianceSettingStateState_NotApplicable),
		string(DeviceComplianceSettingStateState_NotAssigned),
		string(DeviceComplianceSettingStateState_Remediated),
		string(DeviceComplianceSettingStateState_Unknown),
	}
}

func (s *DeviceComplianceSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceSettingStateState(input string) (*DeviceComplianceSettingStateState, error) {
	vals := map[string]DeviceComplianceSettingStateState{
		"compliant":     DeviceComplianceSettingStateState_Compliant,
		"conflict":      DeviceComplianceSettingStateState_Conflict,
		"error":         DeviceComplianceSettingStateState_Error,
		"noncompliant":  DeviceComplianceSettingStateState_NonCompliant,
		"notapplicable": DeviceComplianceSettingStateState_NotApplicable,
		"notassigned":   DeviceComplianceSettingStateState_NotAssigned,
		"remediated":    DeviceComplianceSettingStateState_Remediated,
		"unknown":       DeviceComplianceSettingStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceSettingStateState(input)
	return &out, nil
}
