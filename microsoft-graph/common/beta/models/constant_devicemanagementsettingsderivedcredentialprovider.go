package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingsDerivedCredentialProvider string

const (
	DeviceManagementSettingsDerivedCredentialProviderentrustDataCard DeviceManagementSettingsDerivedCredentialProvider = "EntrustDataCard"
	DeviceManagementSettingsDerivedCredentialProviderintercede       DeviceManagementSettingsDerivedCredentialProvider = "Intercede"
	DeviceManagementSettingsDerivedCredentialProvidernotConfigured   DeviceManagementSettingsDerivedCredentialProvider = "NotConfigured"
	DeviceManagementSettingsDerivedCredentialProviderpurebred        DeviceManagementSettingsDerivedCredentialProvider = "Purebred"
	DeviceManagementSettingsDerivedCredentialProviderxTec            DeviceManagementSettingsDerivedCredentialProvider = "XTec"
)

func PossibleValuesForDeviceManagementSettingsDerivedCredentialProvider() []string {
	return []string{
		string(DeviceManagementSettingsDerivedCredentialProviderentrustDataCard),
		string(DeviceManagementSettingsDerivedCredentialProviderintercede),
		string(DeviceManagementSettingsDerivedCredentialProvidernotConfigured),
		string(DeviceManagementSettingsDerivedCredentialProviderpurebred),
		string(DeviceManagementSettingsDerivedCredentialProviderxTec),
	}
}

func parseDeviceManagementSettingsDerivedCredentialProvider(input string) (*DeviceManagementSettingsDerivedCredentialProvider, error) {
	vals := map[string]DeviceManagementSettingsDerivedCredentialProvider{
		"entrustdatacard": DeviceManagementSettingsDerivedCredentialProviderentrustDataCard,
		"intercede":       DeviceManagementSettingsDerivedCredentialProviderintercede,
		"notconfigured":   DeviceManagementSettingsDerivedCredentialProvidernotConfigured,
		"purebred":        DeviceManagementSettingsDerivedCredentialProviderpurebred,
		"xtec":            DeviceManagementSettingsDerivedCredentialProviderxTec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSettingsDerivedCredentialProvider(input)
	return &out, nil
}
