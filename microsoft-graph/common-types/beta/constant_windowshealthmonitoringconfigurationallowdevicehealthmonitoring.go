package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring string

const (
	WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_Disabled      WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring = "disabled"
	WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_Enabled       WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring = "enabled"
	WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_NotConfigured WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring = "notConfigured"
)

func PossibleValuesForWindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring() []string {
	return []string{
		string(WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_Disabled),
		string(WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_Enabled),
		string(WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_NotConfigured),
	}
}

func (s *WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring(input string) (*WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring, error) {
	vals := map[string]WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring{
		"disabled":      WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_Disabled,
		"enabled":       WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_Enabled,
		"notconfigured": WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHealthMonitoringConfigurationAllowDeviceHealthMonitoring(input)
	return &out, nil
}
