package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationCellularExceptionAction string

const (
	AppleVpnAlwaysOnConfigurationCellularExceptionActionallowTrafficOutside AppleVpnAlwaysOnConfigurationCellularExceptionAction = "AllowTrafficOutside"
	AppleVpnAlwaysOnConfigurationCellularExceptionActiondropTraffic         AppleVpnAlwaysOnConfigurationCellularExceptionAction = "DropTraffic"
	AppleVpnAlwaysOnConfigurationCellularExceptionActionforceTrafficViaVPN  AppleVpnAlwaysOnConfigurationCellularExceptionAction = "ForceTrafficViaVPN"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationCellularExceptionAction() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationCellularExceptionActionallowTrafficOutside),
		string(AppleVpnAlwaysOnConfigurationCellularExceptionActiondropTraffic),
		string(AppleVpnAlwaysOnConfigurationCellularExceptionActionforceTrafficViaVPN),
	}
}

func parseAppleVpnAlwaysOnConfigurationCellularExceptionAction(input string) (*AppleVpnAlwaysOnConfigurationCellularExceptionAction, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationCellularExceptionAction{
		"allowtrafficoutside": AppleVpnAlwaysOnConfigurationCellularExceptionActionallowTrafficOutside,
		"droptraffic":         AppleVpnAlwaysOnConfigurationCellularExceptionActiondropTraffic,
		"forcetrafficviavpn":  AppleVpnAlwaysOnConfigurationCellularExceptionActionforceTrafficViaVPN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationCellularExceptionAction(input)
	return &out, nil
}
