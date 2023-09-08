package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction string

const (
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionaudit         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction = "Audit"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionblock         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction = "Block"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActiondeviceDefault Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction = "DeviceDefault"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionaudit),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionblock),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActiondeviceDefault),
	}
}

func parseWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction(input string) (*Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction{
		"audit":         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionaudit,
		"block":         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionblock,
		"devicedefault": Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActiondeviceDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction(input)
	return &out, nil
}
