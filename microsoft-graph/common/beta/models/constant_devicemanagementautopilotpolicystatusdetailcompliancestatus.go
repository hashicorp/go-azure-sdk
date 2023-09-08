package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotPolicyStatusDetailComplianceStatus string

const (
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatuscompliant    DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "Compliant"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatuserror        DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "Error"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatusinstalled    DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "Installed"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatusnotCompliant DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "NotCompliant"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatusnotInstalled DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "NotInstalled"
	DeviceManagementAutopilotPolicyStatusDetailComplianceStatusunknown      DeviceManagementAutopilotPolicyStatusDetailComplianceStatus = "Unknown"
)

func PossibleValuesForDeviceManagementAutopilotPolicyStatusDetailComplianceStatus() []string {
	return []string{
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatuscompliant),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatuserror),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatusinstalled),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatusnotCompliant),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatusnotInstalled),
		string(DeviceManagementAutopilotPolicyStatusDetailComplianceStatusunknown),
	}
}

func parseDeviceManagementAutopilotPolicyStatusDetailComplianceStatus(input string) (*DeviceManagementAutopilotPolicyStatusDetailComplianceStatus, error) {
	vals := map[string]DeviceManagementAutopilotPolicyStatusDetailComplianceStatus{
		"compliant":    DeviceManagementAutopilotPolicyStatusDetailComplianceStatuscompliant,
		"error":        DeviceManagementAutopilotPolicyStatusDetailComplianceStatuserror,
		"installed":    DeviceManagementAutopilotPolicyStatusDetailComplianceStatusinstalled,
		"notcompliant": DeviceManagementAutopilotPolicyStatusDetailComplianceStatusnotCompliant,
		"notinstalled": DeviceManagementAutopilotPolicyStatusDetailComplianceStatusnotInstalled,
		"unknown":      DeviceManagementAutopilotPolicyStatusDetailComplianceStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotPolicyStatusDetailComplianceStatus(input)
	return &out, nil
}
