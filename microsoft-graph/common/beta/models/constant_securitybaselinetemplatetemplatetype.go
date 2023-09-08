package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineTemplateTemplateType string

const (
	SecurityBaselineTemplateTemplateTypeadvancedThreatProtectionSecurityBaseline  SecurityBaselineTemplateTemplateType = "AdvancedThreatProtectionSecurityBaseline"
	SecurityBaselineTemplateTemplateTypecloudPC                                   SecurityBaselineTemplateTemplateType = "CloudPC"
	SecurityBaselineTemplateTemplateTypecustom                                    SecurityBaselineTemplateTemplateType = "Custom"
	SecurityBaselineTemplateTemplateTypedeviceCompliance                          SecurityBaselineTemplateTemplateType = "DeviceCompliance"
	SecurityBaselineTemplateTemplateTypedeviceConfiguration                       SecurityBaselineTemplateTemplateType = "DeviceConfiguration"
	SecurityBaselineTemplateTemplateTypedeviceConfigurationForOffice365           SecurityBaselineTemplateTemplateType = "DeviceConfigurationForOffice365"
	SecurityBaselineTemplateTemplateTypefirewallSharedSettings                    SecurityBaselineTemplateTemplateType = "FirewallSharedSettings"
	SecurityBaselineTemplateTemplateTypemicrosoftEdgeSecurityBaseline             SecurityBaselineTemplateTemplateType = "MicrosoftEdgeSecurityBaseline"
	SecurityBaselineTemplateTemplateTypemicrosoftOffice365ProPlusSecurityBaseline SecurityBaselineTemplateTemplateType = "MicrosoftOffice365ProPlusSecurityBaseline"
	SecurityBaselineTemplateTemplateTypesecurityBaseline                          SecurityBaselineTemplateTemplateType = "SecurityBaseline"
	SecurityBaselineTemplateTemplateTypesecurityTemplate                          SecurityBaselineTemplateTemplateType = "SecurityTemplate"
	SecurityBaselineTemplateTemplateTypespecializedDevices                        SecurityBaselineTemplateTemplateType = "SpecializedDevices"
)

func PossibleValuesForSecurityBaselineTemplateTemplateType() []string {
	return []string{
		string(SecurityBaselineTemplateTemplateTypeadvancedThreatProtectionSecurityBaseline),
		string(SecurityBaselineTemplateTemplateTypecloudPC),
		string(SecurityBaselineTemplateTemplateTypecustom),
		string(SecurityBaselineTemplateTemplateTypedeviceCompliance),
		string(SecurityBaselineTemplateTemplateTypedeviceConfiguration),
		string(SecurityBaselineTemplateTemplateTypedeviceConfigurationForOffice365),
		string(SecurityBaselineTemplateTemplateTypefirewallSharedSettings),
		string(SecurityBaselineTemplateTemplateTypemicrosoftEdgeSecurityBaseline),
		string(SecurityBaselineTemplateTemplateTypemicrosoftOffice365ProPlusSecurityBaseline),
		string(SecurityBaselineTemplateTemplateTypesecurityBaseline),
		string(SecurityBaselineTemplateTemplateTypesecurityTemplate),
		string(SecurityBaselineTemplateTemplateTypespecializedDevices),
	}
}

func parseSecurityBaselineTemplateTemplateType(input string) (*SecurityBaselineTemplateTemplateType, error) {
	vals := map[string]SecurityBaselineTemplateTemplateType{
		"advancedthreatprotectionsecuritybaseline": SecurityBaselineTemplateTemplateTypeadvancedThreatProtectionSecurityBaseline,
		"cloudpc":                                   SecurityBaselineTemplateTemplateTypecloudPC,
		"custom":                                    SecurityBaselineTemplateTemplateTypecustom,
		"devicecompliance":                          SecurityBaselineTemplateTemplateTypedeviceCompliance,
		"deviceconfiguration":                       SecurityBaselineTemplateTemplateTypedeviceConfiguration,
		"deviceconfigurationforoffice365":           SecurityBaselineTemplateTemplateTypedeviceConfigurationForOffice365,
		"firewallsharedsettings":                    SecurityBaselineTemplateTemplateTypefirewallSharedSettings,
		"microsoftedgesecuritybaseline":             SecurityBaselineTemplateTemplateTypemicrosoftEdgeSecurityBaseline,
		"microsoftoffice365proplussecuritybaseline": SecurityBaselineTemplateTemplateTypemicrosoftOffice365ProPlusSecurityBaseline,
		"securitybaseline":                          SecurityBaselineTemplateTemplateTypesecurityBaseline,
		"securitytemplate":                          SecurityBaselineTemplateTemplateTypesecurityTemplate,
		"specializeddevices":                        SecurityBaselineTemplateTemplateTypespecializedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineTemplateTemplateType(input)
	return &out, nil
}
