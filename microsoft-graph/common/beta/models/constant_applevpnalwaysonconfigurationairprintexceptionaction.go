package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationAirPrintExceptionAction string

const (
	AppleVpnAlwaysOnConfigurationAirPrintExceptionActionallowTrafficOutside AppleVpnAlwaysOnConfigurationAirPrintExceptionAction = "AllowTrafficOutside"
	AppleVpnAlwaysOnConfigurationAirPrintExceptionActiondropTraffic         AppleVpnAlwaysOnConfigurationAirPrintExceptionAction = "DropTraffic"
	AppleVpnAlwaysOnConfigurationAirPrintExceptionActionforceTrafficViaVPN  AppleVpnAlwaysOnConfigurationAirPrintExceptionAction = "ForceTrafficViaVPN"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationAirPrintExceptionAction() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationAirPrintExceptionActionallowTrafficOutside),
		string(AppleVpnAlwaysOnConfigurationAirPrintExceptionActiondropTraffic),
		string(AppleVpnAlwaysOnConfigurationAirPrintExceptionActionforceTrafficViaVPN),
	}
}

func parseAppleVpnAlwaysOnConfigurationAirPrintExceptionAction(input string) (*AppleVpnAlwaysOnConfigurationAirPrintExceptionAction, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationAirPrintExceptionAction{
		"allowtrafficoutside": AppleVpnAlwaysOnConfigurationAirPrintExceptionActionallowTrafficOutside,
		"droptraffic":         AppleVpnAlwaysOnConfigurationAirPrintExceptionActiondropTraffic,
		"forcetrafficviavpn":  AppleVpnAlwaysOnConfigurationAirPrintExceptionActionforceTrafficViaVPN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationAirPrintExceptionAction(input)
	return &out, nil
}
