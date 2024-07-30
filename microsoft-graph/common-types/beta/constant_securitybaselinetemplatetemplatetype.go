package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineTemplateTemplateType string

const (
	SecurityBaselineTemplateTemplateType_AdvancedThreatProtectionSecurityBaseline  SecurityBaselineTemplateTemplateType = "advancedThreatProtectionSecurityBaseline"
	SecurityBaselineTemplateTemplateType_CloudPC                                   SecurityBaselineTemplateTemplateType = "cloudPC"
	SecurityBaselineTemplateTemplateType_Custom                                    SecurityBaselineTemplateTemplateType = "custom"
	SecurityBaselineTemplateTemplateType_DeviceCompliance                          SecurityBaselineTemplateTemplateType = "deviceCompliance"
	SecurityBaselineTemplateTemplateType_DeviceConfiguration                       SecurityBaselineTemplateTemplateType = "deviceConfiguration"
	SecurityBaselineTemplateTemplateType_DeviceConfigurationForOffice365           SecurityBaselineTemplateTemplateType = "deviceConfigurationForOffice365"
	SecurityBaselineTemplateTemplateType_FirewallSharedSettings                    SecurityBaselineTemplateTemplateType = "firewallSharedSettings"
	SecurityBaselineTemplateTemplateType_MicrosoftEdgeSecurityBaseline             SecurityBaselineTemplateTemplateType = "microsoftEdgeSecurityBaseline"
	SecurityBaselineTemplateTemplateType_MicrosoftOffice365ProPlusSecurityBaseline SecurityBaselineTemplateTemplateType = "microsoftOffice365ProPlusSecurityBaseline"
	SecurityBaselineTemplateTemplateType_SecurityBaseline                          SecurityBaselineTemplateTemplateType = "securityBaseline"
	SecurityBaselineTemplateTemplateType_SecurityTemplate                          SecurityBaselineTemplateTemplateType = "securityTemplate"
	SecurityBaselineTemplateTemplateType_SpecializedDevices                        SecurityBaselineTemplateTemplateType = "specializedDevices"
)

func PossibleValuesForSecurityBaselineTemplateTemplateType() []string {
	return []string{
		string(SecurityBaselineTemplateTemplateType_AdvancedThreatProtectionSecurityBaseline),
		string(SecurityBaselineTemplateTemplateType_CloudPC),
		string(SecurityBaselineTemplateTemplateType_Custom),
		string(SecurityBaselineTemplateTemplateType_DeviceCompliance),
		string(SecurityBaselineTemplateTemplateType_DeviceConfiguration),
		string(SecurityBaselineTemplateTemplateType_DeviceConfigurationForOffice365),
		string(SecurityBaselineTemplateTemplateType_FirewallSharedSettings),
		string(SecurityBaselineTemplateTemplateType_MicrosoftEdgeSecurityBaseline),
		string(SecurityBaselineTemplateTemplateType_MicrosoftOffice365ProPlusSecurityBaseline),
		string(SecurityBaselineTemplateTemplateType_SecurityBaseline),
		string(SecurityBaselineTemplateTemplateType_SecurityTemplate),
		string(SecurityBaselineTemplateTemplateType_SpecializedDevices),
	}
}

func (s *SecurityBaselineTemplateTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineTemplateTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineTemplateTemplateType(input string) (*SecurityBaselineTemplateTemplateType, error) {
	vals := map[string]SecurityBaselineTemplateTemplateType{
		"advancedthreatprotectionsecuritybaseline": SecurityBaselineTemplateTemplateType_AdvancedThreatProtectionSecurityBaseline,
		"cloudpc":                                   SecurityBaselineTemplateTemplateType_CloudPC,
		"custom":                                    SecurityBaselineTemplateTemplateType_Custom,
		"devicecompliance":                          SecurityBaselineTemplateTemplateType_DeviceCompliance,
		"deviceconfiguration":                       SecurityBaselineTemplateTemplateType_DeviceConfiguration,
		"deviceconfigurationforoffice365":           SecurityBaselineTemplateTemplateType_DeviceConfigurationForOffice365,
		"firewallsharedsettings":                    SecurityBaselineTemplateTemplateType_FirewallSharedSettings,
		"microsoftedgesecuritybaseline":             SecurityBaselineTemplateTemplateType_MicrosoftEdgeSecurityBaseline,
		"microsoftoffice365proplussecuritybaseline": SecurityBaselineTemplateTemplateType_MicrosoftOffice365ProPlusSecurityBaseline,
		"securitybaseline":                          SecurityBaselineTemplateTemplateType_SecurityBaseline,
		"securitytemplate":                          SecurityBaselineTemplateTemplateType_SecurityTemplate,
		"specializeddevices":                        SecurityBaselineTemplateTemplateType_SpecializedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineTemplateTemplateType(input)
	return &out, nil
}
