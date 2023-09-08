package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialSettingsNotificationType string

const (
	DeviceManagementDerivedCredentialSettingsNotificationTypecompanyPortal DeviceManagementDerivedCredentialSettingsNotificationType = "CompanyPortal"
	DeviceManagementDerivedCredentialSettingsNotificationTypeemail         DeviceManagementDerivedCredentialSettingsNotificationType = "Email"
	DeviceManagementDerivedCredentialSettingsNotificationTypenone          DeviceManagementDerivedCredentialSettingsNotificationType = "None"
)

func PossibleValuesForDeviceManagementDerivedCredentialSettingsNotificationType() []string {
	return []string{
		string(DeviceManagementDerivedCredentialSettingsNotificationTypecompanyPortal),
		string(DeviceManagementDerivedCredentialSettingsNotificationTypeemail),
		string(DeviceManagementDerivedCredentialSettingsNotificationTypenone),
	}
}

func parseDeviceManagementDerivedCredentialSettingsNotificationType(input string) (*DeviceManagementDerivedCredentialSettingsNotificationType, error) {
	vals := map[string]DeviceManagementDerivedCredentialSettingsNotificationType{
		"companyportal": DeviceManagementDerivedCredentialSettingsNotificationTypecompanyPortal,
		"email":         DeviceManagementDerivedCredentialSettingsNotificationTypeemail,
		"none":          DeviceManagementDerivedCredentialSettingsNotificationTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDerivedCredentialSettingsNotificationType(input)
	return &out, nil
}
