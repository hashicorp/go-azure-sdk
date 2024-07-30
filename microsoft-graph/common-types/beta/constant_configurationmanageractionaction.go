package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionAction string

const (
	ConfigurationManagerActionAction_AppEvaluation                   ConfigurationManagerActionAction = "appEvaluation"
	ConfigurationManagerActionAction_FullScan                        ConfigurationManagerActionAction = "fullScan"
	ConfigurationManagerActionAction_QuickScan                       ConfigurationManagerActionAction = "quickScan"
	ConfigurationManagerActionAction_RefreshMachinePolicy            ConfigurationManagerActionAction = "refreshMachinePolicy"
	ConfigurationManagerActionAction_RefreshUserPolicy               ConfigurationManagerActionAction = "refreshUserPolicy"
	ConfigurationManagerActionAction_WakeUpClient                    ConfigurationManagerActionAction = "wakeUpClient"
	ConfigurationManagerActionAction_WindowsDefenderUpdateSignatures ConfigurationManagerActionAction = "windowsDefenderUpdateSignatures"
)

func PossibleValuesForConfigurationManagerActionAction() []string {
	return []string{
		string(ConfigurationManagerActionAction_AppEvaluation),
		string(ConfigurationManagerActionAction_FullScan),
		string(ConfigurationManagerActionAction_QuickScan),
		string(ConfigurationManagerActionAction_RefreshMachinePolicy),
		string(ConfigurationManagerActionAction_RefreshUserPolicy),
		string(ConfigurationManagerActionAction_WakeUpClient),
		string(ConfigurationManagerActionAction_WindowsDefenderUpdateSignatures),
	}
}

func (s *ConfigurationManagerActionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerActionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerActionAction(input string) (*ConfigurationManagerActionAction, error) {
	vals := map[string]ConfigurationManagerActionAction{
		"appevaluation":                   ConfigurationManagerActionAction_AppEvaluation,
		"fullscan":                        ConfigurationManagerActionAction_FullScan,
		"quickscan":                       ConfigurationManagerActionAction_QuickScan,
		"refreshmachinepolicy":            ConfigurationManagerActionAction_RefreshMachinePolicy,
		"refreshuserpolicy":               ConfigurationManagerActionAction_RefreshUserPolicy,
		"wakeupclient":                    ConfigurationManagerActionAction_WakeUpClient,
		"windowsdefenderupdatesignatures": ConfigurationManagerActionAction_WindowsDefenderUpdateSignatures,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionAction(input)
	return &out, nil
}
