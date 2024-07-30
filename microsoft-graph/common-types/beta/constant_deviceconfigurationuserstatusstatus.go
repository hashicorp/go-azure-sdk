package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationUserStatusStatus string

const (
	DeviceConfigurationUserStatusStatus_Compliant     DeviceConfigurationUserStatusStatus = "compliant"
	DeviceConfigurationUserStatusStatus_Conflict      DeviceConfigurationUserStatusStatus = "conflict"
	DeviceConfigurationUserStatusStatus_Error         DeviceConfigurationUserStatusStatus = "error"
	DeviceConfigurationUserStatusStatus_NonCompliant  DeviceConfigurationUserStatusStatus = "nonCompliant"
	DeviceConfigurationUserStatusStatus_NotApplicable DeviceConfigurationUserStatusStatus = "notApplicable"
	DeviceConfigurationUserStatusStatus_NotAssigned   DeviceConfigurationUserStatusStatus = "notAssigned"
	DeviceConfigurationUserStatusStatus_Remediated    DeviceConfigurationUserStatusStatus = "remediated"
	DeviceConfigurationUserStatusStatus_Unknown       DeviceConfigurationUserStatusStatus = "unknown"
)

func PossibleValuesForDeviceConfigurationUserStatusStatus() []string {
	return []string{
		string(DeviceConfigurationUserStatusStatus_Compliant),
		string(DeviceConfigurationUserStatusStatus_Conflict),
		string(DeviceConfigurationUserStatusStatus_Error),
		string(DeviceConfigurationUserStatusStatus_NonCompliant),
		string(DeviceConfigurationUserStatusStatus_NotApplicable),
		string(DeviceConfigurationUserStatusStatus_NotAssigned),
		string(DeviceConfigurationUserStatusStatus_Remediated),
		string(DeviceConfigurationUserStatusStatus_Unknown),
	}
}

func (s *DeviceConfigurationUserStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationUserStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationUserStatusStatus(input string) (*DeviceConfigurationUserStatusStatus, error) {
	vals := map[string]DeviceConfigurationUserStatusStatus{
		"compliant":     DeviceConfigurationUserStatusStatus_Compliant,
		"conflict":      DeviceConfigurationUserStatusStatus_Conflict,
		"error":         DeviceConfigurationUserStatusStatus_Error,
		"noncompliant":  DeviceConfigurationUserStatusStatus_NonCompliant,
		"notapplicable": DeviceConfigurationUserStatusStatus_NotApplicable,
		"notassigned":   DeviceConfigurationUserStatusStatus_NotAssigned,
		"remediated":    DeviceConfigurationUserStatusStatus_Remediated,
		"unknown":       DeviceConfigurationUserStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationUserStatusStatus(input)
	return &out, nil
}
