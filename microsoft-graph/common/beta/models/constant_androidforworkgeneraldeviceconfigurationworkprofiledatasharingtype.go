package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypeallowPersonalToWork AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "AllowPersonalToWork"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypedeviceDefault       AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "DeviceDefault"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypenoRestrictions      AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "NoRestrictions"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypepreventAny          AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "PreventAny"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypeallowPersonalToWork),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypedeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypenoRestrictions),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypepreventAny),
	}
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType{
		"allowpersonaltowork": AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypeallowPersonalToWork,
		"devicedefault":       AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypedeviceDefault,
		"norestrictions":      AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypenoRestrictions,
		"preventany":          AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingTypepreventAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType(input)
	return &out, nil
}
