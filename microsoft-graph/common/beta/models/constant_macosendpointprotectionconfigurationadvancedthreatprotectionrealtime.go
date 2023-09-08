package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimedisabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime = "Disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimeenabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime = "Enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimenotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime = "NotConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimedisabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimeenabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimenotConfigured),
	}
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimedisabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimeenabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTimenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime(input)
	return &out, nil
}
