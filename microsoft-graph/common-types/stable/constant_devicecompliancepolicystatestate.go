package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyStateState string

const (
	DeviceCompliancePolicyStateState_Compliant     DeviceCompliancePolicyStateState = "compliant"
	DeviceCompliancePolicyStateState_Conflict      DeviceCompliancePolicyStateState = "conflict"
	DeviceCompliancePolicyStateState_Error         DeviceCompliancePolicyStateState = "error"
	DeviceCompliancePolicyStateState_NonCompliant  DeviceCompliancePolicyStateState = "nonCompliant"
	DeviceCompliancePolicyStateState_NotApplicable DeviceCompliancePolicyStateState = "notApplicable"
	DeviceCompliancePolicyStateState_NotAssigned   DeviceCompliancePolicyStateState = "notAssigned"
	DeviceCompliancePolicyStateState_Remediated    DeviceCompliancePolicyStateState = "remediated"
	DeviceCompliancePolicyStateState_Unknown       DeviceCompliancePolicyStateState = "unknown"
)

func PossibleValuesForDeviceCompliancePolicyStateState() []string {
	return []string{
		string(DeviceCompliancePolicyStateState_Compliant),
		string(DeviceCompliancePolicyStateState_Conflict),
		string(DeviceCompliancePolicyStateState_Error),
		string(DeviceCompliancePolicyStateState_NonCompliant),
		string(DeviceCompliancePolicyStateState_NotApplicable),
		string(DeviceCompliancePolicyStateState_NotAssigned),
		string(DeviceCompliancePolicyStateState_Remediated),
		string(DeviceCompliancePolicyStateState_Unknown),
	}
}

func (s *DeviceCompliancePolicyStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCompliancePolicyStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCompliancePolicyStateState(input string) (*DeviceCompliancePolicyStateState, error) {
	vals := map[string]DeviceCompliancePolicyStateState{
		"compliant":     DeviceCompliancePolicyStateState_Compliant,
		"conflict":      DeviceCompliancePolicyStateState_Conflict,
		"error":         DeviceCompliancePolicyStateState_Error,
		"noncompliant":  DeviceCompliancePolicyStateState_NonCompliant,
		"notapplicable": DeviceCompliancePolicyStateState_NotApplicable,
		"notassigned":   DeviceCompliancePolicyStateState_NotAssigned,
		"remediated":    DeviceCompliancePolicyStateState_Remediated,
		"unknown":       DeviceCompliancePolicyStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyStateState(input)
	return &out, nil
}
