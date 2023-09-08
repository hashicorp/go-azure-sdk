package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationVoicemailExceptionAction string

const (
	AppleVpnAlwaysOnConfigurationVoicemailExceptionActionallowTrafficOutside AppleVpnAlwaysOnConfigurationVoicemailExceptionAction = "AllowTrafficOutside"
	AppleVpnAlwaysOnConfigurationVoicemailExceptionActiondropTraffic         AppleVpnAlwaysOnConfigurationVoicemailExceptionAction = "DropTraffic"
	AppleVpnAlwaysOnConfigurationVoicemailExceptionActionforceTrafficViaVPN  AppleVpnAlwaysOnConfigurationVoicemailExceptionAction = "ForceTrafficViaVPN"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationVoicemailExceptionAction() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationVoicemailExceptionActionallowTrafficOutside),
		string(AppleVpnAlwaysOnConfigurationVoicemailExceptionActiondropTraffic),
		string(AppleVpnAlwaysOnConfigurationVoicemailExceptionActionforceTrafficViaVPN),
	}
}

func parseAppleVpnAlwaysOnConfigurationVoicemailExceptionAction(input string) (*AppleVpnAlwaysOnConfigurationVoicemailExceptionAction, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationVoicemailExceptionAction{
		"allowtrafficoutside": AppleVpnAlwaysOnConfigurationVoicemailExceptionActionallowTrafficOutside,
		"droptraffic":         AppleVpnAlwaysOnConfigurationVoicemailExceptionActiondropTraffic,
		"forcetrafficviavpn":  AppleVpnAlwaysOnConfigurationVoicemailExceptionActionforceTrafficViaVPN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationVoicemailExceptionAction(input)
	return &out, nil
}
