package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRecordSeverity string

const (
	DeviceManagementAlertRecordSeveritycritical      DeviceManagementAlertRecordSeverity = "Critical"
	DeviceManagementAlertRecordSeverityinformational DeviceManagementAlertRecordSeverity = "Informational"
	DeviceManagementAlertRecordSeverityunknown       DeviceManagementAlertRecordSeverity = "Unknown"
	DeviceManagementAlertRecordSeveritywarning       DeviceManagementAlertRecordSeverity = "Warning"
)

func PossibleValuesForDeviceManagementAlertRecordSeverity() []string {
	return []string{
		string(DeviceManagementAlertRecordSeveritycritical),
		string(DeviceManagementAlertRecordSeverityinformational),
		string(DeviceManagementAlertRecordSeverityunknown),
		string(DeviceManagementAlertRecordSeveritywarning),
	}
}

func parseDeviceManagementAlertRecordSeverity(input string) (*DeviceManagementAlertRecordSeverity, error) {
	vals := map[string]DeviceManagementAlertRecordSeverity{
		"critical":      DeviceManagementAlertRecordSeveritycritical,
		"informational": DeviceManagementAlertRecordSeverityinformational,
		"unknown":       DeviceManagementAlertRecordSeverityunknown,
		"warning":       DeviceManagementAlertRecordSeveritywarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRecordSeverity(input)
	return &out, nil
}
