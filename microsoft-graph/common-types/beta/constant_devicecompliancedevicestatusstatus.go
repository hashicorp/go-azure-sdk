package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceDeviceStatusStatus string

const (
	DeviceComplianceDeviceStatusStatus_Compliant     DeviceComplianceDeviceStatusStatus = "compliant"
	DeviceComplianceDeviceStatusStatus_Conflict      DeviceComplianceDeviceStatusStatus = "conflict"
	DeviceComplianceDeviceStatusStatus_Error         DeviceComplianceDeviceStatusStatus = "error"
	DeviceComplianceDeviceStatusStatus_NonCompliant  DeviceComplianceDeviceStatusStatus = "nonCompliant"
	DeviceComplianceDeviceStatusStatus_NotApplicable DeviceComplianceDeviceStatusStatus = "notApplicable"
	DeviceComplianceDeviceStatusStatus_NotAssigned   DeviceComplianceDeviceStatusStatus = "notAssigned"
	DeviceComplianceDeviceStatusStatus_Remediated    DeviceComplianceDeviceStatusStatus = "remediated"
	DeviceComplianceDeviceStatusStatus_Unknown       DeviceComplianceDeviceStatusStatus = "unknown"
)

func PossibleValuesForDeviceComplianceDeviceStatusStatus() []string {
	return []string{
		string(DeviceComplianceDeviceStatusStatus_Compliant),
		string(DeviceComplianceDeviceStatusStatus_Conflict),
		string(DeviceComplianceDeviceStatusStatus_Error),
		string(DeviceComplianceDeviceStatusStatus_NonCompliant),
		string(DeviceComplianceDeviceStatusStatus_NotApplicable),
		string(DeviceComplianceDeviceStatusStatus_NotAssigned),
		string(DeviceComplianceDeviceStatusStatus_Remediated),
		string(DeviceComplianceDeviceStatusStatus_Unknown),
	}
}

func (s *DeviceComplianceDeviceStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceDeviceStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceDeviceStatusStatus(input string) (*DeviceComplianceDeviceStatusStatus, error) {
	vals := map[string]DeviceComplianceDeviceStatusStatus{
		"compliant":     DeviceComplianceDeviceStatusStatus_Compliant,
		"conflict":      DeviceComplianceDeviceStatusStatus_Conflict,
		"error":         DeviceComplianceDeviceStatusStatus_Error,
		"noncompliant":  DeviceComplianceDeviceStatusStatus_NonCompliant,
		"notapplicable": DeviceComplianceDeviceStatusStatus_NotApplicable,
		"notassigned":   DeviceComplianceDeviceStatusStatus_NotAssigned,
		"remediated":    DeviceComplianceDeviceStatusStatus_Remediated,
		"unknown":       DeviceComplianceDeviceStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceDeviceStatusStatus(input)
	return &out, nil
}
