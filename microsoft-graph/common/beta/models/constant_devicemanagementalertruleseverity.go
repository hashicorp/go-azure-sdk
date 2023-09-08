package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRuleSeverity string

const (
	DeviceManagementAlertRuleSeveritycritical      DeviceManagementAlertRuleSeverity = "Critical"
	DeviceManagementAlertRuleSeverityinformational DeviceManagementAlertRuleSeverity = "Informational"
	DeviceManagementAlertRuleSeverityunknown       DeviceManagementAlertRuleSeverity = "Unknown"
	DeviceManagementAlertRuleSeveritywarning       DeviceManagementAlertRuleSeverity = "Warning"
)

func PossibleValuesForDeviceManagementAlertRuleSeverity() []string {
	return []string{
		string(DeviceManagementAlertRuleSeveritycritical),
		string(DeviceManagementAlertRuleSeverityinformational),
		string(DeviceManagementAlertRuleSeverityunknown),
		string(DeviceManagementAlertRuleSeveritywarning),
	}
}

func parseDeviceManagementAlertRuleSeverity(input string) (*DeviceManagementAlertRuleSeverity, error) {
	vals := map[string]DeviceManagementAlertRuleSeverity{
		"critical":      DeviceManagementAlertRuleSeveritycritical,
		"informational": DeviceManagementAlertRuleSeverityinformational,
		"unknown":       DeviceManagementAlertRuleSeverityunknown,
		"warning":       DeviceManagementAlertRuleSeveritywarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRuleSeverity(input)
	return &out, nil
}
