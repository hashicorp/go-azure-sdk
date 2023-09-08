package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientHealthStateState string

const (
	ConfigurationManagerClientHealthStateStatecommunicationError ConfigurationManagerClientHealthStateState = "CommunicationError"
	ConfigurationManagerClientHealthStateStatehealthy            ConfigurationManagerClientHealthStateState = "Healthy"
	ConfigurationManagerClientHealthStateStateinstallFailed      ConfigurationManagerClientHealthStateState = "InstallFailed"
	ConfigurationManagerClientHealthStateStateinstalled          ConfigurationManagerClientHealthStateState = "Installed"
	ConfigurationManagerClientHealthStateStateunknown            ConfigurationManagerClientHealthStateState = "Unknown"
	ConfigurationManagerClientHealthStateStateupdateFailed       ConfigurationManagerClientHealthStateState = "UpdateFailed"
)

func PossibleValuesForConfigurationManagerClientHealthStateState() []string {
	return []string{
		string(ConfigurationManagerClientHealthStateStatecommunicationError),
		string(ConfigurationManagerClientHealthStateStatehealthy),
		string(ConfigurationManagerClientHealthStateStateinstallFailed),
		string(ConfigurationManagerClientHealthStateStateinstalled),
		string(ConfigurationManagerClientHealthStateStateunknown),
		string(ConfigurationManagerClientHealthStateStateupdateFailed),
	}
}

func parseConfigurationManagerClientHealthStateState(input string) (*ConfigurationManagerClientHealthStateState, error) {
	vals := map[string]ConfigurationManagerClientHealthStateState{
		"communicationerror": ConfigurationManagerClientHealthStateStatecommunicationError,
		"healthy":            ConfigurationManagerClientHealthStateStatehealthy,
		"installfailed":      ConfigurationManagerClientHealthStateStateinstallFailed,
		"installed":          ConfigurationManagerClientHealthStateStateinstalled,
		"unknown":            ConfigurationManagerClientHealthStateStateunknown,
		"updatefailed":       ConfigurationManagerClientHealthStateStateupdateFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerClientHealthStateState(input)
	return &out, nil
}
