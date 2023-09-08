package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationdisabled      AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "Disabled"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationhide          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "Hide"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationnotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationshow          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "Show"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationdisabled),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationhide),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationshow),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration{
		"disabled":      AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationdisabled,
		"hide":          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationhide,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationnotConfigured,
		"show":          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfigurationshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration(input)
	return &out, nil
}
