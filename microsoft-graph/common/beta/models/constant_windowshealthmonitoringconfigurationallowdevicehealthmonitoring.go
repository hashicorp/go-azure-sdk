package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring string

const (
	WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringdisabled      WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring = "Disabled"
	WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringenabled       WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring = "Enabled"
	WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringnotConfigured WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring = "NotConfigured"
)

func PossibleValuesForWindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring() []string {
	return []string{
		string(WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringdisabled),
		string(WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringenabled),
		string(WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringnotConfigured),
	}
}

func parseWindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring(input string) (*WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring, error) {
	vals := map[string]WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring{
		"disabled":      WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringdisabled,
		"enabled":       WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringenabled,
		"notconfigured": WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoringnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring(input)
	return &out, nil
}
