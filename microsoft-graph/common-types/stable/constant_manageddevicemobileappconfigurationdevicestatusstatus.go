package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationDeviceStatusStatus string

const (
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Compliant     ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "compliant"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Conflict      ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "conflict"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Error         ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "error"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NonCompliant  ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "nonCompliant"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NotApplicable ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "notApplicable"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NotAssigned   ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "notAssigned"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Remediated    ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "remediated"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Unknown       ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationDeviceStatusStatus() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Compliant),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Conflict),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Error),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NonCompliant),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NotApplicable),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NotAssigned),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Remediated),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Unknown),
	}
}

func (s *ManagedDeviceMobileAppConfigurationDeviceStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationDeviceStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationDeviceStatusStatus(input string) (*ManagedDeviceMobileAppConfigurationDeviceStatusStatus, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationDeviceStatusStatus{
		"compliant":     ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Compliant,
		"conflict":      ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Conflict,
		"error":         ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Error,
		"noncompliant":  ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationDeviceStatusStatus_NotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Remediated,
		"unknown":       ManagedDeviceMobileAppConfigurationDeviceStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationDeviceStatusStatus(input)
	return &out, nil
}
