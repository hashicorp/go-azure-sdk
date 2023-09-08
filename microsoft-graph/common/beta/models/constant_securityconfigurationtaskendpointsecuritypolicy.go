package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskEndpointSecurityPolicy string

const (
	SecurityConfigurationTaskEndpointSecurityPolicyaccountProtection            SecurityConfigurationTaskEndpointSecurityPolicy = "AccountProtection"
	SecurityConfigurationTaskEndpointSecurityPolicyantivirus                    SecurityConfigurationTaskEndpointSecurityPolicy = "Antivirus"
	SecurityConfigurationTaskEndpointSecurityPolicyattackSurfaceReduction       SecurityConfigurationTaskEndpointSecurityPolicy = "AttackSurfaceReduction"
	SecurityConfigurationTaskEndpointSecurityPolicydiskEncryption               SecurityConfigurationTaskEndpointSecurityPolicy = "DiskEncryption"
	SecurityConfigurationTaskEndpointSecurityPolicyendpointDetectionAndResponse SecurityConfigurationTaskEndpointSecurityPolicy = "EndpointDetectionAndResponse"
	SecurityConfigurationTaskEndpointSecurityPolicyfirewall                     SecurityConfigurationTaskEndpointSecurityPolicy = "Firewall"
	SecurityConfigurationTaskEndpointSecurityPolicyunknown                      SecurityConfigurationTaskEndpointSecurityPolicy = "Unknown"
)

func PossibleValuesForSecurityConfigurationTaskEndpointSecurityPolicy() []string {
	return []string{
		string(SecurityConfigurationTaskEndpointSecurityPolicyaccountProtection),
		string(SecurityConfigurationTaskEndpointSecurityPolicyantivirus),
		string(SecurityConfigurationTaskEndpointSecurityPolicyattackSurfaceReduction),
		string(SecurityConfigurationTaskEndpointSecurityPolicydiskEncryption),
		string(SecurityConfigurationTaskEndpointSecurityPolicyendpointDetectionAndResponse),
		string(SecurityConfigurationTaskEndpointSecurityPolicyfirewall),
		string(SecurityConfigurationTaskEndpointSecurityPolicyunknown),
	}
}

func parseSecurityConfigurationTaskEndpointSecurityPolicy(input string) (*SecurityConfigurationTaskEndpointSecurityPolicy, error) {
	vals := map[string]SecurityConfigurationTaskEndpointSecurityPolicy{
		"accountprotection":            SecurityConfigurationTaskEndpointSecurityPolicyaccountProtection,
		"antivirus":                    SecurityConfigurationTaskEndpointSecurityPolicyantivirus,
		"attacksurfacereduction":       SecurityConfigurationTaskEndpointSecurityPolicyattackSurfaceReduction,
		"diskencryption":               SecurityConfigurationTaskEndpointSecurityPolicydiskEncryption,
		"endpointdetectionandresponse": SecurityConfigurationTaskEndpointSecurityPolicyendpointDetectionAndResponse,
		"firewall":                     SecurityConfigurationTaskEndpointSecurityPolicyfirewall,
		"unknown":                      SecurityConfigurationTaskEndpointSecurityPolicyunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskEndpointSecurityPolicy(input)
	return &out, nil
}
