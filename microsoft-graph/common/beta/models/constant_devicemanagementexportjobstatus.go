package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJobStatus string

const (
	DeviceManagementExportJobStatuscompleted  DeviceManagementExportJobStatus = "Completed"
	DeviceManagementExportJobStatusfailed     DeviceManagementExportJobStatus = "Failed"
	DeviceManagementExportJobStatusinProgress DeviceManagementExportJobStatus = "InProgress"
	DeviceManagementExportJobStatusnotStarted DeviceManagementExportJobStatus = "NotStarted"
	DeviceManagementExportJobStatusunknown    DeviceManagementExportJobStatus = "Unknown"
)

func PossibleValuesForDeviceManagementExportJobStatus() []string {
	return []string{
		string(DeviceManagementExportJobStatuscompleted),
		string(DeviceManagementExportJobStatusfailed),
		string(DeviceManagementExportJobStatusinProgress),
		string(DeviceManagementExportJobStatusnotStarted),
		string(DeviceManagementExportJobStatusunknown),
	}
}

func parseDeviceManagementExportJobStatus(input string) (*DeviceManagementExportJobStatus, error) {
	vals := map[string]DeviceManagementExportJobStatus{
		"completed":  DeviceManagementExportJobStatuscompleted,
		"failed":     DeviceManagementExportJobStatusfailed,
		"inprogress": DeviceManagementExportJobStatusinProgress,
		"notstarted": DeviceManagementExportJobStatusnotStarted,
		"unknown":    DeviceManagementExportJobStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExportJobStatus(input)
	return &out, nil
}
