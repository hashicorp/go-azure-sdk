package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskEndpointSecurityPolicyProfile string

const (
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_AccountProtection            SecurityConfigurationTaskEndpointSecurityPolicyProfile = "accountProtection"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_Antivirus                    SecurityConfigurationTaskEndpointSecurityPolicyProfile = "antivirus"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_AppAndBrowserIsolation       SecurityConfigurationTaskEndpointSecurityPolicyProfile = "appAndBrowserIsolation"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_ApplicationControl           SecurityConfigurationTaskEndpointSecurityPolicyProfile = "applicationControl"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_AttackSurfaceReductionRules  SecurityConfigurationTaskEndpointSecurityPolicyProfile = "attackSurfaceReductionRules"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_BitLocker                    SecurityConfigurationTaskEndpointSecurityPolicyProfile = "bitLocker"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_DeviceControl                SecurityConfigurationTaskEndpointSecurityPolicyProfile = "deviceControl"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_EndpointDetectionAndResponse SecurityConfigurationTaskEndpointSecurityPolicyProfile = "endpointDetectionAndResponse"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_ExploitProtection            SecurityConfigurationTaskEndpointSecurityPolicyProfile = "exploitProtection"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_FileVault                    SecurityConfigurationTaskEndpointSecurityPolicyProfile = "fileVault"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_Firewall                     SecurityConfigurationTaskEndpointSecurityPolicyProfile = "firewall"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_FirewallRules                SecurityConfigurationTaskEndpointSecurityPolicyProfile = "firewallRules"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_Unknown                      SecurityConfigurationTaskEndpointSecurityPolicyProfile = "unknown"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_WebProtection                SecurityConfigurationTaskEndpointSecurityPolicyProfile = "webProtection"
	SecurityConfigurationTaskEndpointSecurityPolicyProfile_WindowsSecurity              SecurityConfigurationTaskEndpointSecurityPolicyProfile = "windowsSecurity"
)

func PossibleValuesForSecurityConfigurationTaskEndpointSecurityPolicyProfile() []string {
	return []string{
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_AccountProtection),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_Antivirus),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_AppAndBrowserIsolation),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_ApplicationControl),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_AttackSurfaceReductionRules),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_BitLocker),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_DeviceControl),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_EndpointDetectionAndResponse),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_ExploitProtection),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_FileVault),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_Firewall),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_FirewallRules),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_Unknown),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_WebProtection),
		string(SecurityConfigurationTaskEndpointSecurityPolicyProfile_WindowsSecurity),
	}
}

func (s *SecurityConfigurationTaskEndpointSecurityPolicyProfile) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityConfigurationTaskEndpointSecurityPolicyProfile(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityConfigurationTaskEndpointSecurityPolicyProfile(input string) (*SecurityConfigurationTaskEndpointSecurityPolicyProfile, error) {
	vals := map[string]SecurityConfigurationTaskEndpointSecurityPolicyProfile{
		"accountprotection":            SecurityConfigurationTaskEndpointSecurityPolicyProfile_AccountProtection,
		"antivirus":                    SecurityConfigurationTaskEndpointSecurityPolicyProfile_Antivirus,
		"appandbrowserisolation":       SecurityConfigurationTaskEndpointSecurityPolicyProfile_AppAndBrowserIsolation,
		"applicationcontrol":           SecurityConfigurationTaskEndpointSecurityPolicyProfile_ApplicationControl,
		"attacksurfacereductionrules":  SecurityConfigurationTaskEndpointSecurityPolicyProfile_AttackSurfaceReductionRules,
		"bitlocker":                    SecurityConfigurationTaskEndpointSecurityPolicyProfile_BitLocker,
		"devicecontrol":                SecurityConfigurationTaskEndpointSecurityPolicyProfile_DeviceControl,
		"endpointdetectionandresponse": SecurityConfigurationTaskEndpointSecurityPolicyProfile_EndpointDetectionAndResponse,
		"exploitprotection":            SecurityConfigurationTaskEndpointSecurityPolicyProfile_ExploitProtection,
		"filevault":                    SecurityConfigurationTaskEndpointSecurityPolicyProfile_FileVault,
		"firewall":                     SecurityConfigurationTaskEndpointSecurityPolicyProfile_Firewall,
		"firewallrules":                SecurityConfigurationTaskEndpointSecurityPolicyProfile_FirewallRules,
		"unknown":                      SecurityConfigurationTaskEndpointSecurityPolicyProfile_Unknown,
		"webprotection":                SecurityConfigurationTaskEndpointSecurityPolicyProfile_WebProtection,
		"windowssecurity":              SecurityConfigurationTaskEndpointSecurityPolicyProfile_WindowsSecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskEndpointSecurityPolicyProfile(input)
	return &out, nil
}
