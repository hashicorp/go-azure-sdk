package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope string

const (
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopebootPerformance     WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "BootPerformance"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopehealthMonitoring    WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "HealthMonitoring"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopeprivilegeManagement WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "PrivilegeManagement"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopeundefined           WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "Undefined"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopewindowsUpdates      WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "WindowsUpdates"
)

func PossibleValuesForWindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope() []string {
	return []string{
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopebootPerformance),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopehealthMonitoring),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopeprivilegeManagement),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopeundefined),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopewindowsUpdates),
	}
}

func parseWindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope(input string) (*WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope, error) {
	vals := map[string]WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope{
		"bootperformance":     WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopebootPerformance,
		"healthmonitoring":    WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopehealthMonitoring,
		"privilegemanagement": WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopeprivilegeManagement,
		"undefined":           WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopeundefined,
		"windowsupdates":      WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScopewindowsUpdates,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope(input)
	return &out, nil
}
