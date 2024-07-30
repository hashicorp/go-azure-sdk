package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRecordSeverity string

const (
	DeviceManagementAlertRecordSeverity_Critical      DeviceManagementAlertRecordSeverity = "critical"
	DeviceManagementAlertRecordSeverity_Informational DeviceManagementAlertRecordSeverity = "informational"
	DeviceManagementAlertRecordSeverity_Unknown       DeviceManagementAlertRecordSeverity = "unknown"
	DeviceManagementAlertRecordSeverity_Warning       DeviceManagementAlertRecordSeverity = "warning"
)

func PossibleValuesForDeviceManagementAlertRecordSeverity() []string {
	return []string{
		string(DeviceManagementAlertRecordSeverity_Critical),
		string(DeviceManagementAlertRecordSeverity_Informational),
		string(DeviceManagementAlertRecordSeverity_Unknown),
		string(DeviceManagementAlertRecordSeverity_Warning),
	}
}

func (s *DeviceManagementAlertRecordSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertRecordSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertRecordSeverity(input string) (*DeviceManagementAlertRecordSeverity, error) {
	vals := map[string]DeviceManagementAlertRecordSeverity{
		"critical":      DeviceManagementAlertRecordSeverity_Critical,
		"informational": DeviceManagementAlertRecordSeverity_Informational,
		"unknown":       DeviceManagementAlertRecordSeverity_Unknown,
		"warning":       DeviceManagementAlertRecordSeverity_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRecordSeverity(input)
	return &out, nil
}
