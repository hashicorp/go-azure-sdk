package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskEndpointSecurityPolicy string

const (
	SecurityConfigurationTaskEndpointSecurityPolicy_AccountProtection            SecurityConfigurationTaskEndpointSecurityPolicy = "accountProtection"
	SecurityConfigurationTaskEndpointSecurityPolicy_Antivirus                    SecurityConfigurationTaskEndpointSecurityPolicy = "antivirus"
	SecurityConfigurationTaskEndpointSecurityPolicy_AttackSurfaceReduction       SecurityConfigurationTaskEndpointSecurityPolicy = "attackSurfaceReduction"
	SecurityConfigurationTaskEndpointSecurityPolicy_DiskEncryption               SecurityConfigurationTaskEndpointSecurityPolicy = "diskEncryption"
	SecurityConfigurationTaskEndpointSecurityPolicy_EndpointDetectionAndResponse SecurityConfigurationTaskEndpointSecurityPolicy = "endpointDetectionAndResponse"
	SecurityConfigurationTaskEndpointSecurityPolicy_Firewall                     SecurityConfigurationTaskEndpointSecurityPolicy = "firewall"
	SecurityConfigurationTaskEndpointSecurityPolicy_Unknown                      SecurityConfigurationTaskEndpointSecurityPolicy = "unknown"
)

func PossibleValuesForSecurityConfigurationTaskEndpointSecurityPolicy() []string {
	return []string{
		string(SecurityConfigurationTaskEndpointSecurityPolicy_AccountProtection),
		string(SecurityConfigurationTaskEndpointSecurityPolicy_Antivirus),
		string(SecurityConfigurationTaskEndpointSecurityPolicy_AttackSurfaceReduction),
		string(SecurityConfigurationTaskEndpointSecurityPolicy_DiskEncryption),
		string(SecurityConfigurationTaskEndpointSecurityPolicy_EndpointDetectionAndResponse),
		string(SecurityConfigurationTaskEndpointSecurityPolicy_Firewall),
		string(SecurityConfigurationTaskEndpointSecurityPolicy_Unknown),
	}
}

func (s *SecurityConfigurationTaskEndpointSecurityPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityConfigurationTaskEndpointSecurityPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityConfigurationTaskEndpointSecurityPolicy(input string) (*SecurityConfigurationTaskEndpointSecurityPolicy, error) {
	vals := map[string]SecurityConfigurationTaskEndpointSecurityPolicy{
		"accountprotection":            SecurityConfigurationTaskEndpointSecurityPolicy_AccountProtection,
		"antivirus":                    SecurityConfigurationTaskEndpointSecurityPolicy_Antivirus,
		"attacksurfacereduction":       SecurityConfigurationTaskEndpointSecurityPolicy_AttackSurfaceReduction,
		"diskencryption":               SecurityConfigurationTaskEndpointSecurityPolicy_DiskEncryption,
		"endpointdetectionandresponse": SecurityConfigurationTaskEndpointSecurityPolicy_EndpointDetectionAndResponse,
		"firewall":                     SecurityConfigurationTaskEndpointSecurityPolicy_Firewall,
		"unknown":                      SecurityConfigurationTaskEndpointSecurityPolicy_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskEndpointSecurityPolicy(input)
	return &out, nil
}
