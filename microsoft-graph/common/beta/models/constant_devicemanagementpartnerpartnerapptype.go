package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerPartnerAppType string

const (
	DeviceManagementPartnerPartnerAppTypemultiTenantApp  DeviceManagementPartnerPartnerAppType = "MultiTenantApp"
	DeviceManagementPartnerPartnerAppTypesingleTenantApp DeviceManagementPartnerPartnerAppType = "SingleTenantApp"
	DeviceManagementPartnerPartnerAppTypeunknown         DeviceManagementPartnerPartnerAppType = "Unknown"
)

func PossibleValuesForDeviceManagementPartnerPartnerAppType() []string {
	return []string{
		string(DeviceManagementPartnerPartnerAppTypemultiTenantApp),
		string(DeviceManagementPartnerPartnerAppTypesingleTenantApp),
		string(DeviceManagementPartnerPartnerAppTypeunknown),
	}
}

func parseDeviceManagementPartnerPartnerAppType(input string) (*DeviceManagementPartnerPartnerAppType, error) {
	vals := map[string]DeviceManagementPartnerPartnerAppType{
		"multitenantapp":  DeviceManagementPartnerPartnerAppTypemultiTenantApp,
		"singletenantapp": DeviceManagementPartnerPartnerAppTypesingleTenantApp,
		"unknown":         DeviceManagementPartnerPartnerAppTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPartnerPartnerAppType(input)
	return &out, nil
}
