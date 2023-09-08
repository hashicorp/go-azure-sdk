package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypeallowPersonalToWork AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "AllowPersonalToWork"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypedeviceDefault       AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "DeviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypenoRestrictions      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "NoRestrictions"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypepreventAny          AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "PreventAny"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypeallowPersonalToWork),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypedeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypenoRestrictions),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypepreventAny),
	}
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType{
		"allowpersonaltowork": AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypeallowPersonalToWork,
		"devicedefault":       AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypedeviceDefault,
		"norestrictions":      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypenoRestrictions,
		"preventany":          AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingTypepreventAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType(input)
	return &out, nil
}
