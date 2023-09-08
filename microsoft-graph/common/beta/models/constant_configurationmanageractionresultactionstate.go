package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionResultActionState string

const (
	ConfigurationManagerActionResultActionStateactive       ConfigurationManagerActionResultActionState = "Active"
	ConfigurationManagerActionResultActionStatecanceled     ConfigurationManagerActionResultActionState = "Canceled"
	ConfigurationManagerActionResultActionStatedone         ConfigurationManagerActionResultActionState = "Done"
	ConfigurationManagerActionResultActionStatefailed       ConfigurationManagerActionResultActionState = "Failed"
	ConfigurationManagerActionResultActionStatenone         ConfigurationManagerActionResultActionState = "None"
	ConfigurationManagerActionResultActionStatenotSupported ConfigurationManagerActionResultActionState = "NotSupported"
	ConfigurationManagerActionResultActionStatepending      ConfigurationManagerActionResultActionState = "Pending"
)

func PossibleValuesForConfigurationManagerActionResultActionState() []string {
	return []string{
		string(ConfigurationManagerActionResultActionStateactive),
		string(ConfigurationManagerActionResultActionStatecanceled),
		string(ConfigurationManagerActionResultActionStatedone),
		string(ConfigurationManagerActionResultActionStatefailed),
		string(ConfigurationManagerActionResultActionStatenone),
		string(ConfigurationManagerActionResultActionStatenotSupported),
		string(ConfigurationManagerActionResultActionStatepending),
	}
}

func parseConfigurationManagerActionResultActionState(input string) (*ConfigurationManagerActionResultActionState, error) {
	vals := map[string]ConfigurationManagerActionResultActionState{
		"active":       ConfigurationManagerActionResultActionStateactive,
		"canceled":     ConfigurationManagerActionResultActionStatecanceled,
		"done":         ConfigurationManagerActionResultActionStatedone,
		"failed":       ConfigurationManagerActionResultActionStatefailed,
		"none":         ConfigurationManagerActionResultActionStatenone,
		"notsupported": ConfigurationManagerActionResultActionStatenotSupported,
		"pending":      ConfigurationManagerActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionResultActionState(input)
	return &out, nil
}
