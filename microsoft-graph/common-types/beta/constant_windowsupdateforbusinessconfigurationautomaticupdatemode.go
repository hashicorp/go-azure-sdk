package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationAutomaticUpdateMode string

const (
	WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootAtMaintenanceTime     WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "autoInstallAndRebootAtMaintenanceTime"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootAtScheduledTime       WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "autoInstallAndRebootAtScheduledTime"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootWithoutEndUserControl WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "autoInstallAndRebootWithoutEndUserControl"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAtMaintenanceTime              WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "autoInstallAtMaintenanceTime"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_NotifyDownload                            WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "notifyDownload"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_UserDefined                               WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "userDefined"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_WindowsDefault                            WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "windowsDefault"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationAutomaticUpdateMode() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootAtMaintenanceTime),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootAtScheduledTime),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootWithoutEndUserControl),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAtMaintenanceTime),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_NotifyDownload),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_UserDefined),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_WindowsDefault),
	}
}

func (s *WindowsUpdateForBusinessConfigurationAutomaticUpdateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationAutomaticUpdateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationAutomaticUpdateMode(input string) (*WindowsUpdateForBusinessConfigurationAutomaticUpdateMode, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationAutomaticUpdateMode{
		"autoinstallandrebootatmaintenancetime":     WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootAtMaintenanceTime,
		"autoinstallandrebootatscheduledtime":       WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootAtScheduledTime,
		"autoinstallandrebootwithoutendusercontrol": WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAndRebootWithoutEndUserControl,
		"autoinstallatmaintenancetime":              WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_AutoInstallAtMaintenanceTime,
		"notifydownload":                            WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_NotifyDownload,
		"userdefined":                               WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_UserDefined,
		"windowsdefault":                            WindowsUpdateForBusinessConfigurationAutomaticUpdateMode_WindowsDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationAutomaticUpdateMode(input)
	return &out, nil
}
