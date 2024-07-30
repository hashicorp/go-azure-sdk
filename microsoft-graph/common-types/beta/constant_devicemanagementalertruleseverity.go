package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRuleSeverity string

const (
	DeviceManagementAlertRuleSeverity_Critical      DeviceManagementAlertRuleSeverity = "critical"
	DeviceManagementAlertRuleSeverity_Informational DeviceManagementAlertRuleSeverity = "informational"
	DeviceManagementAlertRuleSeverity_Unknown       DeviceManagementAlertRuleSeverity = "unknown"
	DeviceManagementAlertRuleSeverity_Warning       DeviceManagementAlertRuleSeverity = "warning"
)

func PossibleValuesForDeviceManagementAlertRuleSeverity() []string {
	return []string{
		string(DeviceManagementAlertRuleSeverity_Critical),
		string(DeviceManagementAlertRuleSeverity_Informational),
		string(DeviceManagementAlertRuleSeverity_Unknown),
		string(DeviceManagementAlertRuleSeverity_Warning),
	}
}

func (s *DeviceManagementAlertRuleSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertRuleSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertRuleSeverity(input string) (*DeviceManagementAlertRuleSeverity, error) {
	vals := map[string]DeviceManagementAlertRuleSeverity{
		"critical":      DeviceManagementAlertRuleSeverity_Critical,
		"informational": DeviceManagementAlertRuleSeverity_Informational,
		"unknown":       DeviceManagementAlertRuleSeverity_Unknown,
		"warning":       DeviceManagementAlertRuleSeverity_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRuleSeverity(input)
	return &out, nil
}
