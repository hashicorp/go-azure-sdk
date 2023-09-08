package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationAutomaticUpdateMode string

const (
	WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootAtMaintenanceTime     WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "AutoInstallAndRebootAtMaintenanceTime"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootAtScheduledTime       WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "AutoInstallAndRebootAtScheduledTime"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootWithoutEndUserControl WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "AutoInstallAndRebootWithoutEndUserControl"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAtMaintenanceTime              WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "AutoInstallAtMaintenanceTime"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateModenotifyDownload                            WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "NotifyDownload"
	WindowsUpdateForBusinessConfigurationAutomaticUpdateModeuserDefined                               WindowsUpdateForBusinessConfigurationAutomaticUpdateMode = "UserDefined"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationAutomaticUpdateMode() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootAtMaintenanceTime),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootAtScheduledTime),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootWithoutEndUserControl),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAtMaintenanceTime),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateModenotifyDownload),
		string(WindowsUpdateForBusinessConfigurationAutomaticUpdateModeuserDefined),
	}
}

func parseWindowsUpdateForBusinessConfigurationAutomaticUpdateMode(input string) (*WindowsUpdateForBusinessConfigurationAutomaticUpdateMode, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationAutomaticUpdateMode{
		"autoinstallandrebootatmaintenancetime":     WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootAtMaintenanceTime,
		"autoinstallandrebootatscheduledtime":       WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootAtScheduledTime,
		"autoinstallandrebootwithoutendusercontrol": WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAndRebootWithoutEndUserControl,
		"autoinstallatmaintenancetime":              WindowsUpdateForBusinessConfigurationAutomaticUpdateModeautoInstallAtMaintenanceTime,
		"notifydownload":                            WindowsUpdateForBusinessConfigurationAutomaticUpdateModenotifyDownload,
		"userdefined":                               WindowsUpdateForBusinessConfigurationAutomaticUpdateModeuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationAutomaticUpdateMode(input)
	return &out, nil
}
