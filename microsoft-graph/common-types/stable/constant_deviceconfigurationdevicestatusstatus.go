package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationDeviceStatusStatus string

const (
	DeviceConfigurationDeviceStatusStatus_Compliant     DeviceConfigurationDeviceStatusStatus = "compliant"
	DeviceConfigurationDeviceStatusStatus_Conflict      DeviceConfigurationDeviceStatusStatus = "conflict"
	DeviceConfigurationDeviceStatusStatus_Error         DeviceConfigurationDeviceStatusStatus = "error"
	DeviceConfigurationDeviceStatusStatus_NonCompliant  DeviceConfigurationDeviceStatusStatus = "nonCompliant"
	DeviceConfigurationDeviceStatusStatus_NotApplicable DeviceConfigurationDeviceStatusStatus = "notApplicable"
	DeviceConfigurationDeviceStatusStatus_NotAssigned   DeviceConfigurationDeviceStatusStatus = "notAssigned"
	DeviceConfigurationDeviceStatusStatus_Remediated    DeviceConfigurationDeviceStatusStatus = "remediated"
	DeviceConfigurationDeviceStatusStatus_Unknown       DeviceConfigurationDeviceStatusStatus = "unknown"
)

func PossibleValuesForDeviceConfigurationDeviceStatusStatus() []string {
	return []string{
		string(DeviceConfigurationDeviceStatusStatus_Compliant),
		string(DeviceConfigurationDeviceStatusStatus_Conflict),
		string(DeviceConfigurationDeviceStatusStatus_Error),
		string(DeviceConfigurationDeviceStatusStatus_NonCompliant),
		string(DeviceConfigurationDeviceStatusStatus_NotApplicable),
		string(DeviceConfigurationDeviceStatusStatus_NotAssigned),
		string(DeviceConfigurationDeviceStatusStatus_Remediated),
		string(DeviceConfigurationDeviceStatusStatus_Unknown),
	}
}

func (s *DeviceConfigurationDeviceStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationDeviceStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationDeviceStatusStatus(input string) (*DeviceConfigurationDeviceStatusStatus, error) {
	vals := map[string]DeviceConfigurationDeviceStatusStatus{
		"compliant":     DeviceConfigurationDeviceStatusStatus_Compliant,
		"conflict":      DeviceConfigurationDeviceStatusStatus_Conflict,
		"error":         DeviceConfigurationDeviceStatusStatus_Error,
		"noncompliant":  DeviceConfigurationDeviceStatusStatus_NonCompliant,
		"notapplicable": DeviceConfigurationDeviceStatusStatus_NotApplicable,
		"notassigned":   DeviceConfigurationDeviceStatusStatus_NotAssigned,
		"remediated":    DeviceConfigurationDeviceStatusStatus_Remediated,
		"unknown":       DeviceConfigurationDeviceStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationDeviceStatusStatus(input)
	return &out, nil
}
