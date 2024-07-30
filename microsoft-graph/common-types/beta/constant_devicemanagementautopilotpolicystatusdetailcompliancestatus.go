package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotPolicyStatusDetailComplianceStatus string

const (
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Compliant    DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "compliant"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Error        DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "error"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Installed    DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "installed"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_NotCompliant DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "notCompliant"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_NotInstalled DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "notInstalled"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Unknown      DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotPolicyStatusDetailComplianceStatus() []string {
	return []string{
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Compliant),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Error),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Installed),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_NotCompliant),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_NotInstalled),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Unknown),
	}
}

func (s *DeviceManagementAutopilotPolicyStatusDetailComplianceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotPolicyStatusDetailComplianceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotPolicyStatusDetailComplianceStatus(input string) (*DeviceManagementAutopilotPolicyStatusDetailComplianceStatus, error) {
	vals := map[string]DeviceManagementAutopilotPolicyStatusDetailComplianceStatus{
		"compliant":    DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Compliant,
		"error":        DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Error,
		"installed":    DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Installed,
		"notcompliant": DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_NotCompliant,
		"notinstalled": DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_NotInstalled,
		"unknown":      DeviceManagementAutopilotPolicyStatusDetailComplianceStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotPolicyStatusDetailComplianceStatus(input)
	return &out, nil
}
