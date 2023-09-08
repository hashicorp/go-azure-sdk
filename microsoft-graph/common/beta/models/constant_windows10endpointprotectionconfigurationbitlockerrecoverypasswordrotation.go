package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation string

const (
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationdisabled                   Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "Disabled"
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationenabledForAzureAd          Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "EnabledForAzureAd"
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationenabledForAzureAdAndHybrid Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "EnabledForAzureAdAndHybrid"
	Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationnotConfigured              Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation = "NotConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationdisabled),
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationenabledForAzureAd),
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationenabledForAzureAdAndHybrid),
		string(Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationnotConfigured),
	}
}

func parseWindows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation(input string) (*Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation{
		"disabled":                   Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationdisabled,
		"enabledforazuread":          Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationenabledForAzureAd,
		"enabledforazureadandhybrid": Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationenabledForAzureAdAndHybrid,
		"notconfigured":              Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotationnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationBitLockerRecoveryPasswordRotation(input)
	return &out, nil
}
