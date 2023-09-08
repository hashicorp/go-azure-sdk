package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConfigurationProfileTarget string

const (
	Windows10VpnConfigurationProfileTargetautoPilotDevice Windows10VpnConfigurationProfileTarget = "AutoPilotDevice"
	Windows10VpnConfigurationProfileTargetdevice          Windows10VpnConfigurationProfileTarget = "Device"
	Windows10VpnConfigurationProfileTargetuser            Windows10VpnConfigurationProfileTarget = "User"
)

func PossibleValuesForWindows10VpnConfigurationProfileTarget() []string {
	return []string{
		string(Windows10VpnConfigurationProfileTargetautoPilotDevice),
		string(Windows10VpnConfigurationProfileTargetdevice),
		string(Windows10VpnConfigurationProfileTargetuser),
	}
}

func parseWindows10VpnConfigurationProfileTarget(input string) (*Windows10VpnConfigurationProfileTarget, error) {
	vals := map[string]Windows10VpnConfigurationProfileTarget{
		"autopilotdevice": Windows10VpnConfigurationProfileTargetautoPilotDevice,
		"device":          Windows10VpnConfigurationProfileTargetdevice,
		"user":            Windows10VpnConfigurationProfileTargetuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnConfigurationProfileTarget(input)
	return &out, nil
}
