package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplateTemplateType string

const (
	DeviceManagementTemplateTemplateTypeadvancedThreatProtectionSecurityBaseline  DeviceManagementTemplateTemplateType = "AdvancedThreatProtectionSecurityBaseline"
	DeviceManagementTemplateTemplateTypecloudPC                                   DeviceManagementTemplateTemplateType = "CloudPC"
	DeviceManagementTemplateTemplateTypecustom                                    DeviceManagementTemplateTemplateType = "Custom"
	DeviceManagementTemplateTemplateTypedeviceCompliance                          DeviceManagementTemplateTemplateType = "DeviceCompliance"
	DeviceManagementTemplateTemplateTypedeviceConfiguration                       DeviceManagementTemplateTemplateType = "DeviceConfiguration"
	DeviceManagementTemplateTemplateTypedeviceConfigurationForOffice365           DeviceManagementTemplateTemplateType = "DeviceConfigurationForOffice365"
	DeviceManagementTemplateTemplateTypefirewallSharedSettings                    DeviceManagementTemplateTemplateType = "FirewallSharedSettings"
	DeviceManagementTemplateTemplateTypemicrosoftEdgeSecurityBaseline             DeviceManagementTemplateTemplateType = "MicrosoftEdgeSecurityBaseline"
	DeviceManagementTemplateTemplateTypemicrosoftOffice365ProPlusSecurityBaseline DeviceManagementTemplateTemplateType = "MicrosoftOffice365ProPlusSecurityBaseline"
	DeviceManagementTemplateTemplateTypesecurityBaseline                          DeviceManagementTemplateTemplateType = "SecurityBaseline"
	DeviceManagementTemplateTemplateTypesecurityTemplate                          DeviceManagementTemplateTemplateType = "SecurityTemplate"
	DeviceManagementTemplateTemplateTypespecializedDevices                        DeviceManagementTemplateTemplateType = "SpecializedDevices"
)

func PossibleValuesForDeviceManagementTemplateTemplateType() []string {
	return []string{
		string(DeviceManagementTemplateTemplateTypeadvancedThreatProtectionSecurityBaseline),
		string(DeviceManagementTemplateTemplateTypecloudPC),
		string(DeviceManagementTemplateTemplateTypecustom),
		string(DeviceManagementTemplateTemplateTypedeviceCompliance),
		string(DeviceManagementTemplateTemplateTypedeviceConfiguration),
		string(DeviceManagementTemplateTemplateTypedeviceConfigurationForOffice365),
		string(DeviceManagementTemplateTemplateTypefirewallSharedSettings),
		string(DeviceManagementTemplateTemplateTypemicrosoftEdgeSecurityBaseline),
		string(DeviceManagementTemplateTemplateTypemicrosoftOffice365ProPlusSecurityBaseline),
		string(DeviceManagementTemplateTemplateTypesecurityBaseline),
		string(DeviceManagementTemplateTemplateTypesecurityTemplate),
		string(DeviceManagementTemplateTemplateTypespecializedDevices),
	}
}

func parseDeviceManagementTemplateTemplateType(input string) (*DeviceManagementTemplateTemplateType, error) {
	vals := map[string]DeviceManagementTemplateTemplateType{
		"advancedthreatprotectionsecuritybaseline": DeviceManagementTemplateTemplateTypeadvancedThreatProtectionSecurityBaseline,
		"cloudpc":                                   DeviceManagementTemplateTemplateTypecloudPC,
		"custom":                                    DeviceManagementTemplateTemplateTypecustom,
		"devicecompliance":                          DeviceManagementTemplateTemplateTypedeviceCompliance,
		"deviceconfiguration":                       DeviceManagementTemplateTemplateTypedeviceConfiguration,
		"deviceconfigurationforoffice365":           DeviceManagementTemplateTemplateTypedeviceConfigurationForOffice365,
		"firewallsharedsettings":                    DeviceManagementTemplateTemplateTypefirewallSharedSettings,
		"microsoftedgesecuritybaseline":             DeviceManagementTemplateTemplateTypemicrosoftEdgeSecurityBaseline,
		"microsoftoffice365proplussecuritybaseline": DeviceManagementTemplateTemplateTypemicrosoftOffice365ProPlusSecurityBaseline,
		"securitybaseline":                          DeviceManagementTemplateTemplateTypesecurityBaseline,
		"securitytemplate":                          DeviceManagementTemplateTemplateTypesecurityTemplate,
		"specializeddevices":                        DeviceManagementTemplateTemplateTypespecializedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplateTemplateType(input)
	return &out, nil
}
