package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivereddisabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered = "Disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDeliveredenabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered = "Enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDeliverednotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered = "NotConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivereddisabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDeliveredenabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDeliverednotConfigured),
	}
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivereddisabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDeliveredenabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDeliverednotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered(input)
	return &out, nil
}
