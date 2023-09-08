package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCachedReportConfigurationStatus string

const (
	DeviceManagementCachedReportConfigurationStatuscompleted  DeviceManagementCachedReportConfigurationStatus = "Completed"
	DeviceManagementCachedReportConfigurationStatusfailed     DeviceManagementCachedReportConfigurationStatus = "Failed"
	DeviceManagementCachedReportConfigurationStatusinProgress DeviceManagementCachedReportConfigurationStatus = "InProgress"
	DeviceManagementCachedReportConfigurationStatusnotStarted DeviceManagementCachedReportConfigurationStatus = "NotStarted"
	DeviceManagementCachedReportConfigurationStatusunknown    DeviceManagementCachedReportConfigurationStatus = "Unknown"
)

func PossibleValuesForDeviceManagementCachedReportConfigurationStatus() []string {
	return []string{
		string(DeviceManagementCachedReportConfigurationStatuscompleted),
		string(DeviceManagementCachedReportConfigurationStatusfailed),
		string(DeviceManagementCachedReportConfigurationStatusinProgress),
		string(DeviceManagementCachedReportConfigurationStatusnotStarted),
		string(DeviceManagementCachedReportConfigurationStatusunknown),
	}
}

func parseDeviceManagementCachedReportConfigurationStatus(input string) (*DeviceManagementCachedReportConfigurationStatus, error) {
	vals := map[string]DeviceManagementCachedReportConfigurationStatus{
		"completed":  DeviceManagementCachedReportConfigurationStatuscompleted,
		"failed":     DeviceManagementCachedReportConfigurationStatusfailed,
		"inprogress": DeviceManagementCachedReportConfigurationStatusinProgress,
		"notstarted": DeviceManagementCachedReportConfigurationStatusnotStarted,
		"unknown":    DeviceManagementCachedReportConfigurationStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCachedReportConfigurationStatus(input)
	return &out, nil
}
