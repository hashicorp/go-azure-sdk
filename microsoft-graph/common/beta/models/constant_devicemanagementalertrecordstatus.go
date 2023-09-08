package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRecordStatus string

const (
	DeviceManagementAlertRecordStatusactive   DeviceManagementAlertRecordStatus = "Active"
	DeviceManagementAlertRecordStatusresolved DeviceManagementAlertRecordStatus = "Resolved"
)

func PossibleValuesForDeviceManagementAlertRecordStatus() []string {
	return []string{
		string(DeviceManagementAlertRecordStatusactive),
		string(DeviceManagementAlertRecordStatusresolved),
	}
}

func parseDeviceManagementAlertRecordStatus(input string) (*DeviceManagementAlertRecordStatus, error) {
	vals := map[string]DeviceManagementAlertRecordStatus{
		"active":   DeviceManagementAlertRecordStatusactive,
		"resolved": DeviceManagementAlertRecordStatusresolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRecordStatus(input)
	return &out, nil
}
