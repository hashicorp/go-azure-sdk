package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationKioskModeAppType string

const (
	IosGeneralDeviceConfigurationKioskModeAppTypeappStoreApp   IosGeneralDeviceConfigurationKioskModeAppType = "AppStoreApp"
	IosGeneralDeviceConfigurationKioskModeAppTypebuiltInApp    IosGeneralDeviceConfigurationKioskModeAppType = "BuiltInApp"
	IosGeneralDeviceConfigurationKioskModeAppTypemanagedApp    IosGeneralDeviceConfigurationKioskModeAppType = "ManagedApp"
	IosGeneralDeviceConfigurationKioskModeAppTypenotConfigured IosGeneralDeviceConfigurationKioskModeAppType = "NotConfigured"
)

func PossibleValuesForIosGeneralDeviceConfigurationKioskModeAppType() []string {
	return []string{
		string(IosGeneralDeviceConfigurationKioskModeAppTypeappStoreApp),
		string(IosGeneralDeviceConfigurationKioskModeAppTypebuiltInApp),
		string(IosGeneralDeviceConfigurationKioskModeAppTypemanagedApp),
		string(IosGeneralDeviceConfigurationKioskModeAppTypenotConfigured),
	}
}

func parseIosGeneralDeviceConfigurationKioskModeAppType(input string) (*IosGeneralDeviceConfigurationKioskModeAppType, error) {
	vals := map[string]IosGeneralDeviceConfigurationKioskModeAppType{
		"appstoreapp":   IosGeneralDeviceConfigurationKioskModeAppTypeappStoreApp,
		"builtinapp":    IosGeneralDeviceConfigurationKioskModeAppTypebuiltInApp,
		"managedapp":    IosGeneralDeviceConfigurationKioskModeAppTypemanagedApp,
		"notconfigured": IosGeneralDeviceConfigurationKioskModeAppTypenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationKioskModeAppType(input)
	return &out, nil
}
