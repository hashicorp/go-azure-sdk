package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceUserStatusStatus string

const (
	DeviceComplianceUserStatusStatus_Compliant     DeviceComplianceUserStatusStatus = "compliant"
	DeviceComplianceUserStatusStatus_Conflict      DeviceComplianceUserStatusStatus = "conflict"
	DeviceComplianceUserStatusStatus_Error         DeviceComplianceUserStatusStatus = "error"
	DeviceComplianceUserStatusStatus_NonCompliant  DeviceComplianceUserStatusStatus = "nonCompliant"
	DeviceComplianceUserStatusStatus_NotApplicable DeviceComplianceUserStatusStatus = "notApplicable"
	DeviceComplianceUserStatusStatus_NotAssigned   DeviceComplianceUserStatusStatus = "notAssigned"
	DeviceComplianceUserStatusStatus_Remediated    DeviceComplianceUserStatusStatus = "remediated"
	DeviceComplianceUserStatusStatus_Unknown       DeviceComplianceUserStatusStatus = "unknown"
)

func PossibleValuesForDeviceComplianceUserStatusStatus() []string {
	return []string{
		string(DeviceComplianceUserStatusStatus_Compliant),
		string(DeviceComplianceUserStatusStatus_Conflict),
		string(DeviceComplianceUserStatusStatus_Error),
		string(DeviceComplianceUserStatusStatus_NonCompliant),
		string(DeviceComplianceUserStatusStatus_NotApplicable),
		string(DeviceComplianceUserStatusStatus_NotAssigned),
		string(DeviceComplianceUserStatusStatus_Remediated),
		string(DeviceComplianceUserStatusStatus_Unknown),
	}
}

func (s *DeviceComplianceUserStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceUserStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceUserStatusStatus(input string) (*DeviceComplianceUserStatusStatus, error) {
	vals := map[string]DeviceComplianceUserStatusStatus{
		"compliant":     DeviceComplianceUserStatusStatus_Compliant,
		"conflict":      DeviceComplianceUserStatusStatus_Conflict,
		"error":         DeviceComplianceUserStatusStatus_Error,
		"noncompliant":  DeviceComplianceUserStatusStatus_NonCompliant,
		"notapplicable": DeviceComplianceUserStatusStatus_NotApplicable,
		"notassigned":   DeviceComplianceUserStatusStatus_NotAssigned,
		"remediated":    DeviceComplianceUserStatusStatus_Remediated,
		"unknown":       DeviceComplianceUserStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceUserStatusStatus(input)
	return &out, nil
}
