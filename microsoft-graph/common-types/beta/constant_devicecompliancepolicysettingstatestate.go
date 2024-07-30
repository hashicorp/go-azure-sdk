package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateState string

const (
	DeviceCompliancePolicySettingStateState_Compliant     DeviceCompliancePolicySettingStateState = "compliant"
	DeviceCompliancePolicySettingStateState_Conflict      DeviceCompliancePolicySettingStateState = "conflict"
	DeviceCompliancePolicySettingStateState_Error         DeviceCompliancePolicySettingStateState = "error"
	DeviceCompliancePolicySettingStateState_NonCompliant  DeviceCompliancePolicySettingStateState = "nonCompliant"
	DeviceCompliancePolicySettingStateState_NotApplicable DeviceCompliancePolicySettingStateState = "notApplicable"
	DeviceCompliancePolicySettingStateState_NotAssigned   DeviceCompliancePolicySettingStateState = "notAssigned"
	DeviceCompliancePolicySettingStateState_Remediated    DeviceCompliancePolicySettingStateState = "remediated"
	DeviceCompliancePolicySettingStateState_Unknown       DeviceCompliancePolicySettingStateState = "unknown"
)

func PossibleValuesForDeviceCompliancePolicySettingStateState() []string {
	return []string{
		string(DeviceCompliancePolicySettingStateState_Compliant),
		string(DeviceCompliancePolicySettingStateState_Conflict),
		string(DeviceCompliancePolicySettingStateState_Error),
		string(DeviceCompliancePolicySettingStateState_NonCompliant),
		string(DeviceCompliancePolicySettingStateState_NotApplicable),
		string(DeviceCompliancePolicySettingStateState_NotAssigned),
		string(DeviceCompliancePolicySettingStateState_Remediated),
		string(DeviceCompliancePolicySettingStateState_Unknown),
	}
}

func (s *DeviceCompliancePolicySettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicySettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicySettingStateState(input string) (*DeviceCompliancePolicySettingStateState, error) {
	vals := map[string]DeviceCompliancePolicySettingStateState{
		"compliant":     DeviceCompliancePolicySettingStateState_Compliant,
		"conflict":      DeviceCompliancePolicySettingStateState_Conflict,
		"error":         DeviceCompliancePolicySettingStateState_Error,
		"noncompliant":  DeviceCompliancePolicySettingStateState_NonCompliant,
		"notapplicable": DeviceCompliancePolicySettingStateState_NotApplicable,
		"notassigned":   DeviceCompliancePolicySettingStateState_NotAssigned,
		"remediated":    DeviceCompliancePolicySettingStateState_Remediated,
		"unknown":       DeviceCompliancePolicySettingStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicySettingStateState(input)
	return &out, nil
}
