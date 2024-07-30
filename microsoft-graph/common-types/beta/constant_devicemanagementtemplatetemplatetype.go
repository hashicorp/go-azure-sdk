package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplateTemplateType string

const (
	DeviceManagementTemplateTemplateType_AdvancedThreatProtectionSecurityBaseline  DeviceManagementTemplateTemplateType = "advancedThreatProtectionSecurityBaseline"
	DeviceManagementTemplateTemplateType_CloudPC                                   DeviceManagementTemplateTemplateType = "cloudPC"
	DeviceManagementTemplateTemplateType_Custom                                    DeviceManagementTemplateTemplateType = "custom"
	DeviceManagementTemplateTemplateType_DeviceCompliance                          DeviceManagementTemplateTemplateType = "deviceCompliance"
	DeviceManagementTemplateTemplateType_DeviceConfiguration                       DeviceManagementTemplateTemplateType = "deviceConfiguration"
	DeviceManagementTemplateTemplateType_DeviceConfigurationForOffice365           DeviceManagementTemplateTemplateType = "deviceConfigurationForOffice365"
	DeviceManagementTemplateTemplateType_FirewallSharedSettings                    DeviceManagementTemplateTemplateType = "firewallSharedSettings"
	DeviceManagementTemplateTemplateType_MicrosoftEdgeSecurityBaseline             DeviceManagementTemplateTemplateType = "microsoftEdgeSecurityBaseline"
	DeviceManagementTemplateTemplateType_MicrosoftOffice365ProPlusSecurityBaseline DeviceManagementTemplateTemplateType = "microsoftOffice365ProPlusSecurityBaseline"
	DeviceManagementTemplateTemplateType_SecurityBaseline                          DeviceManagementTemplateTemplateType = "securityBaseline"
	DeviceManagementTemplateTemplateType_SecurityTemplate                          DeviceManagementTemplateTemplateType = "securityTemplate"
	DeviceManagementTemplateTemplateType_SpecializedDevices                        DeviceManagementTemplateTemplateType = "specializedDevices"
)

func PossibleValuesForDeviceManagementTemplateTemplateType() []string {
	return []string{
		string(DeviceManagementTemplateTemplateType_AdvancedThreatProtectionSecurityBaseline),
		string(DeviceManagementTemplateTemplateType_CloudPC),
		string(DeviceManagementTemplateTemplateType_Custom),
		string(DeviceManagementTemplateTemplateType_DeviceCompliance),
		string(DeviceManagementTemplateTemplateType_DeviceConfiguration),
		string(DeviceManagementTemplateTemplateType_DeviceConfigurationForOffice365),
		string(DeviceManagementTemplateTemplateType_FirewallSharedSettings),
		string(DeviceManagementTemplateTemplateType_MicrosoftEdgeSecurityBaseline),
		string(DeviceManagementTemplateTemplateType_MicrosoftOffice365ProPlusSecurityBaseline),
		string(DeviceManagementTemplateTemplateType_SecurityBaseline),
		string(DeviceManagementTemplateTemplateType_SecurityTemplate),
		string(DeviceManagementTemplateTemplateType_SpecializedDevices),
	}
}

func (s *DeviceManagementTemplateTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementTemplateTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementTemplateTemplateType(input string) (*DeviceManagementTemplateTemplateType, error) {
	vals := map[string]DeviceManagementTemplateTemplateType{
		"advancedthreatprotectionsecuritybaseline": DeviceManagementTemplateTemplateType_AdvancedThreatProtectionSecurityBaseline,
		"cloudpc":                                   DeviceManagementTemplateTemplateType_CloudPC,
		"custom":                                    DeviceManagementTemplateTemplateType_Custom,
		"devicecompliance":                          DeviceManagementTemplateTemplateType_DeviceCompliance,
		"deviceconfiguration":                       DeviceManagementTemplateTemplateType_DeviceConfiguration,
		"deviceconfigurationforoffice365":           DeviceManagementTemplateTemplateType_DeviceConfigurationForOffice365,
		"firewallsharedsettings":                    DeviceManagementTemplateTemplateType_FirewallSharedSettings,
		"microsoftedgesecuritybaseline":             DeviceManagementTemplateTemplateType_MicrosoftEdgeSecurityBaseline,
		"microsoftoffice365proplussecuritybaseline": DeviceManagementTemplateTemplateType_MicrosoftOffice365ProPlusSecurityBaseline,
		"securitybaseline":                          DeviceManagementTemplateTemplateType_SecurityBaseline,
		"securitytemplate":                          DeviceManagementTemplateTemplateType_SecurityTemplate,
		"specializeddevices":                        DeviceManagementTemplateTemplateType_SpecializedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplateTemplateType(input)
	return &out, nil
}
