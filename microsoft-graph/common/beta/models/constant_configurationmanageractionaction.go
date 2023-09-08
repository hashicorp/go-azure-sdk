package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionAction string

const (
	ConfigurationManagerActionActionappEvaluation                   ConfigurationManagerActionAction = "AppEvaluation"
	ConfigurationManagerActionActionfullScan                        ConfigurationManagerActionAction = "FullScan"
	ConfigurationManagerActionActionquickScan                       ConfigurationManagerActionAction = "QuickScan"
	ConfigurationManagerActionActionrefreshMachinePolicy            ConfigurationManagerActionAction = "RefreshMachinePolicy"
	ConfigurationManagerActionActionrefreshUserPolicy               ConfigurationManagerActionAction = "RefreshUserPolicy"
	ConfigurationManagerActionActionwakeUpClient                    ConfigurationManagerActionAction = "WakeUpClient"
	ConfigurationManagerActionActionwindowsDefenderUpdateSignatures ConfigurationManagerActionAction = "WindowsDefenderUpdateSignatures"
)

func PossibleValuesForConfigurationManagerActionAction() []string {
	return []string{
		string(ConfigurationManagerActionActionappEvaluation),
		string(ConfigurationManagerActionActionfullScan),
		string(ConfigurationManagerActionActionquickScan),
		string(ConfigurationManagerActionActionrefreshMachinePolicy),
		string(ConfigurationManagerActionActionrefreshUserPolicy),
		string(ConfigurationManagerActionActionwakeUpClient),
		string(ConfigurationManagerActionActionwindowsDefenderUpdateSignatures),
	}
}

func parseConfigurationManagerActionAction(input string) (*ConfigurationManagerActionAction, error) {
	vals := map[string]ConfigurationManagerActionAction{
		"appevaluation":                   ConfigurationManagerActionActionappEvaluation,
		"fullscan":                        ConfigurationManagerActionActionfullScan,
		"quickscan":                       ConfigurationManagerActionActionquickScan,
		"refreshmachinepolicy":            ConfigurationManagerActionActionrefreshMachinePolicy,
		"refreshuserpolicy":               ConfigurationManagerActionActionrefreshUserPolicy,
		"wakeupclient":                    ConfigurationManagerActionActionwakeUpClient,
		"windowsdefenderupdatesignatures": ConfigurationManagerActionActionwindowsDefenderUpdateSignatures,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionAction(input)
	return &out, nil
}
