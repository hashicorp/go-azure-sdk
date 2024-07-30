package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope string

const (
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_BootPerformance     WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "bootPerformance"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_HealthMonitoring    WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "healthMonitoring"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_PrivilegeManagement WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "privilegeManagement"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_Undefined           WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "undefined"
	WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_WindowsUpdates      WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope = "windowsUpdates"
)

func PossibleValuesForWindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope() []string {
	return []string{
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_BootPerformance),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_HealthMonitoring),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_PrivilegeManagement),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_Undefined),
		string(WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_WindowsUpdates),
	}
}

func (s *WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope(input string) (*WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope, error) {
	vals := map[string]WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope{
		"bootperformance":     WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_BootPerformance,
		"healthmonitoring":    WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_HealthMonitoring,
		"privilegemanagement": WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_PrivilegeManagement,
		"undefined":           WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_Undefined,
		"windowsupdates":      WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope_WindowsUpdates,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHealthMonitoringConfigurationConfigDeviceHealthMonitoringScope(input)
	return &out, nil
}
